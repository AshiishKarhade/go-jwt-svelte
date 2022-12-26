-- CREATE USER 'auth_user'@'localhost' IDENTIFIED BY 'Auth123';
--
-- CREATE DATABASE auth;

-- GRANT ALL PRIVILEGES ON auth.* TO 'auth_user'@'localhost';

-- USE auth;

CREATE TABLE jwt_go_users (
                      id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
                      email VARCHAR(255) NOT NULL UNIQUE,
                      password VARCHAR(255) NOT NULL
);

DROP TABLE jwt_go_users;




