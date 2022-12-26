package participants

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Model estructura de ParticipantsTable
type Participants struct {
	ID         string    `json:"id" db:"id" valid:"required,uuid"`
	LotteryId  string    `json:"lottery_id" db:"lottery_id" valid:"required"`
	WalletId   string    `json:"wallet_id" db:"wallet_id" valid:"required"`
	Amount     float64   `json:"amount" db:"amount" valid:"required"`
	Accepted   bool      `json:"accepted" db:"accepted" valid:"-"`
	TypeCharge int       `json:"type_charge" db:"type_charge" valid:"required"`
	Returned   bool      `json:"returned" db:"returned" valid:"-"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

func NewParticipants(id string, lotteryId string, walletId string, amount float64, accepted bool, typeCharge int, returned bool) *Participants {
	return &Participants{
		ID:         id,
		LotteryId:  lotteryId,
		WalletId:   walletId,
		Amount:     amount,
		Accepted:   accepted,
		TypeCharge: typeCharge,
		Returned:   returned,
	}
}

func (m *Participants) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
