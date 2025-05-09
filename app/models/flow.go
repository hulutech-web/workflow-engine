package models

type Flow struct {
	Model
	FlowNo      string       `gorm:"column:flow_no;not null" json:"flow_no" form:"flow_no"`
	FlowName    string       `gorm:"column:flow_name;not null;default:''" json:"flow_name" form:"flow_name"`
	TemplateID  int          `gorm:"column:template_id;not null;default:0" json:"template_id" form:"template_id"`
	Flowchart   string       `gorm:"column:flowchart" json:"flowchart" form:"flowchart"`
	Jsplumb     string       `gorm:"column:jsplumb;comment:'jsplumb流程图数据'" json:"jsplumb" form:"jsplumb"`
	TypeID      int          `gorm:"column:type_id;not null;default:0" json:"type_id" form:"type_id"`
	IsPublish   bool         `gorm:"column:is_publish;not null;default:0" json:"is_publish" form:"is_publish"`
	IsShow      bool         `gorm:"column:is_show;not null;default:1" json:"is_show" form:"is_show"`
	Processes   []Process    `gorm:"foreignKey:FlowID"`     // HasMany Process
	ProcessVars []ProcessVar `gorm:"foreignKey:FlowID"`     // HasMany ProcessVar
	Template    Template     `gorm:"foreignKey:TemplateID"` // BelongsTo Template
	Flowtype    Flowtype     `gorm:"foreignKey:TypeID"`     // BelongsTo FlowType
}
