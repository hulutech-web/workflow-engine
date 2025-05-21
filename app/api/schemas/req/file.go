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
