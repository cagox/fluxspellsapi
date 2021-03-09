CREATE TABLE IF NOT EXISTS spells(
    spell_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    spellname VARCHAR(90) CHARACTER SET utf8mb4 NOT NULL UNIQUE,
    spelldescription TEXT NOT NULL,
    summary TEXT NOT NULL,
    cost varchar(90) CHARACTER SET utf8mb4 NOT NULL,
    difficulty varchar(90) CHARACTER SET utf8mb4 NOT NULL,
    prerequisites varchar(255) CHARACTER SET utf8mb4 NOT NULL,
    components varchar(255) CHARACTER SET utf8mb4 NOT NULL,
    spellrange varchar(90) CHARACTER SET utf8mb4 NOT NULL
);
