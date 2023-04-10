package main

import (
    "fmt"
    "os"
    "strings"
    flag "github.com/ogier/pflag"
    color "github.com/fatih/color"
)


var (
    user string
)

func main() {
    // parse the flag
    flag.Parse()

    // when no flag is specified
    if flag.NFlag() == 0 {
        fmt.Printf("Usage: %s [options]\n", os.Args[0])
        fmt.Println("Options")
        flag.PrintDefaults()
        os.Exit(1)
    }

    // split the multiple users into individual name
    users := strings.Split(user, ",")

    // Print users to the screen
    fmt.Printf("Searching user(s): %s\n", users)
    for _, u := range users {
        result := getUsers(u)
        color.Cyan(`Username:  %s`, result.Login)
        color.Blue(`Name:      %s`, result.Name)
        color.Green(`Email:        %s`, result.Email)
        color.HiMagenta(`Bio:      %s`, result.Bio)
        color.Cyan(`Location:   %s`, result.Location)
        color.Green(`Created On:   %s`, result.CreatedAt)
    }
}

func init() {
    flag.StringVarP(&user, "user", "u", "", "Search Users")
}

