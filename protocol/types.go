package protocol

type login_resp struct {
	Result int `json:"result"`
	Data   struct {
		UserId    string `json:"userid"`
		UserName  string `json:"usrname"`
		UserNewno string `json:"usernewno"`
		UserType  string `json:"usrtype"`
		Nickname  string `json:"nickname"`
		UserNick  string `json:"usernick"`
		VipState  string `json:"vipstate"`
	} `json:"data"`
}

type _task_resp struct {
	Rtcode   int       `json:"rtcode"`
	Info     _info     `json:"info"`
	UserInfo _userinfo `json:"userinfo"`
	// GlobalNew interface{} `json:"global_new"`
	// Time      interface{} `json:"time"` // v1: float64, v-current: string
}

type _info struct {
	Tasks    []Task `json:"tasks"`
	User     _user  `json:"user"`
	ShowArc  int    `json:"show_arc"`
	TotalNum string `json:"total_num"`
}

type Task struct {
	Id   string `json:"id"`
	Flag string `json:"flag"`
	// Database            string `json:"database"`
	// ClassValue          string `json:"class_value"`
	GlobalId       string  `json:"global_id"`
	ResType        string  `json:"restype"`
	FileSize       string  `json:"filesize"`
	FileType       string  `json:"filetype"`
	Cid            string  `json:"cid"`
	GCid           string  `json:"gcid"`
	TaskName       string  `json:"taskname"`
	DownloadStatus string  `json:"download_status"`
	Speed          string  `json:"speed"`
	Progress       float32 `json:"progress"`
	// UsedTime            string `json:"used_time"`
	LeftLiveTime string `json:"left_live_time"`
	LixianURL    string `json:"lixian_url"`
	URL          string `json:"url"`
	ReferURL     string `json:"refer_url"`
	Cookie       string `json:"cookie"`
	// Vod                 string `json:"vod"`
	Status string `json:"status"`
	// Message             string `json:"message"`
	// DtCommitted         string `json:"dt_committed"`
	// DtDeleted           string `json:"dt_deleted"`
	// ListSum             string `json:"list_sum"`
	// FinishSum           string `json:"finish_sum"`
	// FlagKilledInASecond string `json:"flag_killed_in_a_second"`
	// ResCount            string `json:"res_count"`
	// UsingResCount       string `json:"using_res_count"`
	// VerifyFlag          string `json:"verify_flag"`
	// VerifyTime          string `json:"verify_time"`
	// ProgressText 			 string `json:"progress_text"`
	// ProgressImg         string `json:"progress_img"`
	// ProgressClass       string `json:"progress_class"`
	LeftTime string `json:"left_time"`
	UserId   int    `json:"userid"`
	// OpenFormat string `json:"openformat"`
	// TaskNameShow        string `json:"taskname_show"`
	// Ext                 string `json:"ext"`
	// ExtShow             string `json:"ext_show"`
	TaskType byte `json:"tasktype"`
	// FormatImg           string `json:"format_img"`
	// ResCountDegree      byte   `json:"res_count_degree"`
	YsFileSize string `json:"ysfilesize"`
	// BtMovie             byte   `json:"bt_movie"`
	UserType string `json:"user_type"`
}

type _user struct {
	ExpireDate          string `json:"expire_date"`
	MaxTaskNum          string `json:"max_task_num"`
	MaxStore            string `json:"max_store"`
	VipStore            string `json:"vip_store"`
	BuyStore            string `json:"buy_store"`
	XzStore             string `json:"xz_store"`
	BuyNumTask          string `json:"buy_num_task"`
	BuyNumConn          string `json:"buy_num_connection"`
	BuyBandwith         string `json:"buy_bandwidth"`
	BuyTaskLiveTime     string `json:"buy_task_live_time"`
	XpExpireDate        string `json:"experience_expire_date"`
	AvailableSpace      string `json:"available_space"`
	TotalNum            string `json:"total_num"`
	HistoryTaskTotalNum string `json:"history_task_total_num"`
	SuspendingNum       string `json:"suspending_num"`
	DownloadingNum      string `json:"downloading_num"`
	WaitingNum          string `json:"waiting_num"`
	CompleteNum         string `json:"complete_num"`
	StorePeriod         string `json:"store_period"`
	Cookie              string `json:"cookie"`
	VipLevel            string `json:"vip_level"`
	UserType            string `json:"user_type"`
	GoldbeanNum         string `json:"goldbean_num"`
	ConvertFlag         string `json:"convert_flag"`
	SilverbeanNum       string `json:"silverbean_num"`
	SpecialNet          string `json:"special_net"`
	TotalFilterNum      string `json:"total_filter_num"`
}

type _userinfo struct {
	AllSpace       string `json:"all_space"`
	AllUsedStore   int64  `json:"all_used_store"`
	AllSpaceFormat string `json:"all_space_format"`
	AllUsedFormat  string `json:"all_used_format"`
	Isp            bool   `json:"isp"`
	Percent        string `json:"percent"`
}

type bt_list struct {
	Id     string       `json:"Tid"`
	InfoId string       `json:"Infoid"`
	BtNum  string       `json:"btnum"`
	Record []_bt_record `json:"Record"`
}

type _bt_list struct {
	Id       string       `json:"Tid"`
	InfoId   string       `json:"Infoid"`
	BtNum    string       `json:"btnum"`
	BtPerNum int          `json:"btpernum"`
	NowPage  int          `json:"now_page"`
	Record   []_bt_record `json:"Record"`
}

type _bt_record struct {
	Id           int    `json:"id"`
	FileName     string `json:"title"`
	Status       string `json:"download_status"`
	Cid          string `json:"cid"`
	SizeReadable string `json:"size"`
	Percent      int    `json:"percent"`
	TaskId       string `json:"taskid"`
	LiveTime     string `json:"livetime"`
	DownURL      string `json:"downurl"`
	FileSize     string `json:"filesize"`
	Verify       string `json:"verify"`
	URL          string `json:"url"`
	DirName      string `json:"dirtitle"`
}

type _task_pre struct {
	Cid        string
	GCid       string
	SizeCost   string
	FileName   string
	Goldbean   string
	Silverbean string
	// IsFull     string
	// Random     []byte
}

type _btup_result struct {
	Ret    int           `json:"ret_value"`
	InfoId string        `json:"infoid"`
	Name   string        `json:"ftitle"`
	Size   int           `json:"btsize"`
	IsFull string        `json:"is_full"`
	List   []_bt_record2 `json:"filelist"`
	Random string        `json:"random"`
}

type _bt_record2 struct {
	Id    string `json:"id"`
	Size  string `json:"subsize"`
	Sizef string `json:"subformatsize"`
	Valid byte   `json:"valid"`
	Index string `json:"findex"`
	Name  string `json:"subtitle"`
	Ext   string `json:"ext"`
}

type _bt_qtask struct {
	InfoId string
	Size   string
	Name   string
	IsFull string
	Files  []string
	Sizesf []string
	Sizes  []string
	Picked []string
	Ext    []string
	Index  []string
	Random string
	Ret    string
}

type _ptask_record struct {
	Id             string  `json:"tid"`
	URL            string  `json:"url"`
	Speed          string  `json:"speed"`
	Progress       float32 `json:"fpercent"` // diff between 'percent' and 'fpercent'?
	LeaveTime      string  `json:"leave_time"`
	Size           string  `json:"fsize"`
	DownloadStatus string  `json:"download_status"`
	LixianURL      string  `json:"lixian_url"`
	LeftLiveTime   string  `json:"left_live_time"`
	TaskType       string  `json:"tasktype"`
	FileSize       string  `json:"filesize"`
}

type _ptask_resp struct {
	List []_ptask_record `json:"Record"`
	Info struct {
		DownNum string `json:"downloading_num"`
		WaitNum string `json:"waiting_num"`
	} `json:"Task"`
}

type VodHistTask struct {
	GCid     string `json:"gcid"`
	UrlHash  string `json:"url_hash"`
	Cid      string `json:"cid"`
	URL      string `json:"url"`
	Name     string `json:"file_name"`
	Type     byte   `json:"tasktype"`
	Src      string `json:"src_url"`
	Size     int64  `json:"file_size"`
	Duration int    `json:"duration"`
	Played   string `json:"playtime"`
	Created  string `json:"createtime"`
}

type hist_resp struct {
	List   []VodHistTask `json:"history_play_list"`
	MaxNum int           `json:"max_num"`
	Uid    string        `json:"userid"`
	Ret    byte          `json:"ret"`
	Start  string        `json:"start_t"`
	End    string        `json:"end_t"`
	Num    int           `json:"record_num"`
	Type   string        `json:"type"`
}

type subbt_resp struct {
	Uid  string `json:"userid"`
	Ret  byte   `json:"ret"`
	List []struct {
		GCid     string `json:"gcid"`
		UrlHash  string `json:"url_hash"`
		Name     string `json:"name"`
		Index    int    `json:"index"`
		Cid      string `json:"cid"`
		Size     int64  `json:"file_size"`
		Duration int64  `json:"duration"`
	} `json:"subfile_list"`
	UrlHash  string `json:"main_task_url_hash"`
	InfoHash string `json:"info_hash"`
	Num      int    `json:"record_num"`
}

type progress_resp struct {
	List []struct {
		Progress string `json:"progress"`
		UrlHash  string `json:"url_hash"`
	} `json:"progress_info_list"`
	Uid string `json:"userid"`
	Ret byte   `json:"ret"`
}

type VodLXTask struct {
	Vod       int    `json:"vod"`
	GCid      string `json:"gcid"`
	Status    int    `json:"download_status"`
	SecDown   int    `json:"sec_down"`
	FinishNum int    `json:"finish_num"`
	TotalNum  int    `json:"total_num"`
	Cid       string `json:"cid"`
	URL       string `json:"url"`
	FileType  int    `json:"filetype"`
	ResType   int    `json:"restype"`
	Flag      int    `json:"flag"`
	Size      int64  `json:"filesize"`
	Id        string `json:"taskid"`
	Progress  int    `json:"progress"`
	Name      string `json:"taskname"`
	LixianURL string `json:"lixian_url"`
	LeftTime  int    `json:"left_live_time"`
}

type lxtask_resp struct {
	List []VodLXTask `json:"tasklist"`
	Ret  string      `json:"ret"`
	User struct {
		From      int    `json:"from"`
		Name      string `json:"name"`
		IP        string `json:"ip"`
		NewNO     string `json:"newno"`
		Uid       string `json:"userid"`
		VIP       string `json:"vip"`
		Type      string `json:"user_type"`
		Sid       string `json:"sessionid"`
		AvalSpace string `json:"available_space"`
		MaxStore  string `json:"max_store"`
	} `json:"user_info"`
	Count    string `json:"task_total_cnt"`
	ErrMsg   string `json:"error_msg"`
	CountAll string `json:"task_total_cnt_all"`
}

type mb_resp struct {
	//url: http://i.vod.xunlei.com/miaobo/
	// infohash/507625D95D7695F9A1F91EC41A633D250419DF26?jsonp=jsonp1383446499423&t=0.7223064277786762
	InfoHash string `json:"infohash"`
	Result   byte   `json:"result"`
	Index    int    `json:"idx"`
}

type VodInfo struct {
	URLs   []string `json:"vod_urls"`
	Spec   int      `json:"spec_id"`
	URL    string   `json:"vod_url"`
	HasSub byte     `json:"has_subtitle"`
}

type vod_resp struct {
	Status    byte   `json:"status"`
	UrlHash   string `json:"url_hash"`
	TransWait int    `json:"tans_wait"`
	Uid       string `json:"userid"`
	Ret       byte   `json:"ret"`
	SrcInfo   struct {
		Name string `json:"file_name"`
		Cid  string `json:"cid"`
		Size string `json:"file_size"`
		GCid string `json:"gcid"`
	} `json:"src_info"`
	Duration  int64 `json:"duration"`
	VodPermit struct {
		Msg string `json:"msg"`
		Ret byte   `json:"ret"`
	} `json:"vod_permit"`
	ErrMsg  string    `json:"error_msg"`
	VodList []VodInfo `json:"vodinfo_list"`
}
