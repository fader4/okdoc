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
			in: `
loca("foo", "bar")

# @General(
# 	ident = a.b,
# 	a.b = c.d,
# 	int = 123,
# 	foat = 123.123,
# 	bool = True,
# 	null = Null,
# 	str = "b\"a)r",
# 	arr = [1,"2",[1,"2",3],{a = "b", c = "d"}],
# 	obj = {
# 		int = 123,
# 		foat = 123.123,
# 		bool = True,
# 		null = Null,
# 		arr = [1,"2",[1,"2",3],{a.b = "c", d = "e", f = g.z}]
# 	}
# )
def main():
	@Return("OK")
	return True
`,
			want: `[{"fields":[[null,"ident"],[null,"a.b"],["int",123],["foat",123.123],["bool",true],["null",null],["str","b\\\"a)r"],["arr",[1,"2",[1,"2",3],{"a":"b","c":"d"}]],["obj",{"arr":[1,"2",[1,"2",3],{"a.b":"c","d":"e","f":"g.z"}],"bool":true,"foat":123.123,"int":123,"null":null}]],"name":"General","num_chars":317,"pos":{"end_line":19,"start_left_chars":1,"start_left_spaces":1,"start_line":3},"source_file_md5":"6ec9bcef78c725142a2a2d4fa9c0942d"},{"fields":[[null,"OK"]],"name":"Return","num_chars":13,"pos":{"end_line":21,"start_left_chars":1,"start_left_spaces":1,"start_line":21},"source_file_md5":"6ec9bcef78c725142a2a2d4fa9c0942d"}]`,
		},
	}
	for _, case_ := range cases {
		t.Run("", func(t *testing.T) {
			res, err := Parse([]byte(case_.in))
			assert.NoError(t, err)
			got, err := json.Marshal(res)
			log.Println(string(got))
			assert.NoError(t, err, "ьarshal to json")
			assert.EqualValues(t, string(case_.want), string(got))
		})
	}
}
