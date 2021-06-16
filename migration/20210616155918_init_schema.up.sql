CREATE TABLE products (
    id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    stock INT(11) NOT NULL,
    created_at timestamp NOT NULL DEFAULT (now()),
    updated_at timestamp DEFAULT NULL
);