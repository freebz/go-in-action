package main

import "fmt"

func main(){

    // Create a map with a eky of type string and a value of type int.
    dict1 := make(map[string]int)

    // Create a map with a key and value of type string.
    dict2 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
    fmt.Println(dict1, dict2)


    // Create a map suing a slice of strings as the key.
    // Compiler Exception: invalid map key type []string
    // dict3 := map[[]string]int{}

    
    // Create a map using a slice of strings as the value.
    dict3 := map[int][]string{}
    fmt.Println(dict3)


    // Create an empty map to store colors and thier color codes.
    colors1 := map[string]string{}

    // Add the Red color code to the map.
    colors1["Red"] = "#da1337"
    fmt.Println(colors1)


    // Create a nil map by just declaring the map.
    var colors2 map[string]string

    // Add the Red color code to the map.
    // Rumtime Error: assignment to entry in nil map
    // colors2["Red"] = "#da1337"
    fmt.Println(colors2)


    // Retrieve the value for the key "Blue".
    value, exists := colors1["Blue"]
    
    // Did this key exist?
    if exists {
        fmt.Println(value)
    }
    fmt.Println(value, exists)


    // Create a map of colors and color hex codes.
    colors3 := map[string]string{
        "AliceBlue":   "#f0f8ff",
        "Coral":       "#ff7F50",
        "DarkGray":    "#a9a9a9",
        "ForestGreen": "#228b22",
    }

    // Display all the colors in the map.
    for key, value := range colors3 {
        fmt.Printf("Key: %s  Value: %s\n", key, value)
    }


    // Remove the key/value pair for the key "Coral".
    delete(colors3, "Coral")
    
    // Display all the colors in the map.
    for key, value := range colors3 {
        fmt.Printf("Key: %s  Value: %s\n", key, value)
    }


    // Create a map of colors and color hex codes.
    colors4 := map[string]string{
        "AliceBlue":   "#f0f8ff",
        "Coral":       "#ff7F50",
        "DarkGray":    "#a9a9a9",
        "ForestGreen": "#228b22",
    }

    // Display all the colors in the map.
    for key, value := range colors4 {
        fmt.Printf("Key: %s  Value: %s\n", key, value)
    }    

    // Call the function to remove the specified key.
    removeColor(colors4, "Coral")
    
    // Display all the colors in the map.
    for key, value := range colors4 {
        fmt.Printf("Key: %s  Value: %s\n", key, value)
    }    
}

// removeColor removes keys from the specified map.
func removeColor(colors map[string]string, key string) {
    delete(colors, key)
}
