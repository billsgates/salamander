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
 rating INT DEFAULT 0 NOT NULL,
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
 room_status VARCHAR(255) NOT NULL DEFAULT 'created',
 created_at TIMESTAMP NOT NULL DEFAULT current_timestamp(),
 updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
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
 payment_status VARCHAR(255) NOT NULL DEFAULT 'unpaid',
 created_at TIMESTAMP NOT NULL DEFAULT current_timestamp(),
 updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
 is_host BOOLEAN NOT NULL,

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

CREATE TABLE rounds (
 room_id INT,
 starting_time TIMESTAMP NOT NULL,
 round_interval INT NOT NULL,
 payment_deadline INT NOT NULL,
 is_add_to_google_calendar BOOLEAN NOT NULL DEFAULT false,
 created_at TIMESTAMP NOT NULL DEFAULT current_timestamp(),
 updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),

 PRIMARY KEY (room_id, starting_time),
 FOREIGN KEY (room_id) REFERENCES rooms(room_id) ON DELETE CASCADE
);

-- password = 'kevin'
INSERT INTO users (name, email, password_digest) VALUES ('Kevin Yu', 'kevin@ntu.im', '61d1f2b7264c447dcdb110f233551e5c51520d5f');
-- password = 'frank'
INSERT INTO users (name, email, password_digest) VALUES ('Frank Chen', 'frank@ntu.im', '9fe148e76ff638747e0e5ca03c28b1391f7597fe');
-- password = 'paul'
INSERT INTO users (name, email, password_digest) VALUES ('Paul Liu', 'paul@ntu.im', 'c955b83937bb4c3f875e093ae14038867ac35493');
-- password = 'jason'
INSERT INTO users (name, email, password_digest) VALUES ('Jason Wang', 'jason@ntu.im', 'e1d20ac5d01c96892298f5f92539d41ebdd28a18');

INSERT INTO service_providers (name) VALUES ('Neflix');
INSERT INTO service_providers (name) VALUES ('Youtube Premium');
INSERT INTO service_providers (name) VALUES ('Spotify');

-- Netflix Plans
INSERT INTO plans (service_id, plan_name, cost, max_count) VALUES ('1', 'Basic', 270, 1);
INSERT INTO plans (service_id, plan_name, cost, max_count) VALUES ('1', 'Standard', 330, 2);
INSERT INTO plans (service_id, plan_name, cost, max_count) VALUES ('1', 'Premium', 390, 4);
-- YouTube Preium Plans
INSERT INTO plans (service_id, plan_name, cost, max_count) VALUES ('2', 'Student', 109, 1);
INSERT INTO plans (service_id, plan_name, cost, max_count) VALUES ('2', 'Individual', 179, 1);
INSERT INTO plans (service_id, plan_name, cost, max_count) VALUES ('2', 'Family', 269, 6);
-- Spotify Plans
INSERT INTO plans (service_id, plan_name, cost, max_count) VALUES ('3', 'Individual', 149, 1);
INSERT INTO plans (service_id, plan_name, cost, max_count) VALUES ('3', 'Duo', 198, 2);
INSERT INTO plans (service_id, plan_name, cost, max_count) VALUES ('3', 'Family', 240, 6);

INSERT INTO `rooms` (`room_id`, `announcement`, `is_public`, `max_count`, `admin_id`, `service_id`, `plan_name`) VALUES
(1,	NULL,	1,	4,	1,	1,	'Premium');

INSERT INTO `invitation_codes` (`room_id`, `invitation_code`, `is_valid`, `created_at`, `updated_at`) VALUES
(1,	'15a447a',	0,	'2021-05-06 07:01:52',	'2021-05-06 07:10:19'),
(1,	'4242227',	1,	'2021-05-06 07:01:48',	'2021-05-06 07:01:48'),
(1,	'8b88a7c',	0,	'2021-05-06 07:01:52',	'2021-05-06 07:02:45'),
(1,	'910891a',	0,	'2021-05-06 07:01:50',	'2021-05-06 07:09:00'),
(1,	'9c7644e',	1,	'2021-05-06 07:01:51',	'2021-05-06 07:01:51'),
(1,	'9f95495',	1,	'2021-05-06 07:01:51',	'2021-05-06 07:01:51');

INSERT INTO `participation` (`user_id`, `room_id`, `payment_status`, `created_at`, `updated_at`, `is_host`) VALUES
(1,	1,	'confirmed',	'2021-05-06 07:01:20',	'2021-05-06 07:01:20',	1),
(2,	1,	'unpaid',	'2021-05-06 07:02:46',	'2021-05-06 07:02:46',	0),
(3,	1,	'unpaid',	'2021-05-06 07:10:20',	'2021-05-06 07:10:20',	0),
(4,	1,	'unpaid',	'2021-05-06 07:09:00',	'2021-05-06 07:09:00',	0);