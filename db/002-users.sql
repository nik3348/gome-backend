create table users
(
	uid int unsigned auto_increment,
	name varchar(255) not null,
	email varchar(255) null,
	cid int unsigned,
    primary key (uid),
    foreign key (cid) references courses(cid)
);
insert into users(name, email, cid)
values ('Yumiko', 'yumi@gmail.com', null),
       ('Darren', 'darrg@gmail.com', null),
       ('Zab', 'z@gmail.com', 1),
       ('Bash', 'bd@gmail.com', null);
