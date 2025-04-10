package core

import "time"

func UseInterval(ctx *Context, interval time.Duration, callback func()) {
	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			callback()
			ctx.MarkDirty()
		}
	}()
}
