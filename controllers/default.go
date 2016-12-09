package controllers

import (
	"encoding/json"
	"github.com/astaxie/beegae"
	"io"
	//"golang.org/x/net/context"
	//"fmt"
	"google.golang.org/appengine/datastore"
	"techaguru-orangepineapple/models"
)

type MainController struct {
	beegae.Controller
}

type Value struct {
	TimeCreated       string `json:"time created"`
	Name              string `json:"name"`
	Age               int    `json:"age"`
	Municipality      string `json:"municipality"`
	Family            string `json:"family"`
	FarmLocation      string `json:"farm location"`
	VarietyOfRice     string `json:"variety of rice"`
	AppliedFertilizer string `json:"applied fertilizer"`
	LandPreparation   string `json:"land preparation"`
	Seedling          string `json:"seedling"`
	Planting          string `json:"planting"`
	Harvesting        string `json:"harvesting"`
}

func decodeProfile(r io.ReadCloser) (*models.Value, error) {
	defer r.Close()
	var profile models.Value
	err := json.NewDecoder(r).Decode(&profile)
	return &profile, err
}

func (this *MainController) Get() {
	profile := []models.Value{}
	ks, err := datastore.NewQuery("Value").GetAll(this.AppEngineCtx, &profile)
	if err != nil {
		this.Data["Profile"] = err
		return
	}
	for i := 0; i < len(profile); i++ {
		profile[i].Id = ks[i].IntID()
	}
	this.Data["Profile"] = &profile
	//this.ServeJson()
	this.TplNames = "index.html"
}

func (this *MainController) Post() {
	profile, err := decodeProfile(this.Ctx.Input.Request.Body)

	if err != nil {
		this.Data["json"] = err
		return
	}
	p, err := profile.Save(this.AppEngineCtx)

	if err != nil {
		//this.Data["json"] = err
	} else {
		this.Data["json"] = &p
	}
	this.ServeJson()
}
