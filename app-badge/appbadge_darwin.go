package appbadge

type AppBadgeDarwin struct{}

func (*AppBadgeDarwin) SetBadge(value int) error {
	return ErrUnsupportedPlatform
}

func (*AppBadgeDarwin) ClearBadge() error {
	return ErrUnsupportedPlatform
}

func init() {
	api = &AppBadgeDarwin{}
}
