CREATE TABLE products (
    id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    stock INT(11) NOT NULL,
    created_at timestamp NOT NULL DEFAULT (now()),
    updated_at timestamp DEFAULT NULL
);

INSERT INTO `products` (`id`,`name`,`stock`,`created_at`,`updated_at`)
VALUES ('1','Logitech G Pro X','10','2021-06-16 16:05:48',null);