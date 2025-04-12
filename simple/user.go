package simple

type UserRepository struct {

}

func NewUserRepository () *UserRepository {
	return &UserRepository{}
}

type UserService struct {
	Repo *UserRepository
}

func NewUserService(userRepo *UserRepository) *UserService {
	return &UserService{
		Repo: userRepo,
	}
}

