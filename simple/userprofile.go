package simple

type User struct {
}

func NewUser() *User {
	return &User{}
}

type Profile struct {
	Password *string
}

func NewProfile(password string) *Profile {
	return &Profile{
		Password: &password,
	}
}

type UserProfile struct {
	*User
	*Profile
}





