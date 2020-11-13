package models
type CoursesTable struct {
	Id int `json:"id"`
	Name string `json:"name"`
}
func (v CoursesTable) TableName() string {
	return "courses"
}
type StudentsTable struct {
	Id int `json:"id"`
	Name string `json:"name"`
}
func (v StudentsTable) TableName() string {
	return "student"
}
type SelectJson struct {
	Err_code int         `json:"err_code"`
	Err_msg  string      `json:"err_msg"`
	Data     interface{} `json:"data"`
}