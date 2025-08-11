package application_test

import (
	"app/internal/application"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTearDown(t *testing.T) {
	cfg := application.ConfigApplicationDefault{
		Addr: ":8081",
	}
	result := application.NewApplicationDefault(&cfg).TearDown()

	assert.NoError(t, result)
}

func TestSetUp(t *testing.T) {
	var cfg application.ConfigApplicationDefault
	result := application.NewApplicationDefault(&cfg).SetUp()

	assert.NoError(t, result)
}
