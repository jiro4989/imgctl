package subcmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	type TestData struct {
		desc        string
		args        []string
		prefix      []string
		isErrorTest bool
	}
	tds := []TestData{
		{
			desc:   "xは必須である",
			args:   []string{"x=1", "y=2"},
			prefix: []string{"x"},
		},
		{
			desc:   "xとyは必須である",
			args:   []string{"x=1", "y=2"},
			prefix: []string{"x", "y"},
		},
		{
			desc:        "xとzは必須であるが、xが含まれていない",
			args:        []string{"y=2"},
			prefix:      []string{"x", "z"},
			isErrorTest: true,
		},
		{
			desc:        "xとzは必須であるが、zが含まれていない",
			args:        []string{"x=1", "y=2"},
			prefix:      []string{"x", "z"},
			isErrorTest: true,
		},
		{
			desc:        "xとzは必須であるが、xが2つ含まれている",
			args:        []string{"x=1", "y=2", "x=3"},
			prefix:      []string{"x", "z"},
			isErrorTest: true,
		},
		{
			desc:        "xとzは必須であるが、xが2つ含まれている",
			args:        []string{"x=1", "x=3", "y=2"},
			prefix:      []string{"x", "z"},
			isErrorTest: true,
		},
	}
	for _, v := range tds {
		err := validate(v.args, v.prefix...)
		if v.isErrorTest {
			assert.NotNil(t, err, v.desc)
		} else {
			assert.Nil(t, err, v.desc)
		}
	}
}
