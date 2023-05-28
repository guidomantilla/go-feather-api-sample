create table auth_resources
(
    id        number generated always as identity primary key,
    name      varchar2(100) not null,
    enabled   number(1) null,
    constraint auth_resources_uq unique (name)
);

create table auth_roles
(
    id        number generated always as identity primary key,
    name      varchar2(50) not null,
    enabled   number(1) null,
    constraint auth_roles_uq unique (name)
);

create table auth_access_control_list
(
    id          number generated always as identity primary key,
    role_id     number not null,
    resource_id number not null,
    permission  varchar2(20) not null,
    enabled     number(1) not null,
    constraint auth_access_control_list_auth_resources_id_fk
        foreign key (resource_id) references auth_resources (id),
    constraint auth_access_control_list_auth_roles_id_fk
        foreign key (role_id) references auth_roles (id)
);

create table auth_users
(
    id         number generated always as identity primary key,
    role_id    number not null,
    username   varchar2(50) not null,
    password   varchar2(250) not null,
    passphrase varchar2(250) not null,
    enabled    number(1) not null,
    constraint auth_users_pk
        unique (username),
    constraint auth_users_auth_roles_id_fk
        foreign key (role_id) references auth_roles (id)
);

CREATE OR REPLACE VIEW auth_principals AS
SELECT u.username,
       r.name                                             as role,
       ar.name                                            as resource,
    acl.permission,
    u.password,
    u.passphrase,
    (u.enabled = r.enabled = ar.enabled = acl.enabled) as enabled
    FROM auth_access_control_list acl
    JOIN auth_resources ar ON ar.id = acl.resource_id
    JOIN auth_roles r ON r.id = acl.role_id
    JOIN auth_users u ON u.role_id = r.id
    where (u.enabled = r.enabled = ar.enabled = acl.enabled) = 1;