package system

import (
	"github.com/weilaim/wm-admin/server/global"
	"github.com/weilaim/wm-admin/server/model/system"
)

type OperationRecordService struct{}

// 创建记录
func (o *OperationRecordService)CreateSysOperationRecord(sysOperationRecord system.SysOperationRecord)(err error){
	err = global.WM_DB.Create(&sysOperationRecord).Error
	return err
}