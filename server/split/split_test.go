package split_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	split "github.com/mplewis/bondburger/server/split"
)

func TestSplit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Split Suite")
}

var _ = Describe("Split", func() {
	It("splits paragraphs into sentences", func() {
		body := "Foo. Bar. Baz. Quux. Plugh."
		Expect(split.Sentences(body)).To(Equal([]string{"Foo.", "Bar.", "Baz.", "Quux.", "Plugh."}))
		body = "   Foo. Bar.   Baz. Quux. Plugh. "
		Expect(split.Sentences(body)).To(Equal([]string{"Foo.", "Bar.", "Baz.", "Quux.", "Plugh."}))
		body = `He survives by stealing a parachute from the pilot, while Jaws lands on a trapeze net within a circus tent.
At the Drax Industries spaceplane-manufacturing complex in California, Bond meets the owner of the company, Hugo Drax, and his henchman Chang.`
		Expect(split.Sentences(body)).To(Equal([]string{
			"He survives by stealing a parachute from the pilot, while Jaws lands on a trapeze net within a circus tent.",
			"At the Drax Industries spaceplane-manufacturing complex in California, Bond meets the owner of the company, Hugo Drax, and his henchman Chang.",
		}))
		body = `Contacting Leiter, the pair gets the U.S. Navy to intercept Disco Volante, to engage in an underwater battle, and to recover one of the bombs. Bond pursues Largo, and grabs hold of Disco Volante as she sheds the rear half to become a hydrofoil to escape.`
		Expect(split.Sentences(body)).To(Equal([]string{
			"Contacting Leiter, the pair gets the U.S. Navy to intercept Disco Volante, to engage in an underwater battle, and to recover one of the bombs.",
			"Bond pursues Largo, and grabs hold of Disco Volante as she sheds the rear half to become a hydrofoil to escape.",
		}))
		body = `While he is gone, Carver's assassin Dr. Kaufman kills Paris.`
		Expect(split.Sentences(body)).To(Equal([]string{
			`While he is gone, Carver's assassin Dr. Kaufman kills Paris.`,
		}))
	})
})
