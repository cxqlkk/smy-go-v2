package controllers

import (
	"github.com/kataras/iris"
	"smyappTwo/services"
)

type FileInfoController struct {
	Service services.FileInfoService
	Ctx     iris.Context
}

func (c *FileInfoController) Add() interface{} {
	//f, _, _ := c.Ctx.FormFile("icon")
//todo

	return nil
}


