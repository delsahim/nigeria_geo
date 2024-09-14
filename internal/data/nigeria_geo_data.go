package data

import "github.com/delsahim/nigeria_geo/internal/models"


var States = map[string]models.State{
    "Lagos": {
        Name: "Lagos",
        Code: "LA",
        LGAs: []models.LGA{
            {Name: "Alimosho", State: "Lagos"},
            {Name: "Ajeromi-Ifelodun", State: "Lagos"},
            {Name: "Kosofe", State: "Lagos"},
            {Name: "Mushin", State: "Lagos"},
            {Name: "Oshodi-Isolo", State: "Lagos"},
        },
    },
    "Kano": {
        Name: "Kano",
        Code: "KN",
        LGAs: []models.LGA{
            {Name: "Nassarawa", State: "Kano"},
            {Name: "Kano Municipal", State: "Kano"},
            {Name: "Dala", State: "Kano"},
            {Name: "Gwale", State: "Kano"},
            {Name: "Tarauni", State: "Kano"},
        },
    },
    // Add more states
}