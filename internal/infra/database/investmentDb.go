package database

import (
	"log"
	"time"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
	"gorm.io/gorm"
)

type InvestmentDb struct {
	DB *gorm.DB
}

func NewInvestmentDB(db *gorm.DB) *InvestmentDb {
	return &InvestmentDb{
		DB: db,
	}
}

func (invDb *InvestmentDb) CreateInvestment(invest *entity.Investment) (int, error) {
	err := invDb.DB.Create(invest).Error

	if err != nil {
		log.Fatal(err.Error())
		return 500, err
	}

	return 201, nil
}

func (invDb *InvestmentDb) GetTotalOfInvestment(userid string) (float64, int, error) {
	var investments []entity.Investment

	err := invDb.DB.Where("sellPrice = ? AND userid = ?", 0.0, userid).Find(&investments).Error

	if err != nil {
		log.Fatal(err.Error())
		return 0, 404, err
	}

	var totalInvested float64

	for _, value := range investments {
		totalInvested += value.Value
	}

	return totalInvested, 200, nil
}

func (invDb *InvestmentDb) GetAllOfInvestment(pageNumber int, sort, userid string) ([]entity.Investment, int, error) {
	var investments []entity.Investment
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if pageNumber != 0 {
		err = invDb.DB.Where("userid = ?", userid).Limit(5).Offset((pageNumber - 1) * 5).Order("buyDate " + sort).Find(&investments).Error
	} else {
		err = invDb.DB.Where("userid = ?", userid).Limit(5).Order("buyDate " + sort).Find(&investments).Error
	}

	if err != nil {
		log.Fatal(err.Error())
		return nil, 404, err
	}

	return investments, 200, nil
}

func (invDb *InvestmentDb) GetInvestmentByName(name, userid string) (entity.Investment, int, error) {
	var investment entity.Investment

	err := invDb.DB.First(&investment, "stockCode = ? AND userid = ?", name, userid).Error

	if err != nil {
		log.Fatal(err.Error())
		return investment, 404, err
	}

	return investment, 200, nil
}

func (invDb *InvestmentDb) GetInvestmentByCategory(pageNumber int, category, userid string) ([]entity.Investment, int, error) {
	var investments []entity.Investment
	var err error

	if pageNumber != 0 {
		err = invDb.DB.Where("category = ? AND sellPrice = ? AND userid = ?", category, 0.0, userid).Limit(5).Offset((pageNumber - 1) * 5).Find(&investments).Error
	} else {
		err = invDb.DB.Where("category = ? AND sellPrice = ? AND userid = ?", category, 0.0, userid).Limit(5).Find(&investments).Error
	}

	if err != nil {
		log.Fatal(err.Error())
		return nil, 404, err
	}

	return investments, 200, nil
}

func (invDb *InvestmentDb) UpdateInvestment(newData *entity.Investment) (int, error) {
	_, status, err := invDb.GetInvestmentByName(newData.StockCode, newData.Userid)

	if status != 200 {
		return 404, err
	}

	err = invDb.DB.Save(newData).Error

	if err != nil {
		log.Fatal(err.Error())
		return 500, err
	}

	return 200, nil
}

func (invDb *InvestmentDb) GetAssetGrowth(userid string) (entity.Metrics, int, error) {
	return entity.Metrics{}, 200, nil
}

func (invDb *InvestmentDb) GetPortfolioDiversification(userid string) (entity.Metrics, int, error) {
	// var investments []entity.Investment

	return entity.Metrics{}, 200, nil

}

func (invDb *InvestmentDb) GetMonthInvestment(userid string, month time.Month) (entity.Metrics, int, error) {
	return entity.Metrics{}, 200, nil
}
