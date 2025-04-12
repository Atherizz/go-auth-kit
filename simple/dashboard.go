package simple

type ShowDashboard interface {
	Show(name string) string
}

type ShowDashboardImpl struct {
	
}

func NewShowDashboardImpl() *ShowDashboardImpl {
	return &ShowDashboardImpl{}
}


func (dashboard ShowDashboardImpl) Show (name string) string {
	return "show " + name
}

type DashboardService struct {
	ShowDashboard ShowDashboard
}

func NewDashboardService (showDashboard ShowDashboard) *DashboardService {
	return &DashboardService{
		ShowDashboard: showDashboard,
	}
}
