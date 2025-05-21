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
    // 从 Viper 获取配置
    endpoint := viper.GetString("oss.endpoint")
    accessKey := viper.GetString("oss.access_key_id")
    secretKey := viper.GetString("oss.access_key_secret")
    bucketName := viper.GetString("oss.bucket")

    // 创建客户端
    client, err := oss.New(endpoint, accessKey, secretKey)
    if err != nil {
        return nil, err
    }

    // 获取存储桶
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
	// 正确初始化方式
	// ossClient, err := music.NewOSSClient()
	_, err := NewOSSClient()
	if err != nil {
		panic(fmt.Errorf("OSS 客户端初始化失败: %v", err))
	}

	// // 需要传入实际文件路径
	// localPath := "./test/111.mp3"
	// if err := ossClient.UploadMusicFile(localPath); err != nil {
	// 	panic(fmt.Errorf("文件上传失败: %v", err))
	// }

	// fmt.Println("文件上传成功！")
}