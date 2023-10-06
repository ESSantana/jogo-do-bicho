CREATE TYPE doc_type AS ENUM ('CPF', 'RG');

CREATE TABLE IF NOT EXISTS gamblers (
    id SERIAL PRIMARY KEY,
    gambler_name VARCHAR(50) NOT NULL,
    document VARCHAR(20) NOT NULL,
    document_type doc_type NOT NULL,
    birth_date VARCHAR(10) NOT NULL
);