CREATE TABLE oauth_accounts (
  id UUID PRIMARY KEY,
  account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
  provider TEXT NOT NULL,
  provider_user_id TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  UNIQUE (provider, provider_user_id),
  UNIQUE (account_id, provider)
);

CREATE INDEX idx_oauth_accounts_account_id
  ON oauth_accounts(account_id);
