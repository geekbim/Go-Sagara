CREATE TABLE IF NOT EXISTS offices
(
  id uuid NOT NULL,
  name varchar(255) DEFAULT NULL,
  room_type varchar(255) DEFAULT NULL,
  description TEXT DEFAULT NULL,
  image_reference uuid DEFAULT NULL REFERENCES data_objects(id),
  region_id BIGINT DEFAULT NULL REFERENCES indonesia_provinces(id),
  additional_data JSONB DEFAULT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);