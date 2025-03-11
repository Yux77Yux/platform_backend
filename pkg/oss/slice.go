package oss

import (
	"fmt"
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// 新增结构体管理上传状态
type MultipartUploader struct {
	imur   oss.InitiateMultipartUploadResult
	parts  []oss.UploadPart
	bucket *oss.Bucket
}

// 初始化分片上传
func (o *OssClient) InitMultipartUpload(objectName string) (*MultipartUploader, error) {
	bucket, err := o.client.Bucket(o.BucketName)
	if err != nil {
		return nil, fmt.Errorf("get bucket failed: %w", err)
	}

	imur, err := bucket.InitiateMultipartUpload(objectName)
	if err != nil {
		return nil, fmt.Errorf("init upload failed: %w", err)
	}

	return &MultipartUploader{
		imur:   imur,
		bucket: bucket,
		parts:  make([]oss.UploadPart, 0),
	}, nil
}

// 上传单个分片
func (m *MultipartUploader) UploadPart(file io.Reader, partSize int64, partNumber int) error {
	part, err := m.bucket.UploadPart(m.imur, file, partSize, partNumber)
	if err != nil {
		return fmt.Errorf("upload part %d failed: %w", partNumber, err)
	}

	m.parts = append(m.parts, part)
	return nil
}

// 完成上传
func (m *MultipartUploader) Complete() (string, error) {
	cmur, err := m.bucket.CompleteMultipartUpload(m.imur, m.parts)
	if err != nil {
		return "", fmt.Errorf("complete failed: %w", err)
	}
	return cmur.Location, nil
}

// 终止上传
func (m *MultipartUploader) Abort() error {
	return m.bucket.AbortMultipartUpload(m.imur)
}
