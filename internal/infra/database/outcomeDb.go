package database

import (
	"fmt"
	"log"
	"time"

	"github.com/rafaapcode/finance-app-backend/internal/entity"
	"gorm.io/gorm"
)

type OutcomeDb struct {
	DB *gorm.DB
}

func NewOutcomeDb(db *gorm.DB) *OutcomeDb {
	return &OutcomeDb{
		DB: db,
	}
}

func (outDb *OutcomeDb) CreateOutcome(outcome *entity.Outcome) (int, error) {
	err := outDb.DB.Create(outcome).Error
	if err != nil {
		log.Fatal(err.Error())
		return 500, err
	}

	return 201, nil
}

func (outDb *OutcomeDb) GetOutcomeById(id string) (*entity.Outcome, int, error) {
	var outcome entity.Outcome
	err := outDb.DB.First(&outcome, "id = ?", id).Error
	if err != nil {
		log.Fatal(err.Error())
		return nil, 404, err
	}
	return &outcome, 200, nil
}

func (outDb *OutcomeDb) GetAllOutcomeOfMonth(month time.Time, userId string) ([]entity.Outcome, int, error) {
	var outcomes []entity.Outcome
	monthDate := int(month.Month())
	yearDate := int(month.Year())

	err := outDb.DB.Where("DATE_TRUNC('month', createdAt) = DATE_TRUNC('month', TO_DATE(? || '-01', 'YYYY-MM-DD'))", fmt.Sprintf("%04d-%02d", yearDate, monthDate)).Preload("User", "id = ?", userId).Find(&outcomes).Error

	if err != nil {
		log.Fatal(err.Error())
		return nil, 404, err
	}

	return outcomes, 200, nil
}

func (outDb *OutcomeDb) GetAllFixedOutcome(userId string) ([]entity.Outcome, int, error) {
	var outcomes []entity.Outcome

	err := outDb.DB.Where("type = ?", "fixo").Preload("User", "id = ?", userId).Find(&outcomes).Error
	if err != nil {
		log.Fatal(err.Error())
		return nil, 404, err
	}

	return outcomes, 200, nil
}

func (outDb *OutcomeDb) GetAllOutcomeByCategory(category string, userId string) ([]entity.Outcome, int, error) {
	var outcomes []entity.Outcome

	err := outDb.DB.Where("category = ?", category).Preload("User", "id = ?", userId).Find(&outcomes).Error
	if err != nil {
		log.Fatal(err.Error())
		return nil, 404, err
	}

	return outcomes, 200, nil
}

func (outDb *OutcomeDb) GetAllOutcomeByPaymentMethod(method string, userId string) ([]entity.Outcome, int, error) {
	var outcomes []entity.Outcome

	err := outDb.DB.Where("paymentMethod = ?", method).Preload("User", "id = ?", userId).Find(&outcomes).Error
	if err != nil {
		log.Fatal(err.Error())
		return nil, 404, err
	}

	return outcomes, 200, nil
}

func (outDb *OutcomeDb) GetAllOutcomeByType(typeOutcome string, userId string) ([]entity.Outcome, int, error) {
	var outcomes []entity.Outcome

	err := outDb.DB.Where("type = ?", typeOutcome).Preload("User", "id = ?", userId).Find(&outcomes).Error
	if err != nil {
		log.Fatal(err.Error())
		return nil, 404, err
	}

	return outcomes, 200, nil
}

func (outDb *OutcomeDb) GetOutcomeAboutToExpire(daysToExpire int, userId string) ([]entity.Outcome, int, error) {
	var outcomes []entity.Outcome
	query := fmt.Sprintf("expireDate BETWEEN CURRENT_DATE AND CURRENT_DATE + INTERVAL '%d days'", daysToExpire)

	err := outDb.DB.Where(query).Preload("User", "id = ?", userId).Find(&outcomes).Error
	if err != nil {
		log.Fatal(err.Error())
		return nil, 404, err
	}

	return outcomes, 200, nil
}

func (outDb *OutcomeDb) GetOutcomeLessThan(value float64, userId string) ([]entity.Outcome, int, error) {
	var outcomes []entity.Outcome

	err := outDb.DB.Where("value < ?", value).Preload("User", "id = ?", userId).Find(&outcomes).Error
	if err != nil {
		log.Fatal(err.Error())
		return nil, 404, err
	}

	return outcomes, 200, nil
}

func (outDb *OutcomeDb) GetOutcomeMoreThan(value float64, userId string) ([]entity.Outcome, int, error) {
	var outcomes []entity.Outcome

	err := outDb.DB.Where("value > ?", value).Preload("User", "id = ?", userId).Find(&outcomes).Error
	if err != nil {
		log.Fatal(err.Error())
		return nil, 404, err
	}

	return outcomes, 200, nil
}

func (outDb *OutcomeDb) UpdateOutcome(newData *entity.Outcome) (int, error) {
	_, status, err := outDb.GetOutcomeById(newData.Id.String())
	if status != 200 {
		log.Fatal(err.Error())
		return 404, err
	}

	err = outDb.DB.Save(newData).Error

	if err != nil {
		log.Fatal(err.Error())
		return 500, err
	}

	return 200, nil
}

func (outDb *OutcomeDb) DeleteOutcome(id string) (string, int, error) {
	outcome, status, err := outDb.GetOutcomeById(id)
	if status != 200 {
		log.Fatal(err.Error())
		return "", 404, err
	}

	err = outDb.DB.Delete(outcome).Error

	if err != nil {
		log.Fatal(err.Error())
		return "", 500, err
	}

	return "Saída deletada com sucesso", 200, nil
}
