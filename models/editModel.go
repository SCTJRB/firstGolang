package models

type CoursesEdit struct {
	Id   int    `gorm:"AUTO_INCREMENT"`
	Name string `gorm:"name"`
}

func (c CoursesEdit) TableName() string {
	return "courses"
}

type StudentEdit struct {
	Id   int    `gorm:"AUTO_INCREMENT",json:"id"`
	Name string `gorm:"name",json:"name"`
}

func (c StudentEdit) TableName() string {
	return "student"
}

type EditJson struct {
	Err_code int         `json:"err_code"`
	Err_msg  string      `json:"err_msg"`
	Data     interface{} `json:"data"`
}