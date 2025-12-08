CREATE TABLE IF NOT EXISTS provider_credentials
(
    id SERIAL PRIMARY KEY,
    project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    provider_id INT NOT NULL REFERENCES providers(id) ON DELETE CASCADE,

    client_id TEXT NOT NULL,
    client_secret TEXT NOT NULL,

    scopes TEXT[],
    is_active BOOLEAN DEFAULT TRUE,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),

    UNIQUE (project_id, provider_id)
)
