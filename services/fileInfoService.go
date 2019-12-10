package services

import (
	"smyappTwo/dao"
	"smyappTwo/pojo"
)

type FileInfoService interface {
	Create(f *pojo.FileInfo)error
}

type fileInfoService struct {
	dao dao.FileInfoDao
}

func NewFileInfoService(d dao.FileInfoDao)FileInfoService{
	return &fileInfoService{d}
}


func (s *fileInfoService)Create(f *pojo.FileInfo)error{
	return s.dao.Create(f)
}