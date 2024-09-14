package validators

import "github.com/delsahim/nigeria_geo/internal/data"



func IsValidState(state string) bool {
    _, exists := data.States[state]
    return exists
}

func IsLGAInState(lga, state string) bool {
    stateData, stateExists := data.States[state]
    if !stateExists {
        return false
    }
    return stateData.ContainsLGA(lga)
}