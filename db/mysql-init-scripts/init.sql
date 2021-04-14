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
 user_id INT NOT NULL AUTO_INCREMENT,
 name VARCHAR(255) NOT NULL,
 email VARCHAR(200) NOT NULL,
 rating INT,
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,

 PRIMARY KEY (user_id)
);

CREATE TABLE service_providers (
 service_id INT NOT NULL AUTO_INCREMENT,
 name VARCHAR(255) NOT NULL,
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,

 PRIMARY KEY (service_id)
);
 
CREATE TABLE plans (
 service_id INT,
 plan_name VARCHAR(255) NOT NULL,
 cost INT NOT NULL,
 detail VARCHAR(255) NOT NULL,
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,

 PRIMARY KEY (service_id, plan_name),
 FOREIGN KEY (service_id) REFERENCES service_providers(service_id)
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
 FOREIGN KEY (admin_id) REFERENCES users(user_id),
 FOREIGN KEY (service_id) REFERENCES service_providers(service_id),
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
 FOREIGN KEY (user_id) REFERENCES users(user_id),
 FOREIGN KEY (room_id) REFERENCES rooms(room_id)
);

CREATE TABLE invitation (
 user_id INT,
 room_id INT,
 is_accepted BOOLEAN NOT NULL,

 PRIMARY KEY (user_id, room_id),
 FOREIGN KEY (user_id) REFERENCES users(user_id),
 FOREIGN KEY (room_id) REFERENCES rooms(room_id)
);
