package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPeopleApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PeopleApi Suite")
}
