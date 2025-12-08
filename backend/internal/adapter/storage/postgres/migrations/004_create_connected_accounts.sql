CREATE TABLE IF NOT EXISTS connected_accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    provider_id INT NOT NULL REFERENCES providers(id) ON DELETE CASCADE,
    user_id TEXT NOT NULL,

    -- Account details
    external_user_id TEXT,               -- Slack user ID, GitHub user ID etc
    external_team_id TEXT,               -- Slack workspace ID, GitHub org ID
    account_email TEXT,

    -- Token info
    access_token TEXT NOT NULL,
    refresh_token TEXT,
    expires_at TIMESTAMP,
    token_type TEXT,
    scope TEXT,

    -- Metadata
    raw_response TEXT,                  -- full OAuth response
    connected_at TIMESTAMP DEFAULT NOW(),

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
