package uirecover

import (
	"github.com/ying32/govcl/vcl"
	"strings"
	"time"
)

func UseConfigOrDefault(mconfig *config.Config) *config.Config {
	if mconfig == nil {
		return config.Default()
	}
	return mconfig
}

func TEdit(config *config.Config, ui *vcl.TEdit, key string, defVal ...string) {
	cfg := UseConfigOrDefault(config)
	defer func() {
		cfg.Set(key, ui.Text())
	}()
	v := cfg.String(key, defVal...)
	ui.SetText(v)
	ui.SetOnChange(func(sender vcl.IObject) {
		cfg.Set(key, ui.Text())
	})
}

func TLabeledEdit(config *config.Config, ui *vcl.TLabeledEdit, key string, defVal ...string) {
	cfg := UseConfigOrDefault(config)
	defer func() {
		cfg.Set(key, ui.Text())
	}()
	v := cfg.String(key, defVal...)
	ui.SetText(v)
	ui.SetOnChange(func(sender vcl.IObject) {
		cfg.Set(key, ui.Text())
	})
}
func TComboBox(config *config.Config, ui *vcl.TComboBox, key string, defVal ...any) {
	cfg := UseConfigOrDefault(config)
	defer func() {
		cfg.Set(key, ui.Text())
	}()
	v := cfg.String(key)
	if v != "" {
		ui.SetText(v)
	} else {
		if len(defVal) > 0 {
			switch t := defVal[0].(type) {
			case string:
				ui.SetText(t)
			case int:
				for i := 0; i < int(ui.Items().Count()); i++ {
					if i == t {
						ui.SetText(ui.Items().Strings(int32(i)))
						break
					}
				}
			}
		} else {
			for i := 0; i < int(ui.Items().Count()); i++ {
				if i == 0 {
					ui.SetText(ui.Items().Strings(int32(i)))
					break
				}
			}
		}
	}
	ui.SetOnChange(func(sender vcl.IObject) {
		cfg.Set(key, ui.Text())
	})
}
func TDateTimePicker(config *config.Config, ui *vcl.TDateTimePicker, key string, defVal ...time.Time) {
	cfg := UseConfigOrDefault(config)
	defer func() {
		cfg.Set(key, ui.DateTime().Format("2006-01-02 15:04:05"))
	}()
	vs := cfg.String(key)
	if vs != "" {
		if strings.Contains(vs, " ") {
			t, err := time.Parse("2006-01-02 15:04:05", vs)
			if err == nil {
				ui.SetDateTime(t)
			}
		} else {
			t, err := time.Parse("2006-01-02", vs)
			if err == nil {
				ui.SetDateTime(t)
			}
		}
	} else {
		if len(defVal) > 0 {
			ui.SetDateTime(defVal[0])
		}
	}
	ui.SetOnChange(func(sender vcl.IObject) {
		cfg.Set(key, ui.DateTime().Format("2006-01-02 15:04:05"))
	})
}
