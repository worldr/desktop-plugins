package appbadge

type AppBadgeWindows struct{}

func (*AppBadgeWindows) SetBadge(value int32) error {
	return ErrNotImplemented
}

func (*AppBadgeWindows) ClearBadge() error {
	return ErrNotImplemented
}

func init() {
	Api = &AppBadgeWindows{}
}
