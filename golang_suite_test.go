package main_test

import (
	"log"
	"testing"

	"github.com/farismnrr/golang-authorization-api/handler"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGolang(t *testing.T) {
	if !handler.CopyrightHandler() {
		log.Fatal("Unauthorized! Please contact the owner's code")
	}

	RegisterFailHandler(Fail)
	RunSpecs(t, "Golang Suite")
}
