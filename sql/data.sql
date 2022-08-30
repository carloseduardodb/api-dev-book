insert into users(id, name, nick, email, password) 
values
(30, 'user', 'user', 'user@gmail.com', '$2a$10$gtmHbEL9S5ffoW1z0tMEZ.ipMYTaSrgsjIxKbkZGKyeKFbjgeJjwO'),
(31, 'user2', 'user2', 'user2@gmail.com', '$2a$10$gtmHbEL9S5ffoW1z0tMEZ.ipMYTaSrgsjIxKbkZGKyeKFbjgeJjwO'),
(32, 'user3', 'user3', 'user3@gmail.com', '$2a$10$gtmHbEL9S5ffoW1z0tMEZ.ipMYTaSrgsjIxKbkZGKyeKFbjgeJjwO'),
(33, 'user4', 'user4', 'user4@gmail.com', '$2a$10$gtmHbEL9S5ffoW1z0tMEZ.ipMYTaSrgsjIxKbkZGKyeKFbjgeJjwO'),
(34, 'user5', 'user5', 'user5@gmail.com', '$2a$10$gtmHbEL9S5ffoW1z0tMEZ.ipMYTaSrgsjIxKbkZGKyeKFbjgeJjwO'),
(34, 'user5', 'user5', 'user5@gmail.com', '$2a$10$gtmHbEL9S5ffoW1z0tMEZ.ipMYTaSrgsjIxKbkZGKyeKFbjgeJjwO'),
(35, 'user6', 'user6', 'user6@gmail.com', '$2a$10$gtmHbEL9S5ffoW1z0tMEZ.ipMYTaSrgsjIxKbkZGKyeKFbjgeJjwO'),
(36, 'user7', 'user7', 'user7@gmail.com', '$2a$10$gtmHbEL9S5ffoW1z0tMEZ.ipMYTaSrgsjIxKbkZGKyeKFbjgeJjwO'),
(37, 'user8', 'user8', 'user8@gmail.com', '$2a$10$gtmHbEL9S5ffoW1z0tMEZ.ipMYTaSrgsjIxKbkZGKyeKFbjgeJjwO'),
(38, 'user9', 'user9', 'user9@gmail.com', '$2a$10$gtmHbEL9S5ffoW1z0tMEZ.ipMYTaSrgsjIxKbkZGKyeKFbjgeJjwO'),
(39, 'user10', 'user10', 'user10@gmail.com', '$2a$10$gtmHbEL9S5ffoW1z0tMEZ.ipMYTaSrgsjIxKbkZGKyeKFbjgeJjwO');

insert into follows(user_id, following_id) 
values
(31, 32),
(31, 33),
(31, 34),
(31, 35),
(31, 36),
(31, 37),
(31, 38),
(31, 39),
(32, 31),
(33, 31),
(34, 31),
(35, 31),
(36, 31),
(37, 31),
(38, 31),
(39, 31);
