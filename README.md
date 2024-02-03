
# GENZ Compiler

A simple compiler written in Go. Right now it only prints a 1 byte integer :)


## Required
Make sure you are not using a soy dev operating system.

Install nvim. 

Go to a random person on the street and say " I use vim btw."

Install Go.

Make sure nasm is installed.
## Write your Code!

Create a *.genz file.

``` GENZ
return 69;
```

What else do you want? It's a language that only returns a 1 byte integer.

Maybe next version will be capable of printing 420.
## Usage/Examples

```Bash
$ go build -o gz main.go

$./gz <filename.genz>

$ nasm -felf64 out.asm

$ ld -o out out.o

$ echo $?
```

Yes, yes I know I can make it like this or like that etc etc. IMTRYING.


## Output

```
69
```

Congratulations you're a GENZ developer! 
