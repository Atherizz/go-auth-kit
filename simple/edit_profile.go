package simple



type username string
type password string

func ProvideUsername(s string) *username {
	u := username(s)
	return &u
}

func ProvidePassword(s string) *password {
	pw := password(s)
	return &pw
}

type Profil struct {
	Username username
	Password password
}

func NewProfil(username *username, password *password) *Profil {
	return &Profil{
		Username: *username,
		Password: *password,
	}
}

type ProfilAction interface {
	ShowData() (string, string) 
	ChangePassword(username string, oldPassword string, newPassword string)
}

type ProfilActionImpl struct {
	Profil Profil
}

func (profilImpl *ProfilActionImpl) ShowData() (string, string) {
	return string(profilImpl.Profil.Username), string(profilImpl.Profil.Password)
}

func (profilImpl *ProfilActionImpl) ChangePassword(username string, oldPassword string, newPassword string) {
	if username == string(profilImpl.Profil.Username) && oldPassword == string(profilImpl.Profil.Password) {
		profilImpl.Profil.Password = password(newPassword)
	}
}

func NewProfilActionImpl(profil *Profil) *ProfilActionImpl {
	return &ProfilActionImpl{
		Profil: *profil,
	}
}

type ProfilService struct {
	ProfilAction ProfilAction
}

func NewProfilService(profilAction ProfilAction) *ProfilService {
	return &ProfilService{
		ProfilAction: profilAction,
	}
}
