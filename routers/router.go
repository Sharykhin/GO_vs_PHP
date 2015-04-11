package routers

import (
	"todos/controllers"
	"github.com/astaxie/beego"
	todo_module "todos/modules/todo/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/todo",&todo_module.TodoController{})
}
