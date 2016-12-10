package routers

import (
	"github.com/astaxie/beegae"
	"greenbox/controllers"
	_ "net/http"
)

func init() {

	beegae.DirectoryIndex = true
	beegae.SetViewsPath("../views")
	beegae.Router("/", &controllers.MainController{})
	beegae.Router("/features", &controllers.FeatureController{})
	beegae.Router("/contact", &controllers.ContactController{})
	beegae.Router("/about", &controllers.AboutController{})
	beegae.Router("/post", &controllers.PostController{})
	beegae.Router("/news", &controllers.NewsController{})

}
