CREATE DATABASE hermit_crab_db;
USE hermit_crab_db;
 
DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS room;
DROP TABLE IF EXISTS service_provider;
DROP TABLE IF EXISTS participation;
DROP TABLE IF EXISTS invitation;
DROP TABLE IF EXISTS plan;
 
CREATE TABLE user (
 user_id INT AUTO_INCREMENT,
 name VARCHAR(255) NOT NULL,
 email VARCHAR(200) NOT NULL,
 rating INT,
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,

 PRIMARY KEY (user_id)
)
 
CREATE TABLE room (
 room_id INT AUTO_INCREMENT,
 title VARCHAR(255) NOT NULL,
 account_name VARCHAR(255) NOT NULL,
 account_password VARCHAR(1000) NOT NULL,
 starting_time TIMESTAMP NOT NULL,
 ending_time TIMESTAMP NOT NULL,
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,
 interval INT NOT NULL,
 admin_id INT NOT NULL,
 service_id INT NOT NULL,
 plan_id INT NOT NULL,

 PRIMARY KEY (room_id),
 FOREIGN KEY (admin_id) REFERENCES user(user_id),
 FOREIGN KEY (service_id) REFERENCES service_provider(service_id),
 FOREIGN KEY (service_id, plan_id) REFERENCES plan(service_id, plan_id)
)

CREATE TABLE service_provider (
 service_id INT AUTO_INCREMENT,
 name VARCHAR(255) NOT NULL,
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,

 PRIMARY KEY (service_id)
)

CREATE TABLE participation (
 user_id INT,
 room_id INT,
 joined_at TIMESTAMP NOT NULL,
 left_at TIMESTAMP,
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,
 is_host BOOLEAN NOT NULL,

 PRIMARY KEY (user_id, room_id),
 FOREIGN KEY (user_id) REFERENCES user(user_id),
 FOREIGN KEY (room_id) REFERENCES room(room_id)
)

CREATE TABLE invitation (
 user_id INT,
 room_id INT,
 is_accepted BOOLEAN NOT NULL,

 PRIMARY KEY (user_id, room_id),
 FOREIGN KEY (user_id) REFERENCES user(user_id),
 FOREIGN KEY (room_id) REFERENCES user(room_id)
)

CREATE TABLE plan (
 plan_id INT AUTO_INCREMENT,
 service_id INT,
 cost INT NOT NULL,
 detail VARCHAR(255) NOT NULL,
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,

 PRIMARY KEY (service_id, plan_id),
 FOREIGN KEY (service_id) REFERENCES service_provider(service_id)
)
