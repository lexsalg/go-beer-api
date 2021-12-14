package beers

import (
	"beerapi/bussiness/beers"
	"beerapi/drivers/mysql/users"
	"beerapi/helper"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Suite struct {
	suite.Suite
	DB         *gorm.DB
	repository beers.Repository
}

func (s *Suite) SetupSuite() {

	s.DB = helper.InitTestingDB()

	_ = s.DB.AutoMigrate(&users.Users{}, &Beers{})

	s.DB.Logger.LogMode(logger.Info)
	s.repository = NewRepositoryMySQL(s.DB)
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))

}
func (s *Suite) Test_BeerRepository_Save_Succeed() {
	beer := beers.Domain{
		Name:     "Pilsen",
		Brewery:  "Backus",
		Country:  "Perú",
		Price:    1,
		Currency: "PEN",
	}
	_, err := s.repository.Create(&beer)
	require.NoError(s.T(), err)

}

func (s *Suite) Test_BeerRepository_Create_Succeed() {
	beer := beers.Domain{
		Name:     "Pilsen",
		Brewery:  "Backus",
		Country:  "Perú",
		Price:    1,
		Currency: "PEN",
	}
	_, err := s.repository.Create(&beer)
	require.NoError(s.T(), err)

}

func (s *Suite) Test_BeerRepository_Update_Succeed() {
	beerID := int64(1)
	beer := beers.Domain{
		Name:     "Pilsen",
		Brewery:  "Backus",
		Country:  "Perú",
		Price:    1,
		Currency: "PEN",
	}
	beerUpdated, err := s.repository.Update(&beer, beerID)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), beerUpdated.Country, beer.Country)
	assert.Equal(s.T(), beerUpdated.Price, beer.Price)
}

func (s *Suite) Test_BeerRepository_Update_Fail() {
	beerID := int64(100000)
	beer := beers.Domain{
		Name:     "Pilsen",
		Brewery:  "Backus",
		Country:  "Perú",
		Price:    1,
		Currency: "PEN",
	}
	_, err := s.repository.Update(&beer, beerID)
	require.Error(s.T(), err)

}

func (s *Suite) Test_BeerRepository_FindByID_Fail() {
	beerID := int64(100000)
	_, err := s.repository.FindByID(beerID)
	require.Error(s.T(), err)

}
func (s *Suite) Test_BeerRepository_FindByID_Success() {
	beerID := int64(1)
	beer := beers.Domain{
		Name:     "Pilsen",
		Brewery:  "Backus",
		Country:  "Perú",
		Price:    1,
		Currency: "PEN",
	}

	beerBD, err := s.repository.FindByID(beerID)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), beer.Name, beerBD.Name)
	assert.Equal(s.T(), beer.Price, beerBD.Price)

}

func (s *Suite) Test_BeerRepository_All_Success() {

	beersBD, err := s.repository.All()
	size := len(*beersBD)
	require.NoError(s.T(), err)
	assert.NotEqual(s.T(), size, 0)

}
