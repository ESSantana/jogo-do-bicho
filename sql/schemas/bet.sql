CREATE TABLE IF NOT EXISTS bets (
    id SERIAL,
    animal varchar(20) NOT NULL,
    bet_number int NOT NULL,
    bet_price decimal NOT NULL
);