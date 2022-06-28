CREATE TABLE IF NOT EXISTS sensors (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    type CHAR(20) NOT NULL,
    name CHAR(20) NOT NULL
);

INSERT INTO sensors (type, name) VALUES
("0001", "Water level"),
("0002", "Temperature"),
("0004", "Battery voltage"),
("OD", "Ordnance datum");
