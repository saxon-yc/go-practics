package dbsvc

import (
	"fmt"
	"go-practics/internal/model"
	"time"

	"gorm.io/gorm"
)

type IComponent interface {
	AddComponents(components []model.Items) error
}

// AddComponents Add components to the component_db
func (db *DbService) AddComponents(components []model.Items) error {
	fmt.Printf("AddComponents components: %v\n", components)
	var data model.QkeDbComponent
	return db.Gdb.Transaction(func(tx *gorm.DB) error {
		for _, v := range components {
			err := tx.Model(&data).Create(&model.QkeDbComponent{
				ComponentName:    v.Name,
				ComponentType:    v.Type,
				ComponentVersion: v.Version,
				PreVersions:      v.PreVersions,
				K8sVersions:      v.K8sVersions,
				Description:      v.Description,
				EnableInstalled:  v.EnableInstalled,
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
				Deleted:          false,
			}).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}
