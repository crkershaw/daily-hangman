DROP TABLE IF EXISTS wordlist;
CREATE TABLE wordlist (
    id_num	int NOT NULL AUTO_INCREMENT,
    wordlist_name	VARCHAR(128) NOT NULL,
    word_num	INT NOT NULL,
    word	varchar(255),
    message	varchar(255),
    creation_date	datetime,
    PRIMARY KEY (id_num)
);