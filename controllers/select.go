package controllers

import (
	"firstProject/DB"
	"firstProject/models"
	"fmt"
	"github.com/astaxie/beego"
)

type SelectController struct {
	beego.Controller
}





func (this *SelectController) SelectPage() {
	this.TplName = "selectPage.html"
}

func (this *SelectController) SelectCourses() {
	db := DB.GetDB()
	var c [] models.CoursesTable
	if err := db.Find(&c).Error;err != nil {
		fmt.Println("查询失败", err)
	}else {
		res := models.SelectJson{
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
	var s [] models.StudentsTable
	if err := db.Find(&s).Error;err != nil {
		fmt.Println("查询失败", err)
	}else {
		res := models.SelectJson{
			Err_code: 0,
			Err_msg:  "请求成功",
			Data:     s,
		}
		this.Data["json"] = res
		this.ServeJSON()
	}
}