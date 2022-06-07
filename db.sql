CREATE TABLE authors(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE books(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

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
);

INSERT INTO authors VALUES (NULL,'Гоголь');
INSERT INTO authors VALUES (NULL,'Толстой');
INSERT INTO authors VALUES (NULL,'А. Стругацкий');
INSERT INTO authors VALUES (NULL,'Б. Стругацкий');

INSERT INTO books VALUES (NULL,'Ревизор');
INSERT INTO books VALUES (NULL,'Мёртвые души');
INSERT INTO books VALUES (NULL,'Анна Каренина');
INSERT INTO books VALUES (NULL,'Война и мир');
INSERT INTO books VALUES (NULL,'Пикник на обочине');
INSERT INTO books VALUES (NULL,'Улитка на склоне');

INSERT INTO authors_books VALUES (NULL,1,1);
INSERT INTO authors_books VALUES (NULL,1,2);
INSERT INTO authors_books VALUES (NULL,2,3);
INSERT INTO authors_books VALUES (NULL,2,4);
INSERT INTO authors_books VALUES (NULL,3,5);
INSERT INTO authors_books VALUES (NULL,3,6);
INSERT INTO authors_books VALUES (NULL,4,5);
INSERT INTO authors_books VALUES (NULL,4,6);

