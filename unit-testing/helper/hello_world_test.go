package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("before test")

	m.Run()

	fmt.Println("after testing")
}

func TestHello(t *testing.T) {
	sayHello := HelloWorld("Dyaksa")
	assert.Equal(t, "Hello World Dyaksa", sayHello, "result must be 'Hello World Dyaksa'")
}

func TestSayHello(t *testing.T) {
	sayHai := SayHello()
	require.Equal(t, "Say Hello!", sayHai, "result must be 'Say Hello!'")
}

func BenchMarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Dyaksa")
	}
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("cannot run on macos")
	}
	SayHello := SayHello()
	assert.Equal(t, "Say Hello!", SayHello, "result must be 'Say Hello!'")
}

func TestSubTest(t *testing.T) {
	t.Run("Dyaksa", func(t *testing.T) {
		sayHello := HelloWorld("Dyaksa")
		assert.Equal(t, "Hello World Dyaksa", sayHello, "result must be 'Hello World Dyaksa'")
	})

	t.Run("Tifani", func(t *testing.T) {
		sayHello := HelloWorld("Tifani")
		assert.Equal(t, "Hello World Tifani", sayHello, "result must be 'Hello World Tifani'")
	})
}

func TestHelloWorldTable(t *testing.T) {

	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Dyaksa)",
			request:  "Dyaksa",
			expected: "Hello World Dyaksa",
		},
		{
			name:     "HelloWorld(Tifani)",
			request:  "Tifani",
			expected: "Hello World Tifani",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sayHello := HelloWorld(test.request)
			assert.Equal(t, test.expected, sayHello, "result must be Hello World "+test.request)
		})
	}

}
