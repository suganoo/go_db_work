drop table if exists members;

create table members (
  id          serial primary key,
  first_name  varchar(255)
  last_name   varchar(255)
  email       varchar(255)
);
