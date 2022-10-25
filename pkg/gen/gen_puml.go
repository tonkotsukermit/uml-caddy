package gen

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"compress/flate"
)

const pumlEncode = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"

//GetPlantUMLPng retrieves the png from a known plant uml generator url plantuml.com/plantuml/png/
func GetPlantUMLPng(puml string, url string) ([]byte , error) {

	//compress
	compressed, err := Deflate([]byte(puml))
	if err != nil {
		return nil, err
	}

	//encode and request
 	res, err := http.Get(url + string(Encode(compressed)))
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	return body, err
}


// Deflate will run the Deflate compression algorithm on the input.
func Deflate(input []byte) (_ []byte, err error) {
	var b bytes.Buffer

	zw, err := flate.NewWriter(&b, flate.BestCompression)
	if err != nil {
		err = fmt.Errorf("couldn't create a new flate writer: %w", err)
		return
	}

	if _, err = io.Copy(zw, bytes.NewReader(input)); err != nil {
		err = fmt.Errorf("couldn't copy input into writer: %w", err)
		return
	}

	if err = zw.Close(); err != nil {
		err = fmt.Errorf("couldn't close writer: %w", err)
		return
	}

	return b.Bytes(), nil
}

// Encode will encode the input in a similar way as base64.
func Encode(input []byte) []byte {
	inputLength := len(input)
	adjustedInputLength := inputLength + 3 - inputLength%3 // nolint: gomnd

	adjustedInput := make([]byte, adjustedInputLength)
	copy(adjustedInput, input)

	bufferLength := adjustedInputLength * (4 / 3) // nolint: gomnd
	buffer := bytes.NewBuffer(make([]byte, 0, bufferLength))

	for i := 0; i < inputLength; i += 3 {
		b1, b2, b3 := adjustedInput[i], adjustedInput[i+1], adjustedInput[i+2]

		b4 := b3 & 0x3f                    // nolint: gomnd
		b3 = ((b2 & 0xf) << 2) | (b3 >> 6) // nolint: gomnd
		b2 = ((b1 & 0x3) << 4) | (b2 >> 4) // nolint: gomnd
		b1 >>= 2

		for _, b := range []byte{b1, b2, b3, b4} {
			buffer.WriteByte(pumlEncode[b])
		}
	}

	return buffer.Bytes()
}