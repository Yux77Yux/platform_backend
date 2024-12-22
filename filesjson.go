package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

const (
	Endpoint     = "https://oss-cn-guangzhou.aliyuncs.com"
	EndpointName = "oss-cn-guangzhou.aliyuncs.com"
	Region       = "cn-guangzhou"
	BucketName   = "platform-user"
	PrefixName   = "Media/"
)

func main() {
	Endpoint := "https://oss-cn-guangzhou.aliyuncs.com"
	Region := "cn-guangzhou"
	// 初始化OSS客户端
	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		log.Fatalf("Error getting credentials: %v", err)
	}
	clientOptions := []oss.ClientOption{oss.SetCredentialsProvider(&provider)}
	clientOptions = append(clientOptions, oss.Region(Region))
	// 设置签名版本
	clientOptions = append(clientOptions, oss.AuthVersion(oss.AuthV4))
	client, err := oss.New(Endpoint, "", "", clientOptions...)
	if err != nil {
		log.Printf("error: oss create client failed : %v", err)
	}

	// 指定目标Bucket和文件夹
	bucket, err := client.Bucket(BucketName)
	if err != nil {
		log.Fatalf("Error getting bucket: %v", err)
	}

	// 遍历文件夹并收集信息
	var fileList []struct {
		FileName   string `json:"fileName"`
		FileType   string `json:"fileType"`
		AccessLink string `json:"accessLink"`
	}
	marker := ""
	for {
		objects, err := bucket.ListObjects(oss.Prefix(PrefixName), oss.Marker(marker))
		if err != nil {
			log.Fatalf("Error listing objects: %v", err)
		}

		for _, object := range objects.Objects {
			fileName := object.Key
			fileType := getFileType(fileName)
			accessLink := fmt.Sprintf("https://%s.%s/%s", BucketName, EndpointName, url.QueryEscape(object.Key))

			fileList = append(fileList, struct {
				FileName   string `json:"fileName"`
				FileType   string `json:"fileType"`
				AccessLink string `json:"accessLink"`
			}{
				FileName:   fileName,
				FileType:   fileType,
				AccessLink: accessLink,
			})
		}

		if !objects.IsTruncated {
			break
		}
		marker = objects.NextMarker
	}

	// 将信息写入JSON文件
	jsonData, err := json.MarshalIndent(fileList, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling to JSON: %v", err)
	}
	err = os.WriteFile("files_info.json", jsonData, 0644)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}
	fmt.Println("Files information exported to files_info.json")
}

// 简单函数根据文件名后缀推断文件类型
func getFileType(fileName string) string {
	extensions := map[string]string{
		// 文本文件
		".txt":  "text/plain",
		".html": "text/html",
		".css":  "text/css",
		".js":   "application/javascript",
		".json": "application/json",
		".xml":  "application/xml",

		// 图片文件
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
		".png":  "image/png",
		".gif":  "image/gif",
		".bmp":  "image/bmp",
		".svg":  "image/svg+xml",
		".ico":  "image/vnd.microsoft.icon",
		".tiff": "image/tiff",
		".webp": "image/webp",

		// 视频文件
		".mp4":  "video/mp4",
		".avi":  "video/x-msvideo",
		".mov":  "video/quicktime",
		".wmv":  "video/x-ms-wmv",
		".flv":  "video/x-flv",
		".mkv":  "video/x-matroska",
		".webm": "video/webm",
		".3gp":  "video/3gpp",
		".ogv":  "video/ogg",

		// 音频文件
		".mp3":  "audio/mpeg",
		".wav":  "audio/wav",
		".ogg":  "audio/ogg",
		".flac": "audio/flac",
		".aac":  "audio/aac",
		".m4a":  "audio/mp4",
		".wma":  "audio/x-ms-wma",
		".amr":  "audio/amr",

		// 字幕文件
		".srt": "application/x-subrip",
		".vtt": "text/vtt",
		".ass": "text/x-ssa",
		".sub": "text/plain",

		// 压缩文件
		".zip": "application/zip",
		".rar": "application/vnd.rar",
		".tar": "application/x-tar",
		".gz":  "application/gzip",
		".7z":  "application/x-7z-compressed",

		// 文档文件
		".pdf":  "application/pdf",
		".doc":  "application/msword",
		".docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		".xls":  "application/vnd.ms-excel",
		".xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		".ppt":  "application/vnd.ms-powerpoint",
		".pptx": "application/vnd.openxmlformats-officedocument.presentationml.presentation",

		// 其他常用文件类型
		".csv":  "text/csv",
		".rtf":  "application/rtf",
		".epub": "application/epub+zip",
		".azw":  "application/vnd.amazon.ebook",
		".apk":  "application/vnd.android.package-archive",
		".bin":  "application/octet-stream",
		".exe":  "application/vnd.microsoft.portable-executable",
	}
	ext := filepath.Ext(fileName)
	if fileType, ok := extensions[ext]; ok {
		return fileType
	}
	return "unknown"
}
