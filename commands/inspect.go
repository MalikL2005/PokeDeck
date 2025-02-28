package commands

import(
    "fmt"
    "errors"
)


func CommandInspect (c *Config) error {
    if c.PokeDeck == nil {
        return errors.New("you have not caught that pokemon")
    }

    if len(c.Args) < 1 {
        return errors.New("Command Inspect needs a name as an argument")
    }
    p, ok := (*c.PokeDeck)[c.Args[0]]
    if !ok {
        return errors.New("you have not caught that pokemon")
    }
    fmt.Printf("Name: %s\n", p.Name)
    fmt.Printf("Height: %d\n", p.Height)
    fmt.Printf("Weight: %d\n", p.Weight)
    displayStats(p)
    displayTypes(p)
    return nil
}



func displayTypes (p Pokemon){
    fmt.Println("Types:")
    for _, t := range p.Types {
        fmt.Printf(" - %s\n", t.Tp.TypeName)
    }
}



func displayStats (p Pokemon) {
    fmt.Println("Stats:")
    for _, stat := range p.Stats {
        fmt.Printf(" -%s: %d\n", stat.StatName.Name, stat.BaseStat)
    }
}

