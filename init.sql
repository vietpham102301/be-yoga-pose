CREATE DATABASE IF NOT EXISTS yoga_support;


create table history
(
    id         int auto_increment
        primary key,
    user_id    int                                not null,
    pose_name  text                               not null,
    score      int                                not null,
    path       text                               not null,
    created_at datetime default CURRENT_TIMESTAMP null,
    is_saved   tinyint(1)                         not null
);

create table users
(
    id            int auto_increment
        primary key,
    username      varchar(255)                        not null,
    email         varchar(255)                        not null,
    password_hash varchar(60)                         not null,
    created_at    timestamp default CURRENT_TIMESTAMP null,
    updated_at    timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    constraint email
        unique (email),
    constraint username
        unique (username)
);

create table yoga_poses
(
    id         int auto_increment
        primary key,
    name       varchar(255)                        not null,
    path       varchar(255)                        not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);

