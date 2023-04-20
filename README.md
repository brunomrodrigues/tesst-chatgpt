# Go Money Converter
This is a simple Go package for converting money values between Brazilian Reals and other currencies using the ExchangeRate-API.com service.

## Usage
To use this package in your Go project, you can import it using the following import statement:

```go
import "github.com/YOUR_USERNAME/go-money-converter"
```

### Convert(brl float64, currency string) (float64, error)
This function allows you to convert money values between Brazilian Reals (BRL) and other currencies.

#### Parameters
- brl (required): The amount of money in Brazilian Reals (BRL) to be converted
- currency (required): The code of the currency to convert to (e.g. USD, EUR, JPY, etc.)

#### Example
To convert 100 Brazilian Reals to US Dollars, you can use the following code:

```go
package main

import (
	"fmt"
	"github.com/YOUR_USERNAME/go-money-converter"
)

func main() {
	converted, err := money.Convert(100, "USD")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("100 BRL = %.2f USD\n", converted)
}
```
This will output:

100 BRL = 18.37 USD

#### Dependencies
This package uses the following dependencies:

net/http - Go standard HTTP client library.
encoding/json - Go standard JSON encoding and decoding library.

#### License
This package is released under the MIT License.
