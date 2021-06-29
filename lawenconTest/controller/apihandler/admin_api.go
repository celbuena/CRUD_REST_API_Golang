package apihandler

import (
	"fmt"
	"io/ioutil"
	"lawenconTest/utils"
	"log"
	"net/http"
)

func (c *ctrl) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	re := utils.GetValueFromContext(r.Context())
	fmt.Println(re, "ïni context roles")
	role := re.Roles
	body, _ := ioutil.ReadAll(r.Body)
	c.admin.Signup(body, role).Default(w)
}

func (c *ctrl) UpdateAdminHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	c.admin.UpdateInfoAdmin(r.FormValue("id"), body).Default(w)
}
func (c *ctrl) LoginHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	log.Println(string(body))
	c.admin.Login(body).Default(w)
}

func (c *ctrl) GetAllAdminHandler(w http.ResponseWriter, r *http.Request) {
	value := utils.GetValueFromContext(r.Context()).UserId
	log.Println(string(value))

	c.admin.GetAllAdmin(value).Default(w)
}

func (c *ctrl) GetAllTransactionHandler(w http.ResponseWriter, r *http.Request) {
	value := utils.GetValueFromContext(r.Context()).UserId
	log.Println(string(value))

	c.admin.GetAllTransaction(value).Default(w)
}

func (c *ctrl) DeleteAdminHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	c.admin.DeleteAdmin(body).Default(w)
}