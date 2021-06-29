package usecase

import (
	"fmt"
	"lawenconTest/domain"
	"lawenconTest/middleware"
	"lawenconTest/response"
	"lawenconTest/utils"
	"net/http"
	"strconv"
	"strings"
)

func (s *usecase) Signup(body []byte, role string) response.Message {
	var admin domain.Admin
	admin.FromJSON(body)

	if !middleware.CheckRole(role){
		return response.Errors(http.StatusForbidden, "method forbidden")
	}
	if s.validation(&admin).IsError(){
		return s.validation(&admin)
	}
	return response.Success("Success")
}

func checkAttribute(admin *domain.Admin) response.Message {

	fmt.Println(strings.EqualFold(admin.Roles, "superadmin"))
	if len(admin.Name) < 3 {
		return response.Errors(http.StatusNotFound, "name is empty")
	} else if len(admin.Email) < 8 {
		return response.Errors(http.StatusNotFound, "Email is empty")
	} else if !strings.EqualFold(admin.Roles, "superadmin") && !strings.EqualFold(admin.Roles, "content") && !strings.EqualFold(admin.Roles, "finance") {
		return response.Errors(http.StatusNotFound, "roles is empty")
	} else if len(admin.Password) < 8 {
		return response.Errors(http.StatusNotFound, "password must contain 8 character")
	}

	return response.Success(admin)
}

func (s *usecase) validation(admin *domain.Admin) response.Message {

	checkResult := checkAttribute(admin)
	if checkResult.IsError() {
		return checkResult
	}
	queryResult, err1 := s.admin.QueryByEmail(admin.Email)

	fmt.Print(err1)
	if queryResult.Email != "" {
		return response.Errors(http.StatusBadRequest, "email already registered")
	}

	err := s.admin.StoreAdmin(admin)
	if err != nil{
		response.Errors(http.StatusInternalServerError, err.Error())
	}
	return response.Success("Success")
}

func (s *usecase) Login(body []byte) response.Message {
	us := new(domain.Admin)
	us.FromJSON(body)

	admin, err := s.admin.Find(us.Email, 0)
	if err == domain.ErrNotAdminNotFound {
		return response.NotFound()
	}
	if err != nil {
		return response.InternalError()
	}
	if !admin.IsMatchPassword(us.Password) {
		return response.Errors(http.StatusNotAcceptable, response.ErrLogin)
	}

	us.ID = admin.ID

	return s.setToken(admin.ID, admin.Roles)
}

func (s *usecase) setToken(adminid int, roles string ) response.Message {
	token, err := utils.NewClaim(adminid, "", roles).SetJWT()
	if err != nil {
		return response.InternalError()
	}

	return response.Success(token)
}

func (s *usecase) GetAllAdmin(adminid int) response.Message {
	gt, err := s.admin.GetAllAdmin(adminid)
	if err != nil{
		return response.Errors(http.StatusInternalServerError, err.Error())
	}
	return response.Success(gt)
}

func (s *usecase) GetAllTransaction(transactionid int) response.Message {
	gt, err := s.admin.GetAllTransaction(transactionid)
	if err != nil {
		return response.Errors(http.StatusInternalServerError, err.Error())
	}
	return response.Success(gt)
}


func (s *usecase) UpdateInfoAdmin(id string, body []byte) response.Message {
	var admin domain.UpdateAdminInfo
	admin.FromJson(body)

	if s.validationUpdate(&admin).IsError() {
		return response.Errors(http.StatusNotFound, "request not complete")
	}
	// update field
	s.admin.UpdateAdmin(id, &admin)
	newId, _ := strconv.Atoi(id)
	admin.SetRole(strings.ToLower(admin.Roles))
	fmt.Println(newId, admin.SetRole)

	return response.Success("success")
}

func checkAttributeUpdate(admin *domain.UpdateAdminInfo) response.Message {

	if len(admin.Name) < 3 {
		return response.Errors(http.StatusNotFound, "name must contain minimal 3 character")
	} else if len(admin.Email) < 8 {
		return response.Errors(http.StatusNotFound, "Email is empty")
	} else if !strings.EqualFold(admin.Roles, "superadmin") && !strings.EqualFold(admin.Roles, "content") && !strings.EqualFold(admin.Roles, "finance") {
		return response.Errors(http.StatusNotFound, "roles is empty")
	}
	return response.Success(admin)
}

func (s *usecase) validationUpdate(admin *domain.UpdateAdminInfo) response.Message {
	checkResult := checkAttributeUpdate(admin)
	if checkResult.IsError() {
		return response.Errors(http.StatusNotFound, "validation not complete")
	}
	return response.Success("Success")
}

func (u *usecase) DeleteAdmin(body []byte) response.Message {
	var admin domain.AdminInfo
	admin.FromJSON(body)
	if admin.ID == 0 {
		response.Errors(http.StatusNotFound,"id not found")
	}
	idAdmin, err := u.admin.QueryByIdAdmin(admin.ID)
	if err != nil{
		response.Errors(http.StatusInternalServerError, err.Error())
	}
	fmt.Println(idAdmin)
	err = u.admin.DeleteAdmin(admin.ID)
	return response.Success("Success")
}
