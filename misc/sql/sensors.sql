DELETE TABLE sensors;

CREATE TABLE sensors (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    sensor_type CHAR(20) NOT NULL,
    sensor_name CHAR(20) NOT NULL,
);

INSERT INTO sensors (TYPE, NAME) VALUES
("0001", "Water level"),
("0002", "Temperature"),
("0004", "Battery voltage"),
("OD", "Ordnance datum");
