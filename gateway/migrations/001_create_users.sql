create table users(
  id serial primary key,
  name varchar not null,
  email varchar not null unique,
  hash varchar not null,
  salt varchar not null,
  created timestamp default now()
);

---- create above / drop below ----

drop table users;
