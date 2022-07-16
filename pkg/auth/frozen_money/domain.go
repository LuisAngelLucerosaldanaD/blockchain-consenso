package frozen_money

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Model estructura de FrozenMoney
type FrozenMoney struct {
	ID        string    `json:"id" db:"id" valid:"required,uuid"`
	WalletId  string    `json:"wallet_id" db:"wallet_id" valid:"required"`
	Amount    int64     `json:"amount" db:"amount" valid:"required"`
	LotteryId string    `json:"lottery_id" db:"lottery_id" valid:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewFrozenMoney(id string, walletId string, amount int64, lotteryId string) *FrozenMoney {
	return &FrozenMoney{
		ID:        id,
		WalletId:  walletId,
		Amount:    amount,
		LotteryId: lotteryId,
	}
}

func (m *FrozenMoney) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
