package enums

import "testing"

type interactionTest struct {
	arg  Interaction
	want Interaction
}

var interactionTests = []interactionTest{
	{Interaction(0), Dislike},
	{Interaction(1), Like},
	{Interaction(2), Resqueak},
	{Interaction(3), UndoDislike},
	{Interaction(4), UndoLike},
	{Interaction(5), UndoResqueak},
}

func TestInteraction(t *testing.T) {
	for _, test := range interactionTests {
		if got := Interaction(test.arg); got != test.want {
			t.Errorf("got %q not equal to expected %q", got, test.want)
		}
	}
}
