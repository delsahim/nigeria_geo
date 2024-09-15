package nigeriageo

import (
	"sort"

	"github.com/delsahim/nigeria_geo/internal/data"
	"github.com/delsahim/nigeria_geo/internal/models"
	"github.com/delsahim/nigeria_geo/internal/validators"
)

// Returns an array of all the states arranged in alphabetical order
func GetAllStates() []string {
    states := make([]string, 0, len(data.States))
    for _, state := range data.States {
        states = append(states, state.Name)
    }
    sort.Strings(states)
    return states
}

// Returns a State object and a boolean value to show if the state exists
func GetState(stateName string) (models.State, bool) {
    state, exists := data.States[stateName]
    return state, exists
}

// Takes a state string as the parameter and returns an array 
//of all the lgas in the state ordered alphabetically
func GetLGAsInState(stateName string) ([]string, bool) {
    state, exists := data.States[stateName]
    if !exists {
        return nil, false
    }

    var lgas []string
    for _ ,value := range state.LGAs {
        lgas = append(lgas,value.Name)
    }
    return lgas, true
}

func ValidateState(stateName string) bool {
    return validators.IsValidState(stateName)
}

func ValidateLGAInState(lgaName, stateName string) bool {
    return validators.IsLGAInState(lgaName, stateName)
}