package sms

import "github.com/denverdino/aliyungo/common"

const (
	SmsEndPoint      = "https://dysmsapi.aliyuncs.com/"
	SmsAPIVersion    = "2017-05-25"
	SendSMS          = "SendSms"
	QuerySendDetails = "QuerySendDetails"
)

type Client struct {
	common.Client
}

func NewClient(accessKeyId, accessKeySecret string) *Client {
	client := new(Client)
	client.Init(SmsEndPoint, SmsAPIVersion, accessKeyId, accessKeySecret)
	return client
}
