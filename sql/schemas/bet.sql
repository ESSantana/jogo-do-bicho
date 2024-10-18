CREATE TABLE IF NOT EXISTS bets (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    gambler_id BIGINT NOT NULL,
    bet_type VARCHAR(30) NOT NULL,
    bet_price FLOAT NOT NULL,
    bet_choice VARCHAR(50) NOT NULL,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_gambler FOREIGN KEY (gambler_id) REFERENCES gamblers(id)
);

