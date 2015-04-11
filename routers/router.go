package routers

import (
	"todos/controllers"
	"github.com/astaxie/beego"
	todo_module "todos/modules/todo/controllers"
)

func init() {

    beego.Router("/", &controllers.MainController{})
    beego.Router("/todo",&todo_module.TodoController{})
    beego.Router("/todo/:id:int",&todo_module.TodoController{},"get:ReadTodo")
    beego.Router("/todo/delete/:id:int",&todo_module.TodoController{},"get:DeleteTodo")
    beego.SetStaticPath("/todo/static","static")
}
