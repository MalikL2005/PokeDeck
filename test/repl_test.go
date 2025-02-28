package main
import (
    "testing"
)

func TestCleanInput (t * testing.T){
    cases := []struct {
        input string
        expected [] string
    } {
        {
            input: " hello world ",
            expected: []string {"hello", "world"},
        },
        {
            input: " Edos whoo ",
            expected: []string {"Edos", "whoo"},
        },
    }

    for _, c := range cases {
        actual := cleanInput(c.input)
        if len(actual) != len(c.expected){
            t.Errorf("Length of slices do not match. Expected %d, got: %d", len(c.expected), len(actual))
            t.Fail()
        }

        for i := range actual {
            word := actual[i]
            expected_word := c.expected[i]
            if word != expected_word {
                t.Errorf("Expected: '%s'; Got: '%s'", expected_word, word)
                t.Fail()
            }

        }

    }


}
