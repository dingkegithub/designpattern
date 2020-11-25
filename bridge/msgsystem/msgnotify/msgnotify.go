package msgnotify

import "github.com/dingkegithub/designpattern/bridge/msgsystem/msgsender"

//
// 消息发送抽象
//
type Notification interface {
	Notify(message string)
}

type SevereNotification struct {
	// 组合实现
	msgSender msgsender.MsgSender
}

func (s *SevereNotification) Notify(message string) {
	s.msgSender.Send(message)
}
