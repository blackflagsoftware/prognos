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

create table if not exists account_column (
	id serial,
	account_id integer,
	position integer,
	column_name varchar(50),
	primary key(id)
);

create table if not exists account_transaction (
	account_id int,
	file_name varchar(100),
	date_loaded timestamp
);

create table if not exists budget_allocation (
	category_id int,
	amount numeric
);

create table if not exists category (
	id serial,
	category_name varchar(50),
	primary key(id)
);

create table if not exists transaction (
	id serial,
	account_id int,
	category_id int,
	txn_date timestamp,
	amount numeric,
	description varchar(100),
	primary key(id)
);