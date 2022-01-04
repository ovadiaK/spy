package test

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/ovadiaK/spy/pkg/spy"
	"testing"
)

func TestSpy(t *testing.T) {
	ginkgo.Describe("set and get should get called", func() {
		ginkgo.It("spy reports functions called", func() {
			Spy := spy.New()
			Spy.Set("hello", "world")
			res := Spy.Get("hello")
			gomega.Expect(res).To(gomega.Equal("world"))
			gomega.Expect(Spy.GetFunc.WasCalled()).To(gomega.BeTrue())
		})
		ginkgo.It("spy reports functions not called", func() {
			Spy := spy.New()
			gomega.Expect(Spy.GetFunc.WasCalled()).To(gomega.BeFalse())
			gomega.Expect(Spy.SetFunc.WasCalled()).To(gomega.BeFalse())
		})
	})
}
