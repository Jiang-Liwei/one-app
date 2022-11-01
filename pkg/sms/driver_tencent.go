package sms

import (
	"encoding/json"
	"fmt"
	"forum/pkg/config"
	"forum/pkg/logger"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111" // 引入sms
)

// Tencent 实现 sms.Driver interface
type Tencent struct{}

// Send 实现 sms.Driver interface 的 Send 方法
func (s Tencent) Send(phone []string, message Message) bool {

	if !(len(message.Data) > 0) {
		logger.ErrorString("短信-腾讯云", "未检测到短信参数", "短信参数为空")
		return false
	}

	var templateParams []string
	for _, templateParam := range message.Data {
		templateParams = append(templateParams, templateParam)
	}
	tencentConfig := config.Get[map[string]string]("sms.tencent")
	logger.DebugJSON("短信-腾讯云", "配置信息", tencentConfig)

	credential := common.NewCredential(
		tencentConfig["secret_id"],
		tencentConfig["secret_key"],
	)
	cpf := profile.NewClientProfile()

	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, err := sms.NewClient(credential, "ap-nanjing", cpf)
	if err != nil {
		logger.ErrorString("短信-腾讯云", "实例化链接失败", err.Error())
		return false
	}

	/* 实例化一个请求对象 */
	request := sms.NewSendSmsRequest()

	request.SmsSdkAppId = common.StringPtr(tencentConfig["app_id"])
	request.SignName = common.StringPtr(tencentConfig["sign_name"])
	println(message.Template)
	request.TemplateId = common.StringPtr(message.Template)

	request.TemplateParamSet = common.StringPtrs(templateParams)

	/* 下发手机号码，采用 E.164 标准，+[国家或地区码][手机号]
	 * 示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号，最多不要超过200个手机号*/
	request.PhoneNumberSet = common.StringPtrs(phone)

	/* 用户的 session 内容（无需要可忽略）: 可以携带用户侧 ID 等上下文信息，server 会原样返回 */
	request.SessionContext = common.StringPtr("")
	/* 国际/港澳台短信 SenderId（无需要可忽略）: 国内短信填空，默认未开通，如需开通请联系 [腾讯云短信小助手] */
	request.SenderId = common.StringPtr("")

	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.SendSms(request)
	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		logger.DebugJSON("短信-腾讯云", "SDK错误", err)
		return false
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		logger.DebugJSON("短信-腾讯云", "非SDK错误", err)
		return false
	}
	b, _ := json.Marshal(response.Response)
	logger.DebugJSON("短信-腾讯云", "发送结果", string(b))
	// 打印返回的json字符串
	fmt.Printf("%s", b)
	return true
}
