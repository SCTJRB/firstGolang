package controllers

import (
	"firstProject/DB"
	"firstProject/models"
	"fmt"
	"github.com/astaxie/beego"
)

type EditController struct {
	beego.Controller
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

	db := DB.GetDB()
	c := models.CoursesEdit{}
	db.Where("id = ?", id).Take(&c)

	if err := db.Model(&c).Update("name", name).Error; err != nil {
		fmt.Println("更新失败", err)
		return
	} else {
		res := models.EditJson{
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
	db := DB.GetDB()
	c := models.StudentEdit{}
	db.Where("id = ?", id).Take(&c)
	if err := db.Model(&c).Update("name", name).Error; err != nil {
		fmt.Println("更新失败", err)
		return
	} else {
		res := models.EditJson{
			Err_code: 0,
			Err_msg:  "更新成功",
			Data:     id,
		}
		this.Data["json"] = res
		this.ServeJSON()
	}
}
