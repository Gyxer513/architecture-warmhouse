CREATE TABLE IF NOT EXISTS sensors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    type VARCHAR(32),
    value FLOAT,
    unit VARCHAR(16),
    status VARCHAR(32),
    last_updated TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE sensors ALTER COLUMN value SET DEFAULT 0.0;
UPDATE sensors SET value = 0.0 WHERE value IS NULL;

INSERT INTO sensors (name, location, type, value, unit, status, last_updated)
VALUES
  ('Temperature Sensor 1', 'Living Room', 'temperature', 21.5, 'C', 'active', NOW()),
  ('Temperature Sensor 2', 'Bedroom', 'temperature', 19.2, 'C', 'active', NOW());
