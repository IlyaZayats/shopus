create table Networks(
    id serial not null primary key,
    name varchar
);

create table Dealers(
    id serial not null primary key,
    network_id integer not null references Networks(id) on delete cascade,
    name varchar
);

create table Lists(
    id serial not null primary key,
    dealer_id integer not null references Dealers(id) on delete cascade,
    name varchar,
    price int,
    amount int,
    created_at date,
    info varchar,
    carrier varchar,
    contact_person varchar,
    note varchar
);

create table Users(
    id serial not null primary key,
    login varchar,
    pwd varchar,
    user_type int
);

insert into users (login, pwd, user_type) values ('admin', 'admin', 1), ('user', 'user', 0);
