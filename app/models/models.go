package models

import (
	"echo_shop/database"
	"strconv"
	"time"
)

const (
	// TrueTinyint true
	TrueTinyint = 1
	// FalseTinyint false
	FalseTinyint = 0
)

// BaseModel model 基类
type BaseModel struct {
	ID uint `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	// MySQL的DATE/DATATIME类型可以对应Golang的time.Time
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	// 有 DeletedAt(类型需要是 *time.Time) 即支持 gorm 软删除
	DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index"`
}

// IDString 获取字符串形式的 id
func (m *BaseModel) IDString() string {
	return strconv.Itoa(int(m.ID))
}

// NewRecord model 是否已创建
func (m *BaseModel) NewRecord() bool {
	return m.ID <= 0
}

// Destroy 删除 model
func (m *BaseModel) Destroy() error {
	err := database.DBManager().Delete(&m).Error
	return err
}

// IsDeleted model 是否已被删除了
func (m *BaseModel) IsDeleted() bool {
	return m.DeletedAt != nil
}
