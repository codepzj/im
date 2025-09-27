package pkg

import (
	"encoding/json"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	profile "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

// 发送短信【腾讯】
func SendTecentSmsCode(secretId, secretKey, smsSdkAppId, signName, templateId, phoneCode, phone string) error {
	credential := common.NewCredential(secretId, secretKey)
	// 腾讯云文档：https://cloud.tencent.com/document/product/382/43199
	/* 实例化一个客户端配置对象 */
	cpf := profile.NewClientProfile()

	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.ReqTimeout = 10 // 请求超时时间
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"

	/* SDK默认用TC3-HMAC-SHA256进行签名，非必要请不要修改这个字段 */
	cpf.SignMethod = "HmacSHA1"

	/* 实例化client对象 */
	client, _ := sms.NewClient(credential, "ap-guangzhou", cpf)

	/* 实例化一个请求对象 */
	request := sms.NewSendSmsRequest()

	/* 短信应用ID */
	request.SmsSdkAppId = common.StringPtr(smsSdkAppId)

	/* 短信签名内容 */
	request.SignName = common.StringPtr(signName)

	/* 模板 ID */
	request.TemplateId = common.StringPtr(templateId)

	/* 模板参数 */
	request.TemplateParamSet = common.StringPtrs([]string{phoneCode})

	/* 下发手机号码，采用 E.164 标准，+[国家或地区码][手机号] */
	request.PhoneNumberSet = common.StringPtrs([]string{"+86" + phone})

	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.SendSms(request)

	if err != nil {
		return err
	}
	b, _ := json.Marshal(response.Response)
	// 打印返回的json字符串
	fmt.Printf("%s", b)

	return nil
}
