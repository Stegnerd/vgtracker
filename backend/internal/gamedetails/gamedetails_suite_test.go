package gamedetails_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGameDetailInternal(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GameDetailInternal Suite")
}
