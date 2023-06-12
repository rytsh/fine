package config

import (
	"context"

	"github.com/rytsh/liz/file"
	"github.com/worldline-go/logz"
	"github.com/worldline-go/struct2"
)

type OverrideHold struct {
	Memory *string
	Value  string
}

func Load(ctx context.Context, visit func()) error {
	var configMap map[string]interface{}
	if err := file.New().Load(File, &configMap); err != nil {
		return err
	}

	decoder := struct2.Decoder{
		TagName:               "cfg",
		WeaklyDashUnderscore:  true,
		WeaklyIgnoreSeperator: true,
		WeaklyTypedInput:      true,
	}

	if err := decoder.Decode(configMap, &App); err != nil {
		return err
	}

	// override used cmd values
	visit()

	// set log again to get changes
	if err := logz.SetLogLevel(App.Log.Level); err != nil {
		return err //nolint:wrapcheck // no need
	}

	return nil
}
