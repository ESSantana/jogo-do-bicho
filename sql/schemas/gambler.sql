CREATE TABLE IF NOT EXISTS gamblers (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  gambler_name VARCHAR(50) NOT NULL,
  document VARCHAR(20) NOT NULL,
  document_type ENUM('CPF', 'RG') NOT NULL,
  birth_date DATETIME NOT NULL,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);