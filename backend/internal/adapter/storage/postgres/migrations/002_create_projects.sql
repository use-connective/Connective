CREATE TABLE IF NOT EXISTS projects
  (
	 id                 UUID NOT NULL CONSTRAINT projects_pk PRIMARY KEY,
	 name               TEXT NOT NULL,
	 owner              INTEGER NOT NULL,
     sdk_auth_secret    TEXT NOT NULL,
	 created_at         DATE NOT NULL DEFAULT NOW(),
	 updated_at         DATE NOT NULL DEFAULT NOW()
  );
