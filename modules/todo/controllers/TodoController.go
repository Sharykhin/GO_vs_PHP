package todo_module

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	models "todos/modules/todo/models"
	"fmt"
    "strconv"
)

type TodoController struct {
	beego.Controller
}

func (this *TodoController) Prepare() {
	
	beego.ViewsPath="modules/todo/views"
	this.Layout = "layout.tpl"
}

func (this *TodoController) Get() {
    beego.ReadFromRequest(&this.Controller)
    
    o := orm.NewOrm()
    o.Using("default") 
    var todos []models.Todo
    o.QueryTable("todo").All(&todos)
    this.TplNames = "list.tpl"
    this.LayoutSections = make(map[string]string)
    this.LayoutSections["FormTodo"] = "form.tpl"
    this.Data["todos"]=todos
    
}

func (this *TodoController) Post() {
	
	o := orm.NewOrm()
    o.Using("default") 

    todo := models.Todo{}
  
    if err := this.ParseForm(&todo); err != nil {
    	this.Ctx.WriteString(fmt.Sprintf("%v",err))    	
    }

    if todo.Id != 0 {

         o.Update(&todo)         

    } else {

         if _,err := o.Insert(&todo); err != nil {
            this.Ctx.WriteString(fmt.Sprintf("%v",err))     
        } 
    } 
    
    flash:=beego.NewFlash()
    flash.Notice("Todo has been deleted successfully")
    flash.Store(&this.Controller)

    this.Ctx.Redirect(302, "/todo")
    return
}

func (this *TodoController) ReadTodo() {
    o := orm.NewOrm()
    o.Using("default") 

    id := this.Ctx.Input.Param(":id")
    todoid,err := strconv.Atoi(id)
    if err != nil {
        this.Ctx.WriteString(fmt.Sprintf("%v",err)) 
    }
    todo := models.Todo{}
    todo.Id = todoid

    err = o.Read(&todo)

    if err == orm.ErrNoRows {
        this.Ctx.WriteString("No result found.")       
    } else if err == orm.ErrMissPK {
        this.Ctx.WriteString("No primary key found.")         
    } 
    this.Data["todo"] = todo
    this.TplNames = "form.tpl"

}

func (this *TodoController) DeleteTodo() {
    o := orm.NewOrm()
    o.Using("default") 

    id := this.Ctx.Input.Param(":id")
    todoid,err := strconv.Atoi(id)
    if err != nil {
        this.Ctx.WriteString(fmt.Sprintf("%v",err))
    }

    if _,err := o.Delete(&models.Todo{Id:todoid}); err != nil {
        this.Ctx.WriteString(fmt.Sprintf("%v",err))
    }

    flash:=beego.NewFlash()
    flash.Notice("Todo has been deleted successfully")
    flash.Store(&this.Controller)

    this.Ctx.Redirect(302,"/todo")
    return

}