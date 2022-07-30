package rewards

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Model estructura de RewardTable
type Reward struct {
	ID        string    `json:"id" db:"id" valid:"required,uuid"`
	LotteryId string    `json:"lottery_id" db:"lottery_id" valid:"required"`
	IdWallet  string    `json:"id_wallet" db:"id_wallet" valid:"required"`
	Amount    float64   `json:"amount" db:"amount" valid:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewReward(id string, lotteryId string, idWallet string, amount float64) *Reward {
	return &Reward{
		ID:        id,
		LotteryId: lotteryId,
		IdWallet:  idWallet,
		Amount:    amount,
	}
}

func (m *Reward) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
