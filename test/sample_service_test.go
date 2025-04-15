package test

import (
	"golang-restful-api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSampleServiceError(t *testing.T) {
	sampleService, err:= simple.InitializeService(true)
	assert.Nil(t, sampleService)
	assert.NotNil(t, err)
}

func TestSampleServiceSuccess(t *testing.T) {
	sampleService, err:= simple.InitializeService(false)
	assert.Nil(t, err)
	assert.NotNil(t, sampleService)
}

func TestEditProfile(t *testing.T) {
	provideUsername := simple.ProvideUsername("atherizz")
	providePassword := simple.ProvidePassword("admin123")

	profilService := simple.BuildProfilService(provideUsername, providePassword)

	profilService.ProfilAction.ChangePassword("atherizz", "admin123", "sekut767")
	_, password := profilService.ProfilAction.ShowData()

	assert.Equal(t, "sekut767", password)
}

