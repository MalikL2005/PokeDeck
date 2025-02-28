package pokecache

import (
	"time"
    "errors"
    "sync"
    "fmt"
)



type Cache struct {
    entries map[string] CacheEntry
    tk * time.Ticker
    mut sync.Mutex
}


type CacheEntry struct {
    createdAt time.Time
    val []byte
}



func NewCache () Cache {
    m := make(map[string] CacheEntry)
    tk := time.NewTicker(5*time.Second)
    ch := Cache{m, tk, sync.Mutex{}}
    // ch.ReadLoop()
    return ch
}



func (c * Cache) Add (key string, val []byte) error {
    cE := CacheEntry{time.Now(), val}
    _, ok := c.entries[key]
    if ok {
        return errors.New("Cannot add entry to cache. Key is already in cache.")
    }
    c.mut.Lock()
    defer c.mut.Unlock()
    if len(c.entries) == 0{
        c.entries = make(map[string]CacheEntry, 1)
    }
    c.entries[key] = cE
    return nil
}



func (c * Cache) Get (key string) ([]byte, bool) {
    cE, ok := c.entries[key]
    if !ok {
        return []byte{}, false
    }

    // display Cache metadata
    fmt.Print("Created at: ")
    fmt.Println(cE.createdAt)
    ago := time.Now().Sub(cE.createdAt)
    fmt.Print(ago)
    fmt.Println(" ago")

    return cE.val, true
}



func (c * Cache) ReadLoop (){
    go func (){
        for t := range c.tk.C{
            for url, entry := range c.entries {
                fmt.Print(url)
                fmt.Println(t.Sub(entry.createdAt))
                if t.Sub(entry.createdAt) > 5 {
                    delete(c.entries, url)
                }
            }
        }
    }()
}



