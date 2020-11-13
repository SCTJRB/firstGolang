package controllers

import (
	"firstProject/DB"
	"firstProject/models"
	"fmt"
	"github.com/astaxie/beego"
)

type DelController struct {
	beego.Controller
}


func (this DelController) DelCourses()  {
	id, err := this.GetInt("id")
	if err != nil {
		fmt.Println(err)
		return
	}else {
		db := DB.GetDB()
		if err := db.Where("id = ?", id).Delete(&models.CoursesDel{}).Error;err !=nil{
			fmt.Println("删除失败", err)
			return
		}else {
			res := models.DelJson{
				Err_code: 0,
				Err_msg:  "删除成功",
				Data:     id,
			}
			this.Data["json"] = res
			this.ServeJSON()
		}
	}
}
func (this DelController) DelStudents()  {
	id, err := this.GetInt("id")
	if err != nil {
		fmt.Println(err)
		return
	}else {
		db := DB.GetDB()
		if err := db.Where("id = ?", id).Delete(&models.StudentDel{}).Error;err !=nil{
			fmt.Println("删除失败", err)
			return
		}else {
			res := models.DelJson{
				Err_code: 0,
				Err_msg:  "删除成功",
				Data:     id,
			}
			this.Data["json"] = res
			this.ServeJSON()
		}
	}
}