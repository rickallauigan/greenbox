package controllers

import (
	//"encoding/json"
	"github.com/astaxie/beegae"
	//_ "greenbox/views"
	//"io"
	//"golang.org/x/net/context"
	//"fmt"
	//"google.golang.org/appengine/datastore"
)

type MainController struct {
	beegae.Controller
}

type FeatureController MainController
type AboutController MainController
type ContactController MainController
type NewsController MainController
type PostController MainController

func (this *MainController) Get() {
	this.TplNames = "index.html"
}

func (this *FeatureController) Get() {
	this.TplNames = "features.html"
}

func (this *AboutController) Get() {
	this.TplNames = "about.html"
}

func (this *ContactController) Get() {
	this.TplNames = "contact.html"
}

func (this *NewsController) Get() {
	this.TplNames = "news.html"
}

func (this *PostController) Get() {
	this.TplNames = "post.html"
}
