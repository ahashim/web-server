package enums

import "testing"

type statusTest struct {
	arg  Status
	want Status
}

var statusTests = []statusTest{
	{Status(0), Unknown},
	{Status(1), Active},
	{Status(2), Suspended},
	{Status(3), Banned},
}

func TestStatus(t *testing.T) {
	for _, test := range statusTests {
		if got := Status(test.arg); got != test.want {
			t.Errorf("got %q not equal to expected %q", got, test.want)
		}
	}
}
