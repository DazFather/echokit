package echokit

import (
	"unicode/utf16"

	"github.com/NicoNex/echotron/v3"
)

// EntityFilter is a function checks an entity, returning true if pass false otherwise
type EntityFilter func(entity echotron.MessageEntity) bool

// ExtractEntities extracts the content of the Entities contained inside given
// message and filtered by given filter function
func ExtractEntities(message echotron.Message, filter EntityFilter) (extracted []string) {
	var text = StirngToUft16(message.Text)

	for _, entity := range message.Entities {
		if entity != nil && filter(*entity) {
			extracted = append(extracted, grabEntityContent(text, *entity))
		}
	}
	return
}

// ExtractEntitiesOfType extracts the content of the Entities contained filtered by their type
func ExtractEntitiesOfType(message echotron.Message, allowedTypes ...echotron.MessageEntityType) []string {
	return ExtractEntities(message, FilterEntityByType(allowedTypes...))
}

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

// FilterEntityByType generates an EntityFilter that will return true only if
// entity's type match at least one of the given ones
func FilterEntityByType(allowedTypes ...echotron.MessageEntityType) EntityFilter {
	return func(entity echotron.MessageEntity) bool {
		return contains(allowedTypes, entity.Type)
	}
}

// contains checks if the value of given elem is present on the given list
func contains[T comparable](list []T, elem T) bool {
	for _, value := range list {
		if value == elem {
			return true
		}
	}
	return false
}
