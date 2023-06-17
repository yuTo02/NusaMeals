package usecase

import (
	"reglog/internal/dtos"
	"reglog/internal/model"
	"reglog/internal/repository"
)

type TableUseCase interface {
	CreateTable(tableDTO *dtos.TableDTO) (*dtos.TableDTO, error)
	UpdateTable(tableID uint, tableDTO *dtos.TableDTO) (*dtos.TableDTO, error)
	DeleteTable(tableID uint) error
	GetTableByID(tableID uint) (*dtos.TableDTO, error)
	GetAllTables() ([]dtos.TableDTO, error)
}

type tableUseCase struct {
	tableRepo repository.TableRepository
}

func NewTableUseCase(tableRepo repository.TableRepository) TableUseCase {
	return &tableUseCase{
		tableRepo: tableRepo,
	}
}

func (t *tableUseCase) CreateTable(tableDTO *dtos.TableDTO) (*dtos.TableDTO, error) {
	table := &model.Table{
		Name: tableDTO.Name,
		// Set other fields accordingly
	}

	err := t.tableRepo.CreateTable(table)
	if err != nil {
		return nil, err
	}

	return tableDTO, nil
}

func (t *tableUseCase) UpdateTable(tableID uint, tableDTO *dtos.TableDTO) (*dtos.TableDTO, error) {
	table, err := t.tableRepo.GetTableByID(tableID)
	if err != nil {
		return nil, err
	}

	// Update the table entity with the new data
	table.Name = tableDTO.Name
	// Update other fields accordingly

	err = t.tableRepo.UpdateTable(table)
	if err != nil {
		return nil, err
	}

	updatedTable, err := t.tableRepo.GetTableByID(tableID)
	if err != nil {
		return nil, err
	}

	updatedTableDTO := &dtos.TableDTO{
		ID:   updatedTable.ID,
		Name: updatedTable.Name,
		// Set other fields accordingly
	}

	return updatedTableDTO, nil
}

func (t *tableUseCase) DeleteTable(tableID uint) error {
	return t.tableRepo.DeleteTable(tableID)
}

func (t *tableUseCase) GetTableByID(tableID uint) (*dtos.TableDTO, error) {
	table, err := t.tableRepo.GetTableByID(tableID)
	if err != nil {
		return nil, err
	}

	tableDTO := &dtos.TableDTO{
		ID:   table.ID,
		Name: table.Name,
		// Set other fields accordingly
	}

	return tableDTO, nil
}

func (t *tableUseCase) GetAllTables() ([]dtos.TableDTO, error) {
	tables, err := t.tableRepo.GetAllTables()
	if err != nil {
		return nil, err
	}

	var tableDTOs []dtos.TableDTO
	for _, table := range tables {
		tableDTO := dtos.TableDTO{
			ID:   table.ID,
			Name: table.Name,
			// Set other fields accordingly
		}
		tableDTOs = append(tableDTOs, tableDTO)
	}

	return tableDTOs, nil
}
