package main

import (
        "fmt"
        "github.com/coreos/go-semver/semver"
        "os"
)

func main() {
        current, err := semver.NewVersion(os.Args[1])
        if err != nil {
                fmt.Println(err.Error())
        }
        switch os.Args[2] {
        case "patch":
                current.BumpPatch()
        case "minor":
                current.BumpMinor()
        case "major":
                current.BumpMajor()
        default:
                fmt.Printf("You must provide patch, minor or major")
        }
        fmt.Println(current)
}
