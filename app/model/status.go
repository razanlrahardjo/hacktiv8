package model

type Status struct {
	Base
	StatusText *string `json:"status_text,omitempty" gorm:"type:varchar(10)"`
}

func (Status) TableName() string {
	return "status"
}

func (Status *Status) Validation(c string) string {
	switch c {
	case "create":
		if Status.StatusText == nil {
			return "Required Status Text"
		}
	}
	return ""
}
