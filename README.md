[![](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/aletheia7/errors)

#### Documentation

errors augments a golang error with a file and line number of where the error
was made.

#### Example

```go
package main

import (
	"fmt"
	"github.com/aletheia7/errors"
	"log"
)

func main() {
	log.SetFlags(0)
	err := fmt.Errorf("basic golang error")
	err1 := errors.Wrap(err)
	log.Println("err1:", err1)
	log.Println("remove file/line from err1:", errors.Cause(err1))
	err2 := errors.New("a better error")
	log.Println("err2:", err2)
	err3 := errors.Errorf("error number %v", 101)
	log.Println("err3:", err3)
}
```
#### Output

err1: basic golang error t/t.go:12
remove file/line from err1: basic golang error
err2: a better error t/t.go:15
err3: error number 101 t/t.go:17

#### License 

Use of this source code is governed by a BSD-2-Clause license that can be
found in the LICENSE file.

[![BSD-2-Clause License](img/osi_logo_100X133_90ppi_0.png)](https://opensource.org/)
