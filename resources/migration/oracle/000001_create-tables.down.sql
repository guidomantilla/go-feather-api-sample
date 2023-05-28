begin
    execute immediate 'drop table auth_access_control_list cascade constraints';
exception
    when others then
        if sqlcode != -942 then
            raise;
        end if;
end;

begin
    execute immediate 'drop table auth_resources cascade constraints';
exception
    when others then
        if sqlcode != -942 then
            raise;
        end if;
end;

begin
    execute immediate 'drop table auth_roles cascade constraints';
exception
    when others then
        if sqlcode != -942 then
            raise;
        end if;
end;

begin
    execute immediate 'drop table auth_users cascade constraints';
exception
    when others then
        if sqlcode != -942 then
            raise;
        end if;
end;

begin
    execute immediate 'drop view auth_principals cascade constraints';
exception
    when others then
        if sqlcode != -942 then
            raise;
        end if;
end;
