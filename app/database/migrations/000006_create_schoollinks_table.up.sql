CREATE TABLE IF NOT EXISTS typelinks(
    school_id INT NOT NULL,
    spell_id INT NOT NULL,
    FOREIGN KEY (school_id) REFERENCES schools (school_id) ON DELETE CASCADE,
    FOREIGN KEY (spell_id) REFERENCES spells (spell_id) ON DELETE CASCADE,
    PRIMARY KEY (school_id, spell_id)
    );