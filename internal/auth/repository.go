package auth

import (
	"context"
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) CreateAccount(
	ctx context.Context,
	account *Account,
) error {
	_, err := r.DB.ExecContext(ctx, `
		INSERT INTO accounts (
			id, email, password_hash, email_verified, two_factor_enabled
		) VALUES ($1, $2, $3, false, false)
	`,
		account.ID,
		account.Email,
		account.PasswordHash,
	)
	return err
}

func (r *Repository) GetAccountByEmail(
	ctx context.Context,
	email string,
) (*Account, error) {
	row := r.DB.QueryRowContext(ctx, `
		SELECT id, email, password_hash, email_verified, two_factor_enabled,
		       created_at, updated_at
		FROM accounts
		WHERE email = $1
	`, email)

	var a Account
	err := row.Scan(
		&a.ID,
		&a.Email,
		&a.PasswordHash,
		&a.EmailVerified,
		&a.TwoFactorEnabled,
		&a.CreatedAt,
		&a.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *Repository) CreateSession(
	ctx context.Context,
	s *Session,
) error {
	_, err := r.DB.ExecContext(ctx, `
		INSERT INTO sessions (id, account_id, expires_at)
		VALUES ($1, $2, $3)
	`,
		s.ID,
		s.AccountID,
		s.ExpiresAt,
	)
	return err
}

func (r *Repository) GetSession(
	ctx context.Context,
	sessionID string,
) (*Session, error) {
	row := r.DB.QueryRowContext(ctx, `
		SELECT id, account_id, expires_at, created_at
		FROM sessions
		WHERE id = $1
	`, sessionID)

	var s Session
	err := row.Scan(
		&s.ID,
		&s.AccountID,
		&s.ExpiresAt,
		&s.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *Repository) DeleteSession(
	ctx context.Context,
	sessionID string,
) error {
	_, err := r.DB.ExecContext(ctx, `
		DELETE FROM sessions
		WHERE id = $1
	`, sessionID)
	return err
}

// email_verifications から有効なレコードを取得
func (r *Repository) GetEmailVerificationByCode(
	ctx context.Context,
	code string,
) (*EmailVerification, error) {
	row := r.DB.QueryRowContext(ctx, `
		SELECT id, account_id, expires_at, used_at
		FROM email_verifications
		WHERE code = $1
	`, code)

	var ev EmailVerification
	if err := row.Scan(
		&ev.ID,
		&ev.AccountID,
		&ev.ExpiresAt,
		&ev.UsedAt,
	); err != nil {
		return nil, err
	}
	return &ev, nil
}

// 認証を確定（account 更新 + verification 無効化）
func (r *Repository) VerifyEmailTx(
	ctx context.Context,
	accountID string,
	verificationID string,
) error {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, `
		UPDATE accounts
		SET email_verified = true, updated_at = now()
		WHERE id = $1
	`, accountID); err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, `
		UPDATE email_verifications
		SET used_at = now()
		WHERE id = $1
	`, verificationID); err != nil {
		return err
	}

	return tx.Commit()
}
