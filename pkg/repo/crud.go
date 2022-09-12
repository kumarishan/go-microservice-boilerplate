package repo

import (
	"context"

	"github.com/kumarishan/errors"
	"gorm.io/gorm"
)

type CrudRepository[E any, ID any, M GormModel[E, ID]] interface {
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
func interfaceCheck[E any, ID any, M GormModel[E, ID]]() CrudRepository[E, ID, M] {
	return &CrudRepositoryImpl[E, ID, M]{
		db: nil,
	}
}

func NewCrudRepositoryImpl[E any, ID any, M GormModel[E, ID]](db *gorm.DB) *CrudRepositoryImpl[E, ID, M] {
	return &CrudRepositoryImpl[E, ID, M]{
		db: db,
	}
}

type CrudRepositoryImpl[E any, ID any, M GormModel[E, ID]] struct {
	db *gorm.DB
}

func (r *CrudRepositoryImpl[E, ID, M]) Save(ctx context.Context, entity *E) (*E, error) {
	var model M
	model = model.FromEntity(entity).(M)

	err := r.db.WithContext(ctx).Create(model).Error
	if err != nil {
		return nil, errors.Return(err, nil, "")
	}

	return model.ToEntity(), nil
}

func (r *CrudRepositoryImpl[E, ID, M]) FindById(ctx context.Context, id ID) (*E, error) {
	var model M
	err := r.db.WithContext(ctx).First(&model, id).Error
	if err != nil {
		return nil, errors.Return(err, nil, "")
	}
	return model.ToEntity(), nil
}

func (r *CrudRepositoryImpl[E, ID, M]) FindAll(ctx context.Context, limit int, offset int) ([]*E, error) {
	return nil, nil
}

func (r *CrudRepositoryImpl[E, ID, M]) Scopes(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) *gorm.DB {
	return r.db.Scopes(scopes...)
}
