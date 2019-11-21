# Golang

Is an open source programming language that makes it easy to build simple, reliable,and efficient software.

The purpose of this repository is to be able to understand and perform different golang exercises, so I can perform in different languages.

For this I use the Holberton methodology, which has helped me to understand low and high level languages.

# Data types

## Signed Integers

| Type | Size | Range |
| ---- | ---- | ----- |
| int8 | 8 bits | -128 to 127 |
| int16 | 16 bits | -2^15 to 2^15 -1 |
| int32 | 32 bits | -2^31 to 2^31 -1 |
| int64 | 64 bits | -2^63 to 2^63 -1 |
| int | Plataform dependent | Plataform dependent |

The size of the generic int type is platform dependent. It is 32 bits wide on a 32-bit system and 64-bits wide on a 64-bit system.

## Unsigned Integers

| Type | Size | Range |
| ---- | ---- | ----- |
| uint8 | 8 bits | 0 to 127 |
| uint16 | 16 bits | 0 to 2^16 -1 |
| uint32 | 32 bits | 0 to 2^32 -1 |
| uint64 | 64 bits | 0 to 2^64 -1 |
| uint | Plataform dependent | Plataform dependent |

## Integer Type aliases

| Type | Alias For |
| ---- | --------- |
| byte | uint8 |
| rune | int32 |

for example:

``` golang
package main
import "fmt"

func main() {
	var myByte byte = 'a'
	var myRune rune = '♥'

	fmt.Printf("%c = %d and %c = %U\n", myByte, myByte, myRune, myRune)
}
```

Output

``` bash
a = 97 and ♥ = U+2665
```

## Floating Point Types

Floating point types are used to store numbers with a decimal component (ex - 1.24, 4.50000). Go has two floating point types - float32 and float64.

* float32 occupies 32 bits in memory and stores values in single-precision floating point format
* float64 occupies 64 bits in memory and stores values in double-precision floating point format

## Operations on Numeric Types

Go provides several operators for performing operations on numeric types

* Arithmetic Operators: +, -, *, /, %
* Comparation Operators: ==, !=, <, >, <=, >=
* Bitwise Operators: &, |, ^, <<, >>
* Increment and Decrement Operators: ++, --
* Assignment Operators: +=, -=, *=, /=, %=, <<=, >>=, &=, |=, ^=
