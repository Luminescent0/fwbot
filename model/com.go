/**
 * @Author: Hardews
 * @Date: 2022/10/30 0:52
**/

package model

// 事件相关结构体，详情见https://docs.go-cqhttp.org/event

type Com struct {
	Time     int64  `json:"time,omitempty"`      //事件发生的时间戳
	SelfId   int64  `json:"self_id,omitempty"`   //收到事件的机器人的qq
	PostType string `json:"post_type,omitempty"` //上报类型
}

// Message post_type为message的上报
type Message struct {
	Com
	MessageType string `json:"message_type"`
	SubType     string `json:"sub_type"` //消息的子类型 可能为group,public
	MessageId   int32  `json:"message_id"`
	UserId      int64  `json:"user_id"` //发送者的qq
	Messages    string `json:"message"`
	RawMessage  string `json:"raw_message"` //CQ code格式的消息
	NoticeType  string `json:"notice_type"`
	Font        string `json:"font"`
	GroupId     int64  `json:"group_id"`
	Sender
}

type Sender struct {
	UserId   int64  `json:"user_id"`
	NickName string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
}
