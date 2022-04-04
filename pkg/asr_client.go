package pkg

import (
	"github.com/spf13/viper"
	asr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/asr/v20190614"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func GetClient() *asr.Client {
	//credential := common.NewCredential(
	//	"AKIDxjFVUChvtaeS0GFxOZT6RVIo31i8ICpB",
	//	"5uTlmVxCArHQQAftqVu2tSklUw83h5KI",
	//)
	credential := common.NewCredential(
		viper.GetString("tx.secretId"),
		viper.GetString("tx.secretKey"),
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = viper.GetString("tx.endpoint")
	client, _ := asr.NewClient(credential, "", cpf)
	return client
}
