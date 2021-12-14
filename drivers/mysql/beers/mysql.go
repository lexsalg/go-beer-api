package beers

import (
	"beerapi/bussiness/beers"

	"gorm.io/gorm"
)

type repositoryBeers struct {
	DB *gorm.DB
}

func NewRepositoryMySQL(db *gorm.DB) beers.Repository {
	return &repositoryBeers{
		DB: db,
	}
}

func (repository repositoryBeers) Create(beer *beers.Domain) (*beers.Domain, error) {
	record := fromDomain(*beer)
	if err := repository.DB.Create(&record).Error; err != nil {
		return &beers.Domain{}, err
	}
	result := toDomain(record)
	return &result, nil
}

func (repository repositoryBeers) Update(beer *beers.Domain, id int64) (*beers.Domain, error) {
	record := fromDomain(*beer)
	if err := repository.DB.Where("id = ?", id).Updates(&record).Error; err != nil {
		return &beers.Domain{}, err
	}
	if err := repository.DB.Where("id = ?", id).First(&record).Error; err != nil {
		return &beers.Domain{}, err
	}
	result := toDomain(record)
	return &result, nil
}

func (repository repositoryBeers) FindByID(id int64) (*beers.Domain, error) {
	record := Beers{}
	if err := repository.DB.Where("id = ?", id).First(&record).Error; err != nil {
		return &beers.Domain{}, err
	}
	result := toDomain(record)
	return &result, nil
}

func (repository repositoryBeers) All() (*[]beers.Domain, error) {
	records := []Beers{}
	if err := repository.DB.Find(&records).Error; err != nil {
		return &[]beers.Domain{}, err
	}
	result := toDomainList(records)
	return &result, nil
}
