// Code generated by 'yaegi extract samhofi.us/x/keybase/v2'. DO NOT EDIT.

package main

import (
	"reflect"
	"samhofi.us/x/keybase/v2"
)

func init() {
	Symbols["samhofi.us/x/keybase/v2/keybase"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"CHAT":        reflect.ValueOf(keybase.CHAT),
		"DEV":         reflect.ValueOf(keybase.DEV),
		"New":         reflect.ValueOf(keybase.New),
		"NewKeybase":  reflect.ValueOf(keybase.NewKeybase),
		"SetExePath":  reflect.ValueOf(keybase.SetExePath),
		"SetHomePath": reflect.ValueOf(keybase.SetHomePath),
		"TEAM":        reflect.ValueOf(keybase.TEAM),
		"USER":        reflect.ValueOf(keybase.USER),

		// type definitions
		"AdvertiseCommandsOptions": reflect.ValueOf((*keybase.AdvertiseCommandsOptions)(nil)),
		"Chat":                     reflect.ValueOf((*keybase.Chat)(nil)),
		"ChatAPI":                  reflect.ValueOf((*keybase.ChatAPI)(nil)),
		"DownloadOptions":          reflect.ValueOf((*keybase.DownloadOptions)(nil)),
		"Error":                    reflect.ValueOf((*keybase.Error)(nil)),
		"ExplodingLifetime":        reflect.ValueOf((*keybase.ExplodingLifetime)(nil)),
		"Handlers":                 reflect.ValueOf((*keybase.Handlers)(nil)),
		"KVOptions":                reflect.ValueOf((*keybase.KVOptions)(nil)),
		"Keybase":                  reflect.ValueOf((*keybase.Keybase)(nil)),
		"KeybaseOpt":               reflect.ValueOf((*keybase.KeybaseOpt)(nil)),
		"ListMembersOptions":       reflect.ValueOf((*keybase.ListMembersOptions)(nil)),
		"ReadMessageOptions":       reflect.ValueOf((*keybase.ReadMessageOptions)(nil)),
		"RequestPayment":           reflect.ValueOf((*keybase.RequestPayment)(nil)),
		"RunOptions":               reflect.ValueOf((*keybase.RunOptions)(nil)),
		"SendMessageBody":          reflect.ValueOf((*keybase.SendMessageBody)(nil)),
		"SendMessageOptions":       reflect.ValueOf((*keybase.SendMessageOptions)(nil)),
		"SendPayment":              reflect.ValueOf((*keybase.SendPayment)(nil)),
		"Team":                     reflect.ValueOf((*keybase.Team)(nil)),
		"TeamAPI":                  reflect.ValueOf((*keybase.TeamAPI)(nil)),
		"UserAPI":                  reflect.ValueOf((*keybase.UserAPI)(nil)),
		"UserCardAPI":              reflect.ValueOf((*keybase.UserCardAPI)(nil)),
		"Wallet":                   reflect.ValueOf((*keybase.Wallet)(nil)),
		"WalletAPI":                reflect.ValueOf((*keybase.WalletAPI)(nil)),

		// interface wrapper definitions
		"_KeybaseOpt": reflect.ValueOf((*_samhofi_us_x_keybase_v2_KeybaseOpt)(nil)),
	}
}

// _samhofi_us_x_keybase_v2_KeybaseOpt is an interface wrapper for KeybaseOpt type
type _samhofi_us_x_keybase_v2_KeybaseOpt struct {
	IValue interface{}
}