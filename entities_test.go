package echokit

import (
	"fmt"

	"github.com/NicoNex/echotron/v3"
)

func ExampleExtractEntities() {
	var (
		filterBig EntityFilter = func(entity echotron.MessageEntity) bool {
			return entity.Length > 6
		}

		message = echotron.Message{
			Text: "Hi @DazFather! Did you ever hear about the #DigitalResistance ?",
			Entities: []*echotron.MessageEntity{
				&echotron.MessageEntity{
					Offset: 3,
					Length: 10,
					Type:   echotron.MentionEntity,
				},
				&echotron.MessageEntity{
					Offset: 43,
					Length: 18,
					Type:   echotron.HashtagEntity,
				},
			},
		}
	)

	fmt.Println(ExtractEntities(message, filterBig)) // Output: [@DazFather #DigitalResistance]
}

func ExampleExtractEntitiesOfType() {
	var message = echotron.Message{
		Text: "Hi @DazFather! Did you ever hear about the #DigitalResistance ?",
		Entities: []*echotron.MessageEntity{
			&echotron.MessageEntity{
				Offset: 3,
				Length: 10,
				Type:   echotron.MentionEntity,
			},
			&echotron.MessageEntity{
				Offset: 43,
				Length: 18,
				Type:   echotron.HashtagEntity,
			},
		},
	}

	tags := ExtractEntitiesOfType(message, echotron.HashtagEntity)
	extracted := ExtractEntitiesOfType(message, echotron.MentionEntity, echotron.HashtagEntity)

	fmt.Println(tags)         // Output: [#DigitalResistance]
	fmt.Println(extracted[0]) // [@DazFather #DigitalResistance]
}

func ExampleExtractEntity() {
	var entity = echotron.MessageEntity{
		Offset: 3,
		Length: 10,
		Type:   echotron.MentionEntity,
	}

	fmt.Println(ExtractEntity("Hi @DazFather!", entity)) // Output: @DazFather
}

func ExampleStirngToUft16() {
	var encoded []uint16 = StirngToUft16("Hello, 世界")
	fmt.Println(encoded) // Output: [72 101 108 108 111 44 32 19990 30028]
}

func ExampleUft16ToStirng() {
	var encoded = []uint16{72, 101, 108, 108, 111, 44, 32, 19990, 30028}
	fmt.Println(Uft16ToStirng(encoded)) // Output: Hello, 世界
}

func ExampleFilterEntityByType() {
	var (
		filter EntityFilter = FilterEntityByType(echotron.HashtagEntity, echotron.CashtagEntity)

		hashtag = echotron.MessageEntity{Type: echotron.HashtagEntity}
		cashtag = echotron.MessageEntity{Type: echotron.CashtagEntity}
		link    = echotron.MessageEntity{Type: echotron.UrlEntity}
	)

	fmt.Println(filter(hashtag)) // Output: true
	fmt.Println(filter(cashtag)) // true
	fmt.Println(filter(link))    // false
}
