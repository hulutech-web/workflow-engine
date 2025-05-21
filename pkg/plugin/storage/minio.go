package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"mime/multipart"
	"path"
	"strings"
)

func (d *Driver) uploadMinio(file *multipart.FileHeader, key string, folder string, cfg map[string]interface{}) error {
	client, err := createMinioClient(cfg)
	if err != nil {
		return err
	}
	bucket, ok := cfg["bucket"].(string)
	if !ok {
		return errors.New("Minio bucket 未正确配置！")
	}
	objectName := path.Join(folder, key)
	fileReader, err := file.Open()
	if err != nil {
		return fmt.Errorf("打开文件失败：%v", err)
	}
	defer fileReader.Close()

	// 上传文件
	_, err = client.PutObject(
		context.Background(),
		bucket,
		objectName,
		fileReader,
		file.Size,
		minio.PutObjectOptions{
			ContentType: file.Header.Get("Content-Type"),
		},
	)
	if err != nil {
		return fmt.Errorf("上传文件到 Minio 失败：%v", err)
	}
	return nil
}

func createMinioClient(c map[string]interface{}) (*minio.Client, error) {
	if c == nil {
		return nil, errors.New("Minio 配置不正确！")
	}
	endpoint, ok := c["endpoint"].(string)
	if !ok {
		return nil, errors.New("Minio endpoint 未正确配置！")
	}
	accessKey, ok := c["accessKey"].(string)
	if !ok {
		return nil, errors.New("Minio accessKey 未正确配置！")
	}
	secretKey, ok := c["secretKey"].(string)
	if !ok {
		return nil, errors.New("Minio secretKey 未正确配置！")
	}
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: strings.HasPrefix(endpoint, "https://"),
	})
	if err != nil {
		return nil, fmt.Errorf("创建 Minio 客户端失败：%v", err)
	}
	return client, nil
}

func (d *Driver) removeMinio(filePath string, cfg map[string]interface{}) error {
	client, err := createMinioClient(cfg)
	if err != nil {
		return err
	}
	bucket, ok := cfg["bucket"].(string)
	if !ok {
		return errors.New("Minio bucket 未正确配置！")
	}
	err = client.RemoveObject(
		context.Background(),
		bucket,
		filePath,
		minio.RemoveObjectOptions{},
	)
	if err != nil {
		return fmt.Errorf("删除 Minio 文件失败：%v", err)
	}
	return nil
}
