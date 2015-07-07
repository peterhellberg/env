/*

Package env loads environment variables into Go types, with fallback values.

Installation

Just go get the package:

    go get -u github.com/peterhellberg/env

Usage

A small usage example

    package main

    import (
    	"fmt"

    	"github.com/peterhellberg/env"
    )

    func main() {
    	fmt.Println(
    		env.Bool("BOOL", false),
    		env.Bytes("BYTES", []byte{4, 2}),
    		env.Duration("DURATION", 250000),
    		env.Int("INT", 1337),
    		env.String("STRING", "Foobar"),
    	)
    }

*/
package env

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// Bool returns a bool from the ENV, or fallback variable
func Bool(key string, fallback bool) bool {
	if b, err := strconv.ParseBool(os.Getenv(key)); err == nil {
		return b
	}

	return fallback
}

// Bytes returns a slice of bytes from the ENV, or fallback variable
func Bytes(key string, fallback []byte) []byte {
	if v := os.Getenv(key); v != "" {
		return []byte(v)
	}

	return fallback
}

// Float64 returns a float64 from the ENV, or a fallback variable
func Float64(key string, fallback float64) float64 {
	if f, err := strconv.ParseFloat(os.Getenv(key), 64); err == nil {
		return f
	}

	return fallback
}

// Duration returns a duration from the ENV, or fallback variable
func Duration(key string, fallback time.Duration) time.Duration {
	if d, err := time.ParseDuration(os.Getenv(key)); err == nil {
		return d
	}

	return fallback
}

// Int returns an int from the ENV, or fallback variable
func Int(key string, fallback int) int {
	if i, err := strconv.Atoi(os.Getenv(key)); err == nil {
		return i
	}

	return fallback
}

// String returns a string from the ENV, or fallback variable
func String(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}

// Strings returns a slice of strings from the ENV, or fallback variable
func Strings(key string, fallback []string) []string {
	if v := os.Getenv(key); v != "" {
		return strings.Split(v, ",")
	}

	return fallback
}
