package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateSalia(base *Salia, meter func() (float64, error), meterEnergy func() (float64, error), phaseCurrents func() (float64, float64, float64, error)) api.Charger {
	switch {
	case meter == nil && meterEnergy == nil && phaseCurrents == nil:
		return base

	case meter != nil && meterEnergy == nil && phaseCurrents == nil:
		return &struct {
			*Salia
			api.Meter
		}{
			Salia: base,
			Meter: &decorateSaliaMeterImpl{
				meter: meter,
			},
		}

	case meter != nil && meterEnergy != nil && phaseCurrents == nil:
		return &struct {
			*Salia
			api.Meter
			api.MeterEnergy
		}{
			Salia: base,
			Meter: &decorateSaliaMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateSaliaMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter != nil && meterEnergy == nil && phaseCurrents != nil:
		return &struct {
			*Salia
			api.Meter
			api.PhaseCurrents
		}{
			Salia: base,
			Meter: &decorateSaliaMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateSaliaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case meter != nil && meterEnergy != nil && phaseCurrents != nil:
		return &struct {
			*Salia
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
		}{
			Salia: base,
			Meter: &decorateSaliaMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateSaliaMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateSaliaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}
	}

	return nil
}

type decorateSaliaMeterImpl struct {
	meter func() (float64, error)
}

func (impl *decorateSaliaMeterImpl) CurrentPower() (float64, error) {
	return impl.meter()
}

type decorateSaliaMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decorateSaliaMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}

type decorateSaliaPhaseCurrentsImpl struct {
	phaseCurrents func() (float64, float64, float64, error)
}

func (impl *decorateSaliaPhaseCurrentsImpl) Currents() (float64, float64, float64, error) {
	return impl.phaseCurrents()
}
