package reward_table

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Model estructura de RewardTable
type RewardTable struct {
	ID        string    `json:"id" db:"id" valid:"required,uuid"`
	LotteryId string    `json:"lottery_id" db:"lottery_id" valid:"required"`
	IdWallet  string    `json:"id_wallet" db:"id_wallet" valid:"required"`
	Amount    int64     `json:"amount" db:"amount" valid:"required"`
	BlockId   int64     `json:"block_id" db:"block_id" valid:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewRewardTable(id string, lotteryId string, idWallet string, amount int64, blockId int64) *RewardTable {
	return &RewardTable{
		ID:        id,
		LotteryId: lotteryId,
		IdWallet:  idWallet,
		Amount:    amount,
		BlockId:   blockId,
	}
}

func (m *RewardTable) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
