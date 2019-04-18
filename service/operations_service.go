package service

import (
	"ConfCenter_web_Admin/initialization"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"gopkg.in/square/go-jose.v1/json"
	"io/ioutil"
	"net/http"
)

var (
	scheme = "http://"
	client = &http.Client{}
)

type OperationsResult struct {
	Result []*Operations `json:"result"`
}

type Operations struct {
	Id           uint64   `json:"id"`
	Route        string   `json:"route"`
	Service      string   `json:"service"`
	ServiceName  string   `json:"servicename"`
	ServiceAddr  []string `json:"serviceaddr"`
	RegisterTime string   `json:"registertime"`
	Balance      string   `json:"Balance"`
}

type Service struct {
	ServiceAddr []string `json:"serviceaddr"` //服务地址  [ip:port]
	//RegisterName string   `json:"registername"` //谁注册的服务  这里的名字是登录的名字，不能让人填写，这里先空着，等登录注册完成之后再说补充
	RegisterTime string `json:"registertime"` //注册时间
	//AltTime      string   `json:"alttime"`      //修改时间  这里的名字是登录的名字，不能让人填写，这里先空着，等登录注册完成之后再说补充
	AltReason   string `json:"altreason"` //修改原因
	ServiceName string `json:"servicename"`
	Balance     string `json:"Balance"`
}

func NewOperationsResult() *OperationsResult {
	return &OperationsResult{
		Result: make([]*Operations, 0, 60),
	}
}

func (o *OperationsResult) Get() error {
	var service = &Service{}
	req, err := http.NewRequest("GET", scheme+initialization.ConfCenterService.Addr+"/quick_operation", nil)
	if err != nil {
		logs.Error("get confCenter service err %v", err)
		return errors.New(fmt.Sprintf("get ConfCenter err %v", err))
	}
	resp, err := client.Do(req)
	if err != nil {
		logs.Error("get ConfCenter response err %v", err)
		return errors.New(fmt.Sprintf("get ConfCenter response err %v", err))
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		logs.Error("readAll Operations service config body err %v", err)
		return errors.New(fmt.Sprintf("readAll Operations service config body err %v", err))
	}
	if o == nil {
		o = NewOperationsResult()
	}
	err = json.Unmarshal(body, o)
	if err != nil {
		logs.Error("unmarshal configCenter service body err %v", err)
		return errors.New(fmt.Sprintf("marshal configCenter service body err %v", err))
	}
	for _, v := range o.Result {
		err = json.Unmarshal([]byte(v.Service), service)
		if err != nil {
			logs.Error("unmarshal o.result.service err %v", err)
			return errors.New(fmt.Sprintf("unmarshal o.result.service err %v", err))
		}
		v.ServiceAddr = service.ServiceAddr
		v.RegisterTime = service.RegisterTime
		v.Balance = service.Balance
	}
	for _, v := range o.Result {
		logs.Debug("operations-----", v.RegisterTime, v.ServiceAddr, v.Balance)
	}
	for _, v := range o.Result {
		switch {
		case v.Balance == "random":
			v.Balance = "轮询"
		case v.Balance == "poling":
			v.Balance = "随机"
		}
	}
	return nil
}
