
-- Enable required extensions
CREATE EXTENSION IF NOT EXISTS btree_gist;

-- ACCOUNTS
CREATE TYPE account_type AS ENUM ('USER', 'ORGANIZATION');
CREATE TABLE IF NOT EXISTS accounts (
  id SERIAL PRIMARY KEY,
  uuid UUID NOT NULL UNIQUE,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now(),
  type account_type NOT NULL
);

-- USERS
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  uuid UUID NOT NULL UNIQUE,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now(),
  account_id INTEGER NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL
);

-- ORGANIZATIONS
CREATE TABLE organizations (
  id SERIAL PRIMARY KEY,
  uuid UUID NOT NULL UNIQUE,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now(),
  account_id INTEGER NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
  name VARCHAR(255) NOT NULL
);

-- TAGS
CREATE TABLE tags (
  id SERIAL PRIMARY KEY,
  uuid UUID NOT NULL UNIQUE,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now(),
  account_id INTEGER NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  color VARCHAR(255),
  UNIQUE (account_id, name)
);
CREATE INDEX idx_tags_account_id ON tags(account_id);
CREATE UNIQUE INDEX idx_tags_account_id_name ON tags(account_id, name);

-- TASKS
CREATE TYPE task_status AS ENUM ('TODO', 'IN_PROGRESS', 'DONE', 'BLOCKED');
CREATE TABLE tasks (
  id SERIAL PRIMARY KEY,
  uuid UUID NOT NULL UNIQUE,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now(),
  account_id INTEGER NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  status task_status NOT NULL,
  estimated_duration INTERVAL,
  UNIQUE (account_id, name)
);
CREATE INDEX idx_tasks_account_id ON tasks(account_id);
CREATE UNIQUE INDEX idx_tasks_account_id_name ON tasks(account_id, name);

-- SESSIONS
CREATE TABLE sessions (
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now(),
  time_range TSRANGE NOT NULL,
  task_id INTEGER NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
  user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE
);
CREATE INDEX idx_sessions_task_id ON sessions(task_id);
CREATE INDEX idx_sessions_user_id ON sessions(user_id);

-- Enforce no overlaps for a given user
CREATE INDEX idx_sessions_user_id_time_range ON sessions USING gist (user_id, time_range);
ALTER TABLE sessions
ADD CONSTRAINT sessions_user_id_no_overlap EXCLUDE USING gist (user_id WITH =, time_range WITH &&);

-- ORGANIZATION USERS
CREATE TABLE organization_users (
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now(),
  user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  organization_id INTEGER NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
  UNIQUE (user_id, organization_id)
);
CREATE INDEX idx_organization_users_organization_id ON organization_users(organization_id);
CREATE INDEX idx_organization_users_user_id ON organization_users(user_id);

-- TASK USERS
CREATE TABLE task_users (
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now(),
  user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  task_id INTEGER NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
  UNIQUE (user_id, task_id)
);
CREATE INDEX idx_task_users_task_id ON task_users(task_id);
CREATE INDEX idx_task_users_user_id ON task_users(user_id);

-- TASK TAGS
CREATE TABLE task_tags (
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now(),
  task_id INTEGER NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
  tag_id INTEGER NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
  UNIQUE (task_id, tag_id)
);
CREATE INDEX idx_task_tags_task_id ON task_tags(task_id);
CREATE INDEX idx_task_tags_tag_id ON task_tags(tag_id);