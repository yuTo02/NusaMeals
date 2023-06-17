package repository

import (
	"github.com/jinzhu/gorm"
	"reglog/internal/model"
)

type TableRepository interface {
	CreateTable(table *model.Table) (*model.Table, error)
	UpdateTable(table *model.Table) (*model.Table, error)
	DeleteTable(tableID uint) error
	GetTableByID(tableID uint) (*model.Table, error)
	GetAllTables() ([]model.Table, error)
}

type tableRepository struct {
	db *gorm.DB
}

func NewTableRepository(db *gorm.DB) TableRepository {
	return &tableRepository{
		db: db,
	}
}

func (r *tableRepository) CreateTable(table *model.Table) (*model.Table, error) {
	result := r.db.Create(table)
	if result.Error != nil {
		return nil, result.Error
	}
	return table, nil
}

func (r *tableRepository) UpdateTable(table *model.Table) (*model.Table, error) {
	result := r.db.Save(table)
	if result.Error != nil {
		return nil, result.Error
	}
	return table, nil
}

func (r *tableRepository) DeleteTable(tableID uint) error {
	return r.db.Delete(&model.Table{}, tableID).Error
}

func (r *tableRepository) GetTableByID(tableID uint) (*model.Table, error) {
	var table model.Table
	err := r.db.First(&table, tableID).Error
	if err != nil {
		return nil, err
	}
	return &table, nil
}

func (r *tableRepository) GetAllTables() ([]model.Table, error) {
	var tables []model.Table
	err := r.db.Find(&tables).Error
	if err != nil {
		return nil, err
	}
	return tables, nil
}
