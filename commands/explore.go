package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)


type exploreMapResponse struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Encounters []struct{Pokemon explorePokemon `json:"pokemon"`} `json:"pokemon_encounters"`
}

type explorePokemon struct {
    Name string `json:"name"`
    Url string `json:"url"`
}



func CommandExplore (c *Config) error{
    fmt.Println(c.Args)
    url := "http://pokeapi.co/api/v2/location-area/"
    if len(c.Args) < 1 {
        return errors.New("Error: Need to pass an Argument to explore-command")
    }
    url += c.Args[0]
    fmt.Println(url)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return err
    }
    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return err
    }

    if res.StatusCode != http.StatusOK {
        return errors.New("Unwell statuscode")
    }

    body, err := io.ReadAll(res.Body)
    if err != nil {
        return err
    }

    var eMR exploreMapResponse
    err = json.Unmarshal(body, &eMR)
    if err != nil {
        return err
    }

    fmt.Println("Found Pokemon:")
    for _, poke := range eMR.Encounters {
        fmt.Printf(" - %s\n", poke.Pokemon.Name)
    }


    return nil
}

