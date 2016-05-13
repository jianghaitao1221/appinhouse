// ModifyData
package controllers

import (
	. "appinhouse/server/constants"
	"appinhouse/server/models"

	"github.com/astaxie/beego"
)

type ModifyDataController struct {
	beego.Controller
}

func (c *ModifyDataController) ModifyAppData() {

	dto := NewSuccessResponseDto()
	apps, err := models.AppListDao.GetList(0, -1)
	if err != nil {
		dto.SetCode(ErrDB)
		return
	}
	if apps == nil {
		c.Data["json"] = dto
		return
	}
	if apps != nil {
		for _, app := range apps {
			info, err := models.AppDao.Get(app)
			if err != nil {
				dto.SetCode(ErrDB)
				return
			}
			if info == nil {
				info = &models.AppInfo{}
				info.App = app
				info.Description = app
				models.AppDao.Save(info)
			}
		}
	}

	c.Data["json"] = dto
	c.ServeJSON()
}
