package model

import (
	"github.com/astaxie/beego/orm"
)		

type Todo struct {
    Id          int 	`form:"id"`
    Title       string	`form:"title"`
    Isdone      bool	
}

func init() {
    // Need to register model in init
    orm.RegisterModel(new(Todo))
}

