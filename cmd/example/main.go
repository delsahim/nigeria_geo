package main

import (
	"fmt"

	"github.com/delsahim/nigeria_geo/internal/data"
	"github.com/delsahim/nigeria_geo/pkg/nigeriageo"
)

func main() {
    fmt.Println("All states:", nigeriageo.GetAllStates())
	fmt.Println(len(data.States))

    state := "Lagos"
    lgas, exists := nigeriageo.GetLGAsInState(state)
    if exists {
        fmt.Printf("LGAs in %s: %v\n", state, lgas)
    } else {
        fmt.Printf("State %s not found\n", state)
    }

    fmt.Println("Is 'Lagos' a valid state?", nigeriageo.ValidateState("Lagos"))
    fmt.Println("Is 'Alimosho' in Lagos?", nigeriageo.ValidateLGAInState("Alimosho", "Lagos"))
    fmt.Println("Is 'InvalidLGA' in Lagos?", nigeriageo.ValidateLGAInState("InvalidLGA", "Lagos"))
}