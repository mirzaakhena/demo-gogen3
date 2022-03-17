package repository

import (
	"context"
	"your/path/project/domain_mydomain/model/entity"
)

type SavePersonRepo interface {
	SavePerson(ctx context.Context, obj *entity.Person) error
}

type FindAllPersonRepo interface {
	FindAllPerson(ctx context.Context, someID string) ([]*entity.Person, error)
}
