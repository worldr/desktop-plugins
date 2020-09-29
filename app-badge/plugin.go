package appbadge

import (
	"errors"
	"log"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

const channelName = "com.worldr.desktop.plugins.app_badge"
const (
	METHOD_SET = "setBadge"
)

type AppBadgePlugin struct{}

var _ flutter.Plugin = &AppBadgePlugin{}
var errorFormat string = "[badge] %v"

func (p *AppBadgePlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc(METHOD_SET, p.setBadge)
	return nil
}

func (p *AppBadgePlugin) setBadge(args interface{}) (reply interface{}, err error) {
	counter, ok := args.(int)
	if !ok {
		return nil, errors.New("invalid args")
	}
	if counter <= 0 {
		return nil, nil
	}
	return nil, nil
}

func newError(message string) error {
	log.Printf(errorFormat, message)
	return errors.New(message)
}
