package core

type ToastOpt interface {
	Apply(*ToastConfig)
}

type ToastConfig struct {
	Duration int // ms
	Style    *Style
}
type toastFunc func(*ToastConfig)

func (f toastFunc) Apply(c *ToastConfig) { f(c) }
func ShowToast(msg string, opts ...ToastOpt) {
	conf := ToastConfig{Duration: 2000}
	for _, opt := range opts {
		opt.Apply(&conf)
	}

	payload := map[string]any{
		"message":  msg,
		"duration": conf.Duration,
	}
	if conf.Style != nil {
		payload["style"] = conf.Style
	}

	SendSystemEvent("toast", payload)
}

func Duration(ms int) ToastOpt {
	return toastFunc(func(c *ToastConfig) {
		c.Duration = ms
	})
}

func UseToastStyle(s Style) ToastOpt {
	return toastFunc(func(c *ToastConfig) {
		c.Style = &s
	})
}
