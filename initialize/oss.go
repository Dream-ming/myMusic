package initialize

import (
	"fmt"
    "github.com/aliyun/aliyun-oss-go-sdk/oss"
    "github.com/spf13/viper"
)

type OSSClient struct {
    client *oss.Client
    bucket *oss.Bucket
}

func NewOSSClient() (*OSSClient, error) {
    endpoint := viper.GetString("oss.endpoint")
    accessKey := viper.GetString("oss.access_key_id")
    secretKey := viper.GetString("oss.access_key_secret")
    bucketName := viper.GetString("oss.bucket")

    client, err := oss.New(endpoint, accessKey, secretKey)
    if err != nil {
        return nil, err
    }

    bucket, err := client.Bucket(bucketName)
    if err != nil {
        return nil, err
    }

    return &OSSClient{
        client: client,
        bucket: bucket,
    }, nil
}

func oss_init() {
	_, err := NewOSSClient()
	if err != nil {
		panic(fmt.Errorf("OSS 客户端初始化失败: %v", err))
	}
}