package client

import (
	"gitee.com/cristiane/go-common/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

const (
	baseUrlProd    = "https://xxx.xxx.xx/api"
	baseUrlTestAli = "http://xx.xx.xx.xx:xxx/api"
	baseUrlDev     = "http://xx.xx.xx.56:xx/api"
	baseUrlLocal   = "http://localhost:8080/api"
)

const (
	noticeList          = "/xxx/notice/list"
	voteUserExtractList = "aaa"
)

const (
	apiV1 = "/v1"
	apiV2 = "/v2"
)

var apiVersion = apiV2
var qToken = token
var baseUrl = baseUrlTestAli + apiVersion

func TestCDNPoolGateway(t *testing.T) {
	t.Run("公告列表", TestNoticeList)
}

const (
	SuccessBusinessCode = 0
)

type HttpCommonRsp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func TestNoticeList(t *testing.T) {
	r := baseUrl + noticeList + "?xx=xx&xx=3"
	t.Logf("request url: %s", r)
	req, err := http.NewRequest("GET", r, nil)
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("token", qToken)
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("req url: %v status : %v", r, rsp.Status)
	if rsp.StatusCode != http.StatusOK {
		t.Error("StatusCode != 200")
		return
	}
	body, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("req url: %v body : \n%s", r, body)
	var obj HttpCommonRsp
	err = json.Unmarshal(string(body), &obj)
	if err != nil {
		t.Error(err)
		return
	}
	if obj.Code != SuccessBusinessCode {
		t.Errorf("business code != %v", SuccessBusinessCode)
		return
	}
}

func TestVoteUserExtractListV2(t *testing.T) {
	r := baseUrl + voteUserExtractList
	t.Logf("request url: %s", r)
	data := url.Values{}
	data.Set("page_size", "20")
	data.Set("page_num", "1")
	t.Logf("req data: %v", data)
	req, err := http.NewRequest("POST", r, strings.NewReader(data.Encode()))
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("token", qToken)
	t.Logf("request token=%v", qToken)
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("req url: %v status : %v", r, rsp.Status)
	if rsp.StatusCode != http.StatusOK {
		t.Error("StatusCode != 200")
		return
	}
	body, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("req url: %v body : \n%s", r, body)
	var obj HttpCommonRsp
	err = json.Unmarshal(string(body), &obj)
	if err != nil {
		t.Error(err)
		return
	}
	if obj.Code != SuccessBusinessCode {
		t.Errorf("business code != %v", SuccessBusinessCode)
		return
	}
}
