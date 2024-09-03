CREATE TABLE user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(256) UNIQUE,
    username VARCHAR(50) UNIQUE,
    password VARCHAR(255),
    photo TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE level (
    id INT AUTO_INCREMENT PRIMARY KEY,
    description ENUM('eazy', 'medium', 'hard', 'upnormal')
);

CREATE TABLE team (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE role (
    id INT AUTO_INCREMENT PRIMARY KEY,
    description VARCHAR(50)
);


CREATE TABLE task (
    id INT AUTO_INCREMENT PRIMARY KEY,
    assign_user_id INT,
    title VARCHAR(255),
    description TEXT,
    level_id INT,
    weight INT,
    status ENUM('pending', 'in progress', 'completed') DEFAULT "pending",
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (assign_user_id) REFERENCES user(id),
    FOREIGN KEY (level_id) REFERENCES level(id),
    INDEX (assign_user_id),
    INDEX (level_id)
);



CREATE TABLE member (
    id INT AUTO_INCREMENT PRIMARY KEY,
    team_id INT,
    user_id INT,
    role_id INT,
    FOREIGN KEY (team_id) REFERENCES team(id),
    FOREIGN KEY (user_id) REFERENCES user(id),
    FOREIGN KEY (role_id) REFERENCES role(id)
);




