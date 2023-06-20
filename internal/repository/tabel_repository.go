package repository

import (
	"reglog/internal/dto/request"
	"reglog/internal/model"

	"github.com/jinzhu/gorm"
)

type TableRepository interface {
	GetTable() ([]model.Table, error)
	GetTableByID(id uint) (model.Table, error)
	GetTableByStatus() ([]model.Table, error)
	CreateTable(req *request.Table) error
	UpdateTable(table *model.Table) error
	DeleteTableByID(id uint) error
}

type tableRepository struct {
	db *gorm.DB
}

func NewTableRepository(db *gorm.DB) *tableRepository {
	return &tableRepository{db: db}
}

func (r *tableRepository) GetTables() ([]model.Table, error) {
	var tables []model.Table
	err := r.db.Preload("Position").Find(&tables).Error
	if err != nil {
		return []model.Table{}, err
	}
	return tables, nil
}

func (r *tableRepository) GetTableByID(id uint) (model.Table, error) {
	var table model.Table
	err := r.db.Preload("Table").Where("id = ?", id).First(&table).Error
	if err != nil {
		return model.Table{}, err
	}
	return table, nil
}

func (r *tableRepository) GetTablesByStatus(name string) ([]model.Table, error) {
	var tables []model.Table
	err := r.db.Preload("Position").Where("status = ?", "availabe").Find(&tables).Error
	if err != nil {
		return nil, err
	}
	return tables, nil
}

func (r *tableRepository) CreateTable(req *request.Table) error {
	table := model.Table{
		Number_Table: req.Number_Table,
		Seat:         req.Seat,
		PositionID:   req.PositionID,
		Status:       req.Status,
		Location:     req.Location,
		Images:       req.Images,
	}
	err := r.db.Create(&table).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *tableRepository) UpdateTable(table *model.Table) error {
	err := r.db.Save(&table).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *tableRepository) DeleteTableByID(id uint) error {
	err := r.db.Where("id = ?", id).Delete(&model.Table{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *tableRepository) UpdateAvailbleStatus(id uint) error {
	err := r.db.Where("id = ?", id).Delete(&model.Table{}).Error
	if err != nil {
		return err
	}
	return nil
}
