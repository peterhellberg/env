package env_test

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/peterhellberg/env"
)

func TestBool(t *testing.T) {
	tests := []struct {
		env string
		in  bool
		out bool
	}{
		{"t", false, true},
		{"1", false, true},
		{"", false, false},
		{"FALSE", true, false},
	}

	for _, tt := range tests {
		os.Setenv("BOOL", tt.env)

		if got := env.Bool("BOOL", tt.in); got != tt.out {
			t.Errorf(`Bool("BOOL", %v) = %v, want %v`, tt.in, got, tt.out)
		}
	}
}

func TestBoolDefault(t *testing.T) {
	in, out := true, true

	os.Clearenv()

	if got := env.Bool("BOOL_DEFAULT", in); got != out {
		t.Errorf(`Bool("BOOL_DEFAULT", %v) = %v, want %v`, in, got, out)
	}
}

func ExampleBool() {
	os.Setenv("BOOL", "t")

	fmt.Println(env.Bool("BOOL", false))
	// Output: true
}

func TestBytes(t *testing.T) {
	in, out := []byte("baz"), []byte("bar")

	os.Setenv("BYTES", string(out))

	if got := env.Bytes("BYTES", in); !bytes.Equal(got, out) {
		t.Errorf(`Bytes("BYTES", "%s") = %s, want %s`, in, got, out)
	}
}

func TestBytesDefault(t *testing.T) {
	in, out := []byte("baz"), []byte("baz")

	os.Clearenv()

	if got := env.Bytes("BYTES_DEFAULT", in); !bytes.Equal(got, out) {
		t.Errorf(`Bytes("BYTES_DEFAULT", "%s") = %s, want %s`, in, got, out)
	}
}

func ExampleBytes() {
	os.Setenv("BYTES", "foo")

	fmt.Printf("%s", env.Bytes("BYTES", nil))
	// Output: foo
}

func TestDuration(t *testing.T) {
	in, fallback, out := "5s", 10*time.Minute, 5*time.Second

	os.Setenv("DURATION", in)

	if got := env.Duration("DURATION", fallback); got != out {
		t.Errorf(`Duration("DURATION", %#v) = %v, want %v`, in, got, out)
	}
}

func TestDurationDefault(t *testing.T) {
	fallback := 123 * time.Second

	os.Clearenv()

	if got := env.Duration("DURATION_DEFAULT", fallback); got != fallback {
		t.Errorf(`Duration("DURATION_DEFAULT", %#v) = %v, want %v`, fallback, got, fallback)
	}
}

func ExampleDuration() {
	os.Setenv("DURATION", "23s")

	fmt.Printf("%s", env.Duration("DURATION", 0))
	// Output: 23s
}

func TestFloat64(t *testing.T) {
	in, out := float64(1.1), float64(2.5)

	os.Setenv("FLOAT64", "2.5")

	if got := env.Float64("FLOAT64", in); got != out {
		t.Errorf(`Float64("FLOAT64", %v) = %v, want %v`, in, got, out)
	}
}

func TestFloat64Default(t *testing.T) {
	in, out := float64(5.2), float64(5.2)

	os.Clearenv()

	if got := env.Float64("FLOAT64_DEFAULT", in); got != out {
		t.Errorf(`Float64("FLOAT64_DEFAULT", %v) = %v, want %v`, in, got, out)
	}
}

func ExampleFloat64() {
	os.Setenv("FLOAT64", "1.23")

	fmt.Println(env.Float64("FLOAT64", 0))
	// Output: 1.23
}

func TestInt(t *testing.T) {
	in, out := 1, 2

	os.Setenv("INT", strconv.Itoa(out))

	if got := env.Int("INT", in); got != out {
		t.Errorf(`Int("INT", %v) = %v, want %v`, in, got, out)
	}
}

func TestIntDefault(t *testing.T) {
	in, out := 3, 3

	os.Clearenv()

	if got := env.Int("INT_DEFAULT", in); got != out {
		t.Errorf(`Int("INT_DEFAULT", %v) = %v, want %v`, in, got, out)
	}
}

func ExampleInt() {
	os.Setenv("INT", "345")

	fmt.Println(env.Int("INT", 0))
	// Output: 345
}

func TestString(t *testing.T) {
	in, out := "baz", "bar"

	os.Setenv("STRING", out)

	if got := env.String("STRING", in); got != out {
		t.Errorf(`String("STRING", "%v") = %v, want %v`, in, got, out)
	}
}

func TestStringDefault(t *testing.T) {
	in, out := "baz", "baz"

	os.Clearenv()

	if got := env.String("STRING_DEFAULT", in); got != out {
		t.Errorf(`String("STRING_DEFAULT", "%v") = %v, want %v`, in, got, out)
	}
}

func ExampleString() {
	os.Setenv("STRING", "foo bar")

	fmt.Println(env.String("STRING", ""))
	// Output: foo bar
}
