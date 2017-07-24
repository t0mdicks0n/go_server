CREATE TABLE IF NOT EXISTS chats (
  id SERIAL PRIMARY KEY,
  timestamp TIMESTAMP default current_timestamp,
  username VARCHAR(40),
  message VARCHAR(500),
  groupname VARCHAR(40)
);

INSERT INTO chats(username, message, groupname) VALUES (
	'Tom', 'Hey everybody from psql', 'lobby'
);

INSERT INTO chats(username, message, groupname) VALUES (
	'Ebba', 'This is Toms sister Ebba', 'lobby'
);