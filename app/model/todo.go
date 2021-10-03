package model

type Todo struct {
	Base
	Title          *string `json:"title,omitempty" gorm:"type:text"`
	Description    *string `json:"description,omitempty" gorm:"type:text"`
	DueDate        *string `json:"due_date,omitempty" gorm:"type:date"`
	PersonInCharge *string `json:"person_in_charge,omitempty" gorm:"type:varchar(256)"`
	Status         *string `json:"status,omitempty" gorm:"type:varchar(10)"`
}

func (Todo) TableName() string {
	return "todo"
}

func (todo *Todo) Validation(c string) string {
	switch c {
	case "update":
		if todo.Status != nil {
			if *todo.Status == "Done" {
				return "Can't Change Status When is Done"
			}
			if *todo.Status == "Delete" {
				return "Can't Change Status When is Delete"
			}
		}
	case "create":
		if todo.Title == nil {
			return "Required Title"
		} else if todo.DueDate == nil {
			return "Required Due Date"
		} else if todo.Description == nil {
			return "Required Description"
		} else if todo.PersonInCharge == nil {
			return "Required Person in Charge"
		} else if todo.Status == nil {
			return "Required Status"
		}
	}
	return ""
}
