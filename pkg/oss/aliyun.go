package oss

import (
	"io"
	"log"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

const (
	Endpoint = "https://oss-cn-guangzhou.aliyuncs.com"
	Region   = "cn-guangzhou"
)

type OssInterface interface {
	CreateBucket() error
	UploadFile(file io.Reader, objectName string) (string, error)
	DeleteObject(objectName string) error
}

type OssClient struct {
	client     *oss.Client
	BucketName string
}

func GetClient(_bucketName string) *OssClient {
	endpoint := Endpoint

	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		log.Printf("error: oss get environment variable failed: %v", err)
		return nil
	}
	log.Println("info: oss get environment variable success")

	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// yourRegion填写Bucket所在地域，以华东1（杭州）为例，填写为cn-hangzhou。其它Region请按实际情况填写。
	clientOptions := []oss.ClientOption{oss.SetCredentialsProvider(&provider)}
	clientOptions = append(clientOptions, oss.Region(Region))
	// 设置签名版本
	clientOptions = append(clientOptions, oss.AuthVersion(oss.AuthV4))
	_client, err := oss.New(endpoint, "", "", clientOptions...)
	if err != nil {
		log.Printf("error: oss create client failed : %v", err)
	}

	return &OssClient{client: _client, BucketName: _bucketName}
}

func (o *OssClient) CreateBucket() error {
	bucketName := o.BucketName

	// 创建存储空间。
	err := o.client.CreateBucket(bucketName)
	if err != nil {
		log.Printf("error: create bucket failed : %v", err)
		return err
	}

	// 设置存储空间的读写权限为公共读。
	err = o.client.SetBucketACL(bucketName, oss.ACLPublicRead)
	if err != nil {
		log.Fatalf("Failed to set bucket ACL for '%s': %v", bucketName, err)
	}

	// 获取存储空间的读写权限。
	aclRes, err := o.client.GetBucketACL(bucketName)
	if err != nil {
		log.Fatalf("Failed to get bucket ACL for '%s': %v", bucketName, err)
	}

	// 存储空间创建成功后，记录日志。
	log.Printf("Bucket created successfully: %s ACL: %v", bucketName, aclRes.ACL)
	return nil
}

func (o *OssClient) UploadFile(file io.Reader, objectName string) (string, error) {
	bucketName := o.BucketName

	// 获取存储空间。
	bucket, err := o.client.Bucket(bucketName)
	if err != nil {
		return "", err
	}
	// 上传文件。
	err = bucket.PutObject(objectName, file)
	if err != nil {
		return "", err
	}

	presignedURL, err := bucket.SignURL(objectName, oss.HTTPGet, 604800) // 31536000 秒有效期
	if err != nil {
		return "", err
	}

	parts := strings.Split(presignedURL, "?")
	return parts[0], nil
}

// deleteObject 用于删除OSS存储空间中的一个对象。
// 参数：
//
//	bucketName - 存储空间名称。
//	objectName - 要删除的对象名称。
//	endpoint - Bucket对应的Endpoint。
//
// 如果成功，记录成功日志；否则，返回错误。
func (o *OssClient) DeleteObject(objectName string) error {
	bucketName := o.BucketName
	// 获取存储空间。
	bucket, err := o.client.Bucket(bucketName)
	if err != nil {
		return err
	}

	// 删除文件。
	err = bucket.DeleteObject(objectName)
	if err != nil {
		return err
	}

	// 文件删除成功后，记录日志。
	log.Printf("Object deleted successfully: %s/%s", bucketName, objectName)
	return nil
}
