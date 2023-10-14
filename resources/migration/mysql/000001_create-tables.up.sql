create table if not exists auth_resources
(
    name        varchar(100)         not null,
    application varchar(100)         not null,
    enabled     tinyint(1) default 1 not null,
    primary key (name, application)
);

create table if not exists auth_roles
(
    name    varchar(100)         not null,
    enabled tinyint(1) default 1 not null,
    primary key (name)
);

create table if not exists auth_access_control_list
(
    role        varchar(100)         not null,
    resource    varchar(100)         not null,
    permission  varchar(20)          not null,
    enabled     tinyint(1) default 1 not null,
    primary key (role, resource, permission),
    constraint auth_access_control_list_auth_resources_fk
        foreign key (resource) references auth_resources (name),
    constraint auth_access_control_list_auth_roles_fk
        foreign key (role) references auth_roles (name)
);

create table if not exists auth_users
(
    username   varchar(50)          not null,
    role       varchar(100)         not null,
    password   varchar(250)         not null,
    passphrase varchar(250)         not null,
    enabled    tinyint(1) default 1 not null,
    primary key (username),
    constraint auth_users_auth_roles_fk
        foreign key (role) references auth_roles (name)
);

create view auth_principals as
select u.username,
       r.name                                             as role,
       ar.application                                     as application,
       ar.name                                            as resource,
    acl.permission,
    u.password,
    u.passphrase,
    (u.enabled = r.enabled = ar.enabled = acl.enabled) as enabled
    from auth_access_control_list acl
    join auth_resources ar on ar.name = acl.resource
    join auth_roles r on r.name = acl.role
    join auth_users u on u.role = r.name
    where (u.enabled = r.enabled = ar.enabled = acl.enabled) = 1;


insert into auth_resources (name, application, enabled) values ('/principals/current', 'go-feather-api-sample', true);
insert into auth_resources (name, application, enabled) values ('/principals/:username', 'go-feather-api-sample', true);
insert into auth_resources (name, application, enabled) values ('/principals', 'go-feather-api-sample', true);
insert into auth_resources (name, application, enabled) values ('/principals/change-password', 'go-feather-api-sample', true);

insert into auth_roles (name, enabled) values ('admin', true);
insert into auth_roles (name, enabled) values ('user', true);

insert into auth_access_control_list (role, resource, permission, enabled) values ('admin', '/principals/current', 'GET', true);
insert into auth_access_control_list (role, resource, permission, enabled) values ('admin', '/principals/:username', 'GET', true);
insert into auth_access_control_list (role, resource, permission, enabled) values ('admin', '/principals', 'POST', true);
insert into auth_access_control_list (role, resource, permission, enabled) values ('admin', '/principals', 'PUT', true);
insert into auth_access_control_list (role, resource, permission, enabled) values ('admin', '/principals', 'DELETE', true);
insert into auth_access_control_list (role, resource, permission, enabled) values ('admin', '/principals/change-password', 'PATCH', true);

insert into auth_access_control_list (role, resource, permission, enabled) values ('user', '/principals/current', 'GET', true);
insert into auth_access_control_list (role, resource, permission, enabled) values ('user', '/principals', 'POST', true);
insert into auth_access_control_list (role, resource, permission, enabled) values ('user', '/principals', 'PUT', true);
insert into auth_access_control_list (role, resource, permission, enabled) values ('user', '/principals/change-password', 'PATCH', true);

insert into auth_users (username, role, password, passphrase, enabled) values ('root', 'admin', '$2a$10$DMrp3hAmPg0EV16AchnF0.rdTiHJ/g3k7J9klzGVZoiZOzSR3u/le', '', true);
insert into auth_users (username, role, password, passphrase, enabled) values ('raven', 'user', '$2a$10$DMrp3hAmPg0EV16AchnF0.rdTiHJ/g3k7J9klzGVZoiZOzSR3u/le', '', true);