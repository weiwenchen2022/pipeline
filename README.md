# pipeline

Package pipeline allows you to use the Pipeline pattern in your processes.

## Install

```bash
go get github.com/weiwenchen2022/pipeline
```

## Examples
```go
package main

import (
	"fmt"

	"github.com/weiwenchen2022/pipeline"
)

func main() {
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
```

## Spec

GoDoc: [https://godoc.org/github.com/weiwenchen2022/pipeline](https://godoc.org/github.com/weiwenchen2022/pipeline)
