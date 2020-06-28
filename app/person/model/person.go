package model

// Person
type Person struct {
	ID        int    `gorm:"column:id;primary_key" json:"id"`
	FirstName string `gorm:"column:first_name" json:"first_name"`
	LastName  string `gorm:"column:last_name" json:"last_name"`
}

func (Person) TableName() string {
	return "person"
}
