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
	baseUrlLocal   = "http://localhost:52001/api"
)

const (
	noticeList              = "/xxx/notice/list"
	verifyCodeSend          = "/common/verify_code/send"
	registerUser            = "/common/register"
	loginUserWithVerifyCode = "/common/login/verify_code"
	loginUserWithPwd        = "/common/login/pwd"
	userPwdReset            = "/user/password/reset"
)

const (
	apiV1 = "/v1"
	apiV2 = "/v2"
)

var apiVersion = apiV1
var qToken = token
var baseUrl = baseUrlLocal + apiVersion

func TestGateway(t *testing.T) {
	t.Run("发送验证码", TestVerifyCodeSend)
	t.Run("注册用户", TestRegisterUser)
	t.Run("登录用户-验证码", TestLoginUserWithVerifyCode)
	t.Run("登录用户-密码", TestLoginUserWithPwd)
	t.Run("重置密码", TestLoginUserPwdReset)
}

const (
	SuccessBusinessCode = 200
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

func TestVerifyCodeSend(t *testing.T) {
	r := baseUrl + verifyCodeSend
	t.Logf("request url: %s", r)
	data := url.Values{}
	data.Set("country_code", "86")
	data.Set("phone", "18319430520")
	data.Set("business_type", "2")
	data.Set("receive_email", "565608463@qq.com")
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

func TestRegisterUser(t *testing.T) {
	r := baseUrl + registerUser
	t.Logf("request url: %s", r)
	data := url.Values{}
	data.Set("user_name", "杨子腾")
	data.Set("password", "07030501310")
	data.Set("sex", "1")
	data.Set("country_code", "86")
	data.Set("phone", "15501707783")
	data.Set("email", "1225807604@qq.com")
	data.Set("verify_code", "044834")
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

func TestLoginUserWithVerifyCode(t *testing.T) {
	r := baseUrl + loginUserWithVerifyCode
	t.Logf("request url: %s", r)
	data := url.Values{}
	data.Set("country_code", "86")
	data.Set("phone", "18319430520")
	data.Set("verify_code", "804814")
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

func TestLoginUserWithPwd(t *testing.T) {
	r := baseUrl + loginUserWithPwd
	t.Logf("request url: %s", r)
	data := url.Values{}
	data.Set("country_code", "86")
	data.Set("phone", "15501707783")
	data.Set("password", "07030501310")
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

func TestLoginUserPwdReset(t *testing.T) {
	r := baseUrl + userPwdReset
	t.Logf("request url: %s", r)
	data := url.Values{}
	data.Set("verify_code", "791945")
	data.Set("password", "12345678")
	t.Logf("req data: %v", data)
	req, err := http.NewRequest("PUT", r, strings.NewReader(data.Encode()))
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
