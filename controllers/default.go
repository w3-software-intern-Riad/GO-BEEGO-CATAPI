package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

type GetAllFavoriteController struct{
	beego.Controller
}

type GetBreedsControllerWeb struct{
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
}
func (f *GetAllFavoriteController) Get(){
	f.TplName="favorite.tpl"
}

func (gb *GetBreedsControllerWeb)Get(){
	gb.TplName="breeds.tpl"
}
