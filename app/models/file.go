package models

type File struct {
	Model
	Cid      uint   `gorm:"not null;default:0;comment:'类目ID';index"`
	UserId   uint   `gorm:"not null;default:0;comment:'用户ID'"`
	Type     string `gorm:"not null;default:'image';comment:'文件类型: [image=图片, video=视频, audio=音频, file=文件]';index"`
	Name     string `gorm:"not null;default:'';comment:'文件名称''"`
	Uri      string `gorm:"not null;comment:'文件路径'"`
	Ext      string `gorm:"not null;default:'';comment:'文件扩展'"`
	Size     int64  `gorm:"not null;default:0;comment:文件大小"`
	Engine   string `gorm:"not null;default:'';comment:'存储引擎'"`
	Path     string `gorm:"not null;default:'';comment:'访问路径'"`
	TenantId uint   `gorm:"not null;default:0;comment:'租户ID'"`
}

type FileCate struct {
	Model
	Name     string `gorm:"not null;default:'';comment:'类目名称'"`
	Sort     int    `gorm:"not null;default:0;comment:'排序'"`
	Type     string `gorm:"not null;default:'';comment:'类目类型: [image=图片, video=视频, audio=音频, file=文件]';index"`
	Pid      uint   `gorm:"not null;default:0;comment:'父类目ID'"`
	TenantId uint   `gorm:"not null;default:0;comment:'租户ID'"`
}
