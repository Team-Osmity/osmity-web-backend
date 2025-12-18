package auth

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repo *Repository
}

func (s *Service) Register(
	ctx context.Context,
	email string,
	password string,
) error {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	account := &Account{
		ID:           generateUUID(),
		Email:        email,
		PasswordHash: ptr(string(hash)),
	}

	return s.Repo.CreateAccount(ctx, account)
}

var ErrInvalidCredentials = errors.New("invalid credentials")

func (s *Service) Login(
	ctx context.Context,
	email string,
	password string,
) (*Session, error) {

	account, err := s.Repo.GetAccountByEmail(ctx, email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if account.PasswordHash == nil {
		return nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(*account.PasswordHash),
		[]byte(password),
	); err != nil {
		return nil, ErrInvalidCredentials
	}

	session := &Session{
		ID:        generateUUID(),
		AccountID: account.ID,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	if err := s.Repo.CreateSession(ctx, session); err != nil {
		return nil, err
	}

	return session, nil
}

func (s *Service) Logout(
	ctx context.Context,
	sessionID string,
) error {
	return s.Repo.DeleteSession(ctx, sessionID)
}

var (
	ErrInvalidCode   = errors.New("invalid_code")
	ErrExpiredCode   = errors.New("expired_code")
	ErrAlreadyUsed   = errors.New("already_used")
)

func (s *Service) VerifyEmail(
	ctx context.Context,
	code string,
) error {
	ev, err := s.Repo.GetEmailVerificationByCode(ctx, code)
	if err != nil {
		return ErrInvalidCode
	}

	if ev.UsedAt != nil {
		return ErrAlreadyUsed
	}
	if time.Now().After(ev.ExpiresAt) {
		return ErrExpiredCode
	}

	return s.Repo.VerifyEmailTx(ctx, ev.AccountID, ev.ID)
}
