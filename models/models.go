package models

import "github.com/jinzhu/gorm"

const  (
	TERMINALS = "terminal"
	Services  = "service"
)

func GetAll(table string) interface{} {
	db := Connect()
	defer db.Close()
	switch table {
	case TERMINALS:
	return db.Order("id  asc").Find(&[]Terminal{}).Value
	case Services:
		return db.Order("id asc").Find(&[]Service{}).Value
	case USER:
		return db.Order("id asc").Find(&[]User{}).Value

	}
	return nil

}