package initialize

import (
	"fmt"
	"github.com/Dream-ming/myMusic/internal/music"
)

func initoss() {
	// 正确初始化方式
	ossClient, err := music.NewOSSClient()
	if err != nil {
		panic(fmt.Errorf("OSS 客户端初始化失败: %v", err))
	}
	
	// 需要传入实际文件路径
	localPath := "./test/111.mp3"
	if err := ossClient.UploadMusicFile(localPath); err != nil {
		panic(fmt.Errorf("文件上传失败: %v", err))
	}
	
	fmt.Println("文件上传成功！")
}