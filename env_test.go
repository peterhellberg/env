package env

import (
	"bytes"
	"os"
	"strconv"
	"testing"
	"time"
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

		if got := Bool("BOOL", tt.in); got != tt.out {
			t.Errorf(`Bool("BOOL", %v) = %v, want %v`, tt.in, got, tt.out)
		}
	}
}

func TestBoolDefault(t *testing.T) {
	in, out := true, true

	os.Clearenv()

	if got := Bool("BOOL_DEFAULT", in); got != out {
		t.Errorf(`Bool("BOOL_DEFAULT", %v) = %v, want %v`, in, got, out)
	}
}

func TestBytes(t *testing.T) {
	in, out := []byte("baz"), []byte("bar")

	os.Setenv("BYTES", string(out))

	if got := Bytes("BYTES", in); !bytes.Equal(got, out) {
		t.Errorf(`Bytes("BYTES", "%s") = %s, want %s`, in, got, out)
	}
}

func TestBytesDefault(t *testing.T) {
	in, out := []byte("baz"), []byte("baz")

	os.Clearenv()

	if got := Bytes("BYTES_DEFAULT", in); !bytes.Equal(got, out) {
		t.Errorf(`Bytes("BYTES_DEFAULT", "%s") = %s, want %s`, in, got, out)
	}
}

func TestDuration(t *testing.T) {
	in, fallback, out := "5s", 10*time.Minute, 5*time.Second

	os.Setenv("DURATION", in)

	if got := Duration("DURATION", fallback); got != out {
		t.Errorf(`Duration("DURATION", %#v) = %v, want %v`, in, got, out)
	}
}

func TestDurationDefault(t *testing.T) {
	fallback := 123 * time.Second

	os.Clearenv()

	if got := Duration("DURATION_DEFAULT", fallback); got != fallback {
		t.Errorf(`Duration("DURATION_DEFAULT", %#v) = %v, want %v`, fallback, got)
	}
}

func TestFloat64(t *testing.T) {
	in, out := float64(1.1), float64(2.5)

	os.Setenv("FLOAT64", "2.5")

	if got := Float64("FLOAT64", in); got != out {
		t.Errorf(`Float64("FLOAT64", %v) = %v, want %v`, in, got, out)
	}
}

func TestFloat64Default(t *testing.T) {
	in, out := float64(5.2), float64(5.2)

	os.Clearenv()

	if got := Float64("FLOAT64_DEFAULT", in); got != out {
		t.Errorf(`Float64("FLOAT64_DEFAULT", %v) = %v, want %v`, in, got, out)
	}
}

func TestInt(t *testing.T) {
	in, out := 1, 2

	os.Setenv("INT", strconv.Itoa(out))

	if got := Int("INT", in); got != out {
		t.Errorf(`Int("INT", %v) = %v, want %v`, in, got, out)
	}
}

func TestIntDefault(t *testing.T) {
	in, out := 3, 3

	os.Clearenv()

	if got := Int("INT_DEFAULT", in); got != out {
		t.Errorf(`Int("INT_DEFAULT", %v) = %v, want %v`, in, got, out)
	}
}

func TestString(t *testing.T) {
	in, out := "baz", "bar"

	os.Setenv("STRING", out)

	if got := String("STRING", in); got != out {
		t.Errorf(`String("STRING", "%v") = %v, want %v`, in, got, out)
	}
}

func TestStringDefault(t *testing.T) {
	in, out := "baz", "baz"

	os.Clearenv()

	if got := String("STRING_DEFAULT", in); got != out {
		t.Errorf(`String("STRING_DEFAULT", "%v") = %v, want %v`, in, got, out)
	}
}
