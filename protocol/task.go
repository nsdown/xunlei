package protocol

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/golang/glog"
	"github.com/zyxar/taipei"
)

var stats []string
var coloring []string

func init() {
	coloring = make([]string, 8)
	stats = make([]string, 8)
	coloring[0] = color_front_yellow
	coloring[1] = color_front_magenta
	coloring[2] = color_front_green
	coloring[3] = color_front_red
	coloring[6] = color_front_cyan
	coloring[5] = color_front_blue
	coloring[4] = color_reset
	stats[0] = color_bg_yellow + "waiting" + color_reset
	stats[1] = color_bg_magenta + "downloading" + color_reset
	stats[2] = color_bg_green + "completed" + color_reset
	stats[3] = color_bg_red + "failed" + color_reset
	stats[5] = color_bg_blue + "pending" + color_reset
	stats[6] = color_bg_cyan + "expired" + color_reset
}

func trim(raw string) string {
	exp := regexp.MustCompile(`<font color='([a-z]*)'>(.*)</font>`)
	s := exp.FindStringSubmatch(raw)
	if s == nil {
		return raw
	}
	return s[2]
}

func (this Task) Coloring() string {
	j, _ := strconv.Atoi(this.DownloadStatus)
	k, _ := strconv.Atoi(this.Flag)
	if k == 4 {
		j += k
	}
	status := stats[j]
	return fmt.Sprintf("%s%s %s %s %s%s %.1f%% %s%s", coloring[j], this.Id, this.TaskName, status, coloring[j], this.FileSize, this.Progress, trim(this.LeftLiveTime), color_reset)
}

func (this Task) String() string {
	return fmt.Sprintf("%s %s [%s] %s %.1f%% %s", this.Id, this.TaskName, this.DownloadStatus, this.FileSize, this.Progress, trim(this.LeftLiveTime))
}

func (this Task) Repr() string {
	j, _ := strconv.Atoi(this.DownloadStatus)
	k, _ := strconv.Atoi(this.Flag)
	if k == 4 {
		j += k
	}
	status := stats[j]
	ret := coloring[j] + this.Id + " " + this.TaskName + " " + status + coloring[j] + " " + this.FileSize + " " + trim(this.LeftLiveTime) + "\n"
	if this.Cid != "" {
		ret += this.Cid + " "
	}
	if this.GCid != "" {
		ret += this.GCid + "\n"
	}
	ret += this.URL
	if this.LixianURL != "" {
		ret += "\n" + this.LixianURL
	}
	return ret + color_reset
}

func (this Task) expired() bool {
	return this.status() == _FLAG_expired
}

func (this Task) purged() bool {
	return this.status() == _FLAG_purged
}

func (this Task) deleted() bool {
	return this.status() == _FLAG_deleted
}

func (this Task) normal() bool {
	return this.status() == _FLAG_normal
}

func (this Task) IsBt() bool {
	return this.TaskType == 0
}

func (this Task) waiting() bool {
	return this.DownloadStatus == "0"
}

func (this Task) completed() bool {
	return this.DownloadStatus == "2"
}

func (this Task) downloading() bool {
	return this.DownloadStatus == "1"
}

func (this Task) failed() bool {
	return this.DownloadStatus == "3"
}

func (this Task) pending() bool {
	return this.DownloadStatus == "5"
}

func (this Task) status() byte {
	switch len(this.Flag) {
	case 0:
		return _FLAG_normal
	case 1:
		t := this.Flag[0] - '0'
		if t < 5 {
			return t
		}
	}
	return _FLAG_invalid
}

func (this *Task) update(t *_ptask_record) {
	if this.Id != t.Id {
		return
	}
	this.Speed = t.Speed
	this.Progress = t.Progress
	this.DownloadStatus = t.DownloadStatus
	this.LixianURL = t.LixianURL
}

func (this *Task) FillBtList() (*bt_list, error) {
	if !this.IsBt() {
		return nil, errors.New("Not BT task.")
	}
	return FillBtList(this.Id, this.Cid)
}

func (this *Task) Remove() error {
	return this.remove(0)
}

func (this *Task) Purge() error {
	if this.deleted() {
		return this.remove(1)
	}
	err := this.remove(0)
	if err != nil {
		return err
	}
	return this.remove(1)
}

func (this *Task) remove(flag byte) error {
	var del_type byte = this.status()
	if del_type == _FLAG_invalid {
		return errors.New("Invalid flag in task.")
	} else if del_type == _FLAG_purged {
		return errors.New("Task already purged.")
	} else if flag == 0 && del_type == _FLAG_deleted {
		return errors.New("Task already deleted.")
	}
	ct := currentTimestamp()
	uri := fmt.Sprintf(TASKDELETE_URL, ct, del_type, ct)
	data := url.Values{}
	data.Add("taskids", this.Id+",")
	data.Add("databases", "0,")
	data.Add("interfrom", "task")
	r, err := post(uri, data.Encode())
	if err != nil {
		return err
	}
	if ok, _ := regexp.Match(`\{"result":1,"type":`, r); ok {
		glog.V(2).Infof("%s\n", r)
		if this.status() == _FLAG_deleted {
			this.Flag = "2"
		} else {
			this.Flag = "1"
		}
		this.Progress = 0
		return nil
	}
	return unexpectedErr
}

func (this *Task) Rename(name string) error {
	return rename_task(this.Id, name, this.TaskType)
}

func (this *Task) Pause() error {
	tids := this.Id + ","
	uri := fmt.Sprintf(TASKPAUSE_URL, tids, M.Uid, currentTimestamp())
	r, err := get(uri)
	if err != nil {
		return err
	}
	if bytes.Compare(r, []byte("pause_task_resp()")) != 0 {
		return invalidResponseErr
	}
	return nil
}

func (this *Task) Readd() error {
	if this.normal() {
		return errors.New("Task already in progress.")
	}
	if this.purged() {
		return addSimpleTask(this.URL)
	}
	return addSimpleTask(this.URL, this.Id)
}

func (this *Task) Resume() error {
	if this.expired() {
		return taskNoRedownCapErr
	}
	status := this.DownloadStatus
	if status != "5" && status != "3" {
		return taskNoRedownCapErr // only valid for `pending` and `failed` tasks
	}
	form := make([]string, 0, 3)
	v := url.Values{}
	v.Add("id[]", this.Id)
	v.Add("url[]", this.URL)
	v.Add("cid[]", this.Cid)
	v.Add("download_status[]", status)
	v.Add("taskname[]", this.TaskName)
	form = append(form, v.Encode())
	form = append(form, "type=1")
	form = append(form, "interfrom=task")
	uri := fmt.Sprintf(REDOWNLOAD_URL, currentTimestamp())
	r, err := post(uri, strings.Join(form, "&"))
	if err != nil {
		return err
	}
	glog.V(2).Infof("%s\n", r)
	return nil
}

func (this *Task) Delay() error {
	return DelayTask(this.Id)
}

func (this Task) GetVodURL() (lurl, hurl string, err error) {
	sid := getCookie("http://xunlei.com", "sessionid")
	v := url.Values{}
	v.Add("url", this.URL)
	v.Add("video_name", this.TaskName)
	v.Add("platform", "0")
	v.Add("userid", M.Uid)
	v.Add("vip", "1")
	v.Add("sessionid", sid)
	v.Add("gcid", this.GCid)
	v.Add("cid", this.Cid)
	v.Add("filesize", this.YsFileSize)
	v.Add("cache", strconv.FormatInt(currentTimestamp(), 10))
	v.Add("from", "lxweb")
	v.Add("jsonp", "XL_CLOUD_FX_INSTANCEqueryBack")
	uri := REQGETMETHODVOD_URL + v.Encode()
	r, err := get(uri)
	if err != nil {
		return
	}
	exp := regexp.MustCompile(`XL_CLOUD_FX_INSTANCEqueryBack\((.*)\)`)
	var res struct {
		Resp vod_resp `json:"resp"`
	}
	s := exp.FindSubmatch(r)
	if s == nil {
		err = invalidResponseErr
		return
	}
	json.Unmarshal(s[1], &res)
	fmt.Printf("%+v\n", res.Resp)
	if res.Resp.Status == 0 { // TODO: also check `TransWait`
		for i, _ := range res.Resp.VodList {
			if res.Resp.VodList[i].Spec == 225536 {
				lurl = res.Resp.VodList[i].URL
			} else if res.Resp.VodList[i].Spec == 282880 {
				hurl = res.Resp.VodList[i].URL
			}
		}
	} else {
		err = errors.New(res.Resp.ErrMsg)
	}
	return
}

func (this _bt_list) String() string {
	r := fmt.Sprintf("%s %s %s/%d\n", this.Id, this.InfoId, this.BtNum, this.BtPerNum)
	for i, _ := range this.Record {
		r += fmt.Sprintf("#%d %s %s %s\n", this.Record[i].Id, this.Record[i].FileName, this.Record[i].SizeReadable, this.Record[i].Status)
	}
	return r
}

func (this bt_list) String() string {
	r := fmt.Sprintf("%s %s %s\n", this.Id, this.InfoId, this.BtNum)
	for i, _ := range this.Record {
		r += fmt.Sprintf("#%d %s %s %s\n", this.Record[i].Id, this.Record[i].FileName, this.Record[i].SizeReadable, this.Record[i].Status)
	}
	return r
}

func (this Task) Verify(path string) bool {
	if this.IsBt() {
		fmt.Println("Verifying [BT]", path)
		if b, err := GetTorrentByHash(this.Cid); err != nil {
			fmt.Println(err)
			return false
		} else {
			if m, err := taipei.DecodeMetaInfo(b); err != nil {
				fmt.Println(err)
				return false
			} else {
				taipei.Iconv(m)
				taipei.SetEcho(true)
				g, err := taipei.VerifyContent(m, path)
				taipei.SetEcho(false)
				if err != nil {
					fmt.Println(err)
				}
				return g
			}
		}
	} else if strings.HasPrefix(this.URL, "ed2k://") {
		fmt.Println("Verifying [ED2K]", path)
		h, err := getEd2kHash(path)
		if err != nil {
			fmt.Println(err)
			return false
		}
		if !strings.EqualFold(h, getEd2kHashFromURL(this.URL)) {
			return false
		}
	}
	return true
}

func (this VodLXTask) String() string {
	name, _ := url.QueryUnescape(this.Name)
	return fmt.Sprintf("%s %s [%d] %dMB %d%% %dDays", this.Id, name, this.Status, this.Size/1024/1204, this.Progress/100, this.LeftTime/3600/24)
}

func (this VodHistTask) String() string {
	name, _ := url.QueryUnescape(this.Name)
	return fmt.Sprintf("%s %dMB %d", name, this.Size/1024/1204, this.Duration)
}
