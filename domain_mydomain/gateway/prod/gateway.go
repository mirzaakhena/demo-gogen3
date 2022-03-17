package prod

import (
	"context"
	"your/path/project/domain_mydomain/model/entity"
	"your/path/project/shared/driver"
	"your/path/project/shared/infrastructure/config"
	"your/path/project/shared/infrastructure/logger"

	"github.com/ostafen/clover"
)

type gateway struct {
	log     logger.Logger
	appData driver.ApplicationData
	config  *config.Config
	db      *clover.DB
}

// NewGateway ...
func NewGateway(log logger.Logger, appData driver.ApplicationData, config *config.Config) *gateway {

	db, err := clover.Open("database2")
	if err != nil {
		panic(err.Error())
	}

	exist, err := db.HasCollection("person")
	if err != nil {
		panic(err.Error())
	}

	if !exist {
		err = db.CreateCollection("person")
		if err != nil {
			panic(err.Error())
		}
	}

	return &gateway{
		log:     log,
		appData: appData,
		config:  config,
		db:      db,
	}
}

func (r *gateway) SavePerson(ctx context.Context, obj *entity.Person) error {
	r.log.Info(ctx, "called")

	doc := clover.NewDocument()
	doc.Set("person", obj)

	docId, err := r.db.InsertOne("person", doc)
	if err != nil {
		return err
	}

	r.log.Info(ctx, "ID %v", docId)

	return nil
}

func (r *gateway) FindAllPerson(ctx context.Context, someID string) ([]*entity.Person, error) {
	r.log.Info(ctx, "called")

	query := r.db.Query("person")

	objs, err := query.FindAll()
	if err != nil {
		return nil, err
	}

	results := make([]*entity.Person, 0)

	for _, obj := range objs {

		document := struct{ Person entity.Person }{}

		err := obj.Unmarshal(&document)
		if err != nil {
			return nil, err
		}

		results = append(results, &document.Person)

	}

	return results, nil
}
