package main

import (
    "testing"
)

func TestHello(t *testing.T) {
    // Define a test table with input and expected output pairs
    tests := []struct {
        input    string
        expected string
    }{
        {"Alice", "Hello, Alice!"},
        {"Bob", "Hello, Bob!"},
        {"", "Hello, !"}, // Testing with an empty string
    }

    // Iterate through the test cases
    for _, test := range tests {
        t.Run(test.input, func(t *testing.T) {
            result := hello(test.input)
            if result != test.expected {
                t.Errorf("Expected %s, but got %s", test.expected, result)
            }
        })
    }
}
