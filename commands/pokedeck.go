package commands

import(
    "fmt"
    "errors"
)

func CommandPokedeck (c *Config) error {
    if c == nil {
        return errors.New("Your Pokedeck is empty. Use catch to fill it!")
    }
    if len(*c.PokeDeck) == 0 {
        return errors.New("Your Pokedeck is empty. Use catch to fill it!")
    }

    fmt.Println("Your Pokedeck:")
    for key := range *c.PokeDeck {
        fmt.Println("- " + key)
    }
    
    return nil
}

