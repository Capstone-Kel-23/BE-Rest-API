package utils

type Role struct {
	slug string
}

func (r Role) String() string {
	return r.slug
}

var prefix = "ROLE_"

var (
	Admin = Role{prefix + "ADMIN"}
	User  = Role{prefix + "USER"}
)
