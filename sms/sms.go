package sms

import (
	"net/http"

	"github.com/denverdino/aliyungo/common"
	"fmt"
)

type SendArgs struct {
	PhoneNumbers  string
	SignName      string
	TemplateCode  string
	TemplateParam string
	OutId         string
}

type SendResponse struct {
	common.ErrorResponse
	BizId string
}

//please set the signature and template in the console of Aliyun before you call this API
func (this *Client) SendSms(args *SendArgs) (*SendResponse, error) {
	res := &SendResponse{}
	err := this.InvokeByAnyMethod(http.MethodPost, SendSMS, "", args, res)
	if err == nil && res.Code != "OK" {
		err = fmt.Errorf("aliyun send SMS error: RequestId: %s Code: %s Message: %s", res.RequestId, res.Code, res.Message)
	}
	return res, err
}

type QueryArgs struct {
	PhoneNumber string
	BizId       string
	SendDate    string
	PageSize    int
	CurrentPage int
}

type QueryResponse struct {
	common.Response
	TotalCount int
	TotalPage  int
	SmsSendDetailDTOs struct {
		SendStatus   int
		ErrCode      string
		TemplateCode string
		Content      string
		SendDate     string
		ReceiveDate  string
		OutId        string
	}
}

func (this *Client) QuerySendDetails(args *QueryArgs) (*QueryResponse, error) {
	res := &QueryResponse{}
	err := this.InvokeByAnyMethod(http.MethodGet, QuerySendDetails, "", args, res)
	return res, err
}
