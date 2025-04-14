package processor

import (
	"github.com/nycu-ucr/ausf/internal/sbi/consumer"
	"github.com/nycu-ucr/ausf/pkg/app"
)

type ProcessorAusf interface {
	app.App

	Consumer() *consumer.Consumer
}

type Processor struct {
	ProcessorAusf
}

func NewProcessor(ausf ProcessorAusf) (*Processor, error) {
	p := &Processor{
		ProcessorAusf: ausf,
	}
	return p, nil
}
