package model

import "time"

type Components struct {
	Components []ComptItem `yaml:"components"`
}

type ComptItem struct {
	Name            string `yaml:"name"`
	Type            string `yaml:"type"`
	Version         string `yaml:"version"`
	PreVersions     string `yaml:"preVersions"`
	K8sVersions     string `yaml:"k8sVersions"`
	Description     string `yaml:"description"`
	EnableInstalled bool   `yaml:"enableInstalled"`
}

type QkeDbComponent struct {
	ComponentId      uint      `gorm:"column:component_id;primaryKey;comment:QKE支持的组件ID" json:"component_id"`
	ComponentName    string    `gorm:"column:component_name;type:varchar(40);not null;comment:组件名称" json:"component_name"`
	ComponentVersion string    `gorm:"column:component_version;type:varchar(10);not null;comment:组件版本" json:"component_version"`
	ComponentType    string    `gorm:"column:component_type;type:varchar(30);not null;comment:组件类型" json:"component_type"`
	InstallType      string    `gorm:"column:install_type;type:varchar(10);comment:升级方式" json:"install_type"`
	PreVersions      string    `gorm:"column:pre_versions;type:text;not null;comment:可从低版本升级到当前版本，插入数据以逗号分割" json:"pre_versions"`
	K8sVersions      string    `gorm:"column:k8s_versions;type:text;not null;comment:支持的K8s版本" json:"k8s_versions"`
	Description      string    `gorm:"column:description;type:text;comment:插件描述信息" json:"description"`
	EnableInstalled  bool      `gorm:"column:enable_installed;type:boolean;comment:能否在集群创建后再选择安装, true支持, false不支持" json:"enable_installed"`
	CreatedAt        time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at;default:now()" json:"updated_at"`
	Deleted          bool      `gorm:"column:deleted;type:boolean;default:false" json:"deleted"`
}
