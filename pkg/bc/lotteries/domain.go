package lotteries

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Model estructura de LotteryTable
type Lottery struct {
	ID                    string     `json:"id" db:"id" valid:"required,uuid"`
	BlockId               int64      `json:"block_id" db:"block_id" valid:"required"`
	RegistrationStartDate time.Time  `json:"registration_start_date" db:"registration_start_date" valid:"required"`
	RegistrationEndDate   *time.Time `json:"registration_end_date" db:"registration_end_date" valid:"required"`
	LotteryStartDate      *time.Time `json:"lottery_start_date" db:"lottery_start_date" valid:"required"`
	LotteryEndDate        *time.Time `json:"lottery_end_date" db:"lottery_end_date" valid:"required"`
	ProcessEndDate        *time.Time `json:"process_end_date" db:"process_end_date" valid:"required"`
	ProcessStatus         int        `json:"process_status" db:"process_status" valid:"required"`
	CreatedAt             time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at" db:"updated_at"`
}

func NewLottery(id string, blockId int64, registrationStartDate time.Time, registrationEndDate *time.Time, lotteryStartDate *time.Time, lotteryEndDate *time.Time, processEndDate *time.Time, processStatus int) *Lottery {
	return &Lottery{
		ID:                    id,
		BlockId:               blockId,
		RegistrationStartDate: registrationStartDate,
		RegistrationEndDate:   registrationEndDate,
		LotteryStartDate:      lotteryStartDate,
		LotteryEndDate:        lotteryEndDate,
		ProcessEndDate:        processEndDate,
		ProcessStatus:         processStatus,
	}
}

func (m *Lottery) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
