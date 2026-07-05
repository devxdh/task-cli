BEGIN;

-- Create Task Type safely
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'task_status') THEN
    CREATE TYPE task_status AS ENUM ('due', 'completed');
  END IF;
END
$$;

-- Creates Task Table
CREATE TABLE IF NOT EXISTS tasks (
  id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  title VARCHAR(100) NOT NULL,
  description TEXT,
  status task_status NOT NULL DEFAULT 'due',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

COMMIT;