package commands

import (
    "net/http"
    "log"
    "fmt"
    "io"
    "encoding/json"
    "errors"
)


type mapResponse struct {
    Count int `json:"count"`
    Next string `json:"next"` 
    Previous string `json:"previous"`
    Results []mapResult `json:"results"`
}

type mapResult struct {
    Name string `json:"name"`
    Url string `json:"url"`
}


func CommandMap(c * Config) error {
    if c.NextUrl == "" && c.PreviousUrl == "" {
        c.NextUrl = "http://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
    }

    // check for cache
    if data, ok := c.Cache.Get(c.NextUrl); ok {
        var mR mapResponse
        err  := json.Unmarshal(data, &mR)
        if err != nil {
            log.Println("Error unmarshalling")
            return err
        }
        c.PreviousUrl = c.NextUrl
        c.NextUrl = mR.Next

        fmt.Println("Got data from cache :)")
        for _, res := range mR.Results {
            fmt.Println(res.Name)

        }
    }

    req, err := http.NewRequest("GET", c.NextUrl, nil)
    if err != nil {
        log.Println("Request error")
        return err
    }

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Println("Response error")
        return err
    }
    defer res.Body.Close()

    res_body, err := io.ReadAll(res.Body)
    if err != nil {
        log.Println("Error Parsing body")
        return err
    }

    var mR mapResponse
    err  = json.Unmarshal(res_body, &mR)
    if err != nil {
        log.Println("Error unmarshalling")
        return err
    }

    c.PreviousUrl = c.NextUrl
    c.NextUrl = mR.Next

    for _, res := range mR.Results {
        fmt.Println(res.Name)

    }

    // add to cache 
    err = c.Cache.Add(c.PreviousUrl, res_body)
    if err != nil {
        return errors.New("Cache-Error: Could not add to cache")
    }

    return nil
}



func CommandMapb (c * Config) error {
    if c.PreviousUrl == ""{
        fmt.Println("you're on the first page")
        return nil
    }

    // check for cache 
    if data, ok := c.Cache.Get(c.PreviousUrl); ok {
        var mR mapResponse
        err  := json.Unmarshal(data, &mR)
        if err != nil {
            log.Println("Error unmarshalling")
            return err
        }
        c.NextUrl = c.PreviousUrl
        c.PreviousUrl = mR.Previous

        fmt.Println("Got data from cache :)")
        for _, res := range mR.Results {
            fmt.Println(res.Name)

        }
        return nil
    }

    req, err := http.NewRequest("GET", c.PreviousUrl, nil)
    if err != nil {
        log.Println("Error creating request")
        return err
    }
    
    res, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Println("Error during response")
        return err
    }
    defer res.Body.Close()

    res_body, err := io.ReadAll(res.Body)
    if err != nil {
        log.Println("Error during reading of res.body")
        return err
    }

    var mR mapResponse
    err = json.Unmarshal(res_body, &mR)
    if err != nil {
        log.Println("Error during unmarshaling")
        return err
    }

    c.NextUrl = c.PreviousUrl
    c.PreviousUrl = mR.Previous

    for _, res := range mR.Results {
        fmt.Println(res.Name)
    }

    // add to cache 
    err = c.Cache.Add(c.PreviousUrl, res_body)
    if err != nil {
        return errors.New("Cache-Error: Could not add to cache")
    }

    return nil
}


