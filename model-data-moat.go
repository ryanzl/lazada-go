package lazada

type DataMoatLoginReq struct {
	Time          	string `json:"time,omitempty"`
	AppName       	string `json:"appName,omitempty"`
	UserId      	string `json:"userId,omitempty"`
	Tid	     		string `json:"userId,omitempty"`
	UserIp	      	string `json:"userIp,omitempty"`
	Ati	      		string `json:"ati,omitempty"`
	LoginResult	    string `json:"loginResult,omitempty"`
	LoginMessage    string `json:"loginMessage,omitempty"`
}

type DataMoatLoginResp struct {
	Code                 string   `json:"code,omitempty"`
	RequestId            string   `json:"request_id,omitempty"`
	Result   DataMoatLoginResult `json:"result,omitempty"`
}
type DataMoatLoginResult struct {
	Msg          	string `json:"msg,omitempty"`
	Success       	bool `json:"success,omitempty"`
}
