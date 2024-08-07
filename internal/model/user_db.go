package model

import (
	"time"
)

// UserDB represents a user in the database.
// type UserDB struct {
// 	CreateTime time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`
// 	UpdateTime time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`
// 	UserName   string    `gorm:"column:username;comment:用户名" json:"username"`
// 	Password   string    `gorm:"column:password;comment:密码" json:"password"`
// 	Token      string    `gorm:"column:token;comment:token" json:"token"`
// 	Nickname   string    `gorm:"column:nickname;"`
// 	UserID     uint32    `gorm:"column:user_id;primaryKey;comment:主键" json:"user_id"`
// 	TenantID   uint32    `gorm:"column:tenant_id;comment:租户ID" json:"tenant_id"`
// }

// mysql 格式
/* type ClusterDB struct {
	ClusterID   uint   `gorm:"column:cluster_id;primaryKey;comment:集群ID" json:"cluster_id"`
	ClusterName string `gorm:"column:cluster_name" json:"cluster_name"`
}
type QKEPluginDB struct {
	PluginID      uint      `gorm:"column:plugin_id;primaryKey;comment:QKE支持的组件ID" json:"plugin_id"`
	PluginName    string    `gorm:"column:plugin_name;type:varbinary(40);comment:组件名称" json:"plugin_name"`
	PluginVersion string    `gorm:"column:plugin_version;type:varbinary(10);comment:组件版本" json:"plugin_version"`
	PluginType    string    `gorm:"column:plugin_type;type:varbinary(30);comment:组件名称" json:"plugin_type"`
	InstallType   string    `gorm:"column:install_type;type:varbinary(10);comment:升级方式 helm、qke_cli、kubectl、API 指定4中升级方式, 在初始化数据时默认指定一种, 目前固定为helm" json:"install_type"`
	PreVersions   string    `gorm:"column:pre_versions;comment:可从低版本升级到当前版本，插入数据以逗号分割" json:"pre_versions"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
	Deleted       bool      `gorm:"column:deleted;type:bool;comment:标记删除false|true" json:"deleted"`
}
type ClusterPluginDB struct {
	ID            uint      `gorm:"column:id;primaryKey;comment:集群使用的组件ID" json:"id"`
	ClusterID     uint      `gorm:"column:cluster_id;comment:关联集群ID" json:"cluster_id"`
	PluginName    string    `gorm:"column:plugin_name;type:varbinary(40);comment:组件名称" json:"plugin_name"`
	PluginVersion string    `gorm:"column:plugin_version;type:varbinary(10);comment:组件版本" json:"plugin_version"`
	PluginType    string    `gorm:"column:plugin_type;type:varbinary(30);comment:组件名称" json:"plugin_type"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
	Deleted       bool      `gorm:"column:deleted;type:bool;comment:标记删除false|true" json:"deleted"`
}
*/

// psql 格式
type ClusterDB struct {
	ClusterID   uint   `gorm:"column:cluster_id;primaryKey;comment:集群ID" json:"cluster_id"`
	ClusterName string `gorm:"column:cluster_name;type:varchar(255);comment:集群名称" json:"cluster_name"`
}

type QKEPluginDB struct {
	PluginID        uint      `gorm:"column:plugin_id;primaryKey;comment:QKE支持的组件ID" json:"plugin_id"`
	PluginName      string    `gorm:"column:plugin_name;type:varchar(40);not null;comment:组件名称" json:"plugin_name"`
	PluginVersion   string    `gorm:"column:plugin_version;type:varchar(10);not null;comment:组件版本" json:"plugin_version"`
	PluginType      string    `gorm:"column:plugin_type;type:varchar(30);not null;comment:组件类型" json:"plugin_type"`
	InstallType     string    `gorm:"column:install_type;type:varchar(10);comment:升级方式" json:"install_type"`
	PreVersions     string    `gorm:"column:pre_versions;type:text;not null;comment:可从低版本升级到当前版本，插入数据以逗号分割" json:"pre_versions"`
	K8sVersions     string    `gorm:"column:k8s_versions;type:text;not null;comment:支持的K8s版本" json:"k8s_versions"`
	Description     string    `gorm:"column:description;type:text;comment:插件描述信息" json:"description"`
	EnableInstalled bool      `gorm:"column:enable_installed;type:boolean;comment:能否在集群创建后在选择安装, true支持, false不支持" json:"enable_installed"`
	CreatedAt       time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;default:now()" json:"updated_at"`
	Deleted         bool      `gorm:"column:deleted;type:boolean;default:false;comment:标记删除false|true" json:"deleted"`
}

type ClusterPluginDB struct {
	ClusterID     string    `gorm:"column:cluster_id;not null;comment:关联集群ID" json:"cluster_id"`
	PluginID      uint      `gorm:"column:plugin_id;comment:QKE支持的组件ID" json:"plugin_id"`
	PluginName    string    `gorm:"column:plugin_name;type:varchar(40);not null;comment:组件名称" json:"plugin_name"`
	PluginVersion string    `gorm:"column:plugin_version;type:varchar(10);not null;comment:组件版本" json:"plugin_version"`
	PluginType    string    `gorm:"column:plugin_type;type:varchar(30);not null;comment:组件类型" json:"plugin_type"`
	CreatedAt     time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;default:now()" json:"updated_at"`
	Deleted       bool      `gorm:"column:deleted;type:boolean;default:false;comment:标记删除false|true" json:"deleted"`
}
