package validator_votes

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Model estructura de ValidatorVotes
type ValidatorVotes struct {
	ID             string    `json:"id" db:"id" valid:"required,uuid"`
	WalletId       string    `json:"wallet_id" db:"wallet_id" valid:"required"`
	LotteryId      string    `json:"lottery_id" db:"lottery_id" valid:"required"`
	ParticipantsId string    `json:"participants_id" db:"participants_id" valid:"required"`
	Hash           string    `json:"hash" db:"hash" valid:"required"`
	Vote           bool      `json:"vote" db:"vote" valid:"required"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

func NewValidatorVotes(id string, walletId string, lotteryId string, participantsId string, hash string, vote bool) *ValidatorVotes {
	return &ValidatorVotes{
		ID:             id,
		WalletId:       walletId,
		LotteryId:      lotteryId,
		ParticipantsId: participantsId,
		Hash:           hash,
		Vote:           vote,
	}
}

func (m *ValidatorVotes) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
