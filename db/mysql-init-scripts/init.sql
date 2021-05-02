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
DROP TABLE IF EXISTS plans;
 
CREATE TABLE users (
 id INT NOT NULL AUTO_INCREMENT,
 name VARCHAR(255) NOT NULL,
 email VARCHAR(200) NOT NULL,
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
 payment_period INT NOT NULL,
 status_type VARCHAR(255) NOT NULL,
 starting_time TIMESTAMP,
 ending_time TIMESTAMP,
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
 status_type VARCHAR(255) NOT NULL,
 joined_at TIMESTAMP NOT NULL DEFAULT current_timestamp(),
 left_at TIMESTAMP,
 created_at TIMESTAMP NOT NULL DEFAULT current_timestamp(),
 updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
 is_host BOOLEAN NOT NULL,

 PRIMARY KEY (user_id, room_id),
 FOREIGN KEY (user_id) REFERENCES users(id),
 FOREIGN KEY (room_id) REFERENCES rooms(room_id)
);

CREATE TABLE invitation_codes (
 room_id INT,
 invitation_code VARCHAR(255) NOT NULL,
 is_valid BOOLEAN NOT NULL DEFAULT true,
 created_at TIMESTAMP NOT NULL DEFAULT current_timestamp(),
 updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),

 PRIMARY KEY (room_id, invitation_code),
 FOREIGN KEY (room_id) REFERENCES rooms(room_id)
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
