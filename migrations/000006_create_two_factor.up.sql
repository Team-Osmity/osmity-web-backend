CREATE TABLE two_factor (
  account_id UUID PRIMARY KEY REFERENCES accounts(id) ON DELETE CASCADE,
  secret TEXT NOT NULL,
  enabled BOOLEAN NOT NULL DEFAULT false,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
