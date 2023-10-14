package models

type AuthResource struct {
	Name        *string `db:"name,pk"`
	Application *string `db:"application,pk"`
	Enabled     *bool   `db:"enabled"`
}

type AuthRole struct {
	Name    *string `db:"name,pk"`
	Enabled *bool   `db:"enabled"`
}

type AuthAccessControlList struct {
	Role       *string `db:"role,pk"`
	Resource   *string `db:"resource,pk"`
	Permission *string `db:"permission,pk"`
	Enabled    *bool   `db:"enabled"`
}

type AuthUser struct {
	Username   *string `db:"username,pk"`
	Role       *string `db:"role"`
	Password   *string `db:"password"`
	Passphrase *string `db:"passphrase"`
	Enabled    *bool   `db:"enabled"`
}

type AuthPrincipal struct {
	Username    *string `db:"username,pk"`
	Role        *string `db:"role"`
	Application *string `db:"application,pk"`
	Resource    *string `db:"resource"`
	Permission  *string `db:"permission"`
	Password    *string `db:"password"`
	Passphrase  *string `db:"passphrase"`
	Enabled     *bool   `db:"enabled"`
}

//
