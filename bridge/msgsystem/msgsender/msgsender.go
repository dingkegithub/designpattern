package msgsender

import "container/list"

//
// 消息发送实现，独立扩展
// 可以扩展各种发送渠道
// 比如在添加微信，钉钉
// 等通知
//
type MsgSender interface {
	Send(message string)
}

//
// telephone sender
//
type TelephoneMsgSender struct {
	telephone *list.List
}

func NewTelephoneMsgSender(tp *list.List) MsgSender {
	return &TelephoneMsgSender{
		telephone: tp,
	}
}

func (t *TelephoneMsgSender) Send(message string) {

}

//
// email sender
//
type EmailMsgSender struct {
	emails *list.List
}

func NewEmailMsgSender(emails *list.List) MsgSender {
	return &EmailMsgSender{
		emails: emails,
	}
}

func (t *EmailMsgSender) Send(message string) {

}
