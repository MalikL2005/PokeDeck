package test

import (
    // "fmt"
    "time"
    "testing"
	"github.com/MalikL2005/pokedexCli/pokecache"
)


func TestCache (t *testing.T){
    const interval = 5 * time.Second
    url := "http://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
    cases := [] struct {
        key string
    }{
        {key: url,},
        {key: url,},
    }

    cache := pokecache.NewCache()
    cache.ReadLoop()

    for _, c := range cases{
        cache.Add(c.key, []byte{})
        _, ok := cache.Get(c.key)
        if !ok {
            t.Errorf("URL (%s) was not added to cache", c.key)
            return
        }
        time.Sleep(6*time.Second)
        _, ok = cache.Get(c.key)
        if ok {
            t.Errorf("URL (%s) was not removed from cache", c.key)
            return
        }
    }

}



