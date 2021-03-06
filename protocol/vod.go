// http://vod.xunlei.com/js/list.js

package protocol

import (
	"encoding/json"
	"fmt"
)

func GetHistoryPlayList() ([]VodHistTask, error) {
	uri := fmt.Sprintf(HISTORY_PLAY_URL, 30, 0, "all", "create", currentTimestamp()) //TODO: eliminate hard-code
	b, err := get(uri)
	if err != nil {
		return nil, err
	}
	var resp struct {
		Data hist_resp `json:"resp"`
	}
	err = json.Unmarshal(b, &resp)
	return resp.Data.List, nil
}

func SubmitBt(infohash string, num int) *subbt_resp {
	return nil
}

func QueryProgress() *progress_resp {
	return nil
}

func GetLxtaskList() ([]VodLXTask, error) {
	var payload struct {
		UserInfo struct {
			Uid   string `json:"userid"`
			Name  string `json:"name"`
			NewNo string `json:"newno"`
			VIP   string `json:"vip"`
			IP    string `json:"ip"`
			Sid   string `json:"sessionid"`
			From  int    `json:"from"`
		} `json:"user_info"`
		Offset int   `json:"offset"`
		Num    int   `json:"req_num"`
		Type   int   `json:"req_type"`
		Attr   int   `json:"fileattribute"`
		Time   int64 `json:"t"`
	}
	domain := "http://xunlei.com"
	payload.UserInfo.Uid = getCookie(domain, "userid")
	payload.UserInfo.Name = getCookie(domain, "usrname")
	payload.UserInfo.NewNo = getCookie(domain, "usernewno")
	payload.UserInfo.VIP = getCookie(domain, "isvip")
	payload.UserInfo.Sid = getCookie(domain, "sessionid")
	payload.Offset = 0
	payload.Num = 30
	payload.Type = 2
	payload.Attr = 1
	payload.Time = currentTimestamp()
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	v, err := post(LXTASK_LIST_URL, string(b))
	if err != nil {
		return nil, err
	}
	var resp lxtask_resp
	err = json.Unmarshal(v, &resp)
	return resp.List, err
}
