package appbadge

type AppBadgeLinux struct{}

func (*AppBadgeLinux) SetBadge(value int) error {
	return ErrNotImplemented
}

func (*AppBadgeLinux) ClearBadge() error {
	return ErrNotImplemented
}

func init() {
	api = &AppBadgeLinux{}
}