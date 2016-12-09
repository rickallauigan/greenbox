package main

import (
	//"github.com/ant0ine/go-json-rest/rest"
	"github.com/astaxie/beegae"
	"techaguru-orangepineapple/controllers"
	//"techaguru-orangepineapple/routers"
	"net/http"
)

func init() {
	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("../views/static/css"))))
	http.Handle("/static/js/", http.StripPrefix("/static/js/", http.FileServer(http.Dir("../views/static/js"))))
	http.Handle("/static/components/js/", http.StripPrefix("/static/components/js/", http.FileServer(http.Dir("../views/static/components/js"))))
	http.Handle("/static/fonts/", http.StripPrefix("/static/fonts/", http.FileServer(http.Dir("../views/static/fonts"))))
	http.Handle("/static/images/",http.StripPrefix("/static/images/",http.FileServer(http.Dir("../views/static/images"))))
	beegae.DirectoryIndex = true
	beegae.SetViewsPath("../views/")
	//beegae.SetStaticPath("/static", "../view/static/")
	//beegae.SetStaticPath("/static/images","../view/static/images")

	beegae.Router("/api/v1/create", &controllers.MainController{}, "post:Post")
	beegae.Router("/", &controllers.MainController{}, "get:Get")
	beegae.Run()

}
