CREATE TABLE IF NOT EXISTS people(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name VARCHAR (128) NOT NULL,
    last_name VARCHAR(128) NOT NULL,
    title VARCHAR(128) DEFAULT NULL,
    Company VARCHAR(128) DEFAULT NULL,
    Age INT DEFAULT NULL,
    Married BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE VIRTUAL TABLE searchable_people USING fts5(first_name, last_name, title, company, content=people, content_rowid=id);

CREATE TRIGGER people_ai AFTER INSERT ON people BEGIN
  INSERT INTO searchable_people(rowid, first_name, last_name, title, company) VALUES (new.id, new.first_name, new.last_name, new.title, new.company);
END;
CREATE TRIGGER people_ad AFTER DELETE ON people BEGIN
  INSERT INTO searchable_people(fts_idx, rowid, first_name, last_name, title, company) VALUES('delete', old.id, old.first_name, old.last_name, old.title, old.company);
END;
CREATE TRIGGER people_au AFTER UPDATE ON people BEGIN
  INSERT INTO searchable_people(fts_idx, rowid, first_name, last_name, title, company) VALUES('delete', old.id, old.first_name, old.last_name, old.title, old.company);
  INSERT INTO searchable_people(rowid, first_name, last_name, title, company) VALUES (new.id, new.first_name, new.last_name, new.title, new.company);
END;
