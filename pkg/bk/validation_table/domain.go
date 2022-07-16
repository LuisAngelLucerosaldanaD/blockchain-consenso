package validation_table

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Model estructura de ValidationTable
type ValidationTable struct {
	ID             string    `json:"id" db:"id" valid:"required,uuid"`
	LotteryId      string    `json:"lottery_id" db:"lottery_id" valid:"required"`
	WalletId       string    `json:"wallet_id" db:"wallet_id" valid:"required"`
	ParticipantsId string    `json:"participants_id" db:"participants_id" valid:"required"`
	Vote           bool      `json:"vote" db:"vote" valid:"required"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

func NewValidationTable(id string, lotteryId string, walletId string, participantsId string, vote bool) *ValidationTable {
	return &ValidationTable{
		ID:             id,
		LotteryId:      lotteryId,
		WalletId:       walletId,
		ParticipantsId: participantsId,
		Vote:           vote,
	}
}

func (m *ValidationTable) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
