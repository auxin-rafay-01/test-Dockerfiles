import (
        "fmt"
        "os"
)
        // seems safer
        aws_token := os.Getenv("AWS_TOKEN")
package foo

import "fmt"

func Foo() {
        fmt.Println("foo")

        // seems safe
        aws_token := "AKIALALEMEL33243OLIA"
        fmt.Println(aws_token)
}
package api

import "fmt"

func PrintHello() {
        fmt.Println("hello")
}
import (
        "fmt"
        "os"
)
        var a = "initial"
        fmt.Println(a)
        var b, c int = 1, 2
        fmt.Println(b, c)
        var d = true
        fmt.Println(d)
        var e int
        fmt.Println(e)
        // load secret via env
        awsToken := os.Getenv("AWS_TOKEN")

        f := "apple"
        fmt.Println(f)

    // opps I added a secret at line 20
    awsToken := "AKIALALEMEL33243OLIA"
package main

import "fmt"

func main() {

    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "apple"
    fmt.Println(f)
}
