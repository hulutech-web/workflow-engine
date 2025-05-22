package req

type UploadReq struct {
	Cid uint `form:"cid" json:"cid" validate:"gte=0"`
}

type FileStorageReq struct {
	EngineType string                 `json:"engine_type" validate:"required"`
	Engine     map[string]interface{} `json:"engine" form:"engine"`
	ImageExt   []string               `json:"image_ext" form:"image_ext"`
	VideoExt   []string               `json:"video_ext" form:"video_ext"`
	AudioExt   []string               `json:"audio_ext" form:"audio_ext"`
	FileExt    []string               `json:"file_ext" form:"file_ext"`
	MaxSize    int64                  `json:"max_size" form:"max_size"`
}

// FileListReq 文件列表参数
type FileListReq struct {
	Cid  int    `form:"cid,default=-1"`                       // 类目ID
	Type int    `form:"type" binding:"omitempty,oneof=10 20"` // 文件类型: [10=图片, 20=视频]
	Name string `form:"keyword"`                              // 文件名称
}

// FileRenameReq 文件重命名参数
type FileRenameReq struct {
	ID   uint   `form:"id" binding:"required,gt=0"`              // 主键
	Name string `form:"keyword" binding:"required,min=1,max=30"` // 文件名称
}

// FileMoveReq 文件移动参数
type FileMoveReq struct {
	Ids []uint `form:"ids" binding:"required"` // 主键
	Cid uint   `form:"cid,default=0"`          // 类目ID
}

// FileAddReq 文件新增参数
type FileAddReq struct {
	Cid  uint   `form:"cid" binding:"gte=0"`        // 类目ID
	Aid  uint   `form:"aid" binding:"gte=0"`        // 管理ID
	Uid  uint   `form:"uid" binding:"gte=0"`        // 用户ID
	Type int    `form:"type" binding:"oneof=10 20"` // 文件类型: [10=图片, 20=视频]
	Name string `form:"name"`                       // 文件名称
	Uri  string `form:"uri"`                        // 文件路径
	Ext  string `form:"ext"`                        // 文件扩展
	Size int64  `form:"size"`                       // 文件大小
}

// FileCateListReq 附件分类列表参数
type FileCateListReq struct {
	Type int    `form:"type" binding:"omitempty,oneof=10 20 30"` // 分类类型: [10=图片,20=视频]
	Name string `form:"keyword"`                                 // 分类名称
}

// FileCateAddReq 附件分类新增参数
type FileCateAddReq struct {
	Pid  uint   `form:"pid" binding:"gte=0"`                    // 父级ID
	Type int    `form:"type" binding:"required,oneof=10 20 30"` // 分类类型: [10=图片,20=视频]
	Name string `form:"name" binding:"required,min=1,max=30"`   // 分类名称
}

// FileCateRenameReq 附件分类重命名参数
type FileCateRenameReq struct {
	ID   uint   `form:"id" binding:"required,gt=0"`              // 主键
	Name string `form:"keyword" binding:"required,min=1,max=30"` // 分类名称
}
