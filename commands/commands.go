package commands
import (
    "fmt"
    "os"
    "github.com/MalikL2005/pokedexCli/pokecache"
)



type CliCommand struct {
    Name string
    Description string
    Callback func(c * Config)error
}

type Config struct {
    PreviousUrl string
    NextUrl string
    Cache pokecache.Cache
    Args []string
    PokeDeck * map[string] Pokemon
}



func CommandExit(c * Config) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}



func CommandHelp(c * Config) error {
    fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
    for _, cmd := range GetCommands() {
        fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
    }
    return nil
}



func GetCommands() map[string] CliCommand {
    var commands = map [string] CliCommand {
        "exit": {
            Name: "exit",
            Description: "Exit the Pokedex",
            Callback: CommandExit,
        },
        "help": {
            Name: "help",
            Description: "Displays a help message",
            Callback: CommandHelp,
        },
        "map": {
            Name: "map", 
            Description: "Displays all maps (20 per page)",
            Callback: CommandMap,
        },
        "mapb": {
            Name: "mapb",
            Description: "Displays all maps (20 per page), moves to last page",
            Callback: CommandMapb,
        },
        "explore": {
            Name: "explore",
            Description: "Lists all pokemon that exist on a map",
            Callback: CommandExplore,
        },
        "catch": {
            Name: "catch",
            Description: "Catches a pokemon",
            Callback: CommandCatch,
        },
        "inspect": {
            Name: "inspect",
            Description: "Inspects a pokemon",
            Callback: CommandInspect,
        },
        "pokedeck": {
            Name: "pokedex",
            Description: "Display all pokemons in deck",
            Callback: CommandPokedeck,
        },
    }
    return commands
}

