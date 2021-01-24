package main

import (
	"strings"

	"github.com/cosmos72/gomacro/fast"
	bot "github.com/kf5grd/keybasebot"
	"samhofi.us/x/keybase/v2/types/chat1"
)

func cmdEval(m chat1.MsgSummary, b *bot.Bot) (bool, error) {
	body := m.Content.Text.Body
	cmd := strings.Replace(body, "!eval", "", 1)
	cmd = strings.TrimSpace(cmd)

	defer func() {
		if r := recover(); r != nil {
			b.KB.ReactByConvID(m.ConvID, m.Id, ":no_entry_sign:")
		}
	}()

	b.Logger.Debug("Evaluating: %s", cmd)
	interp := b.Meta["interp"].(*fast.Interp)
	vals, _ := interp.Eval(cmd)
	b.KB.ReactByConvID(m.ConvID, m.Id, ":white_check_mark:")

	if len(vals) == 0 {
		return true, nil
	}

	for _, val := range vals {
		b.KB.SendMessageByConvID(m.ConvID, "%v", val)
	}

	return true, nil
}
