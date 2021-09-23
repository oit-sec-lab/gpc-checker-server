CREATE DATABASE IF NOT EXISTS testdb;
USE testdb;
CREATE TABLE IF NOT EXISTS sites
(
	id  int auto_increment,
	url varchar(256),
	gpc boolean,
	primary key (id)
);
