package controllers

import (
	"firstProject/DB"
	"fmt"
	"github.com/astaxie/beego"
)

type DelController struct {
	beego.Controller
}

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


type delJson struct {
	Err_code int         `json:"err_code"`
	Err_msg  string      `json:"err_msg"`
	Data     interface{} `json:"data"`
}

func (this DelController) DelCourses()  {
	id, err := this.GetInt("id")
	if err != nil {
		fmt.Println(err)
		return
	}else {
		db := DB.GetDB()
		if err := db.Where("id = ?", id).Delete(&CoursesDel{}).Error;err !=nil{
			fmt.Println("删除失败", err)
			return
		}else {
			res := delJson{
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
		if err := db.Where("id = ?", id).Delete(&StudentDel{}).Error;err !=nil{
			fmt.Println("删除失败", err)
			return
		}else {
			res := delJson{
				Err_code: 0,
				Err_msg:  "删除成功",
				Data:     id,
			}
			this.Data["json"] = res
			this.ServeJSON()
		}
	}
}