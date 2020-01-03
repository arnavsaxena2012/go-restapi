CREATE TABLE Book
(
	id serial NOT NULL,
	 VARCHAR(50),
	Author VARCHAR(50),
	ISBN VARCHAR(25)	

);

LOAD DATA LOCAL INFILE  '/var/lib/mysql-files/books.csv' into table Book
FIELDS TERMINATED BY ',' ENCLOSED BY '"'
LINES TERMINATED BY '\r' 
IGNORE 1 LINES;
