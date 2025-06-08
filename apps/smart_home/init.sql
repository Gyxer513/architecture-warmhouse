CREATE TABLE IF NOT EXISTS sensors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    type VARCHAR(32)
);

INSERT INTO sensors (name, location, type) VALUES
('Temperature Sensor 1', 'Living Room', 'temperature'),
('Temperature Sensor 2', 'Bedroom', 'temperature');
