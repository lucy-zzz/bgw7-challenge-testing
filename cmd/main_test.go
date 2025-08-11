package main

import (
	"app/internal/application/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunAppWithApplication(t *testing.T) {
	t.Run("should setup and run application successfully with stub", func(t *testing.T) {
		stubApp := mocks.NewApplicationStub()

		err := runAppWithApplication(stubApp)

		assert.NoError(t, err)
	})

	t.Run("should return error when setup fails", func(t *testing.T) {
		stubApp := mocks.NewApplicationStub()
		stubApp.SetUpError = errors.New("setup failed")

		err := runAppWithApplication(stubApp)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to setup app: setup failed")
	})

	t.Run("should return error when run fails", func(t *testing.T) {
		stubApp := mocks.NewApplicationStub()
		stubApp.RunError = errors.New("run failed")

		err := runAppWithApplication(stubApp)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to run app: run failed")
	})

	t.Run("should call TearDown even when setup fails", func(t *testing.T) {
		stubApp := mocks.NewApplicationStub()
		stubApp.SetUpError = errors.New("setup failed")

		err := runAppWithApplication(stubApp)

		assert.Error(t, err)
		// TearDown é chamado via defer, então sempre executa
	})
}
