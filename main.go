package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/MalikL2005/pokedexCli/pokecache"
	"github.com/MalikL2005/pokedexCli/commands"
)


func main () {
    var buffer string
    var ok bool
    var command_buff string
    var command commands.CliCommand
    scan := bufio.NewScanner(os.Stdin)
    cmds := commands.GetCommands()
    var c commands.Config
    pokedeck := make(map[string]commands.Pokemon)
    c.PokeDeck = &pokedeck
    c.Cache = pokecache.NewCache()
    c.Cache.ReadLoop()

    for {
        fmt.Print("Pokedex > ")
        c.Args = []string{}
        scan.Scan()
        buffer = strings.TrimSpace(strings.ToLower(scan.Text()))
        command_buff = strings.Split(buffer, " ")[0]
        command, ok = cmds[command_buff]
        c.Args = strings.Split(buffer, " ")[1:]
        if !ok {
            fmt.Printf("Your (unknown) command was: %s\n", command_buff)
            continue
        }
        err := command.Callback(&c)
        if err != nil {
            fmt.Println(err)
            // os.Exit(1)
        }
    }
}



func cleanInput(text string) [] string {
    text = strings.TrimSpace(text)
    strs := strings.SplitAfter(text, " ")
    for i, str := range strs{
        strs[i] = strings.TrimSpace(str)
        if str == ""{
            strs = append(strs[:i], strs[i+1:]...)
        }
    }
    fmt.Println(strs)
    return strs
}




