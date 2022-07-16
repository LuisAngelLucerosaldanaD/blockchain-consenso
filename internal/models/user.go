package models

import (
	"time"
)

type User struct {
	ID                 string    `json:"id" db:"id" valid:"required,uuid"`
	Nickname           string    `json:"nickname" db:"nickname" valid:"required"`
	Email              string    `json:"email" db:"email" valid:"required"`
	Password           string    `json:"password,omitempty" db:"password" valid:"required"`
	Name               string    `json:"name" db:"name" valid:"required"`
	Lastname           string    `json:"lastname" db:"lastname" valid:"required"`
	IdType             int       `json:"id_type" db:"id_type" valid:"required"`
	IdNumber           string    `json:"id_number" db:"id_number" valid:"required"`
	Cellphone          string    `json:"cellphone" db:"cellphone" valid:"required"`
	StatusId           int       `json:"status_id" db:"status_id" valid:"required"`
	FailedAttempts     int       `json:"failed_attempts,omitempty" db:"failed_attempts"`
	BlockDate          time.Time `json:"block_date,omitempty" db:"block_date"`
	DisabledDate       time.Time `json:"disabled_date,omitempty" db:"disabled_date"`
	LastLogin          time.Time `json:"last_login,omitempty" db:"last_login" `
	LastChangePassword time.Time `json:"last_change_password,omitempty" db:"last_change_password"`
	BirthDate          time.Time `json:"birth_date,omitempty" db:"birth_date"`
	VerifiedCode       string    `json:"verified_code,omitempty" db:"verified_code"`
	VerifiedAt         time.Time `json:"verified_at,omitempty" db:"verified_at"`
	IsDeleted          bool      `json:"is_deleted,omitempty" db:"is_deleted"`
	IdUser             string    `json:"id_user,omitempty" db:"id_user"`
	IdRole             int       `json:"id_role" db:"id_role" valid:"required"`
	FullPathPhoto      string    `json:"full_path_photo,omitempty" db:"full_path_photo"`
	RsaPrivate         string    `json:"rsa_private,omitempty" db:"rsa_private"  valid:"required"`
	RsaPublic          string    `json:"rsa_public,omitempty" db:"rsa_public" valid:"required"`
	RecoveryAccountAt  time.Time `json:"recovery_account_at,omitempty" db:"recovery_account_at"`
	RealIP             string    `json:"real_ip,omitempty"`
	DeletedAt          time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	CreatedAt          time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type Wallet struct {
	ID               string    `json:"id" db:"id" valid:"required,uuid"`
	Mnemonic         string    `json:"mnemonic" db:"mnemonic" valid:"required"`
	RsaPublic        string    `json:"rsa_public" db:"rsa_public" valid:"required"`
	RsaPrivate       string    `json:"rsa_private" db:"rsa_private" valid:"required"`
	RsaPublicDevice  string    `json:"rsa_public_device" db:"rsa_public_device" valid:"required"`
	RsaPrivateDevice string    `json:"rsa_private_device" db:"rsa_private_device" valid:"required"`
	IpDevice         string    `json:"ip_device" db:"ip_device" valid:"required"`
	StatusId         int       `json:"status_id" db:"status_id" valid:"required"`
	IdentityNumber   string    `json:"identity_number" db:"identity_number" valid:"required"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}

type UserTemp struct {
	ID           string    `json:"id" db:"id" valid:"required,uuid"`
	Nickname     string    `json:"nickname" db:"nickname" valid:"required"`
	Email        string    `json:"email" db:"email" valid:"required"`
	Password     string    `json:"password" db:"password" valid:"required"`
	Name         string    `json:"name" db:"name" valid:"-"`
	Lastname     string    `json:"lastname" db:"lastname" valid:"-"`
	IdType       int       `json:"id_type" db:"id_type" valid:"-"`
	IdNumber     string    `json:"id_number" db:"id_number" valid:"-"`
	Cellphone    string    `json:"cellphone" db:"cellphone" valid:"required"`
	BirthDate    time.Time `json:"birth_date" db:"birth_date" valid:"required"`
	VerifiedCode string    `json:"verified_code" db:"verified_code" valid:"required"`
	IsDeleted    bool      `json:"is_deleted" db:"is_deleted"`
	IdUser       string    `json:"id_user" db:"id_user"`
	DeletedAt    time.Time `json:"deleted_at" db:"deleted_at"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
