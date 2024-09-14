package models


type State struct {
    Name string
    Code string
    LGAs []LGA
}

func (s *State) AddLGA(lga LGA) {
    s.LGAs = append(s.LGAs, lga)
}

func (s *State) ContainsLGA(lgaName string) bool {
    for _, lga := range s.LGAs {
        if lga.Name == lgaName {
            return true
        }
    }
    return false
}