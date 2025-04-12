package simple

type UserProfileService struct {
	User *UserService
	Profile *ProfileService
}

func NewUserProfileService (user *UserService, profile *ProfileService) *UserProfileService{
	return &UserProfileService{
		User: user,
		Profile: profile,
	}
}