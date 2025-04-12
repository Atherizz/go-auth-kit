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