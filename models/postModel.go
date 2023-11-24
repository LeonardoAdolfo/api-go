package model

import "gorm.io/gorm"

//Modelo da tabela
type Post struct {
	gorm.Model
	Title string
	Body  string
}
