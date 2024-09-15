# Nigeria Geo Package

The Nigeria Geo Package is a Go library that provides functionality for working with Nigerian states and Local Government Areas (LGAs). It offers utilities for retrieving information about states and LGAs, as well as validation functions.

## Installation

To install the Nigeria Geo Package, use the following command:

```bash
go get github.com/delsahim/nigeria_geo
```

## Usage

Import the package in your Go code:

```go
import "github.com/delsahim/nigeria_geo"
```

## Functions

### GetAllStates() []string

Returns an array of all the states arranged in alphabetical order.

Example:
```go
states := nigeriageo.GetAllStates()
fmt.Println(states) // Output: [Abia, Adamawa, Akwa Ibom, ...]
```

### GetState(stateName string) (models.State, bool)

Returns a State object and a boolean value to show if the state exists.

Example:
```go
state, exists := nigeriageo.GetState("Lagos")
if exists {
    fmt.Printf("State: %s, Code: %s\n", state.Name, state.Code)
} else {
    fmt.Println("State not found")
}
```

### GetLGAsInState(stateName string) ([]string, bool)

Takes a state string as the parameter and returns an array of all the LGAs in the state ordered alphabetically, along with a boolean indicating if the state exists.

Example:
```go
lgas, exists := nigeriageo.GetLGAsInState("Lagos")
if exists {
    fmt.Println("LGAs in Lagos:", lgas)
} else {
    fmt.Println("State not found")
}
```

### ValidateState(stateName string) bool

Validates if a given state name exists.

Example:
```go
isValid := nigeriageo.ValidateState("Lagos")
fmt.Println("Is Lagos a valid state?", isValid) // Output: true
```

### ValidateLGAInState(lgaName, stateName string) bool

Validates if a given LGA exists in a specified state.

Example:
```go
isValid := nigeriageo.ValidateLGAInState("Alimosho", "Lagos")
fmt.Println("Is Alimosho in Lagos?", isValid) // Output: true
```

## Contributing

Contributions to the Nigeria Geo Package are welcome! Please feel free to submit a Pull Request.

## Acknowledgments

- Data sourced from https://gist.github.com/chrisidakwo/4ba3a4f03afc442305021be4ca67738e.

## Contact

For any queries or suggestions, please open an issue on the GitHub repository.