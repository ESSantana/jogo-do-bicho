CREATE TABLE IF NOT EXISTS bets (
    id SERIAL PRIMARY KEY,
    gambler_id INT NOT NULL,
    bet_type VARCHAR(30) NOT NULL,
    bet_price DOUBLE PRECISION NOT NULL,
    bet_choice VARCHAR(50) NOT NULL,
    CONSTRAINT fk_gambler FOREIGN KEY (gambler_id) REFERENCES gamblers(id)
);