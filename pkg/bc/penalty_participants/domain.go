package penalty_participants

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Model estructura de PenaltyParticipants
type PenaltyParticipants struct {
	ID                string    `json:"id" db:"id" valid:"required,uuid"`
	LotteryId         string    `json:"lottery_id" db:"lottery_id" valid:"required"`
	ParticipantsId    string    `json:"participants_id" db:"participants_id" valid:"required"`
	Amount            float64   `json:"amount" db:"amount" valid:"required"`
	PenaltyPercentage float64   `json:"penalty_percentage" db:"penalty_percentage" valid:"required"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}

func NewPenaltyParticipants(id string, lotteryId string, participantsId string, amount float64, penaltyPercentage float64) *PenaltyParticipants {
	return &PenaltyParticipants{
		ID:                id,
		LotteryId:         lotteryId,
		ParticipantsId:    participantsId,
		Amount:            amount,
		PenaltyPercentage: penaltyPercentage,
	}
}

func (m *PenaltyParticipants) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
