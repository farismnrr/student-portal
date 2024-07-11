package main_test

import (
	"log"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGolang(t *testing.T) {
	copyrightKey := "COPYRIGHTKEY"
	copyrightUser := "farismnrr"
	copyrightMessage := "COPYRIGHTKEY: farismnrr"

	t.Run(copyrightMessage, func(t *testing.T) {
		got := copyrightMessage
		want := copyrightKey + ": " + copyrightUser

		if got != want {
			log.Fatal("Copyright Unauthorized! Please contact the owner's code!")
		} else {
			log.Println("Copyright Authorized")
		}
	})

	RegisterFailHandler(Fail)
	RunSpecs(t, "Golang Suite")
}
