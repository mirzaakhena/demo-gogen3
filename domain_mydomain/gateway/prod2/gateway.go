package prod2

import (
	"context"
	"your/path/project/domain_mydomain/model/entity"
	"your/path/project/shared/driver"
	"your/path/project/shared/infrastructure/config"
	"your/path/project/shared/infrastructure/logger"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// "github.com/ostafen/clover"
)

type gateway struct {
	log     logger.Logger
	appData driver.ApplicationData
	config  *config.Config
	db      *gorm.DB
}

// NewGateway ...
func NewGateway(log logger.Logger, appData driver.ApplicationData, config *config.Config) *gateway {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&entity.Person{})
	if err != nil {
		panic("cannot create schema")
	}

	return &gateway{
		log:     log,
		appData: appData,
		config:  config,
		db:      db,
	}
}

func (r *gateway) FindAllPerson(ctx context.Context, someID string) ([]*entity.Person, error) {
	r.log.Info(ctx, "called")

	var result []*entity.Person

	err := r.db.Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *gateway) SavePerson(ctx context.Context, obj *entity.Person) error {
	r.log.Info(ctx, "called")

	err := r.db.Create(obj).Error
	if err != nil {
		return err
	}

	return nil
}
