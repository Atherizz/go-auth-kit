package test

import (
	"golang-restful-api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	conn, cleanup  := simple.InitializeConnection("database")
	assert.NotNil(t, conn)

	cleanup()

}