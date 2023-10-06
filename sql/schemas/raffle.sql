CREATE TABLE IF NOT EXISTS raffles (
    id SERIAL PRIMARY KEY,
    raffle_edition INT NOT NULL,
    animal VARCHAR(255) NOT NULL,
    raffle_number VARCHAR(20) NOT NULL,
    raffle_order INT NOT NULL
);