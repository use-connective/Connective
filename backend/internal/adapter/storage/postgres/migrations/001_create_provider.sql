CREATE TABLE IF NOT EXISTS providers (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,        -- e.g. slack, github, notion
    display_name TEXT NOT NULL,       -- Slack, GitHub, Notion
    auth_type TEXT NOT NULL,          -- oauth2, api_key, basic_auth
    image_url TEXT NOT NULL,
    category TEXT NOT NULL,
    description varchar NOT NULL,
    primary_color varchar NOT NULL,

    -- OAuth URLs
    auth_url TEXT NOT NULL,
    token_url TEXT NOT NULL,
    refresh_token_url TEXT NOT NULL,

    -- Redirect URL for this provider
    redirect_url TEXT NOT NULL,


    -- Default Scope
    default_scopes TEXT[],
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
