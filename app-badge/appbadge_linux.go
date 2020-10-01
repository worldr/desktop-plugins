package appbadge

type AppBadgeLinux struct{}

func (*AppBadgeLinux) SetBadge(value int32) error {
	return ErrNotImplemented
}

func (*AppBadgeLinux) ClearBadge() error {
	return ErrNotImplemented
}

func init() {
	Api = &AppBadgeLinux{}
}
