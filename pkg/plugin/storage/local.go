package storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
)

func (d *Driver) uploadLocal(file *multipart.FileHeader, key string, folder string) error {
	directory := cfg.Storage.LocalPath
	if len(directory) == 0 {
		directory = "./public/uploads"
	}
	src, err := file.Open()
	if err != nil {
		return fmt.Errorf("打开文件失败: %v", err)
	}
	defer src.Close()

	savePath := path.Join(directory, folder, path.Dir(key))
	saveFilePath := path.Join(directory, folder, key)

	err = os.MkdirAll(savePath, 0755)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("创建目录失败: %v", err)
	}

	out, err := os.Create(saveFilePath)
	if err != nil {
		return fmt.Errorf("创建文件失败: %v", err)
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	if err != nil {
		return fmt.Errorf("写入文件失败: %v", err)
	}
	return nil
}

func (d *Driver) removeLocal(filePath string) error {
	fullPath := path.Join(cfg.Storage.LocalPath, filePath)

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return fmt.Errorf("文件不存在: %v", err)
	}

	// 删除文件
	if err := os.Remove(fullPath); err != nil {
		return fmt.Errorf("删除文件失败: %v", err)
	}

	return nil
}
