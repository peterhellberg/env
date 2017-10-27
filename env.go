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
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// DefaultClient is the default client backed by os.Getenv
var DefaultClient = NewClient(os.Getenv)

// Client is the env client interface
type Client interface {
	Bool(key string, fallback bool) bool
	Bytes(key string, fallback []byte) []byte
	Float64(key string, fallback float64) float64
	Duration(key string, fallback time.Duration) time.Duration
	Int(key string, fallback int) int
	String(key, fallback string) string
	Strings(key string, fallback []string, seps ...string) []string
	URL(key string, fallback *url.URL) *url.URL
}

type client struct {
	Getenv func(string) string
}

// NewClient creates a new Client backed by provided func(string) string
func NewClient(f func(string) string) Client {
	return &client{f}
}

// Map is a map[string]string
type Map map[string]string

// MapClient creates a new Client backed by provided map[string]string
func MapClient(v map[string]string) Client {
	return NewClient(Func(v))
}

// Func creates a func(string) string backed by a map[string]string
func Func(v map[string]string) func(string) string {
	return func(k string) string {
		return v[k]
	}
}

// Bool returns a bool from the ENV, or fallback variable
func (c *client) Bool(key string, fallback bool) bool {
	if b, err := strconv.ParseBool(c.Getenv(key)); err == nil {
		return b
	}

	return fallback
}

// Bytes returns a slice of bytes from the ENV, or fallback variable
func (c *client) Bytes(key string, fallback []byte) []byte {
	if v := c.Getenv(key); v != "" {
		return []byte(v)
	}

	return fallback
}

// Float64 returns a float64 from the ENV, or a fallback variable
func (c *client) Float64(key string, fallback float64) float64 {
	if f, err := strconv.ParseFloat(c.Getenv(key), 64); err == nil {
		return f
	}

	return fallback
}

// Duration returns a duration from the ENV, or fallback variable
func (c *client) Duration(key string, fallback time.Duration) time.Duration {
	if d, err := time.ParseDuration(c.Getenv(key)); err == nil {
		return d
	}

	return fallback
}

// Int returns an int from the ENV, or fallback variable
func (c *client) Int(key string, fallback int) int {
	if i, err := strconv.Atoi(c.Getenv(key)); err == nil {
		return i
	}

	return fallback
}

// String returns a string from the ENV, or fallback variable
func (c *client) String(key, fallback string) string {
	if v := c.Getenv(key); v != "" {
		return v
	}

	return fallback
}

// Strings returns a slice of strings from the ENV, or fallback variable
func (c *client) Strings(key string, fallback []string, seps ...string) []string {
	if v := c.Getenv(key); v != "" {
		sep := ","

		if len(seps) > 0 {
			sep = seps[0]
		}

		return strings.Split(v, sep)
	}

	return fallback
}

// URL returns a URL from the ENV, or fallback URL if missing/invalid
func (c *client) URL(key string, fallback *url.URL) *url.URL {
	if v := c.Getenv(key); v != "" {
		u, err := url.Parse(v)
		if err != nil {
			return fallback
		}

		return u
	}

	return fallback
}

// Bool returns a bool from the ENV, or fallback variable
func Bool(key string, fallback bool) bool {
	return DefaultClient.Bool(key, fallback)
}

// Bytes returns a slice of bytes from the ENV, or fallback variable
func Bytes(key string, fallback []byte) []byte {
	return DefaultClient.Bytes(key, fallback)
}

// Float64 returns a float64 from the ENV, or a fallback variable
func Float64(key string, fallback float64) float64 {
	return DefaultClient.Float64(key, fallback)
}

// Duration returns a duration from the ENV, or fallback variable
func Duration(key string, fallback time.Duration) time.Duration {
	return DefaultClient.Duration(key, fallback)
}

// Int returns an int from the ENV, or fallback variable
func Int(key string, fallback int) int {
	return DefaultClient.Int(key, fallback)
}

// String returns a string from the ENV, or fallback variable
func String(key, fallback string) string {
	return DefaultClient.String(key, fallback)
}

// Strings returns a slice of strings from the ENV, or fallback variable
func Strings(key string, fallback []string, seps ...string) []string {
	return DefaultClient.Strings(key, fallback, seps...)
}

// URL returns a URL from the ENV, or fallback URL if missing/invalid
func URL(key string, fallback *url.URL) *url.URL {
	return DefaultClient.URL(key, fallback)
}
