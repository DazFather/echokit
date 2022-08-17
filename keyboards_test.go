package echokit

import (
	"fmt"

	"github.com/NicoNex/echotron/v3"
)

func ExampleInlineKeyboard() {
	var markup echotron.InlineKeyboardMarkup = InlineKeyboard([][]InlineButton{
		// First row
		{
			{Text: "Click me", CallbackData: "/command"},
			{Text: "Golang", URL: "https://go.dev"},
		},
		// Second row
		{{Text: "I'm on the 2nd row", CallbackData: "/lol"}},
	})

	fmt.Println(markup.InlineKeyboard[0][0].Text) // Output: Click me
}

func ExampleKeyboard() {
	var (
		// After using it keyboard will be removed
		disposable bool = true
		// A text that will appear as placeholder on the message input box
		inputPlaceholder string = "Use the keyboard down below:"
		buttons                 = [][]KeyButton{
			{{Text: "Hello World!"}, {Text: "Send your location", RequestLocation: true}},
			{{Text: "Send me your contact infos", RequestContact: true}},
		}
	)

	var markup echotron.ReplyKeyboardMarkup = Keyboard(disposable, inputPlaceholder, buttons)
	fmt.Println(markup.Keyboard[0][0].Text)
	// These will be set to: true
	fmt.Println(markup.ResizeKeyboard, markup.Selective)
	// These will be set to the firsts 2 arguments (disposable and inputPlaceholder)
	fmt.Println(markup.OneTimeKeyboard, markup.InputFieldPlaceholder)
	// Output: Hello World!
	// true true
	// true Use the keyboard down below:
}

func ExampleKeyboardRemover() {
	var (
		globally   bool                         = true
		kbdRemover echotron.ReplyKeyboardRemove = KeyboardRemover(globally)
	)

	fmt.Println(kbdRemover.RemoveKeyboard) // Output: true
	fmt.Println(kbdRemover.Selective)      // false
}

func ExampleArrange() {
	var buttons [][]KeyButton = Arrange(2, // Arrange following buttons in 2 columns
		KeyButton{Text: "Hello World!"},
		KeyButton{Text: "Send your location", RequestLocation: true},
		KeyButton{Text: "Send me your contact infos", RequestContact: true},
	)

	// First row
	fmt.Println(buttons[0][0].Text)
	fmt.Println(buttons[0][1].Text)

	// Second row
	fmt.Println(buttons[1][0].Text)

	// Output: Hello World!
	// Send your location
	// Send me your contact infos
}

func ExampleGenInlineKeyboard() {
	// Arrange following buttons in 2 columns and generate an InlineKeyboardMarkup
	var markup echotron.InlineKeyboardMarkup = GenInlineKeyboard(2,
		InlineButton{Text: "Click me", CallbackData: "/command"},
		InlineButton{Text: "Golang", URL: "https://go.dev"},
		InlineButton{Text: "I'm on the 2nd row", CallbackData: "/lol"},
	)

	// First row
	fmt.Println(markup.InlineKeyboard[0][0].Text)
	fmt.Println(markup.InlineKeyboard[0][1].Text)

	// Second row
	fmt.Println(markup.InlineKeyboard[1][0].Text)

	// Output: Click me
	// Golang
	// I'm on the 2nd row
}

func ExampleInlineCaller() {
	var button InlineButton = InlineCaller("Click me", "/command", "first", "second")

	fmt.Println(button.Text)         // Output: Click me
	fmt.Println(button.CallbackData) // /command first second
}

func ExampleWrap() {
	var (
		strings   []string         = Wrap("Hello")
		buttonRow []InlineButton   = Wrap(InlineButton{Text: "Hello!"})
		buttons   [][]InlineButton = Wrap(buttonRow)
	)

	fmt.Println(strings)   // Output: [Hello]
	fmt.Println(buttonRow) // [{Hello!}]
	fmt.Println(buttons)   // [[{Hello!}]]
}
