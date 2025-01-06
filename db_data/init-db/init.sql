CREATE TABLE IF NOT EXISTS currencies(
    cur_id INT NOT NULL,
    date DATE NOT NULL,
    cur_abbreviation VARCHAR(10) NOT NULL,
    cur_scale INT NOT NULL,
    cur_name VARCHAR(100) NOT NULL,
    cur_official_rate DECIMAL(10, 4) NOT NULL,
    PRIMARY KEY (cur_id, date)
);
