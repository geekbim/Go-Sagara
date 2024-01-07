CREATE TABLE IF NOT EXISTS indonesia_provinces
(
  id bigint NOT NULL,
  name varchar(50) DEFAULT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);