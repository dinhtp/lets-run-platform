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

func (r *Repository) FindOne(id int) (*model.Platform, error) {
    var result model.Platform

    query := r.db.Model(&model.Platform{}).Where("id = ?", id).First(&result)

    if err := query.Error; nil != err {
        return nil, err
    }

    return &result, nil
}

func (r *Repository) Create(o *model.Platform) (*model.Platform, error) {
    if err := r.db.Create(o).Error; nil != err {
        return nil, err
    }

    return o, nil
}
