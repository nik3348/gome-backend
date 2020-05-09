create table courses
(
    cid int unsigned auto_increment,
    name varchar(255) not null,
    primary key (cid)
);
insert into courses(name)
values ('Java'),
       ('Math'),
       ('Chemistry'),
       ('Economics');
