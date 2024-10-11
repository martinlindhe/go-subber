package subtitles

import (
	"fmt"
	"log"
	"os"
)

// Parse tries to parse a subtitle
func Parse(b []byte) (Subtitle, error) {
	s := ConvertToUTF8(b)
	if looksLikeCCDBCapture(s) {
		return NewFromCCDBCapture(s)
	} else if looksLikeSSA(s) {
		return NewFromSSA(s)
	} else if looksLikeDCSub(s) {
		return NewFromDCSub(s)
	} else if looksLikeSRT(s) {
		return NewFromSRT(s)
	} else if looksLikeVTT(s) {
		return NewFromVTT(s)
	}
	return Subtitle{}, fmt.Errorf("parse: unrecognized subtitle type")
}

// LooksLikeTextSubtitle returns true i byte stream seems to be of a recognized format
func LooksLikeTextSubtitle(filename string) bool {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	s := ConvertToUTF8(data)
	return looksLikeCCDBCapture(s) || looksLikeSSA(s) || looksLikeDCSub(s) || looksLikeSRT(s) || looksLikeVTT(s)
}
