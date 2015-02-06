package autonomous_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAutonomous(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Autonomous Suite")
}
