# rand
generator for random string

[![Release](https://img.shields.io/github/v/release/dot1024/rand)](github.com/dot1024/rand/releases)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/dot1024/rand?tab=doc)

## Usage

install it by the following command if it is need to use the random library, 

```bash
go get github.com/dot1024/rand@latest
```

random gen tool command with a command such as

```bash
go install github.com/dot1024/rand/rand-gen@latest
```

command introduction:

```
Usage of ./build/rand-gen:
-rl int
        length of random string (default 10)
-rt string
        type of random string, options value is 1000, 0100, 0010, 0001 (default "1")
-vb
        show all additional information
```

command example

```bash
# generate random string by default, source is numeric characters, random length is 10
rand-gen

# args: rt is binary number (as string), base value is belowing:
#   (8) 1000 for special characters
#   (4) 0100 for uppercase letters
#   (2) 0010 for lowercase letters
#   (1) 0001 for numeric characters
# then rt values eg:
#   0001 <- 1
#   0010 <- 2
#   0100 <- 4
#   1000 <- 8
#   1100 <- 8|4
#   0110 <- 4|2
#   0011 <- 2|1
#   1001 <- 8|1
#   1110 <- 8|4|2
#   1101 <- 8|4|1
#   1111 <- 8|4|2|1
#   1111 <- 8|4|2|1
#   ......

# source is special string, random length is 10
rand-gen -rt 0001 -rl 10

# source is lowercase letters, random length is 10
rand-gen -rt 0010 -rl 10

# source is uppercase letters, random length is 10
rand-gen -rt 0100 -rl 10

# source is special characters, random length is 10
rand-gen -rt 1000 -rl 10
```

*I hope the tool will help all those who need it. Good luck!*
