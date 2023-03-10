package global

import (
	"time"

	"gorm.io/gorm"
)

type WM_MODEL struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"` // 删除时间
}
