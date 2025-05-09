package models

type Template struct {
	Model
	TemplateName  string `gorm:"column:template_name;not null;default:''" json:"template_name"`
	TemplateForms []TemplateForm
}
