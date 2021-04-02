package annotation

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	annotationDebug = 1

	cases := []struct {
		in   string
		want string
	}{
		{
			in:   "\n\t@Foo(bar)",
			want: `[{"fields":[[null,"bar"]],"name":"Foo","num_chars":9,"pos":{"end_line":1,"start_left_chars":1,"start_left_spaces":1,"start_line":1}}]`,
		},
		{
			in:   "\n@Foo(bar)",
			want: `[{"fields":[[null,"bar"]],"name":"Foo","num_chars":9,"pos":{"end_line":1,"start_left_chars":0,"start_left_spaces":0,"start_line":1}}]`,
		},
		{
			in:   "@Foo(bar)",
			want: `[{"fields":[[null,"bar"]],"name":"Foo","num_chars":9,"pos":{"end_line":0,"start_left_chars":0,"start_left_spaces":0,"start_line":0}}]`,
		},
		{
			in:   " @Foo(bar)",
			want: `[{"fields":[[null,"bar"]],"name":"Foo","num_chars":9,"pos":{"end_line":0,"start_left_chars":1,"start_left_spaces":1,"start_line":0}}]`,
		},
		{
			in:   "\t@Foo(bar)",
			want: `[{"fields":[[null,"bar"]],"name":"Foo","num_chars":9,"pos":{"end_line":0,"start_left_chars":1,"start_left_spaces":1,"start_line":0}}]`,
		},

		{
			in:   "@Foo(bar,baz)",
			want: `[{"fields":[[null,"bar"],[null,"baz"]],"name":"Foo","num_chars":13,"pos":{"end_line":0,"start_left_chars":0,"start_left_spaces":0,"start_line":0}}]`,
		},
		{
			in:   `@Foo(bar="a")`,
			want: `[{"fields":[["bar","a"]],"name":"Foo","num_chars":13,"pos":{"end_line":0,"start_left_chars":0,"start_left_spaces":0,"start_line":0}}]`,
		},
		{
			in:   `@Foo(bar=a)`,
			want: `[{"fields":[["bar","a"]],"name":"Foo","num_chars":11,"pos":{"end_line":0,"start_left_chars":0,"start_left_spaces":0,"start_line":0}}]`,
		},
		{
			in:   `@Foo(a.b=c.d)`,
			want: `[{"fields":[["a.b","c.d"]],"name":"Foo","num_chars":13,"pos":{"end_line":0,"start_left_chars":0,"start_left_spaces":0,"start_line":0}}]`,
		},
		{
			in:   `@Foo(a.b=[1,2,3])`,
			want: `[{"fields":[["a.b",[1,2,3]]],"name":"Foo","num_chars":17,"pos":{"end_line":0,"start_left_chars":0,"start_left_spaces":0,"start_line":0}}]`,
		},
		{
			in:   `@Foo(a.b={a = "b"})`,
			want: `[{"fields":[["a.b",{"a":"b"}]],"name":"Foo","num_chars":19,"pos":{"end_line":0,"start_left_chars":0,"start_left_spaces":0,"start_line":0}}]`,
		},
		{
			in:   `@Foo(a.b=True)`,
			want: `[{"fields":[["a.b",true]],"name":"Foo","num_chars":14,"pos":{"end_line":0,"start_left_chars":0,"start_left_spaces":0,"start_line":0}}]`,
		},
		{
			in:   `@Foo(a.b=False)`,
			want: `[{"fields":[["a.b",false]],"name":"Foo","num_chars":15,"pos":{"end_line":0,"start_left_chars":0,"start_left_spaces":0,"start_line":0}}]`,
		},
		{
			in:   `@Foo(a.b=Null)`,
			want: `[{"fields":[["a.b",null]],"name":"Foo","num_chars":14,"pos":{"end_line":0,"start_left_chars":0,"start_left_spaces":0,"start_line":0}}]`,
		},
		{
			in:   `@Foo(a.b="")`,
			want: `[{"fields":[["a.b",""]],"name":"Foo","num_chars":12,"pos":{"end_line":0,"start_left_chars":0,"start_left_spaces":0,"start_line":0}}]`,
		},
		{
			in:   `@Foo(a.b=1)`,
			want: `[{"fields":[["a.b",1]],"name":"Foo","num_chars":11,"pos":{"end_line":0,"start_left_chars":0,"start_left_spaces":0,"start_line":0}}]`,
		},
		{
			in:   `@Foo(a.b=-1)`,
			want: `[{"fields":[["a.b",-1]],"name":"Foo","num_chars":12,"pos":{"end_line":0,"start_left_chars":0,"start_left_spaces":0,"start_line":0}}]`,
		},
		{
			in:   `@Foo(a.b=1.1)`,
			want: `[{"fields":[["a.b",1.1]],"name":"Foo","num_chars":13,"pos":{"end_line":0,"start_left_chars":0,"start_left_spaces":0,"start_line":0}}]`,
		},
		{
			in:   `@Foo(a.b=-1.1)`,
			want: `[{"fields":[["a.b",-1.1]],"name":"Foo","num_chars":14,"pos":{"end_line":0,"start_left_chars":0,"start_left_spaces":0,"start_line":0}}]`,
		},

		{
			in:   `@Foo(a.b={a="",b="b",c=d.e,f=1,g=1.2,arr=[],arr2=[1],obj={},obj1={a="c"}})`,
			want: `[{"fields":[["a.b",{"a":"","arr":[],"arr2":[1],"b":"b","c":"d.e","f":1,"g":1.2,"obj":{},"obj1":{"a":"c"}}]],"name":"Foo","num_chars":74,"pos":{"end_line":0,"start_left_chars":0,"start_left_spaces":0,"start_line":0}}]`,
		},
	}
	for _, case_ := range cases {
		t.Run(case_.in, func(t *testing.T) {
			res, err := Parse([]byte(case_.in))
			assert.NoError(t, err)
			got, err := json.Marshal(res)
			log.Println(string(got))
			assert.NoError(t, err, "ьarshal to json")
			assert.EqualValues(t, string(case_.want), string(got))
		})
	}
}
