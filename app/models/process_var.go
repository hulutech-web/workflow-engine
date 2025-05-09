package models

type ProcessVar struct {
	Model
	ProcessID       int      `gorm:"column:process_id;not null" json:"process_id"`
	FlowID          int      `gorm:"column:flow_id;not null;comment:'流程id'" json:"flow_id"`
	ExpressionField string   `gorm:"column:expression_field;not null;comment:'条件表达式字段名称'" json:"expression_field"`
	Process         *Process `gorm:"foreignKey:ProcessID;references:ID"`
}
