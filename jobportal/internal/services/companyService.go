package services

import (
	"context"
	"job-portal-api/internal/models"
)

func (s *Conn) CreateCompany(ctx context.Context, newComp models.Company) (models.Company, error) {
	//store
	comp := models.Company{
		Name: newComp.Name,
		City: newComp.City,
		Jobs: newComp.Jobs,
	}
	err := s.db.Create(&comp).Error
	if err != nil {
		return models.Company{}, err
	}

	// Successfully created the record, return the user.
	return comp, nil

}
func (s *Conn) ViewCompanies(ctx context.Context) ([]models.Company, error) {
	var listComp []models.Company
	tx := s.db.WithContext(ctx)
	err := tx.Find(&listComp).Error
	if err != nil {
		return nil, err
	}

	// Successfully created the record, return the user.
	return listComp, nil

}
func (s *Conn) GetCompanyByID(ctx context.Context, companyId string) (models.Company, error) {
	var comp models.Company
	tx := s.db.WithContext(ctx).Where("ID = ?", companyId)
	err := tx.Find(&comp).Error
	if err != nil {
		return models.Company{}, err
	}

	return comp, nil

}
