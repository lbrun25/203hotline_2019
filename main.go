package main

import (
    "hotline"
    "os"
    "fmt"
)

func printHelp() {
    fmt.Println("USAGE")
    fmt.Println("\t./203hotline [n k | d]")
    fmt.Println("")
    fmt.Println("DESCRIPTION")
    fmt.Println("\tn\t\tn value for the computation of C(n, k)")
    fmt.Println("\tk\t\tk value for the computation of C(n, k)")
    fmt.Println("\td\t\taverage duration of calls (in seconds)")
}

func main() {
    if hotline.CheckHelp() {
        printHelp()
        os.Exit(0)
    }
    if !hotline.CheckArgs() {
        printHelp()
        os.Exit(84)
    }
    hotline.Hotline();
}