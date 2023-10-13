package models

type AuthResource struct {
	Name        *string `db:"name"`
	Application *string `db:"application"`
	Enabled     *bool   `db:"enabled"`
}

type AuthRole struct {
	Name    *string `db:"name"`
	Enabled *bool   `db:"enabled"`
}

type AuthAccessControlList struct {
	Role       *string `db:"role"`
	Resource   *string `db:"resource"`
	Permission *string `db:"permission"`
	Enabled    *bool   `db:"enabled"`
}

type AuthUser struct {
	Username   *string `db:"username"`
	Role       *string `db:"role"`
	Password   *string `db:"password"`
	Passphrase *string `db:"passphrase"`
	Enabled    *bool   `db:"enabled"`
}

type AuthPrincipal struct {
	Username    *string `db:"username"`
	Role        *string `db:"role"`
	Application *string `db:"application"`
	Resource    *string `db:"resource"`
	Permission  *string `db:"permission"`
	Password    *string `db:"password"`
	Passphrase  *string `db:"passphrase"`
	Enabled     *bool   `db:"enabled"`
}
