package main

import "fmt"

func main() {
    ids := []int{33,76,54,12,11,1}

    // Loop through ids
    for i, id := range ids {
        fmt.Printf("%d - ID: %d\n", i, id)
    }

    // Not using index
    for _, id := range ids {
        fmt.Printf("ID: %d\n", id)
    }

    // Add ids together
    sum := 0
    for _, id := range ids {
        sum += id
    }
    fmt.Println("Sum", sum)

    // Range with map
    emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com", "Mike": "mike@gmailcom"}

    for k, v := range emails {
        fmt.Printf("%s: %s\n", k,v)
    }

    // Can also just grab keys
    for k := range emails {
        fmt.Println("Name: " + k)
    }
}