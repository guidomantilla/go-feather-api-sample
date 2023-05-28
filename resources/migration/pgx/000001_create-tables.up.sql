create table if not exists auth_resources
(
    id      serial primary key,
    name    varchar(100) not null,
    enabled boolean null,
    constraint auth_resources_uq unique (name)
);

create table if not exists auth_roles
(
    id      serial primary key,
    name    varchar(50) not null,
    enabled boolean null,
    constraint auth_roles_uq unique (name)
);

create table if not exists auth_access_control_list
(
    id          serial primary key,
    role_id     int not null,
    resource_id int not null,
    permission  varchar(20) not null,
    enabled     boolean not null,
    constraint auth_access_control_list_auth_resources_id_fk
        foreign key (resource_id) references auth_resources (id),
    constraint auth_access_control_list_auth_roles_id_fk
        foreign key (role_id) references auth_roles (id)
);

create table if not exists auth_users
(
    id         serial primary key,
    role_id    int not null,
    username   varchar(50) not null,
    password   varchar(250) not null,
    passphrase varchar(250) not null,
    enabled    boolean not null,
    constraint auth_users_pk unique (username)
    constraint auth_users_auth_roles_id_fk
        foreign key (role_id) references auth_roles (id)
);

create view auth_principals as
select u.username,
       r.name                                             as role,
       ar.name                                            as resource,
    acl.permission,
    u.password,
    u.passphrase,
    (u.enabled = r.enabled = ar.enabled = acl.enabled) as enabled
    from auth_access_control_list acl
    join auth_resources ar on ar.id = acl.resource_id
    join auth_roles r on r.id = acl.role_id
    join auth_users u on u.role_id = r.id
    where (u.enabled = r.enabled = ar.enabled = acl.enabled) = 1;
