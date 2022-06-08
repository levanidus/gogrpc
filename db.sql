CREATE TABLE authors(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE books(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE authors_books(
    id INT AUTO_INCREMENT PRIMARY KEY,
    author_id INT,
    book_id INT,
    CONSTRAINT fk_author_id
    FOREIGN KEY (author_id) 
        REFERENCES authors(id),
    CONSTRAINT fk_book_id
    FOREIGN KEY (book_id) 
        REFERENCES books(id)    
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO authors VALUES (NULL,'Gogol');
INSERT INTO authors VALUES (NULL,'Tolstoi');
INSERT INTO authors VALUES (NULL,'A. Strugatskii');
INSERT INTO authors VALUES (NULL,'B. Strugatskii');

INSERT INTO books VALUES (NULL,'Checker');
INSERT INTO books VALUES (NULL,'Dead souls');
INSERT INTO books VALUES (NULL,'Anna Karenina');
INSERT INTO books VALUES (NULL,'War and peace');
INSERT INTO books VALUES (NULL,'Picnic');
INSERT INTO books VALUES (NULL,'Ulitka na sklone');

INSERT INTO authors_books VALUES (NULL,1,1);
INSERT INTO authors_books VALUES (NULL,1,2);
INSERT INTO authors_books VALUES (NULL,2,3);
INSERT INTO authors_books VALUES (NULL,2,4);
INSERT INTO authors_books VALUES (NULL,3,5);
INSERT INTO authors_books VALUES (NULL,3,6);
INSERT INTO authors_books VALUES (NULL,4,5);
INSERT INTO authors_books VALUES (NULL,4,6);

