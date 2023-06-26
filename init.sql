create table users (
    id bigint primary key auto_increment,
    username varchar(20) not null,
    password varchar(32) not null,
    create_time bigint not null,
    update_time bigint not null,
    delete_time bigint not null,
    unique index(username)
);

create table instance (
    id bigint primary key auto_increment,
    name varchar(50) not null,
    description varchar(1000),
    `condition` longtext not null,
    create_user bigint not null,
    create_time bigint not null,
    update_time bigint not null,
    delete_time bigint not null,
    index(create_user)
);