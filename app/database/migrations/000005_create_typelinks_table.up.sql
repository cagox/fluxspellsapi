CREATE TABLE IF NOT EXISTS typelinks(
    type_id INT NOT NULL,
    spell_id INT NOT NULL,
    FOREIGN KEY (type_id) REFERENCES types (type_id) ON DELETE CASCADE,
    FOREIGN KEY (spell_id) REFERENCES spells (spell_id) ON DELETE CASCADE,
    PRIMARY KEY (type_id, spell_id)
);