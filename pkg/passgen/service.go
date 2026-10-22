package passwords

import (
	"bytes"
)

func NewGenerator(alphabet string, maxLength int) <-chan string {
	out := make(chan string)
	if len(alphabet) == 0 || maxLength <= 0 {
		close(out)
		return out
	}

	go func() {
		defer close(out)

		var buffer bytes.Buffer
		buffer.Grow(maxLength)

		for length := 1; ; length++ {
			generate(&buffer, alphabet, length, out)

			if length == maxLength {
				return
			}
		}
	}()

	return out
}
func generate(buffer *bytes.Buffer, alphabet string, remaining int, out chan<- string) {
	if remaining == 0 {
		out <- buffer.String()
		return
	}
	size := buffer.Len()

	for i := 0; i < len(alphabet); i++ {
		buffer.WriteByte(alphabet[i])
		generate(buffer, alphabet, remaining-1, out)
		buffer.Truncate(size)
	}
}

func NewGeneratorFull(maxLength int) <-chan string {
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890!\"#$%&'()*+,-./:;<=>?@[\\]^_{|}~`"
	return NewGenerator(alphabet, maxLength)
}
