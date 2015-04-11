package todo_module

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	models "todos/modules/todo/models"
	"fmt"
)

type TodoController struct {
	beego.Controller
}

func (this *TodoController) Prepare() {
	
	beego.ViewsPath="modules/todo/views"
	this.Layout = "layout.tpl"
}

func (this *TodoController) Get() {

    this.TplNames = "list.tpl"
    this.LayoutSections = make(map[string]string)
    this.LayoutSections["FormTodo"] = "form.tpl"
    
}

func (this *TodoController) Post() {
	
	o := orm.NewOrm()
    o.Using("default") 

    todo := models.Todo{}
  
    if err := this.ParseForm(&todo); err != nil {
    	this.Ctx.WriteString(fmt.Sprintf("%v",err))    	
    }
    
    if _,err := o.Insert(&todo); err != nil {
    	this.Ctx.WriteString(fmt.Sprintf("%v",err))    	
    } 

    this.Ctx.Redirect(302, "/todo")
}