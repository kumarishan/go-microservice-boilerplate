package repo

import (
	"context"

	"github.com/kumarishan/errors"
	"gorm.io/gorm"
)

type CrudRepository[M GormModel[E, ID], E any, ID any] interface {
	// Count(ctx context.Context) (uint64, error)
	// Delete(ctx context.Context, entity *E) error
	// DeleteAll(ctx context.Context) error
	// DeleteEntities(ctx context.Context, entities []*E) error
	// DeleteAllByIds(ctx context.Context, ids []ID) error
	// DeleteById(ctx context.Context, id ID) error
	// ExistsById(ctx context.Context, id ID) (bool, error)

	FindAll(ctx context.Context, limit int, offset int) ([]*E, error)
	// FindAllByIds(ctx context.Context, ids []ID) ([]*E, error)
	FindById(ctx context.Context, id ID) (*E, error)

	Save(ctx context.Context, entity *E) (*E, error)
	// SaveAll(ctx context.Context, entities []*E) ([]*E, error)

	Scopes(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) *gorm.DB
}

//  for checking if interface is implemented, not for usage
func interfaceCheck[M GormModel[E, ID], E any, ID any]() CrudRepository[M, E, ID] {
	return &CrudRepositoryImpl[M, E, ID]{
		db: nil,
	}
}

func NewCrudRepositoryImpl[M GormModel[E, ID], E any, ID any](db *gorm.DB) *CrudRepositoryImpl[M, E, ID] {
	return &CrudRepositoryImpl[M, E, ID]{
		db: db,
	}
}

type CrudRepositoryImpl[M GormModel[E, ID], E any, ID any] struct {
	db *gorm.DB
}

func (r *CrudRepositoryImpl[M, E, ID]) Save(ctx context.Context, entity *E) (*E, error) {
	var model M
	model = model.FromEntity(entity).(M)

	err := r.db.WithContext(ctx).Create(model).Error
	if err != nil {
		return nil, errors.Return(err, nil, "")
	}

	return model.ToEntity(), nil
}

func (r *CrudRepositoryImpl[M, E, ID]) FindById(ctx context.Context, id ID) (*E, error) {
	var model M
	err := r.db.WithContext(ctx).First(&model, id).Error
	if err != nil {
		return nil, errors.Return(err, nil, "")
	}
	return model.ToEntity(), nil
}

func (r *CrudRepositoryImpl[M, E, ID]) FindAll(ctx context.Context, limit int, offset int) ([]*E, error) {
	return nil, nil
}

func (r *CrudRepositoryImpl[M, E, ID]) Scopes(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) *gorm.DB {
	return r.db.Scopes(scopes...)
}
