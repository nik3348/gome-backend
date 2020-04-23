create table users
(
	uid int unsigned auto_increment,
	name varchar(255) not null,
	email varchar(255) null,
	cid int unsigned,
    primary key (uid),
    foreign key (cid) references courses(cid)
);
