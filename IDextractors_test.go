package echokit_test

import (
	"fmt"

	"github.com/DazFather/echokit"

	"github.com/NicoNex/echotron/v3"
)

func ExampleExtractIDOpt() {
	var (
		message = echotron.Message{
			ID:   1234,
			Chat: echotron.Chat{ID: 5678},
		}

		callback = echotron.CallbackQuery{InlineMessageID: "Hello"}

		update = echotron.Update{ChannelPost: &message}
	)

	fmt.Println(*echokit.ExtractIDOpt(message) == echotron.NewMessageID(message.Chat.ID, message.ID))                      // Output: true
	fmt.Println(*echokit.ExtractIDOpt(callback) == echotron.NewInlineMessageID(callback.InlineMessageID))                  // true
	fmt.Println(*echokit.ExtractIDOpt(update) == echotron.NewMessageID(update.ChannelPost.Chat.ID, update.ChannelPost.ID)) // true

	var invalid *echotron.MessageIDOptions = echokit.ExtractIDOpt(echotron.Message{})
	fmt.Println(invalid) // nil
}
