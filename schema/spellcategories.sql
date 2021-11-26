create table spellcategories(
    spell_id INT,
    category_id INT,
    PRIMARY KEY (spell_id, category_id),
    FOREIGN KEY (spell_id) REFERENCES spells(spell_id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(category_id) ON DELETE CASCADE
)