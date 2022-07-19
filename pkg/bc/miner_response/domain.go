package miner_response

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Model estructura de MinerResponse
type MinerResponse struct {
	ID             string    `json:"id" db:"id" valid:"required,uuid"`
	LotteryId      string    `json:"lottery_id" db:"lottery_id" valid:"required"`
	ParticipantsId string    `json:"participants_id" db:"participants_id" valid:"required"`
	Hash           string    `json:"hash" db:"hash" valid:"required"`
	Status         int       `json:"status" db:"status" valid:"required"`
	Nonce          int64     `json:"nonce"`
	Difficulty     int       `json:"difficulty"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

func NewMinerResponse(id string, lotteryId string, participantsId string, hash string, status int, nonce int64, difficulty int) *MinerResponse {
	return &MinerResponse{
		ID:             id,
		LotteryId:      lotteryId,
		ParticipantsId: participantsId,
		Hash:           hash,
		Status:         status,
		Nonce:          nonce,
		Difficulty:     difficulty,
	}
}

func (m *MinerResponse) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
