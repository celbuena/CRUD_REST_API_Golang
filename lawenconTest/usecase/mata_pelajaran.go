package usecase

import (
	"fmt"
	"net/http"
	"lawenconTest/middleware"
	"lawenconTest/response"
	"lawenconTest/domain"
	"strconv"
)


func (s *usecase) RegisterMapel(body []byte, role string) response.Message {
	var mapel domain.Mata_Pelajaran
		mapel.FromJSON(body)

		if !middleware.CheckRole(role){
			return response.Errors(http.StatusForbidden, "method forbidden")
	}

	if s.validationMapel(&mapel).IsError() {
		return response.Errors(http.StatusNotFound, "request not complete")
	}
	s.admin.StoreMapel(&mapel)
	return response.Success("success")
}

func checkAttributeMapel(mapel *domain.Mata_Pelajaran) response.Message {
	if len(mapel.Mapel) < 0 {
		return response.Errors(http.StatusNotFound, "mapel is empty")
	} else if len(mapel.Icon) < 3 {
		return response.Errors(http.StatusNotFound, "icon is empty")
	}
	return response.Success(mapel)
}
func (s *usecase) validationMapel(mapel *domain.Mata_Pelajaran) response.Message {
	checkResult := checkAttributeMapel(mapel)
	if checkResult.IsError() {
		return response.Errors(http.StatusNotFound, "validation not complete")
	}
	queryResult, _ := s.admin.QueryByMapel(mapel.Mapel)
	if queryResult.Mapel != "" {
		return response.Errors(http.StatusNotFound, "email already registered")
	}
	return response.Success("Success")
}

func (s *usecase) GetAllMapel(idMapel int) response.Message {
	gt, err := s.admin.GetAllMapel(idMapel)
	if err != nil {
		return response.Errors(http.StatusInternalServerError, err.Error())
	}
	return response.Success(gt)
}

func (s *usecase)UpdateInfoMapel(id string, body []byte, roles string)response.Message{
	var mapel domain.UpdateInfoMapel
	mapel.FromJSON(body)

	if !middleware.CheckRole(roles){
		return response.Errors(http.StatusForbidden, "method forbidden")
	}

	checkAttribute := checkAttributeUpdateInfoMapel(&mapel)
	if checkAttribute.IsError(){
		return checkAttribute
	}

	err := s.admin.UpdateMapel(id, &mapel)
	newId, _ := strconv.Atoi(id)
	fmt.Println(newId)
	if err != nil {
		return response.Errors(http.StatusInternalServerError, err.Error())
	}
	return response.Success("Success")
}
func checkAttributeUpdateInfoMapel(mapel *domain.UpdateInfoMapel) response.Message{
	if mapel.IdTingkat == 0  || mapel.Mapel == ""{
		return response.Errors(http.StatusBadRequest, "request not complete")
	}
	return response.Success(mapel)
}

func (s *usecase) DeleteMapel(body []byte, roles string) response.Message {
	var mapel domain.UpdateSchoolInfo
	mapel.FromJson(body)

	if !middleware.CheckRole(roles){
		return response.Errors(http.StatusForbidden, "method forbidden")
	}
	if mapel.Id == 0 {
		return response.Errors(http.StatusNotFound,"id not found")
	}
	idSchool, err := s.admin.QueryByIdMatpel(mapel.Id)
	if err != nil {
		return response.Errors(http.StatusInternalServerError, err.Error())
	}
	fmt.Println(idSchool)
	err = s.admin.DeleteMapel(mapel.Id)
	if err != nil{
		return response.Errors(http.StatusInternalServerError, err.Error())
	}
	return response.Success("Success")
}