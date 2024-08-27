package dbsvc

import (
	"context"
	"errors"
	"go-practics/internal/model"
	"go-practics/internal/util"
	"sort"
	"strings"
	"time"

	"gorm.io/gorm"
)

type IComponent interface {
	AddComponents(components []model.ComptItem) (err error)
	FindComponents() (components []model.QkeDbComponent, err error)
	FindQkeComponents(comptName, comptType string) (output []model.ClusterComponentFields, err error)
}

// AddComponents Add components to the component_db
func (db *DbService) AddComponents(components []model.ComptItem) (err error) {
	var oldCompts []model.QkeDbComponent
	if oldCompts, err = db.FindComponents(); err != nil {
		return err
	}

	var writeCompts []model.QkeDbComponent
	var ctx context.Context

	return db.Gdb.Transaction(func(tx *gorm.DB) error {
		for _, v := range components {
			if len(oldCompts) > 0 && util.ContainsEqualCompt(v, oldCompts) {
				continue
			} else {
				writeCompts = append(writeCompts, model.QkeDbComponent{
					ComponentName:    v.Name,
					ComponentType:    v.Type,
					ComponentVersion: v.Version,
					InstallType:      "helm",
					PreVersions:      v.PreVersions,
					K8sVersions:      v.K8sVersions,
					Description:      v.Description,
					EnableInstalled:  v.EnableInstalled,
					CreatedAt:        time.Now(),
					UpdatedAt:        time.Now(),
					Deleted:          false,
				})
			}
		}

		if len(writeCompts) == 0 {
			return errors.New("can not add the same of components")
		}

		if err := tx.WithContext(ctx).Model(&model.QkeDbComponent{}).Save(&writeCompts).Error; err != nil {
			return err
		}
		return nil
	})
}

func (db *DbService) FindComponents() (components []model.QkeDbComponent, err error) {
	err = db.Gdb.Where("deleted =?", false).Find(&components).Error
	return components, err
}

func (db *DbService) FindQkeComponents(comptName, comptType string) (output []model.ClusterComponentFields, err error) {
	var qkeComponents []model.QkeDbComponent
	if err = db.Gdb.Where("deleted=false AND component_name LIKE ? AND component_type LIKE ?", "%"+comptName, "%"+comptType).Find(&qkeComponents).Error; err != nil {
		return output, err
	}

	for _, v1 := range qkeComponents {
		upgradeVersions := make([]string, 0, len(v1.PreVersions))
		for _, v2 := range qkeComponents {
			if v2.ComponentType == v1.ComponentType && strings.Contains(v2.PreVersions, v1.ComponentVersion) {
				upgradeVersions = append(upgradeVersions, v2.ComponentVersion)
			}
		}
		if len(upgradeVersions) > 1 {
			sort.SliceStable(upgradeVersions, func(i, j int) bool {
				if flag, err := util.CompareVersionStr(upgradeVersions[i], upgradeVersions[j]); flag || err != nil {
					return false
				} else {
					return true
				}

			})
		}
		K8sVersions := strings.Split(v1.K8sVersions, ",")
		if len(K8sVersions) > 1 {
			sort.SliceStable(K8sVersions, func(i, j int) bool {
				if flag, err := util.CompareVersionStr(K8sVersions[i], K8sVersions[j]); flag || err != nil {
					return false
				} else {
					return true
				}
			})
		}
		output = append(output, model.ClusterComponentFields{
			ComponentName:    v1.ComponentName,
			ComponentType:    v1.ComponentType,
			ComponentVersion: v1.ComponentVersion,
			Description:      v1.Description,
			UpgradeVersions:  upgradeVersions,
			K8sVersions:      K8sVersions,
		})
	}
	sort.SliceStable(output, func(i, j int) bool {
		if output[i].ComponentName == output[j].ComponentName && output[i].ComponentType == output[j].ComponentType {
			if flag, err := util.CompareVersionStr(output[i].ComponentVersion, output[j].ComponentVersion); flag || err != nil {
				return false
			} else {
				return true
			}
		} else {
			return false
		}
	})

	return output, nil
}
