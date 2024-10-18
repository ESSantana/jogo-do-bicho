CREATE TABLE IF NOT EXISTS raffles (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    raffle_edition INT NOT NULL,
    animal VARCHAR(255) NOT NULL,
    raffle_number VARCHAR(20) NOT NULL,
    raffle_order INT NOT NULL,
    deleted_at TIMESTAMP
);