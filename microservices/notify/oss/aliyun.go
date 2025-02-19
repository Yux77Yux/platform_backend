package oss

import (
	"io"

	oss "github.com/Yux77Yux/platform_backend/pkg/oss"
)

type OssInterface interface {
	CreateBucket() error
	UploadFile(file io.Reader, objectName string) (string, error)
	DeleteObject(objectName string) error
}

var Client OssInterface

func Init() {
	Client = oss.GetClient("platform-user")
	Client.CreateBucket()
}
