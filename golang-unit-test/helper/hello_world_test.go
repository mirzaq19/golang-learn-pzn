package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Mirzaq",
			request: "Mirzaq",
		},
		{
			name:    "Budi",
			request: "Budi",
		},
		{
			name:    "Joko",
			request: "Joko",
		},
		{
			name:    "M. Auliya Mirzaq Romdloni",
			request: "M. Auliya Mirzaq Romdloni",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Mirzaq", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Mirzaq")
		}
	})
	b.Run("Budi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Budi")
		}
	})
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Mirzaq")
	}
}

func BenchmarkHelloBudi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Budi")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		Name     string
		Request  string
		Expected string
	}{
		{
			Name:     "HelloWorld(Mirzaq)",
			Request:  "Mirzaq",
			Expected: "Hello Mirzaq",
		},
		{
			Name:     "HelloWorld(Budi)",
			Request:  "Budi",
			Expected: "Hello Budi",
		},
		{
			Name:     "HelloWorld(Joko)",
			Request:  "Joko",
			Expected: "Hello Joko",
		},
		{
			Name:     "HelloWorld(Richard)",
			Request:  "Richard",
			Expected: "Hello Richard",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := HelloWorld(test.Request)
			require.Equal(t, test.Expected, result, "Result must be 'Hello"+test.Request+"'")
		})

	}
}

func TestSubTest(t *testing.T) {
	t.Run("Mirzaq", func(t *testing.T) {
		result := HelloWorld("Mirzaq")
		assert.Equal(t, "Hello Mirzaq", result, "Result must be 'Hello Mirzaq'")
	})
	t.Run("Budi", func(t *testing.T) {
		result := HelloWorld("Budi")
		assert.Equal(t, "Hello Budi", result, "Result must be 'Hello Budi'")
	})
}

func TestMain(m *testing.M) {
	// Before
	fmt.Println("Sebelum UNIT TEST")
	m.Run()
	// After
	fmt.Println("Setelah UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on windows")
	}

	result := HelloWorld("Mirzaq")
	assert.Equal(t, "Hello Mirzaq", result, "Result must be 'Hello Mirzaq'")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Mirzaq")
	assert.Equal(t, "Hello Mirzaq", result, "Result must be 'Hello Mirzaq'")

	fmt.Println("TestHelloWorld with assert done!")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Mirzaq")
	require.Equal(t, "Hello Mirzaq", result, "Result must be 'Hello Mirzaq'")

	fmt.Println("TestHelloWorld with require done!")
}

func TestHelloWorldMirzaq(t *testing.T) {
	result := HelloWorld("Mirzaq")
	if result != "Hello Mirzaq" {
		// error
		t.Error("Result must be 'Hello Mirzaq'")
	}

	fmt.Println("TestHelloWorldMirzaq Done!")
}

func TestHelloWorldBudi(t *testing.T) {
	result := HelloWorld("Budi")
	if result != "Hello Budi" {
		// error
		t.Fatal("Result must be 'Hello Budi'")
	}

	fmt.Println("TestHelloWorldBudi Done!")
}
