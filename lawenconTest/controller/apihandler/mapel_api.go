package apihandler

import (
	"fmt"
	"io/ioutil"
	"lawenconTest/utils"
	"log"
	"net/http"
)

func (c *ctrl) RegisterMapelHandler(w http.ResponseWriter, r *http.Request) {
	re := utils.GetValueFromContext(r.Context())
	fmt.Println(re, "ïni context roles")
	role := re.Roles
	body, _ := ioutil.ReadAll(r.Body)
	c.admin.RegisterMapel(body, role).Default(w)
}

func (c *ctrl) GetAllMapelHandler(w http.ResponseWriter, r *http.Request) {
	value := utils.GetValueFromContext(r.Context()).UserId
	log.Println(string(value))
	c.admin.GetAllMapel(value).Default(w)
}

func (c *ctrl) UpdateMapelHandler(w http.ResponseWriter, r *http.Request) {
	re := utils.GetValueFromContext(r.Context())
	fmt.Println(re, "ini context unt update mapel")
	role := re.Roles
	body, _ := ioutil.ReadAll(r.Body)
	log.Println(string(body))
	c.admin.UpdateInfoMapel(r.FormValue("id"), body, role).Default(w)
}

func (c *ctrl) DeleteMapelHandler(w http.ResponseWriter, r *http.Request) {
	re := utils.GetValueFromContext(r.Context())
	fmt.Println(re, "ini context unt delete mapel")
	role := re.Roles
	body,_ := ioutil.ReadAll(r.Body)
	c.admin.DeleteMapel(body, role).Default(w)
}