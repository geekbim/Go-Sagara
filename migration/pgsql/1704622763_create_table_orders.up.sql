CREATE TABLE IF NOT EXISTS orders
(
  id uuid NOT NULL,
  pic_name varchar(255) NOT NULL,
  date_start timestamptz DEFAULT NULL,
  date_end timestamptz DEFAULT NULL,
  payment varchar(50) NOT NULL,
  capacity_employee int DEFAULT 0,
  requirement text NOT NULL,
  other varchar(255) NOT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);