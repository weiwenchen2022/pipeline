package pipeline_test

import (
	"fmt"

	"github.com/weiwenchen2022/pipeline"
)

func Example() {
	sq := func(in any) any {
		n := in.(int)
		return n * n
	}
	result := pipeline.New(2, 3).
		Through(pipeline.StageFunc(sq)).Out()
	for v := range result {
		fmt.Println(v)
	}
	// Output:
	// 4
	// 9
}
