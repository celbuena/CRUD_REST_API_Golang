package apihandler

import (
	"fmt"
	"io/ioutil"
	"lawenconTest/utils"
	"log"
	"net/http"
)

func (c *ctrl) RegisterSchoolHandler(w http.ResponseWriter, r *http.Request) {
	re := utils.GetValueFromContext(r.Context())
	fmt.Println(re, "ïni context roles")
	role := re.Roles
	body, _ := ioutil.ReadAll(r.Body)
	c.admin.RegisterSchool(body, role).Default(w)
}

func (c *ctrl) UpdateSchoolHandler(w http.ResponseWriter, r *http.Request) {
	re := utils.GetValueFromContext(r.Context())
	fmt.Println(re, "ini context unt update admin")
	role := re.Roles
	body, _ := ioutil.ReadAll(r.Body)
	log.Println(string(body))
	c.admin.UpdateInfoSchool(r.FormValue("id"), body, role).Default(w)

}

func (c *ctrl) DeleteSchoolHandler(w http.ResponseWriter, r *http.Request) {
	re := utils.GetValueFromContext(r.Context())
	fmt.Println(re, "ini context unt delete admin")
	role := re.Roles
	body,_ := ioutil.ReadAll(r.Body)
	c.admin.DeleteSchool(r.FormValue("id"), body, role).Default(w)
}

func (c *ctrl) GetAllUserHandler(w http.ResponseWriter, r *http.Request) {
	value := utils.GetValueFromContext(r.Context()).UserId
	log.Println(string(value))
	c.admin.GetAllUser(value).Default(w)

}