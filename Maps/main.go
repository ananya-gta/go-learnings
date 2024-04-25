// In maps the data type of all the keys in a map should be same and similarly for values also.
package main

import (
	"fmt"
)

// 1st way of declaring maps
// func main() {
// 	colors := map[string]string{
//     "red":"color1",
//     "green":"color2", 
//   }

//   fmt.Println(colors)
// }

// 2nd way of declaring maps
// func main() {
//   var colors map[string]string
// empty map 
//   fmt.Println(colors)
// }

// 3rd way of declaring maps
// func main() {
//   colors := make(map[string]string) // empty map

//   // access an individual key, use ["key"]
//   colors["white"] = "color1"

//   // delete a key-value pair
//   delete(colors, "white")
//   fmt.Println(colors)
// }

//iterating over a key-value pair
func main() {
  colors := map[string]string{
       "red":"color1",
      "green":"color2", 
    "white": "color3",
    }
  printMap(colors)
}

func printMap(m map[string]string) {
  for color,num := range m {
    fmt.Printf("Color is %v and color number is %v \n", color,num)
  }
}





