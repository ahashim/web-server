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
		if output := Status(test.arg); output != test.want {
			t.Errorf("Output %q not equal to expected %q", output, test.want)
		}
	}
}
