CREATE DATABASE hermitcrab_db;
USE hermitcrab_db;

CREATE USER 'hermitcrab-admin'@'%' IDENTIFIED BY 'admin-password';
CREATE USER 'hermitcrab-user'@'%' IDENTIFIED BY 'user-password';

GRANT ALL PRIVILEGES ON hermitcrab_db.* TO 'hermitcrab-admin'@'%';
GRANT SELECT, INSERT, DELETE ON hermitcrab_db.* TO 'hermitcrab-user'@'%';
 
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS rooms;
DROP TABLE IF EXISTS service_providers;
DROP TABLE IF EXISTS participation;
DROP TABLE IF EXISTS invitation;
DROP TABLE IF EXISTS plans;
 
CREATE TABLE users (
 id INT NOT NULL AUTO_INCREMENT,
 name VARCHAR(255) NOT NULL,
 email VARCHAR(200) NOT NULL,
 password_digest VARCHAR(1000) NOT NULL,
 rating INT,
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,

 PRIMARY KEY (id)
);

CREATE TABLE service_providers (
 id INT NOT NULL AUTO_INCREMENT,
 name VARCHAR(255) NOT NULL,
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,

 PRIMARY KEY (id)
);
 
CREATE TABLE plans (
 service_id INT,
 plan_name VARCHAR(255) NOT NULL,
 cost INT NOT NULL,
 detail VARCHAR(255) NOT NULL,
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,

 PRIMARY KEY (service_id, plan_name),
 FOREIGN KEY (service_id) REFERENCES service_providers(id)
);

CREATE TABLE rooms (
 room_id INT NOT NULL AUTO_INCREMENT,
 title VARCHAR(255) NOT NULL,
 account_name VARCHAR(255) NOT NULL,
 account_password VARCHAR(1000) NOT NULL,
 starting_time TIMESTAMP NOT NULL,
 ending_time TIMESTAMP NOT NULL,
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,
 payment_interval INT NOT NULL,
 admin_id INT NOT NULL,
 service_id INT NOT NULL,
 plan_name VARCHAR(255) NOT NULL,

 PRIMARY KEY (room_id),
 FOREIGN KEY (admin_id) REFERENCES users(id),
 FOREIGN KEY (service_id) REFERENCES service_providers(id),
 FOREIGN KEY (service_id, plan_name) REFERENCES plans(service_id, plan_name)
);

CREATE TABLE participation (
 user_id INT,
 room_id INT,
 joined_at TIMESTAMP NOT NULL,
 left_at TIMESTAMP,
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,
 is_host BOOLEAN NOT NULL,

 PRIMARY KEY (user_id, room_id),
 FOREIGN KEY (user_id) REFERENCES users(id),
 FOREIGN KEY (room_id) REFERENCES rooms(room_id)
);

CREATE TABLE invitation (
 user_id INT,
 room_id INT,
 is_accepted BOOLEAN NOT NULL,

 PRIMARY KEY (user_id, room_id),
 FOREIGN KEY (user_id) REFERENCES users(id),
 FOREIGN KEY (room_id) REFERENCES rooms(room_id)
);

INSERT INTO users (name, email, email, rating, created_at, updated_at) VALUES ('Kevin Yu', 'kevin@ntu.im', '$2a$10$gVtjNk4YL.O4I//ZBtvfN.YEebwR1Ci3.5OBHan4PWFzniSFqpzce', 3, NOW(), NOW());
INSERT INTO users (name, email, email, rating, created_at, updated_at) VALUES ('Frank Chen', 'frank@ntu.im', '$2a$10$6tsb.2dRzV5gSTEJmtwkgeKpPIMO0VbMv2E6hP9xuAytwFlf0trVm', 5, NOW(), NOW());
INSERT INTO users (name, email, email, rating, created_at, updated_at) VALUES ('Paul Liu', 'paul@ntu.im', '$2a$10$WkWwIpCbMyB1A2OuMC9LI.4LtQZtxNb1djcYqzeP0IayazJQgVkH', 4, NOW(), NOW());
