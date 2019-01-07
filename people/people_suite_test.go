package people_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPeople(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "People Suite")
}
