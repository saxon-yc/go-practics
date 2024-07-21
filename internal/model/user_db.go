package model

import (
	"time"
)

type UserDB struct {
	UserID     uint32    `gorm:"column:user_id;primaryKey;comment:主键" json:"user_id"`
	TenantID   uint32    `gorm:"column:tenant_id;comment:租户ID" json:"tenant_id"`
	CreateTime time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`

	UserName string `gorm:"column:username;comment:用户名" json:"username"`
	Password string `gorm:"column:password;comment:密码" json:"password"`
	Token    string `gorm:"column:token;comment:token" json:"token"`
}
