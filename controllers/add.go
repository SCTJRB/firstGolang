package controllers

import (
	"firstProject/DB"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type AddController struct {
	beego.Controller
}
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

type addJson struct {
	Err_code int         `json:"err_code"`
	Err_msg  string      `json:"err_msg"`
	Data     interface{} `json:"data"`
}

func (this *AddController) AddPage() {
	this.TplName = "addPage.html"
}
//添加课程
func (this *AddController) AddCourses() {
	name := this.GetString("name")
	if name != "" {
		db := DB.GetDB()
		c := &Courses{Name: name}
		if err := db.Create(c).Error; err != nil {
			fmt.Println("插入失败", err)
			return
		} else {
			res := addJson{
				Err_code: 0,
				Err_msg:  "请求成功",
				Data:     c.Id,
			}
			this.Data["json"] = res
			this.ServeJSON()
		}
	}
}
//添加学生
func (this *AddController) AddStudents() {
	name := this.GetString("name")
	if name != "" {
		db :=  DB.GetDB()
		c := &Student{Name: name}
		if err := db.Create(c).Error; err != nil {
			fmt.Println("插入失败", err)
			return
		} else {
			res := addJson{
				Err_code: 0,
				Err_msg:  "请求成功",
				Data:     c.Id,
			}
			this.Data["json"] = res
			this.ServeJSON()
		}
	}
}
