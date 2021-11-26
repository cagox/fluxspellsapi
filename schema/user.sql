create table users (
    user_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    email varchar(255) NOT NULL UNIQUE,
    password_hash char(64) NOT NULL,
    is_admin BOOLEAN,
    is_verified BOOLEAN
)