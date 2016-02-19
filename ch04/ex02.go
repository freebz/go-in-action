package main

import "fmt"

func main(){
    
    // Create a slice of strings.
    // Contains a length and capacity of 5 elements.
    slice1 := make([]string, 5)
    fmt.Println(slice1)

    
    // Create a slice of integers.
    // Contains a length of 3 and has a capacity of 5 elements.
    slice2 := make([]int, 3, 5)
    fmt.Println(slice2)

    
    // Create a slice of strings.
    // Contains a length and capacity of 5 elements.
    slice3 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
    
    // Create a slice of integers.
    // Contains a length and capacity of 3 elements.
    slice4 := []int{10, 20, 30}
    fmt.Println(slice3, slice4)
    

    // Create a slice of strings.
    // Initialize the 100th element with an empty string.
    slice5 := []string{99: ""}
    fmt.Println(slice5)


    // Create an array of three integers.
    array6 := [3]int{10, 20, 30}
    
    // Create a slice of integers with a length and capacity of three.
    slice6 := []int{10, 20, 30}
    fmt.Println(array6, slice6)

    
    // Create a nil slice of integers.
    var slice7 []int
    fmt.Println(slice7)


    // Use make to create an empty slice of integers.
    slice8 := make([]int, 0)

    // Use a slice literal to create an empty slice of integers.
    slice9 := []int{}
    fmt.Println(slice8, slice9)


    // Create a slice of integers.
    // Contains a length and capacity of 5 elements.
    slice10 := []int{10, 20, 30, 40, 50}

    // Change the value of index 1.
    slice10[1] = 25
    fmt.Println(slice10)


    // Create a slice of integers.
    // Contains a length and capacity of 5 elements.
    slice11 := []int{10, 20, 30, 40, 50}
    
    // Create a new slice.
    // Contains a length of 2 and capacity of 4 elements.
    slice12 := slice11[1:3]
    fmt.Println(slice11, slice12)


    // Create a slice of integers.
    // Contains a length and capacity of 5 elements.
    slice13 := []int{10, 20, 30, 40, 50}

    // Create a new slice.
    // Contains a length of 2 and capacity of 4 elements.
    slice14 := slice13[1:3]

    // Change index 1 of newSlice.
    // Change index 2 of the original slice.
    slice14[1] = 35
    fmt.Println(slice13, slice14)


    // Create a slice of integers.
    // Contains a length and capacity of 5 elements.
    slice15 := []int{10, 20, 30, 40, 50}

    // Create a new slice.
    // Contains a length of 2 and capacity of 4 elements.
    slice16 := slice15[1:3]

    // Allocate a new element from capacity.
    // Assign the value of 60 to the new element.
    slice16 = append(slice16, 60)
    fmt.Println(slice15, slice16)


    // Create a slice of integers.
    // Contains a length and capacity of 4 elements.
    slice17 := []int{10, 20, 30, 40}

    // Append a new value to the slice.
    // Assign the value of 50 to the new element.
    slice18 := append(slice17, 50)
    fmt.Println(slice17, slice18)


    // Create a slice of strings.
    // Contains a length and capacity of 5 elements.
    source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
    slice19 := source[2:3:4]
    fmt.Println(slice19)

    // Slice the third element and restrict the capacity.
    // Contains a length and capacity of 1 element.
    slice20 := source[2:3:3]

    // Append a new string to the slice.
    slice20 = append(slice20, "Kiwi")
    fmt.Println(slice20)

    
    // Create two slices each initialized with two integers.
    s1 := []int{1, 2}
    s2 := []int{3, 4}

    // Append the two slices together and display the results.
    fmt.Printf("%v\n", append(s1, s2...))


    // Create a slice of integers.
    // Contains a length and capacity of 4 elements.
    slice21 := []int{10, 20, 30, 40}
    
    // Iterate over each element and display each value.
    for index, value := range slice21 {
        fmt.Printf("Index: %d Value: %d\n", index, value)
    }

    // Iterate over each element and display the value and addresses.
    for index, value := range slice21 {
        fmt.Printf("Value: %d Value-Addr: %X  ElemAddr: %X\n",
            value, &value, &slice21[index])
    }

    // Iterate over each element and display each value.
    for _, value := range slice21 {
        fmt.Printf("Value: %d\n", value)
    }

    // Iterate over each element starting at element 3.
    for index := 2; index < len(slice21); index++ {
        fmt.Printf("Index: %d  Value: %d\n", index, slice21[index])
    }


    // Create a slice of a slice of integers.
    slice22 := [][]int{{10}, {100, 200}}
    fmt.Println(slice22)

    // Append the value of 20 to the first slice of integers.
    slice22[0] = append(slice22[0], 20)
    fmt.Println(slice22)

    
    // Allocate a slice of 1 million integers.
    slice23 := make([]int, 1e6)

    // Pass the slice to the function foo.
    slice23 = foo(slice23)
}

// Function foo accepts a slice of integers and returns the slice back.
func foo(slice []int) []int {
    return slice
}
