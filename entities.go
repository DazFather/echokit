package echokit

import (
	"unicode/utf16"

	"github.com/NicoNex/echotron/v3"
)

// ExtractEntity extract the entity content from a given text
func ExtractEntity(fromText string, entity echotron.MessageEntity) string {
	return grabEntityContent(StirngToUft16(fromText), entity)
}

// StirngToUft16 encode a given string into Utf16 format
func StirngToUft16(toEncode string) []uint16 {
	return utf16.Encode([]rune(toEncode))
}

// Uft16ToStirng decode a given encoded text in Utf16 format into a string
func Uft16ToStirng(encodedText []uint16) string {
	return string(utf16.Decode(encodedText))
}

// grabEntityContent extract the entity content from a given utf16-encoded text
func grabEntityContent(uft16Text []uint16, entity echotron.MessageEntity) string {
	return Uft16ToStirng(uft16Text[entity.Offset : entity.Offset+entity.Length])
}

