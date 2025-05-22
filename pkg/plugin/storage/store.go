package storage

import (
	"fmt"
	"github.com/hulutech-web/workflow-engine/core/config"
	"github.com/hulutech-web/workflow-engine/pkg/util"
	"mime/multipart"
	"path"
	"strings"
	"time"
)

var (
	cfg = config.NewConfig()
)

// Driver 存储引擎
type Driver struct {
	conf   map[string]interface{}
	engine string
}

// UploadFile 文件对象
type UploadFile struct {
	Name   string
	Type   string
	Size   int64
	Ext    string
	Uri    string
	Path   string
	Engine string
}

// NewStorageDriver 实例化存储引擎
func NewStorageDriver(conf map[string]interface{}) *Driver {
	engine := "local"
	if conf == nil {
		engine = conf["engine_type"].(string)
		conf = make(map[string]interface{})
		conf["engine"] = map[string]interface{}{}
		conf["image_ext"] = []string{".jpg", ".jpeg", ".png", ".gif", ".bmp"}
		conf["video_ext"] = []string{".mp4", ".avi", ".rmvb", ".wmv", ".flv"}
		conf["audio_ext"] = []string{".mp3", ".wav", ".wma", ".aac"}
		conf["file_ext"] = []string{".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".pdf", ".txt", ".md", ".zip", ".rar", ".tar", ".gz", ".bz2", ".7z"}
		conf["max_size"] = 1024 * 1024 * 1024 // 1G
	}
	return &Driver{
		conf:   conf,
		engine: engine,
	}
}

// checkFile
func (d *Driver) checkFile(file *multipart.FileHeader, fileType string) error {
	fileName := file.Filename
	fileExt := strings.ToLower(strings.Replace(path.Ext(fileName), ".", "", 1))
	fileSize := file.Size
	cSize, ok := d.conf["max_size"].(int64)
	if !ok {
		cSize = 1024 * 1024 * 1024 // 1G
	}
	switch fileType {
	case "image":
		exts, ok := d.conf["image_ext"].([]string)
		if !ok {
			break
		}
		if !util.ToolsUtil.Contains(exts, fileExt) {
			return fmt.Errorf("不支持的图片格式： %s", fileExt)
		}
		if fileSize > cSize {
			return fmt.Errorf("图片大小超过限制： %dM", fileSize/1024/1024)
		}
	case "video":
		exts, ok := d.conf["video_ext"].([]string)
		if !ok {
			break
		}
		if !util.ToolsUtil.Contains(exts, fileExt) {
			return fmt.Errorf("不支持的视频格式： %s", fileExt)
		}
		if fileSize > cSize {
			return fmt.Errorf("视频大小超过限制： %dM", fileSize/1024/1024)
		}
	case "audio":
		exts, ok := d.conf["audio_ext"].([]string)
		if !ok {
			break
		}
		if !util.ToolsUtil.Contains(exts, fileExt) {
			return fmt.Errorf("不支持的音频格式： %s", fileExt)
		}
		if fileSize > cSize {
			return fmt.Errorf("音频大小超过限制： %dM", fileSize/1024/1024)
		}
	case "file":
		exts, ok := d.conf["file_ext"].([]string)
		if !ok {
			break
		}
		if !util.ToolsUtil.Contains(exts, fileExt) {
			return fmt.Errorf("不支持的文件格式： %s", fileExt)
		}
		if fileSize > cSize {
			return fmt.Errorf("文件大小超过限制： %dM", fileSize/1024/1024)
		}
	default:
		return fmt.Errorf("不支持的文件类型： %s", fileType)
	}
	return nil
}

// buildSaveName
func (d *Driver) buildSaveName(file *multipart.FileHeader) string {
	name := file.Filename
	ext := strings.ToLower(path.Ext(name))
	date := time.Now().Format("20060201")
	return path.Join(date, util.ToolsUtil.MakeUuid()+ext)
}

// Upload 上传文件
func (d *Driver) Upload(file *multipart.FileHeader, folder string, fileType string, engine string, cfg map[string]interface{}) (*UploadFile, error) {
	if e := d.checkFile(file, fileType); e != nil {
		return nil, e
	}
	key := d.buildSaveName(file)
	if engine == "" {
		engine = "local"
	}
	switch engine {
	case "local":
		if err := d.uploadLocal(file, key, folder); err != nil {
			return nil, err
		}
	case "minio":
		if err := d.uploadMinio(file, key, folder, cfg); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("不支持的存储引擎： %s", engine)
	}

	fileRelPath := path.Join(folder, key)
	return &UploadFile{
		Name: file.Filename,
		Type: fileType,
		Size: file.Size,
		Ext:  strings.ToLower(strings.Replace(path.Ext(file.Filename), ".", "", 1)),
		Uri:  fileRelPath,
		Path: util.UrlUtil.ToAbsoluteUrl(fileRelPath, engine, cfg),
	}, nil

}

// Remove 删除文件
func (d *Driver) Remove(filePath string, engine string, cfg map[string]interface{}) error {
	switch engine {
	case "local":
		if err := d.removeLocal(filePath); err != nil {
			return err
		}
	case "minio":
		if err := d.removeMinio(filePath, cfg); err != nil {
			return err
		}
	default:
		return fmt.Errorf("不支持的存储引擎： %s", engine)
	}
	return nil
}
