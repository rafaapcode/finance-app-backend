package dto

type GetJwtInput struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type CreateIncomeDto struct {
	UserId string  `json:"userId"`
	Value  float64 `json:"value"`
}

type CreateExtraIncomeDto struct {
	UserId   string  `json:"userId"`
	Category string  `json:"category"`
	Value    float64 `json:"value"`
}

type CreateOutcomeDto struct {
	OutcomeType   string  `json:"outcome_type"`
	Category      string  `json:"category"`
	PaymentMethod string  `json:"paymentMethod"`
	UserId        string  `json:"userId"`
	Value         float64 `json:"value"`
	Notification  bool    `json:"notification"`
	ExpireDate    int     `json:"expireDate`
}

type CreateGoalDto struct {
	Category   string  `json:"category"`
	UserId     string  `json:"userId"`
	Percentage float64 `json:"percentage"`
}

type CreateInvestmentDto struct {
	UserId    string  `json:"userId"`
	Category  string  `json:"category"`
	Quantity  int     `json:"quantity"`
	StockCode string  `json:"stockCode"`
	BuyPrice  float64 `json:"buyPrice"`
}

type UpdateSellInvestmentDto struct {
	Quantity  int     `json:"quantity"`
	SellPrice float64 `json:"sellPrice"`
}

type UpdateSupplyInvestmentDto struct {
	Quantity    int     `json:"quantity"`
	SupplyPrice float64 `json:"supplyPrice"`
}
