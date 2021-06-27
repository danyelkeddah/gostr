package gostr

import (
	"testing"
)

func TestAfter(t *testing.T) {
	type test struct {
		input  string
		search string
		output string
	}

	tests := []test{
		{
			"test abc",
			"",
			"test abc",
		},
		{
			"test abc",
			"zzz",
			"test abc",
		},
		{
			"test abc",
			"test ",
			"abc",
		},
		{
			"",
			"test ",
			"",
		},
	}

	for _, v := range tests {
		x := After(v.input, v.search)
		if x != v.output {
			t.Error("Expected", v.output, "got", x)
		}
	}
}

func TestAfterLast(t *testing.T) {
	type test struct {
		input  string
		search string
		output string
	}

	tests := []test{
		{
			"test abc abc zzz",
			"",
			"test abc abc zzz",
		},
		{
			"test abc abc zzz",
			"qqq",
			"test abc abc zzz",
		},
		{
			"test abc abc zzz",
			"abc ",
			"zzz",
		},
		{
			"",
			"test ",
			"",
		},
	}

	for _, v := range tests {
		x := AfterLast(v.input, v.search)
		if x != v.output {
			t.Error("Expected", v.output, "got", x)
		}
	}
}

func TestBefore(t *testing.T) {
	type test struct {
		input  string
		search string
		output string
	}

	tests := []test{
		{
			"test abc abc zzz",
			"",
			"test abc abc zzz",
		},
		{
			"test abc abc zzz",
			"qqq",
			"test abc abc zzz",
		},
		{
			"test abc abc zzz",
			"abc ",
			"test ",
		},
		{
			"",
			"test ",
			"",
		},
	}

	for _, v := range tests {
		x := Before(v.input, v.search)
		if x != v.output {
			t.Error("Expected", v.output, "got", x)
		}
	}
}

func TestBeforeLast(t *testing.T) {
	type test struct {
		input  string
		search string
		output string
	}

	tests := []test{
		{
			"test abc abc zzz",
			"",
			"test abc abc zzz",
		},
		{
			"test abc abc zzz",
			"qqq",
			"test abc abc zzz",
		},
		{
			"test qqq abc bbb abc zzz",
			"abc ",
			"test qqq abc bbb ",
		},
		{
			"",
			"test ",
			"",
		},
	}

	for _, v := range tests {
		x := BeforeLast(v.input, v.search)
		if x != v.output {
			t.Error("Expected", v.output, "got", x)
		}
	}
}

func TestBetween(t *testing.T) {
	type test struct {
		input  string
		from   string
		to     string
		output string
	}

	tests := []test{
		{
			"test abc abc zzz",
			"",
			"",
			"test abc abc zzz",
		},
		{
			"test from a to oo oo z",
			"a",
			"z",
			" to oo oo ",
		},
	}

	for _, v := range tests {
		x := Between(v.input, v.from, v.to)
		if x != v.output {
			t.Error("Expected", v.output, "got", x)
		}
	}
}

func TestContains(t *testing.T) {
	type test struct {
		haystack string
		needles  []string
		output   bool
	}

	tests := []test{
		{
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			[]string{"dolor", "***"},
			true,
		},
		{
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			[]string{"****"},
			false,
		},
		{
			"",
			[]string{"****"},
			false,
		},
	}

	for _, v := range tests {
		x := Contains(v.haystack, v.needles)
		if x != v.output {
			t.Error("Expected", v.output, "got", x)
		}
	}
}

func TestContainsAll(t *testing.T) {
	type test struct {
		haystack string
		needles  []string
		output   bool
	}

	tests := []test{
		{
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			[]string{"dolor", "consectetur"},
			true,
		},
		{
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			[]string{"dolor", "****"},
			false,
		},
		{
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			[]string{"****"},
			false,
		},
		{
			"",
			[]string{"****"},
			false,
		},
	}

	for _, v := range tests {
		x := ContainsAll(v.haystack, v.needles)
		if x != v.output {
			t.Error("Expected", v.output, "got", x)
		}
	}
}

func TestEndsWith(t *testing.T) {
	type test struct {
		haystack string
		needles  []string
		output   bool
	}

	tests := []test{
		{
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			[]string{"adipiscing", "elit."},
			true,
		},
		{
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			[]string{"dolor", "****"},
			false,
		},
		{
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			[]string{"****"},
			false,
		},
		{
			"",
			[]string{"****"},
			false,
		},
	}

	for _, v := range tests {
		x := EndsWith(v.haystack, v.needles)
		if x != v.output {
			t.Error("Expected", v.output, "got", x)
		}
	}
}

func TestIsUUID(t *testing.T) {

	type test struct {
		input  string
		output bool
	}

	tests := []test{
		{
			"6ba7b814-9dad-11d1-80b4-00c04fd430c8",
			true,
		},
		{
			"lorem ipsum",
			false,
		},
		{
			"",
			false,
		},
	}

	for _, v := range tests {
		x := IsUUID(v.input)
		if x != v.output {
			t.Error("Expected", v.output, "got", x)
		}
	}
}

func TestLength(t *testing.T) {

	type test struct {
		input  string
		output int
	}

	tests := []test{
		{
			"four",
			4,
		},
	}

	for _, v := range tests {
		x := Length(v.input)
		if x != v.output {
			t.Error("Expected", v.output, "got", x)
		}
	}
}

func TestLower(t *testing.T) {

	type test struct {
		input  string
		output string
	}

	tests := []test{
		{
			"FOUR",
			"four",
		},
		{
			"",
			"",
		},
	}

	for _, v := range tests {
		x := Lower(v.input)
		if x != v.output {
			t.Error("Expected", v.output, "got", x)
		}
	}
}

func TestWords(t *testing.T) {

	type test struct {
		input  string
		n      uint
		end    string
		output string
	}

	tests := []test{
		{
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			4,
			"...",
			"Lorem ipsum dolor sit...",
		},
		{
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			0,
			"...",
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		},
		{
			"",
			5,
			"...",
			"",
		},
	}

	for _, v := range tests {
		x := Words(v.input, v.n, v.end)
		if x != v.output {
			t.Error("Expected", v.output, "got", x)
		}
	}
}

func TestPadRight(t *testing.T) {

	type test struct {
		input  string
		l      uint
		pad    string
		output string
	}

	tests := []test{
		{
			"Lorem",
			10,
			"*",
			"Lorem*****",
		},
		{
			"Lorem",
			3,
			"*",
			"Lorem",
		},
		{
			"Lorem",
			0,
			"*",
			"Lorem",
		},
		{
			"Loremm",
			10,
			"***",
			"Loremm****",
		},
	}

	for _, v := range tests {
		x := PadRight(v.input, v.l, v.pad)
		if x != v.output {
			t.Error("Expected", v.output, "got", x)
		}
	}
}

func TestPadLeft(t *testing.T) {
	type test struct {
		input  string
		l      uint
		pad    string
		output string
	}

	tests := []test{
		{
			"Lorem",
			10,
			"*",
			"*****Lorem",
		},
		{
			"Lorem",
			3,
			"*",
			"Lorem",
		},
		{
			"Lorem",
			0,
			"*",
			"Lorem",
		},
		{
			"Loremm",
			10,
			"***",
			"****Loremm",
		},
	}

	for _, v := range tests {
		x := PadLeft(v.input, v.l, v.pad)
		if x != v.output {
			t.Error("Expected", v.output, "got", x)
		}
	}
}

func TestRandom(t *testing.T) {
	type test struct {
		r uint
	}

	tests := []test{
		{
			5,
		},
		{
			0,
		},
	}

	for _, v := range tests {
		x := Random(v.r)
		if int(v.r) == 0 && len(x) != 16 {
			t.Error("Expected", v.r, "got", x)
		}
		if v.r != 0 && len(x) != int(v.r) {
			t.Error("Expected", v.r, "got", x)
		}
	}
}

func TestRepeat(t *testing.T) {
	type test struct {
		input  string
		t      uint
		output string
	}

	tests := []test{
		{
			"lorem",
			3,
			"loremloremlorem",
		},
		{
			"lorem",
			0,
			"",
		},
		{
			"",
			5,
			"",
		},
	}

	for _, v := range tests {
		x := Repeat(v.input, v.t)
		if x != v.output {
			t.Error("Expected", v.output, "got", x)
		}
	}
}

func TestUpper(t *testing.T) {

	type test struct {
		input  string
		output string
	}

	tests := []test{
		{
			"four",
			"FOUR",
		},
		{
			"",
			"",
		},
	}

	for _, v := range tests {
		x := Upper(v.input)
		if x != v.output {
			t.Error("Expected", v.output, "got", x)
		}
	}
}
