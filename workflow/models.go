package workflow

import "time"

type Process struct {
	ID          uint
	Name        string
	Description string
	Steps       []Step
	CurrentStep int
	Status      string // "draft", "active", "completed", "cancelled"
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Step struct {
	ID          uint
	Name        string
	Description string
	Approvers   []Approver
	Actions     []Action
	Conditions  []Condition
}

type Approver struct {
	UserID    uint
	Role      string
	Approved  bool
	Comment   string
	Timestamp time.Time
}

type Action struct {
	Type    string
	Handler string
	Params  map[string]interface{}
}

type Condition struct {
	Expression string
	Met        bool
}
