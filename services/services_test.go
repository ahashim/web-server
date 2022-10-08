package services

import (
	"os"
	"testing"

	"github.com/ahashim/web-server/config"
	"github.com/ahashim/web-server/ent"
	"github.com/ahashim/web-server/tests"

	"github.com/labstack/echo/v4"
)

var (
	c   *Container
	ctx echo.Context
	usr *ent.User
)

func TestMain(m *testing.M) {
	// Set the environment to test
	config.SwitchEnvironment(config.EnvTest)

	// Create a new container
	c = NewContainer()

	// Create a web context
	ctx, _ = tests.NewContext(c.Web, "/")
	tests.InitSession(ctx)

	// Create a test user
	var err error
	if usr, err = tests.CreateUser(c.ORM); err != nil {
		panic(err)
	}

	// Run tests
	exitVal := m.Run()

	// Shutdown the container
	if err = c.Shutdown(); err != nil {
		panic(err)
	}

	os.Exit(exitVal)
}
