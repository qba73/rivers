CREATE TABLE IF NOT EXISTS temperature_readings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    station_id INT NOT NULL,
    station_name CHAR(50) NOT NULL,
    sensor_ref CHAR(20) NOT NULL,
    datetime TEXT NOT NULL,
    value REAL
);

INSERT INTO temperature_readings (station_id, station_name, sensor_ref, datetime, value) VALUES
(1042, "Sandy Millss", "0001", "2022-06-28T04:45:00Z", 0.384),
(1043, "Ballybofey", "0001", "2022-06-28T04:14:00Z", 1.679),
(3055, "Glaslough", "0001", "2022-06-28T04:45:00Z", 0.478);
