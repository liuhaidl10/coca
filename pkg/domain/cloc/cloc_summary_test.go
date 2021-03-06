package cloc

import (
	"encoding/json"
	"fmt"
	"github.com/boyter/scc/processor"
	. "github.com/onsi/gomega"
	"testing"
)

func Test_parser_json_languages(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var data = `[
  {"Name":"Java","Bytes":21169200,"CodeBytes":0,"Lines":540043,"Code":381028,"Comment":93196,"Blank":65819,"Complexity":43899,"Count":4435,"WeightedComplexity":0,"Files":[]},
  {"Name":"Kotlin","Bytes":6961705,"CodeBytes":0,"Lines":168900,"Code":118448,"Comment":30743,"Blank":19709,"Complexity":7636,"Count":1315,"WeightedComplexity":0,"Files":[]}
]
`

	var f []processor.LanguageSummary
	err := json.Unmarshal([]byte(data), &f)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
	}
	fmt.Println(f)

	g.Expect(len(f)).To(Equal(2))
}
