create table account (
	id serial,
	account_name varchar(50) not null,
	owner_name varchar(50),
	date_format varchar(20),
	reverse_sign bool default false,
	skip_header bool default false,
	line_sep char(2) default '\n',
	element_sep char(2) default ',',
	primary key(id)
);