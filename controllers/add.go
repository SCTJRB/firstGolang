package controllers

import (
	"firstProject/DB"
	"firstProject/models"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type AddController struct {
	beego.Controller
}

func (this *AddController) AddPage() {
	this.TplName = "addPage.html"
}
//添加课程
func (this *AddController) AddCourses() {
	name := this.GetString("name")
	if name != "" {
		db := DB.GetDB()
		c := &models.Courses{Name: name}
		if err := db.Create(c).Error; err != nil {
			fmt.Println("插入失败", err)
			return
		} else {
			res := models.AddJson{
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
		c := &models.Student{Name: name}
		if err := db.Create(c).Error; err != nil {
			fmt.Println("插入失败", err)
			return
		} else {
			res := models.AddJson{
				Err_code: 0,
				Err_msg:  "请求成功",
				Data:     c.Id,
			}
			this.Data["json"] = res
			this.ServeJSON()
		}
	}
}
