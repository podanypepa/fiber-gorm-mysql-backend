package controller

import "github.com/podanypepa/fiber-gorm-mysql-backend/pkg/model/todo"

type resp struct {
	Status  string       `json:""`
	Message string       `json:",omitempty"`
	Data    *[]todo.TODO `json:",omitempty"`
}
