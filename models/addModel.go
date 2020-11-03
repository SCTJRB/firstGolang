package models

type Courses struct {
	Id   int    `gorm:"AUTO_INCREMENT"`
	Name string `gorm:"name"`
}

func (c Courses) TableName() string {
	return "courses"
}

type Student struct {
	Id   int    `gorm:"AUTO_INCREMENT",json:"id"`
	Name string `gorm:"name",json:"name"`
}

func (c Student) TableName() string {
	return "student"
}

type AddJson struct {
	Err_code int         `json:"err_code"`
	Err_msg  string      `json:"err_msg"`
	Data     interface{} `json:"data"`
}