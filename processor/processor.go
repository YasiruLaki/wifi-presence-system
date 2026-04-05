package processor

import "wifi-presence-system/models"

type Processor struct {
	TrusholdStill float64
	TrusholdMoving float64
}

func NewProcessor(trusholdStill, trusholdMoving float64) *Processor {
	return &Processor{
		TrusholdStill: trusholdStill,
		TrusholdMoving: trusholdMoving,
	}
}

func (p *Processor) Process(data models.SensorData) models.PresenceState {
	var state string

	if data.Variance < p.TrusholdStill {
		state = "still"
	} else if data.Variance < p.TrusholdMoving {
		state = "moving"
	} else {
		state = "unknown"
	}

	return models.PresenceState{
		State: state,
		Variance: data.Variance,
		Timestamp: data.Timestamp,
	}
}