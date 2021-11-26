create table spells (
    spell_id INT UNIQUE AUTO_INCREMENT PRIMARY KEY,
    name varchar(255) NOT NULL UNIQUE,
    cost varchar(32) NOT NULL,
    difficulty varchar(32) NOT NULL DEFAULT ('Variable'),
    spellrange varchar(32) NOT NULL DEFAULT('Personal'),
    prerequisites varchar(255) NOT NULL DEFAULT(''),
    ability_score_id INT NOT NULL,
    summary TEXT NOT NULL,
    description TEXT NOT NULL,
    FOREIGN KEY (ability_score_id) REFERENCES ability_scores(ability_score_id) ON DELETE SET NULL
)