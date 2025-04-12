// internal/music/storage.go
package music

import (
	"fmt"
    "github.com/aliyun/aliyun-oss-go-sdk/oss"
    "github.com/spf13/viper"
	"path/filepath"
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

func (o *OSSClient) UploadMusicFile(localPath string) error {
    // 生成OSS路径：music/周杰伦_七里香.mp3
    objectKey := viper.GetString("oss.music_prefix") + filepath.Base(localPath)
    
    err := o.bucket.PutObjectFromFile(objectKey, localPath)
    if err != nil {
        return fmt.Errorf("上传失败: %v", err)
    }
    return nil
}