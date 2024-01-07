CREATE TABLE IF NOT EXISTS users
(
  id uuid NOT NULL,
  name varchar(45) DEFAULT NULL,
  email varchar(45) UNIQUE,
  password varchar(225) DEFAULT NULL,
  position varchar(20) DEFAULT NULL,
  phone_number varchar(15) DEFAULT null,
  company_name varchar(45) DEFAULT null,
  company_address varchar(255) DEFAULT null,
  information varchar(45) DEFAULT null,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);