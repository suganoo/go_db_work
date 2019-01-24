drop table if exists members;

create table members (
  id          serial primary key,
  first_name  varchar(255),
  last_name   varchar(255),
  email       varchar(255),
  accessprev  boolean
);

INSERT INTO members VALUES (1, 'minoru', 'tanaka', 'tanaka@gmail.com', TRUE);
INSERT INTO members VALUES (2, 'tadashi', 'sato', 'sato@gmail.com', FALSE);
INSERT INTO members VALUES (3, 'sachiko', 'suzuki', 'suzuki@yahoo.co.jp', FALSE);
