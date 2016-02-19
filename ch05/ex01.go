package main

import "fmt"

// user defines a user in the program.
type user struct {
    name       string
    email      string
    ext        int
    privileged bool
}

// admin represents an admin user with privileges.
type admin struct {
    person user
    level  string
}

func main(){

    // Declare a variable of type user.
    var bill user
    fmt.Println(bill)
    
    // Declare a variable of type user and initialize all the fields.
    lisa1 := user {
        name:       "Lisa",
        email:      "lisa@email.com",
        ext:        123,
        privileged: true,
    }
    fmt.Println(lisa1)


    // Declare a variable of type user.
    lisa2 := user{"Lisa", "lisa@email.com", 123, true}
    fmt.Println(lisa2)

    
    // Declare a variable of type admin.
    fred1 := admin{
        person: user{
            name:       "Lisa",
            email:      "lisa@email.com",
            ext:        123,
            privileged: true,
        },
        level: "super",
    }
    fmt.Println(fred1)
}
