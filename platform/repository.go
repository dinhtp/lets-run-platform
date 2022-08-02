package platform

import (
    "gorm.io/gorm"

    "github.com/dinhtp/lets-run-platform/model"
)

type Repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
    return &Repository{db: db}
}

func (r *Repository) Create(o *model.Platform) (*model.Platform, error) {
    if err := r.db.Create(o).Error; nil != err {
        return nil, err
    }

    return o, nil
}
