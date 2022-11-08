package sms

import (
	"encoding/json"
	"errors"
	"forum/pkg/config"
	"forum/pkg/logger"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	aliyunclient "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

// Aliyun 实现 sms.Driver interface
type Aliyun struct{}

// Send 实现 sms.Driver interface 的 Send 方法
func (s *Aliyun) Send(phone []string, message Message) bool {

	// 获取配置信息
	aliyunConfigEnv := config.Get[map[string]string]("sms.aliyun")
	logger.DebugJSON("短信-阿里云", "配置信息", aliyunConfigEnv)
	AccessKeyId := aliyunConfigEnv["access_key_id"]
	AccessKeySecret := aliyunConfigEnv["access_key_secret"]
	aliyunConfig := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: &AccessKeyId,
		// 您的 AccessKey Secret
		AccessKeySecret: &AccessKeySecret,
	}
	// 访问的域名
	aliyunConfig.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	client := &aliyunclient.Client{}
	// 实例化一个客户端，从 &dysmsapi.Client 类生成对象 client 。
	client, err := aliyunclient.NewClient(aliyunConfig)
	if err != nil {
		logger.ErrorString("短信-阿里云", "实例化链接失败", err.Error())
		return false
	}

	var stringPhone string

	if !(len(phone) > 0) {
		logger.ErrorString("短信-阿里云", "未检测到手机", "手机参数为空")
		return false
	}

	for _, v := range phone {
		if len(stringPhone) == 0 {
			stringPhone = v
			continue
		}
		stringPhone = stringPhone + "," + v
	}

	if err != nil {
		logger.ErrorString("短信-阿里云", "切片转字符串失败", err.Error())
		return false
	}

	SignName := aliyunConfigEnv["sign_name"]
	byteArray, err := json.Marshal(message.Data)
	if err != nil {
		logger.ErrorString("短信-阿里云", "Map转化为byte数组失败", err.Error())
		return false
	}

	code := string(byteArray)
	sendSmsRequest := &aliyunclient.SendSmsRequest{
		PhoneNumbers:  tea.String(stringPhone),
		SignName:      tea.String(SignName),
		TemplateCode:  &message.Template,
		TemplateParam: tea.String(code),
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				logger.ErrorString("短信-阿里云", "发送意外终止", err.Error())
				_e = r
			}
		}()
		// 创建对应 API 的 Request 。 方法的命名规则为 API 方法名再加上 Request 。例如：
		result, err := client.SendSmsWithOptions(sendSmsRequest, &util.RuntimeOptions{})
		if *result.Body.Code != "OK" {
			err = errors.New(*result.Body.Message)
			return err
		}
		if err != nil {
			logger.ErrorString("短信-阿里云", "发送短信失败", err.Error())
			return err
		}

		return nil
	}()

	if tryErr != nil {
		var smsError = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			smsError = _t
		} else {
			smsError.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		result, _err := util.AssertAsString(smsError.Message)
		logger.DebugJSON("短信-阿里云", "发送短信失败原因", result)
		if _err != nil {
			logger.DebugJSON("短信-阿里云", "获取失败原因失败", _err.Error())
		}
		return false
	}
	return true
}
