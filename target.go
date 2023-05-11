package zerolog

import (
	"fmt"
	"github.com/nooize/lwr"
	"github.com/rs/zerolog"
)

type zerologTarget struct {
	logger zerolog.Logger
}

func (t zerologTarget) Handle(event lwr.Event) error {
	var zeroEvent *zerolog.Event
	switch event.Level() {
	case lwr.Nop, lwr.Disabled:
		return nil
	case lwr.Trace:
		zeroEvent = t.logger.Trace()
	case lwr.Debug:
		zeroEvent = t.logger.Debug()
	case lwr.Info:
		zeroEvent = t.logger.Info()
	case lwr.Warning:
		zeroEvent = t.logger.Warn()
	case lwr.Error:
		zeroEvent = t.logger.Error()
	case lwr.Fatal:
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
