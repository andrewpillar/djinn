CREATE TYPE visibility AS ENUM ('private', 'internal', 'public');

CREATE TABLE namespaces (
	id          SERIAL PRIMARY KEY,
	user_id     INT NOT NULL REFERENCES users(id),
	parent_id   INT NULL,
	name        VARCHAR(32) NOT NULL,
	path        VARCHAR(672) NOT NULL,
	description VARCHAR(255) NULL,
	level       INT NOT NULL,
	visibility  visibility DEFAULT 'private',
	created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at  TIMESTAMP NOT NULL DEFAULT NOW()
);
