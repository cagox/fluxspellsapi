create table categories(
    category_id INT NOT NULL UNIQUE AUTO_INCREMENT PRIMARY KEY,
    name varchar(255) NOT NULL UNIQUE,
    summary TEXT NOT NULL,
    description TEXT NOT NULL
)
