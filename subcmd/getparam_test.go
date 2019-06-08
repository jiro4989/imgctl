package subcmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetParams(t *testing.T) {
	type TestData struct {
		desc        string
		args        []string
		prefixes    []string
		expect      map[string]int
		isErrorTest bool
	}
	tds := []TestData{
		{
			desc:     "xとy",
			args:     []string{"x=1", "y=2"},
			prefixes: []string{"x", "y"},
			expect:   map[string]int{"x": 1, "y": 2},
		},
		{
			desc:     "xだけ",
			args:     []string{"x=1", "y=2"},
			prefixes: []string{"x"},
			expect:   map[string]int{"x": 1},
		},
		{
			desc:     "xとy (ただしyがない)",
			args:     []string{"x=1"},
			prefixes: []string{"x", "y"},
			expect:   map[string]int{"x": 1},
		},
		{
			desc:     "xとy (xもyもない)",
			args:     []string{"z=1"},
			prefixes: []string{"x", "y"},
			expect:   map[string]int{},
		},
		{
			desc:        "不正な値",
			args:        []string{"x=1", "y=a"},
			prefixes:    []string{"x", "y"},
			expect:      nil,
			isErrorTest: true,
		},
		{
			desc:        "valueがない",
			args:        []string{"x=1", "y="},
			prefixes:    []string{"x", "y"},
			expect:      nil,
			isErrorTest: true,
		},
		{
			desc:     "valueがない",
			args:     []string{"x=1", "=2"},
			prefixes: []string{"x", "y"},
			expect:   map[string]int{"x": 1},
		},
	}
	for _, v := range tds {
		got, err := getParams(v.args, v.prefixes...)
		assert.Equal(t, v.expect, got, v.desc)
		if v.isErrorTest {
			assert.NotNil(t, err, v.desc)
		} else {
			assert.Nil(t, err, v.desc)
		}
	}
}

func TestGetKeyValueParams(t *testing.T) {
	type TestData struct {
		desc     string
		args     []string
		prefixes []string
		expect   []string
	}
	tds := []TestData{
		{
			desc:     "xとy",
			args:     []string{"x=1", "y=2"},
			prefixes: []string{"x", "y"},
			expect:   []string{"x=1", "y=2"},
		},
		{
			desc:     "y",
			args:     []string{"x=1", "y=2"},
			prefixes: []string{"y"},
			expect:   []string{"y=2"},
		},
		{
			desc:     "z (存在しない)",
			args:     []string{"x=1", "y=2"},
			prefixes: []string{"z"},
			expect:   []string{},
		},
	}
	for _, v := range tds {
		got := getKeyValueParams(v.args, v.prefixes...)
		assert.Equal(t, v.expect, got, v.desc)
	}
}
