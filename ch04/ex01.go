package main

import "fmt"

func main(){

    // Declare an integer array of five elements.
    var array1 [5]int
    fmt.Println(array1)

    // Declare an integer array of five elements.
    // Initialize each each element with a specific value.
    array2 := [5]int{10, 20, 30, 40, 50}
    fmt.Println(array2)


    // Declare an integer array of five elements.
    // Initialize each each element with a specific value.
    // Capacity is determined based on the number of values initialized.
    array3 := [...]int{10, 20, 30, 40, 50}
    fmt.Println(array3)


    // Declare an integer array of five elements.
    // Initialize index 1 and 2 with specific values.
    // The rest of the elements contain their zero value.
    array4 := [5]int{1: 10, 2: 20}
    fmt.Println(array4)


    // Declare an integer array of five elements.
    // Initialize each each element with a specific value.
    array5 := [5]int{10, 20, 30, 40, 50}
    
    // Change the value at index 2.
    array5[2] = 35
    fmt.Println(array5)


    // Declare an integer array of five elements.
    // Initialize index 0 and 1 of the array with integer pointers.
    array6 := [5]*int{0: new(int), 1: new(int)}

    // Assign values to index 0 and 1.
    *array6[0] = 10
    *array6[1] = 20
    fmt.Println(array6)

    // Declare a string array of five elements.
    var array8 [5]string

    
    // Declare a second string array of five elements.
    // Initialize the array with colors.
    array9 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}

    // Copy the values from array2 into array1.
    array8 = array9

    fmt.Println(array8, array9)


    // Declare a string array of four elements.
    var array10 [4]string

    // Declare a second string array of five elements.
    array11 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}

    // Copy the values from array2 into array1.
    // Compiler Error:
    // cannot use array2 (type [5]string) as type [4] string is assignment
    //array10 = array11
    fmt.Println(array10, array11)


    // Declare a string pointer array of three elements.
    var array12 [3]*string

    // Declare a second string pointer array of the elements.
    // Initialize the array with string pointers.
    array13 := [3]*string{new(string), new(string), new(string)}

    // Add colors to each element
    *array13[0] = "Red"
    *array13[1] = "Blue"
    *array13[2] = "Green"

    // Copy the values from array2 into array1.
    array12 = array13
    fmt.Println(array12, array13)


    // Declare a two demensional integer array of four elements
    // by two elements.
    var array14 [4][2]int

    // Use an array literal to declare and initialize a two
    // demensional integer array.
    array15 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}

    // Declare and initialize index 1 and 3 of the outer array.
    array16 := [4][2]int{1: {20, 21}, 3: {40, 41}}

    // Declare and initialize individual elements of the outer
    // and inner array.
    array17 := [4][2]int{1: {0: 20}, 3: {1: 41}}
    fmt.Println(array14, array15, array16, array17);


    // Declare a two dimensional integer array of two elements.
    var array18 [2][2] int
    
    // Set integer values to each individual element.
    array18[0][0] = 10
    array18[0][1] = 20
    array18[1][0] = 30
    array18[1][1] = 40
    fmt.Println(array18)

    
    // Declare a two dimensional integer array of two elements.
    var array19 [2][2] int
    var array20 [2][2] int
    
    // Add integer values to each individual element.
    array20[0][0] = 10
    array20[0][1] = 20
    array20[1][0] = 30
    array20[1][1] = 40

    // Copy the values from array2 into array1.
    array19 = array20
    fmt.Println(array19, array20)

    
    // Copy index 1 of array1 into a new array of the same type.
    var array21 [2]int = array19[1]

    // Copy the integer found in index 1 of the outer array
    // and index 0 of the interior array into a new variable of
    // type integer.
    var value int = array19[1][0]
    fmt.Println(array21, value)


    // Declare an array of 8 megabytes.
    var array22 [1e6]int
    
    // Pass the array to the function foo.
    foo1(array22)

    // Pass the address of the array to the function foo.
    foo2(&array22)
}

// Function foo accepts an array of one milion integers.
func foo1(array [1e6]int) {
}

// Function foo accepts a pointer to an array of one million integers.
func foo2(array *[1e6]int) {
}
