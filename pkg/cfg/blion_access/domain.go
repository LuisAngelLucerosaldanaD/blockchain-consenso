package blion_access

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// BLionAccess  Model struct BLionAccess
type BLionAccess struct {
	ID        string    `json:"id" db:"id" valid:"required,uuid"`
	Key       string    `json:"key" db:"key" valid:"required"`
	Status    string    `json:"status" db:"status" valid:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewBLionAccess(id string, key string, status string) *BLionAccess {
	return &BLionAccess{
		ID:     id,
		Key:    key,
		Status: status,
	}
}

func (m *BLionAccess) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
