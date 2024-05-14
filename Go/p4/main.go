package main

import (
    "fmt"
)

func main() {

    var x int = 1
    var y int = 2

    var zMin int = min(x,y)
    var zMax int = max(x,y)

    fmt.Printf("Type of zMin = %T\n",zMin)
    fmt.Printf("Min of %d and %d = %d\n",x,y,zMin)
    fmt.Printf("Max of %d and %d = %d\n",x,y,zMax)

}
