package appbadge

import (
	"errors"
	"fmt"
	"log"
	"reflect"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

const channelName = "com.worldr.desktop.plugins.appbadge"
const (
	METHOD_SET   = "setBadge"
	METHOD_CLEAR = "clearBadge"
)

type AppBadgePlugin struct{}

var _ flutter.Plugin = &AppBadgePlugin{}
var errorFormat string = "[badge] %v"

func (p *AppBadgePlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc(METHOD_SET, p.setBadge)
	channel.HandleFunc(METHOD_CLEAR, p.clearBadge)
	return nil
}

func (p *AppBadgePlugin) setBadge(args interface{}) (reply interface{}, err error) {
	fmt.Println("ARGS: %T", args)
	fmt.Println("ARGS: %v", reflect.TypeOf(args))
	counter, ok := args.(int)
	if !ok {
		return nil, errors.New("invalid args")
	}
	if counter <= 0 {
		return nil, Api.ClearBadge()
	}
	return nil, Api.SetBadge(counter)
}

func (p *AppBadgePlugin) clearBadge(args interface{}) (reply interface{}, err error) {
	return nil, Api.ClearBadge()
}

func newError(message string) error {
	log.Printf(errorFormat, message)
	return errors.New(message)
}
