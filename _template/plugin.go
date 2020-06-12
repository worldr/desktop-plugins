package sqflite_sqlcipher

import (
	//"github.com/pkg/errors"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

const channelName = "%%ChannelName%%"

type %%PluginTypeName%% struct{}

var _ flutter.Plugin = &%%PluginTypeName%%{}

// plugin init
func (p *SqlCipherPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})

	// method handlers list
	channel.HandleFunc("test", p.test)


	return nil
}

// plugin methods
func (p *%%PluginTypeName%%) test(arguments interface{}) (reply interface{}, err error) {
	return map[interface{}]interface{}{
		"test": flutter.ProjectName,
	}, nil
}
