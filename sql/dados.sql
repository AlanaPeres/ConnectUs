insert into usuarios (nome, nick, email, senha)
values
("Alana Lima Peres", "alana", "alana@mail.com", "$2a$10$5sBCeV770HrBKa1PkfJTLOvUJEVJRJrIRCdkmWdwePSyzlPri7bWm"),
("Douglas Peres", "douglas", "douglas@gmail.com", "$2a$10$5sBCeV770HrBKa1PkfJTLOvUJEVJRJrIRCdkmWdwePSyzlPri7bWm"),
("Ana Julia Peres", "ana", "ana@gmail.com", "$2a$10$5sBCeV770HrBKa1PkfJTLOvUJEVJRJrIRCdkmWdwePSyzlPri7bWm");
insert into seguidores(usuario_id, seguidor_id)
values
(1, 3),
(3, 1),
(1, 2);

insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicaçao do usuário 1", "publicação do usuário um!", 1),
("Publicaçao do usuário 2", "publicação do usuário dois!", 2),
("Publicaçao do usuário 3", "publicação do usuário três!", 3);