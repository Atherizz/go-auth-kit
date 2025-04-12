//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializeService(isError bool) (*SampleService, error) {
    wire.Build(NewSampleRepository, NewSampleService,)
    return nil, nil
}

func InitializeDatabaseRepository() *DatabaseRepository {
    wire.Build(
        NewDatabaseMySQL, NewDatabasePostgreSQL, NewDatabase,
    )
    return nil
}

var userSet = wire.NewSet(NewUserRepository, NewUserService)
var profileSet = wire.NewSet(NewProfileRepository, NewProfileService)

func InitializeUserProfileService() *UserProfileService {

    wire.Build(
        userSet, profileSet, NewUserProfileService,
    )

    return nil
}

// func InitializeDashboardService() *DashboardService {
//     wire.Build(NewDashboardService, NewShowDashboardImpl)
//     return nil
// }

var dashboardSet = wire.NewSet(
    NewShowDashboardImpl,
    wire.Bind(new(ShowDashboard), new(*ShowDashboardImpl)),
)
// Memberi tahu Wire bahwa jika ada yang membutuhkan interface ShowDashboard,
// maka gunakan implementasi dari *ShowDashboardImpl.

func InitializeDashboardService() *DashboardService {
    wire.Build(dashboardSet, NewDashboardService)
    return nil
}
