package usecase

import (
	"fmt"
	"net/http"
	"lawenconTest/domain"
	"lawenconTest/middleware"
	"lawenconTest/response"
	"strconv"
)


func (s *usecase)CreateMateriHeader(body []byte) response.Message{
	var materi_header domain.Materi_Header
	materi_header.FromJSON(body)

	if s.validationMateriHeader(&materi_header).IsError(){
		return s.validationMateriHeader(&materi_header)
	}
	return response.Success("Success")
}

func checkAttributeMateriHeader(materi_header *domain.Materi_Header) response.Message{
	if materi_header.IdMapel == 0 {
		return response.Errors(http.StatusNotFound,"id_mapel is empty")
	}else if materi_header.Chapter == 0 {
		return response.Errors(http.StatusNotFound,"chapter is empty")
	}else if len(materi_header.Title) < 5{
		return response.Errors(http.StatusNotFound,"title is too short")
	}else if len (materi_header.Label) < 20 {
		return response.Errors(http.StatusNotFound,"label is too short")
	}else if len (materi_header.Materi) < 30 {
		return response.Errors(http.StatusNotFound,"materi is too short")
	}else if materi_header.Detail == ""{
		return response.Errors(http.StatusNotFound,"detail is too short")
	}else if materi_header.StatusMateri == 0 {
		return response.Errors(http.StatusNotFound,"status_materi is empty")
	}

	return response.Success(materi_header)
}

func (s *usecase) validationMateriHeader(materi_header *domain.Materi_Header)response.Message{
	checkResult := checkAttributeMateriHeader(materi_header)
	if checkResult.IsError(){
		return checkResult
	}

	err := s.admin.InsertMateri(materi_header)
	if err != nil {
		return response.Errors(http.StatusInternalServerError,err.Error())
	}
	return response.Success("Success")
}


func (s *usecase)GetMateriHeader(materi_headerId int) response.Message{
	gt, err := s.admin.GetAllMateri(materi_headerId)
	if err != nil{
		return response.Errors(http.StatusInternalServerError, err.Error())
	}

	return response.Success(gt)
}

func (s *usecase) DeleteMateri(body []byte, roles string) response.Message {
	var materi domain.UpdateMateriHeader
	materi.FromJSON(body)

	if !middleware.CheckRole(roles){
		return response.Errors(http.StatusForbidden, "method forbidden")
	}
	if materi.Id == 0 {
		return response.Errors(http.StatusNotFound,"id not found")
	}
	idMateri, err := s.admin.GetIdMaHeader(materi.Id)
	if err != nil {
		return response.Errors(http.StatusInternalServerError, err.Error())
	}
	fmt.Println(idMateri)
	err = s.admin.DeleteMateriHeader(materi.Id)
	if err != nil{
		return response.Errors(http.StatusInternalServerError, err.Error())
	}
	return response.Success("Success")
}

func (s *usecase) UpdateInfoMateriHeader(id string, body []byte) response.Message{
	var materiHeader domain.UpdateMateriHeader
	materiHeader.FromJSON(body)

	if s.validationUpdateMateriHeader(&materiHeader).IsError()	{
		return response.Errors(http.StatusNotFound, "request not complete")
	}
	//update materi header
	s.admin.UpdateMateriHeader(id, &materiHeader)
	newId, _ := strconv.Atoi(id)
	fmt.Println(newId)
	return response.Success("Success")
}

func (s *usecase)validationUpdateMateriHeader(materi_header *domain.UpdateMateriHeader)response.Message{
	checkResult := checkAttributeUpdateMateriHeader(materi_header)
	if checkResult.IsError(){
		return checkResult
	}
	/* dataMapel ,err := s.admin.QueryByIdMatpel(materi_header.Id)
	if err != nil{
		return response.Errors(http.StatusInternalServerError,"id not found")
	}
	fmt.Println(dataMapel)*/
	return response.Success("Success")
}

func checkAttributeUpdateMateriHeader(materi_header *domain.UpdateMateriHeader) response.Message{
	if materi_header.Chapter == 0 {
		return response.Errors(http.StatusNotFound, "chapter is empty")
	}else if len(materi_header.Title) < 5 {
		return response.Errors(http.StatusNotFound,"title is too short")
	}else if len (materi_header.Label) < 20 {
		return response.Errors(http.StatusNotFound,"label is too short")
	}else if len (materi_header.Materi) < 30 {
		return response.Errors(http.StatusNotFound,"materi is too short")
	}else if materi_header.Detail == ""{
		return response.Errors(http.StatusNotFound,"detail is too short")
	}else if materi_header.StatusMateri == "" {
		return response.Errors(http.StatusNotFound,"status_materi is empty")
	}
	return response.Success(materi_header)
}
