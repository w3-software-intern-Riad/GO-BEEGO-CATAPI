package routers

import (
	"catProject/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/get-favorite", &controllers.GetAllFavoriteController{})
	beego.Router("/get-breeds", &controllers.GetBreedsControllerWeb{})
	beego.Router("/cat-images/:breed_id", &controllers.GetCatImagesController{}, "get:GetCatImages")
	beego.Router("/cat", &controllers.CatController{}, "get:GetCatImage")
	beego.Router("/vote", &controllers.VoteController{}, "post:PostVote")
	beego.Router("/favorite", &controllers.FavoriteController{}, "post:PostFavorite")
	beego.Router("/get-favorite-ctl", &controllers.GetFavController{}, "get:GetAllFav")
	beego.Router("/get-breeds-ctl", &controllers.GetBreedsController{}, "get:GetAllBreeds")
}
