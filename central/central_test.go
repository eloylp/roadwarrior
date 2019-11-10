package central_test

import (
	"github.com/DATA-DOG/godog"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	status := godog.RunWithOptions("godog", func(s *godog.Suite) {
		FeatureContext(s)
	}, godog.Options{
		Format:    "progress",
		Paths:     []string{"features"},
		Randomize: time.Now().UTC().UnixNano(),
	})

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^that the central system is running$`, thatTheCentralSystemIsRunning)
	s.Step(`^The central system is up$`, theCentralSystemIsUp)
	s.Step(`^The central system has now (\d+) subsystems up$`, theCentralSystemHasNowSubsystemsUp)
	s.Step(`^I start "([^"]*)" subsystem$`, iStartSubsystem)
	s.Step(`^I stop "([^"]*)" subsystem$`, iStopSubsystem)
	s.Step(`^I stop the central system$`, iStopTheCentralSystem)
	s.Step(`^The central system is down$`, theCentralSystemIsDown)
}

func thatTheCentralSystemIsRunning() error {
	return godog.ErrPending
}

func theCentralSystemIsUp() error {
	return godog.ErrPending
}

func theCentralSystemHasNowSubsystemsUp(arg1 int) error {
	return godog.ErrPending
}

func iStartSubsystem(arg1 string) error {
	return godog.ErrPending
}

func iStopSubsystem(arg1 string) error {
	return godog.ErrPending
}

func iStopTheCentralSystem() error {
	return godog.ErrPending
}

func theCentralSystemIsDown() error {
	return godog.ErrPending
}
