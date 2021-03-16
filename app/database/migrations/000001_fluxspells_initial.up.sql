CREATE TABLE IF NOT EXISTS users(
    user_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    email varchar(255) CHARACTER SET utf8mb4 UNIQUE NOT NULL,
    passwordhash char(60) NOT NULL,
    isadmin BOOL DEFAULT FALSE,
    lastupdated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    timecreated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS spells(
    spell_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(90) CHARACTER SET utf8mb4 NOT NULL UNIQUE,
    description TEXT NOT NULL,
    summary TEXT NOT NULL,
    cost varchar(90) CHARACTER SET utf8mb4 NOT NULL,
    difficulty varchar(90) CHARACTER SET utf8mb4 NOT NULL,
    prerequisites varchar(255) CHARACTER SET utf8mb4 NOT NULL,
    components varchar(255) CHARACTER SET utf8mb4 NOT NULL,
    spellrange varchar(90) CHARACTER SET utf8mb4 NOT NULL
);

CREATE TABLE IF NOT EXISTS schools(
    school_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(90) CHARACTER SET utf8mb4 NOT NULL UNIQUE,
    description TEXT NOT NULL,
    summary TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS types (
    type_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(90) CHARACTER SET utf8mb4 NOT NULL UNIQUE,
    description TEXT NOT NULL,
    summary TEXT NOT NULL
);

  
CREATE TABLE IF NOT EXISTS type_links(
    type_id INT NOT NULL,
    spell_id INT NOT NULL,
    FOREIGN KEY (type_id) REFERENCES types (type_id) ON DELETE CASCADE,
    FOREIGN KEY (spell_id) REFERENCES spells (spell_id) ON DELETE CASCADE,
    PRIMARY KEY (type_id, spell_id)
);


CREATE TABLE IF NOT EXISTS school_links(
    school_id INT NOT NULL,
    spell_id INT NOT NULL,
    FOREIGN KEY (school_id) REFERENCES schools (school_id) ON DELETE CASCADE,
    FOREIGN KEY (spell_id) REFERENCES spells (spell_id) ON DELETE CASCADE,
    PRIMARY KEY (school_id, spell_id)
    );
