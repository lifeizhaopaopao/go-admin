package sys_job

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"go-admin/app/admin/service"
	"go-admin/common/apis"
	"go-admin/common/dto"
	"go-admin/common/global"
	"go-admin/tools/app"
)

type SysJob struct {
	apis.Api
}

// RemoveJobForService 调用service实现
func (e *SysJob) RemoveJobForService(c *gin.Context) {
	log := e.GetLogger(c)
	db, err := e.GetOrm(c)
	if err != nil {
		log.Errorf("get db connection error, %s", err.Error())
		e.Error(c, http.StatusInternalServerError, err, "数据库连接获取失败")
		return
	}
	var v dto.GeneralDelDto
	err = c.BindUri(&v)
	if err != nil {
		log.Warnf("参数验证错误, error: %s", err)
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	s := service.SysJob{}
	s.Log = log
	s.Orm = db
	s.Cron = global.Cfg.GetCrontabKey(c.Request.Host)
	err = s.RemoveJob(&v)
	if err != nil {
		log.Errorf("RemoveJob error, %s", err.Error())
		e.Error(c, http.StatusInternalServerError, err, "")
		return
	}
	app.OK(c, nil, s.Msg)
}

// StartJobForService 启动job service实现
func (e *SysJob) StartJobForService(c *gin.Context) {
	log := e.GetLogger(c)
	db, err := e.GetOrm(c)
	if err != nil {
		log.Errorf("get db connection error, %s", err.Error())
		e.Error(c, http.StatusInternalServerError, err, "数据库连接获取失败")
		return
	}
	var v dto.GeneralGetDto
	err = c.BindUri(&v)
	if err != nil {
		log.Warnf("参数验证错误, error: %s", err)
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	s := service.SysJob{}
	s.Orm = db
	s.Log = log
	s.Cron = global.Cfg.GetCrontabKey(c.Request.Host)
	err = s.StartJob(&v)
	if err != nil {
		log.Errorf("GetCrontabKey error, %s", err.Error())
		e.Error(c, http.StatusInternalServerError, err, "")
		return
	}
	app.OK(c, nil, s.Msg)
}
