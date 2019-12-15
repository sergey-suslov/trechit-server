create table timespans(
  id serial primary key,
  userId integer not null ,
  startTime timestamp default now(),
  stopTime timestamp,
  created timestamp default now() not null
);

---- create above / drop below ----

drop table timespans;
