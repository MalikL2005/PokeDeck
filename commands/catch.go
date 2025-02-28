package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
    "math/rand"
)


type Pokemon struct {
    Id int `json:"id"`
    Exp int `json:"base_experience"`
    Height int `json:"height"`
    Weight int `json:"weight"`
    Name string `json:"name"`
    Types []Type `json:"types"`
    Stats []Stat `json:"stats"`
}


type Type struct {
    Tp struct {TypeName string `json:"name"`} `json:"type"`
}


type Stat struct {
    BaseStat int `json:"base_stat"`
    StatName struct{Name string `json:"name"`} `json:"stat"`
}



func CommandCatch (c * Config) error {
    url := "http://pokeapi.co/api/v2/pokemon/"
    if len(c.Args) < 1 {
        return errors.New("Command catch requires a pokemon´s name")
    }
    fmt.Println(c.Args[0])
    url += c.Args[0]

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return err
    }

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return err
    }
    
    if res.StatusCode != http.StatusOK {
        return errors.New("Unwell status code")
    }

    body, err := io.ReadAll(res.Body)
    if err != nil {
        return err
    }
    
    var p Pokemon
    if err = json.Unmarshal(body, &p); err != nil {
        return err
    }
    fmt.Printf("Throwing a Pokeball at %s...\n", p.Name)

    fmt.Println(p.Exp)
    r := rand.Int()
    r %= p.Exp
    var f float64 = float64(r)/(float64(p.Exp)*float64(p.Exp))*100
    fmt.Println(f)
    if f < 0.1 {
        fmt.Println("It escaped!")
        return nil
    }

    if c.PokeDeck == nil {
        m := make(map[string]Pokemon)
        c.PokeDeck = &m
    }

    if _, ok := (*c.PokeDeck)[p.Name]; ok {
        return errors.New("This pokemon is already in the pokedeck")
    }
    (*c.PokeDeck)[p.Name] = p
    fmt.Println("Added pokemon to deck")
    
    return nil
}


