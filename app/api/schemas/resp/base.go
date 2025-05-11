package resp

type SelectOption struct {
	Label    string      `json:"label" structs:"label"`
	Value    interface{} `json:"value" structs:"value"`
	Disabled bool        `json:"disabled" structs:"disabled"`
}
