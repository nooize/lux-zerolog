package lux_zerolog

import (
	"fmt"
	"github.com/nooize/lux"
	"github.com/rs/zerolog"
)

type zerologTarget struct {
	logger zerolog.Logger
}

func (t zerologTarget) Handle(event lux.Event) error {
	var zeroEvent *zerolog.Event
	switch event.Level() {
	case lux.Nop:
		return nil
	case lux.Trace:
		zeroEvent = t.logger.Trace()
	case lux.Debug:
		zeroEvent = t.logger.Debug()
	case lux.Info:
		zeroEvent = t.logger.Info()
	case lux.Warning:
		zeroEvent = t.logger.Warn()
	case lux.Error:
		zeroEvent = t.logger.Error()
	case lux.Fatal:
		zeroEvent = t.logger.Fatal()
	default:
		return fmt.Errorf("unknown log level: %d", event.Level())
	}
	event.Tags().ForEach(func(key string, val interface{}) {
		zeroEvent = zeroEvent.Interface(key, val)
	})
	zeroEvent = zeroEvent.Time(zerolog.TimestampFieldName, event.Time())
	zeroEvent.Msg(event.Message())
	return nil
}
