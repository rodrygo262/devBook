insert into usuarios ( nome, nick, email, senha )
values
( "usuario 1", "usuario_1", "usuario1@gmail.com", "$2a$10$m5DY80/oWGEl9VwE9ifZY.YrzriXrsWCbP40lBV/msAU0zroXmzCm" ),
( "usuario 2", "usuario_2", "usuario2@gmail.com", "$2a$10$m5DY80/oWGEl9VwE9ifZY.YrzriXrsWCbP40lBV/msAU0zroXmzCm" ),
( "usuario 3", "usuario_3", "usuario3@gmail.com", "$2a$10$m5DY80/oWGEl9VwE9ifZY.YrzriXrsWCbP40lBV/msAU0zroXmzCm" );

insert into seguidores ( usuario_id, seguidor_id )
values
( 1, 2 ),
( 3, 1 ),
( 1, 3 );