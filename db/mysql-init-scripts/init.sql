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
 rating INT DEFAULT 5 NOT NULL,
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
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,

 PRIMARY KEY (service_id, plan_name),
 FOREIGN KEY (service_id) REFERENCES service_providers(id)
);

CREATE TABLE rooms (
 room_id INT NOT NULL AUTO_INCREMENT,
 account_name VARCHAR(255),
 account_password VARCHAR(1000),
 starting_time TIMESTAMP,
 ending_time TIMESTAMP,
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,
 max_count INT NOT NULL,
 admin_id INT,
 service_id INT,
 plan_name VARCHAR(255),

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

-- password = 'passworda'
INSERT INTO users (name, email, password_digest, rating, created_at, updated_at) VALUES ('Kevin Yu', 'kevin@ntu.im', '945b49fe4bd2e575c860b01e0e7d1edbeb0d9cdf', 3, NOW(), NOW());
-- password = 'passwordb'
INSERT INTO users (name, email, password_digest, rating, created_at, updated_at) VALUES ('Frank Chen', 'frank@ntu.im', '07a39ab6f7caf13cc9a9426f1b8fe6048f6c1708', 4, NOW(), NOW());
-- password = 'passwordc'
INSERT INTO users (name, email, password_digest, rating, created_at, updated_at) VALUES ('Paul Liu', 'paul@ntu.im', 'a908c4a44c85cc33a1a5ea74bce2948c89318d52', 5, NOW(), NOW());

INSERT INTO service_providers (name, created_at, updated_at) VALUES ('Neflix', NOW(), NOW());
INSERT INTO service_providers (name, created_at, updated_at) VALUES ('Youtube Premium', NOW(), NOW());
INSERT INTO service_providers (name, created_at, updated_at) VALUES ('Spotify', NOW(), NOW());

INSERT INTO plans (service_id, plan_name, cost, created_at, updated_at) VALUES ('1', 'Basic', 270, NOW(), NOW());
INSERT INTO plans (service_id, plan_name, cost, created_at, updated_at) VALUES ('1', 'Standard', 330, NOW(), NOW());
INSERT INTO plans (service_id, plan_name, cost, created_at, updated_at) VALUES ('1', 'Premium', 390, NOW(), NOW());
