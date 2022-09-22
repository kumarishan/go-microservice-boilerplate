package repo

import (
	"time"

	"gorm.io/gorm"
)

type Model[ID any] struct {
	ID        ID             `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// todo - dirty checkers
// type GormModel[E any, M any] interface {
// 	MapToEntity() *E
// 	MapFromEntity(entity *E) *M
// }
