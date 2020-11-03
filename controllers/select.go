package controllers

import (
	"firstProject/DB"
	"fmt"
	"github.com/astaxie/beego"
)

type SelectController struct {
	beego.Controller
}

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


type selectJson struct {
	Err_code int         `json:"err_code"`
	Err_msg  string      `json:"err_msg"`
	Data     interface{} `json:"data"`
}



func (this *SelectController) SelectPage() {
	this.TplName = "selectPage.html"
}

func (this *SelectController) SelectCourses() {
	db := DB.GetDB()
	var c [] CoursesTable
	if err := db.Find(&c).Error;err != nil {
		fmt.Println("查询失败", err)
	}else {
		res := selectJson{
			Err_code: 0,
			Err_msg:  "请求成功",
			Data:     c,
		}
		this.Data["json"] = res
		this.ServeJSON()
	}
}

func (this *SelectController) SelectStudents() {
	db := DB.GetDB()
	var s [] StudentsTable
	if err := db.Find(&s).Error;err != nil {
		fmt.Println("查询失败", err)
	}else {
		res := selectJson{
			Err_code: 0,
			Err_msg:  "请求成功",
			Data:     s,
		}
		this.Data["json"] = res
		this.ServeJSON()
	}
}