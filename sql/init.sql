GRANT ALL PRIVILEGES ON *.* TO 'root' @'%';

USE `jogo-do-bicho`;

CREATE TABLE IF NOT EXISTS gambler (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(50) NOT NULL,
    document VARCHAR(11) NOT NULL,
    document_type ENUM('cpf', 'rg') NOT NULL,
    birth_date DATETIME NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS bet_group (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    group_name VARCHAR(30) NOT NULL,
    group_number VARCHAR(2) NOT NULL,
    dozen INT NOT NULL
);

CREATE TABLE IF NOT EXISTS raffle (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `edition` INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS raffle_draw (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    raffle_id BIGINT NOT NULL,
    `order` INT NOT NULL,
    `number` VARCHAR(4) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_raffle_draw_raffle FOREIGN KEY (raffle_id) REFERENCES raffle(id)
);

CREATE TABLE IF NOT EXISTS bet (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    gambler_id BIGINT NOT NULL,
    raffle_id BIGINT NOT NULL,
    bet_type ENUM(
        'thousands',
        'hundreds',
        'dozens',
        'group',
        'double_dozens,',
        'double_group'
    ) NOT NULL,
    bet_modifier ENUM('on_top', 'surrounded') NOT NULL,
    bet_price FLOAT NOT NULL,
    bet_combination VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_bet_gambler FOREIGN KEY (gambler_id) REFERENCES gambler(id),
    CONSTRAINT fk_bet_raffle FOREIGN KEY (raffle_id) REFERENCES raffle(id)
);

INSERT INTO
    bet_group (group_name, group_number, dozen)
VALUES
    ('AVESTRUZ', 1, "01"),
    ('AVESTRUZ', 1, "02"),
    ('AVESTRUZ', 1, "03"),
    ('AVESTRUZ', 1, "04"),
    ('ÁGUIA', 2, "05"),
    ('ÁGUIA', 2, "06"),
    ('ÁGUIA', 2, "07"),
    ('ÁGUIA', 2, "08"),
    ('BURRO', 3, "09"),
    ('BURRO', 3, "10"),
    ('BURRO', 3, "11"),
    ('BURRO', 3, "12"),
    ('BORBOLETA', 4, "13"),
    ('BORBOLETA', 4, "14"),
    ('BORBOLETA', 4, "15"),
    ('BORBOLETA', 4, "16"),
    ('CACHORRO', 5, "17"),
    ('CACHORRO', 5, "18"),
    ('CACHORRO', 5, "19"),
    ('CACHORRO', 5, "20"),
    ('CABRA', 6, "21"),
    ('CABRA', 6, "22"),
    ('CABRA', 6, "23"),
    ('CABRA', 6, "24"),
    ('CARNEIRO', 7, "25"),
    ('CARNEIRO', 7, "26"),
    ('CARNEIRO', 7, "27"),
    ('CARNEIRO', 7, "28"),
    ('CAMELO', 8, "29"),
    ('CAMELO', 8, "30"),
    ('CAMELO', 8, "31"),
    ('CAMELO', 8, "32"),
    ('COBRA', 9, "33"),
    ('COBRA', 9, "34"),
    ('COBRA', 9, "35"),
    ('COBRA', 9, "36"),
    ('COELHO', 10, "37"),
    ('COELHO', 10, "38"),
    ('COELHO', 10, "39"),
    ('COELHO', 10, "40"),
    ('CAVALO', 11, "41"),
    ('CAVALO', 11, "42"),
    ('CAVALO', 11, "43"),
    ('CAVALO', 11, "44"),
    ('ELEFANTE', 12, "45"),
    ('ELEFANTE', 12, "46"),
    ('ELEFANTE', 12, "47"),
    ('ELEFANTE', 12, "48"),
    ('GALO', 13, "49"),
    ('GALO', 13, "50"),
    ('GALO', 13, "51"),
    ('GALO', 13, "52"),
    ('GATO', 14, "53"),
    ('GATO', 14, "54"),
    ('GATO', 14, "55"),
    ('GATO', 14, "56"),
    ('JACARÉ', 15, "57"),
    ('JACARÉ', 15, "58"),
    ('JACARÉ', 15, "59"),
    ('JACARÉ', 15, "60"),
    ('LEÃO', 16, "61"),
    ('LEÃO', 16, "62"),
    ('LEÃO', 16, "63"),
    ('LEÃO', 16, "64"),
    ('MACACO', 17, "65"),
    ('MACACO', 17, "66"),
    ('MACACO', 17, "67"),
    ('MACACO', 17, "68"),
    ('PORCO', 18, "69"),
    ('PORCO', 18, "70"),
    ('PORCO', 18, "71"),
    ('PORCO', 18, "72"),
    ('PAVÃO', 19, "73"),
    ('PAVÃO', 19, "74"),
    ('PAVÃO', 19, "75"),
    ('PAVÃO', 19, "76"),
    ('PERU', 20, "77"),
    ('PERU', 20, "78"),
    ('PERU', 20, "79"),
    ('PERU', 20, "80"),
    ('TOURO', 21, "81"),
    ('TOURO', 21, "82"),
    ('TOURO', 21, "83"),
    ('TOURO', 21, "84"),
    ('TIGRE', 22, "85"),
    ('TIGRE', 22, "86"),
    ('TIGRE', 22, "87"),
    ('TIGRE', 22, "88"),
    ('URSO', 23, "89"),
    ('URSO', 23, "90"),
    ('URSO', 23, "91"),
    ('URSO', 23, "92"),
    ('VEADO', 24, "93"),
    ('VEADO', 24, "94"),
    ('VEADO', 24, "95"),
    ('VEADO', 24, "96"),
    ('VACA', 25, "97"),
    ('VACA', 25, "98"),
    ('VACA', 25, "99"),
    ('VACA', 25, "00");

INSERT INTO
    raffle(`edition`, updated_at)
VALUES
    (1, '2024-10-29');