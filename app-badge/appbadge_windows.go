package appbadge

type AppBadgeWindows struct{}

func (*AppBadgeWindows) SetBadge(value int) error {
	return ErrNotImplemented
}

func (*AppBadgeWindows) ClearBadge() error {
	return ErrNotImplemented
}

func init() {
	api = &AppBadgeWindows{}
}
