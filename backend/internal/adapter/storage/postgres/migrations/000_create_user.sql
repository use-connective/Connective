CREATE TABLE IF NOT EXISTS users
  (
	 id         BIGSERIAL PRIMARY KEY,
	 name       TEXT NOT NULL,
	 email      TEXT NOT NULL UNIQUE,
	 password   TEXT NOT NULL,
	 created_at DATE NOT NULL DEFAULT NOW(),
	 updated_at DATE NOT NULL DEFAULT NOW(),
     is_onboarding_completed bool DEFAULT false NOT NULL,
     state      text    default 'PROJECT_PENDING'::text not null
  );
