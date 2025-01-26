package entity

import "errors"

var (
	ErrIdIdRequired            = errors.New("id is required")
	ErrIdInvalidId             = errors.New("id is invalid")
	ErrNameIsRequired          = errors.New("name is required")
	ErrNameIsInvalid           = errors.New("name is invalid")
	ErrEmailIsRequired         = errors.New("email is required")
	ErrEmailIsInvalid          = errors.New("email is invalid")
	ErrPhotoUrlIsRequired      = errors.New("photoUrl is required")
	ErrPhotoUrlIsInvalid       = errors.New("photoUrl is invalid")
	ErrInvestmentIdIsRequired  = errors.New("investment is required")
	ErrInvestmentIdIsInvalid   = errors.New("investment is invalid")
	ErrCategoryIsRequired      = errors.New("category is required")
	ErrStockCodeIsRequired     = errors.New("stockCode is required")
	ErrQuantityIsInvalid       = errors.New("quantity is invalid")
	ErrSupplyPriceIsInvalid    = errors.New("supplyPrice is invalid")
	ErrValueIsInvalid          = errors.New("value is invalid")
	ErrSellPriceIsInvalid      = errors.New("sellPrice is invalid")
	ErrUserIdIsRequired        = errors.New("userid is required")
	ErrUserIdIsInvalid         = errors.New("userid is invalid")
	ErrTypeIsRequired          = errors.New("type is required")
	ErrPaymentMethodIsRequired = errors.New("paymentMethod is required")
	ErrPaymentMethodIsInvalid  = errors.New("paymentMethod is invalid")
	ErrBuyPriceIsInvalid       = errors.New("buyPrice is invalid")
	ErrPercentageIsInvalid     = errors.New("percentage is invalid")
)
