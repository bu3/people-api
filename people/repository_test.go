package people_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/bu3/people-api/people"
)

var _ = Describe("Repository", func() {
	It("should return a list of people", func() {
		people := people.FindPeople()
		gomega.Expect(people).To(gomega.HaveLen(2))
		gomega.Expect(people[0]).To(gomega.Equal("Penn"))
		gomega.Expect(people[1]).To(gomega.Equal("Teller"))
	})
})
