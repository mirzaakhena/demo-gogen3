package entity

import "your/path/project/domain_mydomain/model/errorenum"

type Person struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func NewPerson(Name string, Age int) (*Person, error) {

	if Name == "" {
		return nil, errorenum.NameMustNotEmpty
	}

	if Age <= 0 {
		return nil, errorenum.AgeMustGreaterThanZero
	}

	var obj Person
	obj.Age = Age
	obj.Name = Name

	return &obj, nil
}
