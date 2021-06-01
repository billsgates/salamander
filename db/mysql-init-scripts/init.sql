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
DROP TABLE IF EXISTS applications;
DROP TABLE IF EXISTS invitation_codes;
DROP TABLE IF EXISTS rounds;
DROP TABLE IF EXISTS plans;
 
CREATE TABLE users (
 id INT NOT NULL AUTO_INCREMENT,
 name VARCHAR(255) NOT NULL,
 email VARCHAR(200) NOT NULL UNIQUE,
 phone VARCHAR(200),
 image_url VARCHAR(1000),
 password_digest VARCHAR(1000) NOT NULL,
 rating FLOAT DEFAULT 0 NOT NULL,
 rating_count INT DEFAULT 0 NOT NULL,
 created_at TIMESTAMP NOT NULL DEFAULT current_timestamp(),
 updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),

 PRIMARY KEY (id)
);

CREATE TABLE service_providers (
 id INT NOT NULL AUTO_INCREMENT,
 name VARCHAR(255) NOT NULL,
 created_at TIMESTAMP NOT NULL DEFAULT current_timestamp(),
 updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),

 PRIMARY KEY (id)
);

CREATE TABLE rounds (
 round_id INT NOT NULL AUTO_INCREMENT,
 starting_time TIMESTAMP NOT NULL,
 ending_time TIMESTAMP NOT NULL,
 round_interval INT NOT NULL,
 payment_deadline TIMESTAMP NOT NULL,
 is_add_calendar BOOLEAN NOT NULL DEFAULT false,
 created_at TIMESTAMP NOT NULL DEFAULT current_timestamp(),

 PRIMARY KEY (round_id)
);
 
CREATE TABLE plans (
 service_id INT,
 plan_name VARCHAR(255) NOT NULL,
 cost INT NOT NULL,
 max_count INT NOT NULL,
 created_at TIMESTAMP NOT NULL DEFAULT current_timestamp(),
 updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),

 PRIMARY KEY (service_id, plan_name),
 FOREIGN KEY (service_id) REFERENCES service_providers(id)
);

CREATE TABLE rooms (
 room_id INT NOT NULL AUTO_INCREMENT,
 announcement VARCHAR(1000),
 is_public BOOLEAN NOT NULL,
 room_status VARCHAR(255) DEFAULT 'created',
 round_id INT NULL,
 matching_deadline TIMESTAMP NULL,
 public_message VARCHAR(1000),
 created_at TIMESTAMP NOT NULL DEFAULT current_timestamp(),
 updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
 max_count INT NOT NULL,
 admin_id INT,
 service_id INT,
 plan_name VARCHAR(255),

 PRIMARY KEY (room_id),
 FOREIGN KEY (round_id) REFERENCES rounds(round_id) ON DELETE SET NULL,
 FOREIGN KEY (admin_id) REFERENCES users(id),
 FOREIGN KEY (service_id) REFERENCES service_providers(id),
 FOREIGN KEY (service_id, plan_name) REFERENCES plans(service_id, plan_name)
);

CREATE TABLE participation (
 user_id INT,
 room_id INT,
 payment_status VARCHAR(255) NOT NULL DEFAULT 'unpaid',
 created_at TIMESTAMP NOT NULL DEFAULT current_timestamp(),
 updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
 is_host BOOLEAN NOT NULL,

 PRIMARY KEY (user_id, room_id),
 FOREIGN KEY (user_id) REFERENCES users(id),
 FOREIGN KEY (room_id) REFERENCES rooms(room_id) ON DELETE CASCADE
);

CREATE TABLE applications (
 user_id INT,
 room_id INT,
 application_message VARCHAR(1000),
 is_accepted BOOLEAN NOT NULL DEFAULT false,
 created_at TIMESTAMP NOT NULL DEFAULT current_timestamp(),
 updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),

 PRIMARY KEY (user_id, room_id),
 FOREIGN KEY (user_id) REFERENCES users(id),
 FOREIGN KEY (room_id) REFERENCES rooms(room_id) ON DELETE CASCADE
);

CREATE TABLE invitation_codes (
 room_id INT,
 invitation_code VARCHAR(255) NOT NULL,
 is_valid BOOLEAN NOT NULL DEFAULT true,
 created_at TIMESTAMP NOT NULL DEFAULT current_timestamp(),
 updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),

 PRIMARY KEY (room_id, invitation_code),
 FOREIGN KEY (room_id) REFERENCES rooms(room_id) ON DELETE CASCADE
);


INSERT INTO users (name, email, password_digest, rating, rating_count) VALUES
-- password = 'kevin'
('Kevin Yu', 'kevin.ct.yu@ntu.im', '61d1f2b7264c447dcdb110f233551e5c51520d5f', 0, 0),
-- password = 'frank'
('Frank Chen', 'frankchen93011@gmail.com', '9fe148e76ff638747e0e5ca03c28b1391f7597fe', 0, 0),
-- password = 'paul'
('Paul Liu', 'ee91941387ee@gmail.com', 'c955b83937bb4c3f875e093ae14038867ac35493', 0, 0),
-- password = 'jason'
-- ('Jason Wang', 'jason@ntu.im', 'e1d20ac5d01c96892298f5f92539d41ebdd28a18', 0, 0),
-- password = 'carolyn'
('Carolyn Chen', 'cinnacinna130@ntu.im', 'bc0323299a3f5ad96fd10392a39881a5d3b63ba5', 0, 0),
-- password = 'angela'
('Angela Lee', 'Angyeahyeah6@gmail.com', '957a28d0ff00dce48cfa8187d49df4a1bc71539c', 4.5, 2),
-- password = 'zhen'
('Zhen', 'battle1631@ntu.im', 'd6fcf6ca364a054f2cbdcb008fe9337412a4c2b3', 0, 0);

INSERT INTO service_providers (name) VALUES
('Netflix'),
('Youtube Premium'),
('Spotify');

INSERT INTO plans (service_id, plan_name, cost, max_count) VALUES
-- Netflix Plans
('1', 'Basic', 270, 1),
('1', 'Standard', 330, 2),
('1', 'Premium', 390, 4),
-- YouTube Preium Plans
('2', 'Student', 109, 1),
('2', 'Individual', 179, 1),
('2', 'Family', 269, 6),
-- Spotify Plans
('3', 'Individual', 149, 1),
('3', 'Duo', 198, 2),
('3', 'Family', 240, 6);

INSERT INTO rounds (round_id, starting_time, ending_time, round_interval, payment_deadline, is_add_calendar) VALUES
(1,	'2020-08-01 00:00:00',	'2021-08-01 00:00:00',	12,	'2020-07-25 00:00:00',	1);

INSERT INTO rooms (room_id, announcement, is_public, room_status, round_id, matching_deadline, public_message, max_count, admin_id, service_id, plan_name) VALUES
(1,	'spotify55688/angela',	0,	'start',	1,	NULL,	'',	6,	5,	3,	'Family'),
(2,	'',	1,	'created',	NULL,	'2021-06-17 00:00:00',	'Welcome! This is Kevins Netflix Premium room!',	4,	1,	1,	'Premium'),
(3,	'',	1,	'created',	NULL,	'2021-07-05 00:00:00',	'Franks Youtube Premium room',	6,	2,	2,	'Family'),
(4,	'',	1,	'created',	NULL,	'2021-07-01 00:00:00',	'Hi! welcome to Carolyns Spotify Family room',	6,	5,	3,	'Family');

INSERT INTO participation (user_id, room_id, payment_status, is_host) VALUES
(1,	1,	'unpaid',	0),
(1,	2,	'confirmed',	1),
(2,	1,	'unpaid',	0),
(2,	3,	'confirmed',	1),
(5,	1,	'unpaid',	0),
(5,	4,	'confirmed',	1),
(6,	1,	'confirmed',	1);

INSERT INTO invitation_codes (room_id, invitation_code, is_valid) VALUES
(1,	'5ad0ce4',	0),
(1,	'75ddb97',	0),
(1,	'9bfb3b9',	0);

