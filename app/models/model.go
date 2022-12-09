package models

import (
	"github.com/spf13/cast"
	"time"
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id,omitempty"`
}

// TimestampsField 时间戳
type TimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
}

// GetStringID 获取 ID 的字符串格式
func (m BaseModel) GetStringID() string {
	return cast.ToString(m.ID)
}
