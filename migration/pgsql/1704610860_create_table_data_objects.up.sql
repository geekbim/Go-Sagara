CREATE TABLE IF NOT EXISTS data_objects
(
  id uuid NOT NULL,
  path varchar(255) DEFAULT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);