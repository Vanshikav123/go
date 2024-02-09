package main

import (
	"io"
	"os"
)

//we can create either custom panic or  a system panic can be created

/*Go has the usual mechanisms for control flow: if, for, switch, goto. It also has the go statement to run code in a separate goroutine. Here I’d like to discuss some of the less common ones: defer, panic, and recover.

A defer statement pushes a function call onto a list. The list of saved calls is executed after the surrounding function returns. Defer is commonly used to simplify functions that perform various clean-up actions.

For example, let’s look at a function that opens two files and copies the contents of one file to the other:*/

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}

	written, err = io.Copy(dst, src)
	dst.Close()
	src.Close()
	return
}

/*This works, but there is a bug. If the call to os.Create fails, the function will return without closing the source file.
This can be easily remedied by putting a call to src.Close before the second return statement,
 but if the function were more complex the problem might not be so easily noticed and resolved.
 By introducing defer statements we can ensure that the files are always closed:*/

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

/*Defer statements allow us to think about closing each file right after opening it, guaranteeing that, regardless of the number of return statements in the function, the files will be closed.

The behavior of defer statements is straightforward and predictable. There are three simple rules:

A deferred function’s arguments are evaluated when the defer statement is evaluated.
In this example, the expression “i” is evaluated when the Println call is deferred. The deferred call will print “0” after the function returns.

func a() {
    i := 0
    defer fmt.Println(i)
    i++
    return
}
Deferred function calls are executed in Last In First Out order after the surrounding function returns.
This function prints “3210”:

func b() {
    for i := 0; i < 4; i++ {
        defer fmt.Print(i)
    }
}
Deferred functions may read and assign to the returning function’s named return values.
In this example, a deferred function increments the return value i after the surrounding function returns. Thus, this function returns 2:

func c() (i int) {
    defer func() { i++ }()
    return 1
}
This is convenient for modifying the error return value of a function; we will see an example of this shortly.*/
