create table session_tokens(
    id BIGINT AUTO_INCREMENT UNIQUE NOT NULL PRIMARY KEY,
    user_id INT NOT NULL,
    token varchar(255) NOT NULL UNIQUE,
    date_created DATETIME NOT NULL DEFAULT NOW(),
    expiration_date DATETIME NOT NULL DEFAULT ADDDATE(date_created,1),
    is_expired BOOLEAN DEFAULT false,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
)