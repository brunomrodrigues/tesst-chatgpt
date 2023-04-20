# Go Money Converter
This is a simple Go package for converting money values between Brazilian Reals and other currencies using the ExchangeRate-API.com service.

## Usage

To use this package in your Go project, you can import it using the following import statement:

Import "github.com/YOUR_USERNAME/go-money-converter"

Convert(brl float64, currency string) (float64, error)

This function allows you to convert money values between Brazilian Reals (BRL) and other currencies.

### Parameters
brl (required): The amount of money in Brazilian Reals (BRL) to be converted.

currency (required): The code of the currency to convert to (e.g. USD, EUR, JPY, etc.).





### Dependencies
This package uses the following dependencies:

net/http - Go's standard HTTP client library.
encoding/json - Go's standard JSON encoding and decoding library.
License
This package is released under the MIT License.
