package pkg

import (
	"fmt"
	"log"
	"log/slog"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20180501/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	credential "github.com/aliyun/credentials-go/credentials"
)

// 发送短信【阿里】
func NewAliSmsClient(secretId, secretKey, phone, phoneCode string) error {
	credential, err := credential.NewCredential(nil)
	if err != nil {
		log.Fatalf("创建凭据失败,%s", err.Error())
	}
	config := &openapi.Config{
		Credential:      credential,
		AccessKeyId:     tea.String(secretId),  // 必填
		AccessKeySecret: tea.String(secretKey), // 必填
		Endpoint:        tea.String("dysmsapi.aliyuncs.com"),
	}

	client := &dysmsapi.Client{}
	client, err = dysmsapi.NewClient(config)
	if err != nil {
		log.Fatalf("创建阿里云短信客户端失败,%s", err.Error())
	}

	sendSmsRequest := &dysmsapi.BatchSendMessageToGlobeRequest{
		To:      tea.String(phone),
		Message: tea.String(phoneCode),
	}
	resp, err := client.BatchSendMessageToGlobe(sendSmsRequest)
	if err != nil {
		slog.Error("短信发送失败", "err", err.Error())
		return err
	}
	fmt.Println("resp", resp)
	return nil
}
