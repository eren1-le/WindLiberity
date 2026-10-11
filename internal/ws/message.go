/*
 * @Author: eren dengdengd1222@mail.com
 * @Date: 2026-10-06 14:59:38
 * @LastEditors: eren dengdengd1222@mail.com
 * @LastEditTime: 2026-10-08 10:44:19
 * @FilePath: /WindLiberity/internal/ws/message.go
 * @Description:
 *
 */
package ws

import "encoding/json"

const (
	EventChat   = "chat"   // chat message
	EventRecall = "recall" //withdraw message
	EventNotify = "notify" // notify
	EventClose  = "Close"  // shut connection of client
)

type Chat struct {
	ID       string          `json:"id"`                           // 消息id
	From     *Sender         `json:"from"`                         // 发送人
	To       *Sender         `json:"to"`                           // 接受人/群id
	ChatType int             `json:"chat_type"`                    // 聊天类型
	Type     int             `json:"type"`                         // 消息类型
	Options  json.RawMessage `json:"options" swaggertype:"string"` // 扩展信息
	Content  string          `json:"content"`                      // 消息内容
	T        int64           `json:"t"`                            // 发送时间
}

// Recall
type Recall struct {
	ID       string `json:"id"`        // message id
	FromID   int    `json:"from_id"`   //sender id
	ToID     int    `json:"to_id"`     //receiver id
	ChatType int    `json:"chat_type"` // chat type
}
type Norify struct {
	Type string `json:"type"` // nofity type
}

// Sender
type Sender struct {
	ID     int    `json:"id"`     //user ID
	Name   string `json:"name"`   //user nick
	Avatar string `json:"avatar"` //user avatar
}
