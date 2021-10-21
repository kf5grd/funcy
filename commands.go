package main

import (
	"strings"

	bot "github.com/kf5grd/keybasebot"
	"github.com/traefik/yaegi/interp"
	"samhofi.us/x/keybase/v2/types/chat1"
)

func cmdEval(m chat1.MsgSummary, b *bot.Bot) (bool, error) {
	body := m.Content.Text.Body
	cmd := strings.Replace(body, "!eval", "", 1)
	cmd = strings.TrimSpace(cmd)

	defer func() {
		if r := recover(); r != nil {
			b.Logger.Error("Recovered from panic: %s", r)
			b.KB.ReactByConvID(m.ConvID, m.Id, ":no_entry_sign:")
		}
	}()

	b.Logger.Debug("Evaluating: %s", cmd)
	i := b.Meta["interp"].(*interp.Interpreter)
	val, err := i.Eval(cmd)
	if err != nil {
		b.KB.ReactByConvID(m.ConvID, m.Id, ":no_entry_sign:")
		b.KB.ReplyByConvID(m.ConvID, m.Id, "Error: %v", err)
		return true, nil
	}
	b.KB.ReactByConvID(m.ConvID, m.Id, ":white_check_mark:")
	if val.Interface() != nil {
		b.KB.SendMessageByConvID(m.ConvID, "%v", val)
	}
	return true, nil
}
