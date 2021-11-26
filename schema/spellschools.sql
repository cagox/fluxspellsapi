create table spellschools(
    spell_id INT,
    school_id INT,
    PRIMARY KEY (spell_id, school_id),
    FOREIGN KEY (spell_id) REFERENCES spells(spell_id) ON DELETE CASCADE,
    FOREIGN KEY (school_id) REFERENCES schools(school_id) ON DELETE CASCADE
)