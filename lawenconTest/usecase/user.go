package usecase

import (
	"fmt"
	"lawenconTest/domain"
	"lawenconTest/middleware"
	"lawenconTest/response"
	"strconv"
)

import "net/http"

func (s *usecase) RegisterSchool(body []byte, role string) response.Message {
	var school domain.School
	school.FromJSON(body)

	if !middleware.CheckRole(role){
		return response.Errors(http.StatusForbidden, "method forbidden")
	}

	if s.validationSchool(&school).IsError() {
		return response.Errors(http.StatusNotFound, "request not complete")
	}
	school.SetTime()
	  s.admin.InsertSchool(&school)
	return response.Success("success")
}

func checkAttributeSchool(school *domain.School) response.Message {
	if len(school.SchoolName) < 3 {
		return response.Errors(http.StatusNotFound, "name is empty")
	} else if len(school.SchoolCode) < 8 {
		return response.Errors(http.StatusNotFound, "school code is empty")
	}

	return response.Success(school)
}
func (s *usecase) validationSchool(school *domain.School) response.Message {
	checkResult := checkAttributeSchool(school)
	if checkResult.IsError() {
		return response.Errors(http.StatusNotFound, "validation not complete")
	}
	queryResult, _ := s.admin.QueryByKodeSekolah(school.SchoolCode)
	if queryResult.SchoolCode != "" {
		return response.Errors(http.StatusNotFound, "email already registered")
	}
	// cek id admin
	checkAdmin, _ := s.admin.QueryByIdAdmin(school.CreatedBy)
	if (checkAdmin.ID == 0){
		return response.Errors(http.StatusNotFound, "id not found")
	}
	fmt.Println(checkAdmin)
	return response.Success("Success")
}

//update data school
func (s *usecase) UpdateInfoSchool(id string, body []byte, roles string) response.Message {
	var school domain.UpdateSchoolInfo
	school.FromJson(body)

	if !middleware.CheckRole(roles){
		return response.Errors(http.StatusForbidden, "method forbidden")
	}

	if s.validationUpdateSchool(&school).IsError() {
		return response.Errors(http.StatusNotFound, "request not complete")
	}
	// update field
	s.admin.UpdateSchool(id, &school)
	newId, _ := strconv.Atoi(id)
	fmt.Println(newId, school.SchoolCode)
	return response.Success("success")

}

func checkAttributeUpdateSchool(school *domain.UpdateSchoolInfo) response.Message {

	if len(school.SchoolName) < 3 {
		return response.Errors(http.StatusNotFound, "name must contain minimal 3 character")
	} else if len(school.SchoolCode) < 8 {
		return response.Errors(http.StatusNotFound, "kode_sekolah is empty")
	}
	return response.Success(school)
}

func (s *usecase) validationUpdateSchool(school *domain.UpdateSchoolInfo) response.Message {
	checkResult := checkAttributeUpdateSchool(school)
	if checkResult.IsError() {
		return response.Errors(http.StatusNotFound, "validation not complete")
	}

	return response.Success("Success")
}

func (s *usecase) DeleteSchool(id string, body []byte, roles string) response.Message {
	var school domain.UpdateSchoolInfo
	school.FromJson(body)

	if !middleware.CheckRole(roles){
		return response.Errors(http.StatusForbidden, "method forbidden")
	}
	if school.Id == 0 {
		return response.Errors(http.StatusNotFound,"id not found")
	}
	idSchool, err := s.admin.QueryByIdSchool(school.Id)
	if err != nil {
		return response.Errors(http.StatusInternalServerError, err.Error())
	}
	fmt.Println(idSchool)
	err = s.admin.DeleteSchool(school.Id)
	if err != nil{
		return response.Errors(http.StatusInternalServerError, err.Error())
	}
	return response.Success("Success")
}

func (s *usecase) GetAllUser(userid int) response.Message {
	gt, err := s.admin.GetAllUser(userid)
	if err != nil {
		return response.Errors(http.StatusInternalServerError, err.Error())
	}
	return response.Success(gt)
}