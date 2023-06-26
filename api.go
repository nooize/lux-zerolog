package lux_zerolog

import (
	"github.com/nooize/lux"
	"github.com/rs/zerolog"
)

func NewZerologTarget(logger zerolog.Logger) lux.Target {
	t := &zerologTarget{
		logger: logger,
	}
	return t
}
