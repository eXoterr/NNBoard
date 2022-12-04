CREATE TABLE `Users` (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  login VARCHAR(60),
  password VARCHAR(90),
  role INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE `Posts` (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title VARCHAR(60) NOT NULL,
  body TEXT NOT NULL,
  author INTEGER, 
  FOREIGN KEY(author) REFERENCES Users(id)
);
