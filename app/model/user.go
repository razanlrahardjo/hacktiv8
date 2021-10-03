package model

type User struct {
	Base
	Name *string `json:"name,omitempty" gorm:"type:varchar(64)"`
}

func (User) TableName() string {
	return "user"
}

func (user *User) Validation(c string) string {
	switch c {
	case "create":
		if user.Name == nil {
			return "Required Name"
		}
	}
	return ""
}
