package apihandler

import (
	"fmt"
	"io/ioutil"
	"lawenconTest/utils"
	"log"
	"net/http"
)

func (c *ctrl)CreateMateriHeaderHandler (w http.ResponseWriter, r *http.Request){
	body, _ := ioutil.ReadAll(r.Body)
	log.Println(string(body))
	c.admin.CreateMateriHeader(body).Default(w)
}


func (c *ctrl)GetAllMateriHeaderHandler(w http.ResponseWriter, r *http.Request){
	value := utils.GetValueFromContext(r.Context()).UserId
	log.Println(string(value))
	c.admin.GetMateriHeader(value).Default(w)
}

func (c *ctrl) DeleteMateriHandler(w http.ResponseWriter, r *http.Request) {
	re := utils.GetValueFromContext(r.Context())
	fmt.Println(re, "ini context unt delete materi")
	role := re.Roles
	body,_ := ioutil.ReadAll(r.Body)
	c.admin.DeleteMateri(body, role).Default(w)
}

func (c *ctrl)UpdateMateriHeaderHandler(w http.ResponseWriter, r *http.Request){
	body, _ := ioutil.ReadAll(r.Body)
	c.admin.UpdateInfoMateriHeader(r.FormValue("id"), body).Default(w)
}
