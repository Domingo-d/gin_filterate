package service

import (
	"filterate/global"
	"filterate/initialize"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type (
	FileService struct{}
)

func (fileService *FileService) UpdateFile(fh multipart.File) uint {
	file, err := os.Create(global.FilterateName)
	if err != nil {
		return http.StatusInternalServerError
	}

	defer file.Close()

	if _, err := io.Copy(file, fh); err != nil {
		return http.StatusInternalServerError
	}

	return http.StatusOK
}

func (fileService *FileService) ReLoad() uint {
	tmpCorasick := initialize.NewAhoCorasick()
	if nil != tmpCorasick.ReadPattern(global.FilterateName) {
		return http.StatusInternalServerError
	}

	tmpCorasick.Build()

	global.AhoCorasick = tmpCorasick

	return http.StatusOK
}

//func (fileService *FileService) AddPattern(pattern string) uint {
//	res := ServiceGroupApp.FilterateService.Filter(&request.FilterateReq{Str: pattern})
//	if res.Str != pattern {
//		return http.StatusBadRequest
//	} else {
//
//	}
//
//	return http.StatusOK
//}
