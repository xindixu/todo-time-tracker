-- ACCOUNTS
DROP TABLE IF EXISTS accounts;
DROP TYPE IF EXISTS account_type;

-- USERS
DROP TABLE IF EXISTS users;

-- ORGANIZATIONS
DROP TABLE IF EXISTS organizations;

-- TAGS
DROP INDEX IF EXISTS idx_tags_account_id_name;
DROP INDEX IF EXISTS idx_tags_account_id;
DROP TABLE IF EXISTS tags;

-- TASKS
DROP INDEX IF EXISTS idx_tasks_account_id_name;
DROP INDEX IF EXISTS idx_tasks_account_id;
DROP TABLE IF EXISTS tasks;
DROP TYPE IF EXISTS task_status;

-- SESSIONS
ALTER TABLE sessions DROP CONSTRAINT IF EXISTS sessions_user_id_no_overlap;
DROP INDEX IF EXISTS idx_sessions_user_id_time_range;
DROP TABLE IF EXISTS sessions;

-- ORGANIZATION USERS
DROP INDEX IF EXISTS idx_organization_users_organization_id;
DROP INDEX IF EXISTS idx_organization_users_user_id;
DROP TABLE IF EXISTS organization_users;

-- TASK USERS
DROP INDEX IF EXISTS idx_task_users_task_id;
DROP INDEX IF EXISTS idx_task_users_user_id;
DROP TABLE IF EXISTS task_users;

-- TASK TAGS
DROP INDEX IF EXISTS idx_task_tags_tag_id;
DROP INDEX IF EXISTS idx_task_tags_task_id;
DROP TABLE IF EXISTS task_tags;