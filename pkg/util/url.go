package util

import (
	"fmt"
	"github.com/hulutech-web/workflow-engine/core/config"
	"go.uber.org/zap"
	"log"
	"net/url"
	"path"
	"strings"
)

var (
	UrlUtil      = urlUtil{}
	publicUrl    = config.NewConfig().Server.PublicUrl
	publicPrefix = config.NewConfig().Storage.PublicPrefix
)

// urlUtil 文件路径处理工具
type urlUtil struct{}

// ToAbsoluteUrl 转绝对路径
func (uu urlUtil) ToAbsoluteUrl(u string, engine string, cfg map[string]interface{}) string {
	if u == "" {
		return ""
	}

	// Minio处理
	if engine == "minio" {
		if cfg["prefix"] != "" {
			u = path.Join(cfg["prefix"].(string), u)
		}
		uri := fmt.Sprintf("%s/%s/%s", cfg["endpoint"].(string), cfg["bucket"].(string), u)
		// 如果链接中包含http://或https://，则直接返回
		if strings.HasPrefix(u, "http://") || strings.HasPrefix(u, "https://") {
			return u
		}
		uri = fmt.Sprintf("%s://%s", cfg["http_prefix"].(string), uri)
		return uri
	} else if engine == "local" {
		log.Printf("ToAbsoluteUrl publicUrl=[%+v] publicPrefix=[%+v]", publicUrl, publicPrefix)
		// 本地存储处理
		up, err := url.Parse(publicUrl)
		if err != nil {
			log.Printf("ToAbsoluteUrl Parse publicUrl err: err=[%+v]", err)
			zap.S().Errorf("ToAbsoluteUrl Parse err: err=[%+v]", err)
			return u
		}
		zap.S().Info("ToAbsoluteUrl up=[%+v]", up)
		if strings.HasPrefix(u, "/api/static/") {
			up.Path = path.Join(up.Path, u)
			return up.String()
		}

		up.Path = path.Join(up.Path, publicPrefix, u)
		return up.String()
	}
	return u
}

func (uu urlUtil) ToRelativeUrl(u string, engine string) string {
	// TODO: engine默认local
	if u == "" {
		return ""
	}
	up, err := url.Parse(u)
	if err != nil {
		zap.S().Errorf("ToRelativeUrl Parse err: err=[%+v]", err)
		return u
	}

	if engine == "local" {
		lu := up.String()
		return strings.Replace(
			strings.Replace(lu, publicUrl, "", 1),
			publicPrefix, "", 1)
	}
	// TODO: 其他engine
	return u
}
