package nigeriageo

import (
	"sort"

	"github.com/delsahim/nigeria_geo/internal/data"
	"github.com/delsahim/nigeria_geo/internal/models"
	"github.com/delsahim/nigeria_geo/internal/validators"
)


func GetAllStates() []string {
    states := make([]string, 0, len(data.States))
    for _, state := range data.States {
        states = append(states, state.Name)
    }
    sort.Strings(states)
    return states
}

func GetState(stateName string) (models.State, bool) {
    state, exists := data.States[stateName]
    return state, exists
}

func GetLGAsInState(stateName string) ([]models.LGA, bool) {
    state, exists := data.States[stateName]
    if !exists {
        return nil, false
    }
    return state.LGAs, true
}

func ValidateState(stateName string) bool {
    return validators.IsValidState(stateName)
}

func ValidateLGAInState(lgaName, stateName string) bool {
    return validators.IsLGAInState(lgaName, stateName)
}