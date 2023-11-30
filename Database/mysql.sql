
CREATE DATABASE ClinicApp;

USE ClinicApp;

CREATE TABLE user
(
    id       int          NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name     VARCHAR(255) NOT NULL unique,
    email    VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role     VarChar(20)  NOT NULL
);
-- Slot table
CREATE TABLE slot (
    id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
    doctorid INT NOT NULL REFERENCES user (id),
    date DATETIME UNIQUE NOT NULL
);

CREATE TABLE SlotWithPatient(
    id int not null primary key AUTO_INCREMENT,
    patientid Int NOT NULL REFERENCES user(id),
    slotid Int not null unique REFERENCES slot(id)
);


