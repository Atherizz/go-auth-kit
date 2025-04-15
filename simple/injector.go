//go:build wireinject
// +build wireinject

package simple

import (
	"io"
	"os"

	"github.com/google/wire"
)

func InitializeService(isError bool) (*SampleService, error) {
	wire.Build(NewSampleRepository, NewSampleService)
	return nil, nil
}

// MULTIPLE BINDING
func InitializeDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabaseMySQL, NewDatabasePostgreSQL, NewDatabase,
	)
	return nil
}

// PROVIDER SET
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

// BINDING INTERFACE
var dashboardSet = wire.NewSet(
	NewShowDashboardImpl,
	wire.Bind(new(ShowDashboard), new(*ShowDashboardImpl)),
)

// "Kalau kamu butuh interface ShowDashboard, dan punya *ShowDashboardImpl,
// silakan pakai *ShowDashboardImpl sebagai implementasi dari ShowDashboard."

func InitializeDashboardService() *DashboardService {
	wire.Build(dashboardSet, NewDashboardService)
	return nil
}

var userProfileSet = wire.NewSet(NewUser, NewProfile)

// STRUCT PROVIDER
func InitializeUserProfile(password string) *UserProfile {
	wire.Build(
		userProfileSet,
		wire.Struct(new(UserProfile), "User", "Profile"),
	)
	return nil
}

// BINDING VALUE
var userProfileValueSet = wire.NewSet(
	wire.Value(&User{}),
	wire.Value(&Profile{}),
)

func InitializeUserProfileUsingValue(password string) *UserProfile {
	wire.Build(
		userProfileValueSet,
		wire.Struct(new(UserProfile), "*"),
	)
	return nil
}

// func InitializeProfilWithInterface(u string, p string) *ProfilService {
// 	username := ProvideUsername(u)
// 	password := ProvidePassword(p)
// 	return BuildProfilService(username, password)
// }

func BuildProfilService(username *username, password *password) *ProfilService {
	wire.Build(
		NewProfil,
		NewProfilActionImpl,
		wire.Bind(new(ProfilAction), new(*ProfilActionImpl)),
		NewProfilService,
	)
	return nil
}

func InitializerReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

// Membuat provider dari field yang ada di provider lain
func InitailizeConfiguration() *Configuration {
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Apllication), "Configuration"),
	)
	return nil
}

func InitializeConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile,)
	return nil,nil
}