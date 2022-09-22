package repo

import (
	"context"

	"github.com/kumarishan/errors"
	"gorm.io/gorm"
)

type PtrE[E any] interface {
	*E
}

type CrudRepository[E any, ID any] interface {
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
func interfaceCheck[E any, ID any]() CrudRepository[E, ID] {
	return &CrudRepositoryImpl[E, ID]{
		db: nil,
	}
}

func NewCrudRepositoryImpl[E any, ID any](db *gorm.DB) *CrudRepositoryImpl[E, ID] {
	return &CrudRepositoryImpl[E, ID]{
		db: db,
	}
}

type CrudRepositoryImpl[E any, ID any] struct {
	db *gorm.DB
}

func (r *CrudRepositoryImpl[E, ID]) Save(ctx context.Context, entity *E) (*E, error) {
	err := r.db.WithContext(ctx).Create(entity).Error
	if err != nil {
		return nil, errors.Return(err, nil, "")
	}

	return entity, nil
}

func (r *CrudRepositoryImpl[E, ID]) FindById(ctx context.Context, id ID) (*E, error) {
	var entity = new(E)
	err := r.db.WithContext(ctx).First(entity, "id = ?", id).Error
	if err != nil {
		return nil, errors.Return(err, nil, "")
	}
	return entity, nil
}

func (r *CrudRepositoryImpl[E, ID]) FindAll(ctx context.Context, limit int, offset int) ([]*E, error) {
	var entities []*E
	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&entities).Error
	if err != nil {
		return nil, errors.Return(err, nil, "")
	}
	return entities, nil
}

func (r *CrudRepositoryImpl[E, ID]) Scopes(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) *gorm.DB {
	return r.db.Scopes(scopes...)
}
