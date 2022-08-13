package echokit

import "github.com/NicoNex/echotron/v3"

// ExtractIDOpt extract a MessageIDOptions from different data structures (IDOptionsExtractor), nil if can't
func ExtractIDOpt[IDExtractor IDOptionsExtractor](from IDExtractor) *echotron.MessageIDOptions {
	var extractable interface {
		extractIDOpt() *echotron.MessageIDOptions
	}

	switch data := any(from).(type) {
	case echotron.Message:
		extractable = message(data)
	case echotron.Update:
		extractable = update(data)
	case echotron.CallbackQuery:
		extractable = callback(data)
	default:
		return nil
	}

	return extractable.extractIDOpt()
}

type IDOptionsExtractor interface {
	echotron.Message | echotron.Update | echotron.CallbackQuery | echotron.ChosenInlineResult
}

// Messsage
type message echotron.Message

func (m message) extractIDOpt() *echotron.MessageIDOptions {
	msgID := echotron.NewMessageID(m.Chat.ID, m.ID)
	return &msgID
}

// Update
type update echotron.Update

func (update update) extractMessage() *echotron.Message {
	if update.Message != nil {
		return update.Message
	} else if update.EditedMessage != nil {
		return update.EditedMessage
	} else if update.ChannelPost != nil {
		return update.ChannelPost
	} else if update.EditedChannelPost != nil {
		return update.EditedChannelPost
	} else if update.CallbackQuery != nil {
		return update.CallbackQuery.Message
	}

	return nil
}

func (u update) extractIDOpt() *echotron.MessageIDOptions {
	if c := u.CallbackQuery; c != nil {
		return callback(*c).extractIDOpt()
	}
	if msg := u.extractMessage(); msg != nil {
		return message(*msg).extractIDOpt()
	}
	return nil
}

// CallbackQuery
type callback echotron.CallbackQuery

func (c callback) extractIDOpt() *echotron.MessageIDOptions {
	if c.Message != nil {
		return message(*c.Message).extractIDOpt()
	}
	if c.InlineMessageID != "" {
		msgID := echotron.NewInlineMessageID(c.InlineMessageID)
		return &msgID
	}
	return nil
}

// ChosenInlineResult
type inlineRes echotron.ChosenInlineResult

func (i inlineRes) extractIDOpt() *echotron.MessageIDOptions {
	if i.InlineMessageID != "" {
		msgID := echotron.NewInlineMessageID(i.InlineMessageID)
		return &msgID
	}
	return nil
}
