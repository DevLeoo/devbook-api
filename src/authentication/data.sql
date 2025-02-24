 insert into users (name, nick, email, pwd)
 values 
 ("User 1", "user_1", "user_1@gmail.com", ""),
 ("User 2", "user_2", "user_2@gmail.com", ""),
 ("User 3", "user_3", "user_3@gmail.com", "");

 insert into followers (user_id, follower_id)
 values 
 (1, 2),
 (2, 3),
 (3, 1);

insert into posts(title, content, author_id) 
values
("Publicacao do usuario 1", "essa aqui é a publicacao do usuario 1", 1),
("Publicacao do usuario 4", "essa aqui é a publicacao do usuario 4", 4),
("Publicacao do usuario 3", "essa aqui é a publicacao do usuario 3", 3);
