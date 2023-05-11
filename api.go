package zerolog

import (
	"github.com/nooize/lwr"
	"github.com/rs/zerolog"
)

func NewZerologTarget(logger zerolog.Logger) lwr.Target {
	t := &zerologTarget{
		logger: logger,
	}
	return t
}
