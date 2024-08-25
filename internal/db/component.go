package dbsvc

import (
	"context"
	"errors"
	"go-practics/internal/model"
	"go-practics/internal/util"
	"time"

	"gorm.io/gorm"
)

type IComponent interface {
	AddComponents(components []model.ComptItem) (err error)
	FindComponents() (components []model.QkeDbComponent, err error)
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
