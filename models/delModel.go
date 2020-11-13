package models

type CoursesDel struct {
	Id   int    `gorm:"AUTO_INCREMENT"`
	Name string `gorm:"name"`
}

func (c CoursesDel) TableName() string {
	return "courses"
}

type StudentDel struct {
	Id   int    `gorm:"AUTO_INCREMENT",json:"id"`
	Name string `gorm:"name",json:"name"`
}

func (c StudentDel) TableName() string {
	return "student"
}


type DelJson struct {
	Err_code int         `json:"err_code"`
	Err_msg  string      `json:"err_msg"`
	Data     interface{} `json:"data"`
}
