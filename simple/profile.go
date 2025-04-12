package simple

type ProfileRepository struct {

}

func NewProfileRepository () *ProfileRepository {
	return &ProfileRepository{}
}

type ProfileService struct {
	*ProfileRepository
}

func NewProfileService (profileRepo *ProfileRepository) *ProfileService {
	return &ProfileService{
		ProfileRepository: profileRepo,
	}
}