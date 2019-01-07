package integration_tests_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gexec"
	"net/http"
	"os/exec"
	"io"
	"bytes"
	"io/ioutil"
	"strings"
)

var _ = Describe("Integration", func() {
	var (
		stdout   bytes.Buffer
		stderr   bytes.Buffer
	)

	It("Should return a list of people", func() {
		api, err := gexec.Build("github.com/bu3/people-api")
		Expect(err).ToNot(HaveOccurred())

		cmd := exec.Command(api)
		_, err = gexec.Start(cmd, io.MultiWriter(GinkgoWriter, &stdout), io.MultiWriter(GinkgoWriter, &stderr))
		Expect(err).NotTo(HaveOccurred())

		response, err := http.Get("http://localhost:8000/people")
		Expect(err).ToNot(HaveOccurred())

		body, err := ioutil.ReadAll(response.Body)
			Expect(strings.Trim(string(body), "\n")).To(Equal("[\"Penn\",\"Teller\"]"))
	})
})
