package controllers

import (
	"firstProject/models"
	"fmt"
	"github.com/astaxie/beego"
)

type EditController struct {
	beego.Controller
}

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

type editJson struct {
	Err_code int         `json:"err_code"`
	Err_msg  string      `json:"err_msg"`
	Data     interface{} `json:"data"`
}

func (this *EditController) EditPage() {
	this.TplName = "editPage.html"
}
func (this *EditController) EditCourses() {
	id, err := this.GetInt("id")
	if err != nil {
		fmt.Println(err)
		return
	}
	name := this.GetString("name")
	fmt.Println(name)
	db := models.GetDB()
	c := CoursesEdit{}
	db.Where("id = ?", id).Take(&c)
	if err := db.Model(&c).Update("name", name).Error; err != nil {
		fmt.Println("更新失败", err)
		return
	} else {
		res := editJson{
			Err_code: 0,
			Err_msg:  "更新成功",
			Data:     id,
		}
		this.Data["json"] = res
		this.ServeJSON()
	}
}
func (this *EditController) EditStudents() {
	id, err := this.GetInt("id")
	if err != nil {
		fmt.Println(err)
		return
	}
	name := this.GetString("name")
	fmt.Println(name)
	db := models.GetDB()
	c := StudentEdit{}
	db.Where("id = ?", id).Take(&c)
	if err := db.Model(&c).Update("name", name).Error; err != nil {
		fmt.Println("更新失败", err)
		return
	} else {
		res := editJson{
			Err_code: 0,
			Err_msg:  "更新成功",
			Data:     id,
		}
		this.Data["json"] = res
		this.ServeJSON()
	}
}
