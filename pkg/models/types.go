package models

type Model struct {
	AuthResource          *AuthResource          `db_table:"auth_resources"`
	AuthRole              *AuthRole              `db_table:"auth_roles"`
	AuthAccessControlList *AuthAccessControlList `db_table:"auth_access_control_list"`
	AuthUser              *AuthUser              `db_table:"auth_users"`
	AuthPrincipal         *AuthPrincipal         `db_view:"auth_principals"`
}

type AuthResource struct {
	ID      *int    `db_column:"id,pk"`
	Name    *string `db_column:"name,uq"`
	Enabled *bool   `db_column:"enabled"`
}

type AuthRole struct {
	ID      *int    `db_column:"id,pk"`
	Name    *string `db_column:"name,uq"`
	Enabled *bool   `db_column:"enabled"`
}

type AuthAccessControlList struct {
	ID         *int    `db_column:"id,pk"`
	RoleID     *int    `db_column:"role_id"`
	ResourceID *int    `db_column:"resource_id"`
	Permission *string `db_column:"permission"`
	Enabled    *bool   `db_column:"enabled"`
}

type AuthUser struct {
	ID         *int    `db_column:"id,pk"`
	RoleID     *int    `db_column:"role_id"`
	Username   *string `db_column:"username,uq"`
	Password   *string `db_column:"password"`
	Passphrase *string `db_column:"passphrase"`
	Enabled    *bool   `db_column:"enabled"`
}

type AuthPrincipal struct {
	Username    *string `db_column:"username,uq"`
	Role        *string `db_column:"role"`
	Application *string `db_column:"application"`
	Resource    *string `db_column:"resource"`
	Permission  *string `db_column:"permission"`
	Password    *string `db_column:"password"`
	Passphrase  *string `db_column:"passphrase"`
	Enabled     *bool   `db_column:"enabled"`
}
