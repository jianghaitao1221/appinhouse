package controllers

import (
	. "appinhouse/server/constants"
	"appinhouse/server/models"
	"strings"

	"github.com/astaxie/beego"
)

type GetPlistController struct {
	BaseController
}

func (c *GetPlistController) GetPList() {
	dto := NewSuccessResponseDto()
	version := c.Ctx.Input.Param(":version")
	version = strings.Replace(version, Plist, "", -1)
	if version == "" {
		beego.Info("GetPList Params fail .version:", version)
		c.setError4Dto(ErrorParam, dto)
		return
	}
	var err error
	appParam := c.Ctx.Input.Param(":app")
	app, err := c.appExist(appParam)
	if err != nil {
		beego.Info("GetPList param app  error !app:", appParam)
		c.setError4Dto(err, dto)
		return
	}

	envParam := c.Ctx.Input.Param(":environment")
	env, err := GetEnvironment(envParam)
	if err != nil {
		beego.Info("GetPList param environment  error !environment:", envParam)
		c.setError4Dto(err, dto)
		return
	}
	plist, err := models.PlistDao.Get(env, app, version)
	if err != nil {
		beego.Info("GetPList from redis fail :", err.Error())
		c.setError4Dto(err, dto)
		return
	}
	c.Ctx.Output.Header("Content-Type", "application/x-plist")
	c.Ctx.Output.Body([]byte(plist))
}
