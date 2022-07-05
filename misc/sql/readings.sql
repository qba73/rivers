CREATE TABLE IF NOT EXISTS waterlevel_readings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    station_id INT NOT NULL,
    station_name CHAR(50) NOT NULL,
    datetime TEXT NOT NULL,
    value INTEGER
);
