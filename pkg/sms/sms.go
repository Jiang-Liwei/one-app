package sms

import (
	"forum/pkg/config"
	"sync"
)

// Message 短信结构体
type Message struct {
	Template string
	Data     map[string]string

	ConTent string
}

// SMS 短信操作类
type SMS struct {
	Driver Driver
}

// 单例模式
var once sync.Once

// internalSMS 内部使用的 SMS 对象
var internalSMS *SMS

// NewSMS 单例模式获取
func NewSMS() *SMS {
	once.Do(func() {
		switch config.Get[string]("sms.platform") {
		case "aliyun":
			internalSMS = &SMS{
				Driver: &Aliyun{},
			}
		default:
			internalSMS = &SMS{
				Driver: &Tencent{},
			}
		}

	})

	return internalSMS
}

// Send 发送短信
func (sms *SMS) Send(phone []string, message Message) bool {
	return sms.Driver.Send(phone, message)
}
