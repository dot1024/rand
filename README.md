# rand
generater for random string

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
        type of random string (default "1")
-vb
        show detail
```

command example

```bash
# generate random string by default, source is numeric characters, random length is 10
rand-gen

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
