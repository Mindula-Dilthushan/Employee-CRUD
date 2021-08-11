DROP DATABASE employeegolang;
CREATE DATABASE employeegolang;
USE employeegolang;

CREATE TABLE employee(
    nic varchar (12) NOT NULL,
    fullname varchar (40) NOT NULL,
    address varchar (40) NOT NULL,
    mobile DECIMAL (10) NOT NULL,
    CONSTRAINT PRIMARY KEY (nic)
);

SELECT * FROM employee;
