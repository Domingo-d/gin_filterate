package service

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"server/global"
)

type (
	FileService struct{}
)

func (fileService *FileService) UpdateFile(fh multipart.File) uint {
	file, err := os.Create(global.Config.FilterateName)
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
	tmpCorasick := global.AhoCorasick.NewAhoCorasick()
	if io.EOF != tmpCorasick.ReadPattern(global.Config.FilterateName) {
		return http.StatusInternalServerError
	}

	tmpCorasick.Build()

	// 指针赋值是原子操作
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
