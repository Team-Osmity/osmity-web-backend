package auth

import "time"

type Account struct {
	ID               string
	Email            string
	PasswordHash     *string
	EmailVerified    bool
	TwoFactorEnabled bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Profile struct {
	AccountID   string
	DisplayName *string
	AvatarURL   *string
	Bio         *string
}

type Session struct {
	ID        string
	AccountID string
	ExpiresAt time.Time
	CreatedAt time.Time
}
