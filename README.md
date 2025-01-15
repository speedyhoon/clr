# <img src="/logo.webp" alt="CLR logo" width="200">
[![Go Report Card](https://goreportcard.com/badge/github.com/speedyhoon/numnam)](https://goreportcard.com/report/github.com/speedyhoon/numnam)

Efficient 16-bit and 32-bit color representation types and conversion between hexadecimal, RGB, RGBA and CSS named-colors.

## 16-bit colors
`C16` has four 4-bit channels (red, green, blue and alpha transparency) to represent 3 and 4-digit hexadecimal colors.
* `#f60`
* `#f60f`

All CSS named-colors, RGB and RGBA colors can be converted to `C16` with some degree of accuracy loss for non-web-safe colors.
    
## 32-bit colors
`C32` has four 8-bit channels (red, green, blue and alpha transparency) to represent CSS named-colors, RGB, RGBA or 6 and 8-digit hexadecimal colors.
* `darkviolet`
* `rgba(255,128,0)`
* `rgba(255,128,0,255)`
* `#ff6600`
* `#ff6600ff`

All 16-bit colors can be converted to `C32` without **any** color accuracy loss.

## Example

```go
package main

import (
	"fmt"
	"github.com/speedyhoon/clr"
)

func main() {
	a := clr.BlanchedAlmond
	b, err := clr.FromHex6("#d2b48c")
	fmt.Println(a, b, b.RGBA(), err)
}
```
Outputs:
```text
#ffebcd tan rgb(210,180,140,255) #cfb <nil>
```
