package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)
var _ = Describe("Web App", func() {

    Context("Retrieve the main page", func() {
        It("returns a 200 Status Code", func() {
            Request("GET", "/", HandleIndex)
            Expect(response.Code).To(Equal(200))
        })
    })
})
