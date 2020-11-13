package routers

import (
	"firstProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Index")
	beego.Router("/selectPage", &controllers.SelectController{}, "get:SelectPage")
	beego.Router("/addPage", &controllers.AddController{}, "get:AddPage")
	beego.Router("/editPage", &controllers.EditController{}, "get:EditPage")

	beego.Router("/getCoursesList", &controllers.SelectController{}, "get:SelectCourses")
	beego.Router("/getStudentsList", &controllers.SelectController{}, "get:SelectStudents")

	beego.Router("/addCourses", &controllers.AddController{}, "post:AddCourses")
	beego.Router("/addStudents", &controllers.AddController{}, "post:AddStudents")

	beego.Router("/editCourses", &controllers.EditController{}, "post:EditCourses")
	beego.Router("/editStudents", &controllers.EditController{}, "post:EditStudents")
	beego.Router("/delCourses", &controllers.DelController{}, "post:DelCourses")
	beego.Router("/delStudents", &controllers.DelController{}, "post:DelStudents")
}
