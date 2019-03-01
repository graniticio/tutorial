DROP DATABASE IF EXISTS recordstore;
CREATE DATABASE recordstore;

USE recordstore;

CREATE TABLE artist (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(50) NOT NULL,
  first_active SMALLINT,
  last_active SMALLINT,
  PRIMARY KEY (id)
);

INSERT INTO artist (name, first_active) VALUES ('Younger artist', 2006);
INSERT INTO artist (name, first_active) VALUES ('Older artist', 1976);
INSERT INTO artist (name, first_active, last_active) VALUES ('Retired artist', 1951, 2010);

CREATE TABLE related_artist (
  artist_id INT NOT NULL,
  related_artist_id INT NOT NULL,
  FOREIGN KEY (artist_id) REFERENCES artist(id),
  FOREIGN KEY (related_artist_id) REFERENCES artist(id),
  UNIQUE (artist_id, related_artist_id)
);


CREATE USER IF NOT EXISTS 'grnc'@'%' IDENTIFIED BY 'OKnasd8!k';
GRANT ALL PRIVILEGES ON recordstore.* TO 'grnc'@'%' WITH GRANT OPTION;

