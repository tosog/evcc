package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateKSE(base *KSE, phaseSwitcher func(int) error, phaseGetter func() (int, error), identifier func() (string, error)) api.Charger {
	switch {
	case identifier == nil && phaseGetter == nil && phaseSwitcher == nil:
		return base

	case identifier == nil && phaseGetter == nil && phaseSwitcher != nil:
		return &struct {
			*KSE
			api.PhaseSwitcher
		}{
			KSE: base,
			PhaseSwitcher: &decorateKSEPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case identifier == nil && phaseGetter != nil && phaseSwitcher != nil:
		return &struct {
			*KSE
			api.PhaseGetter
			api.PhaseSwitcher
		}{
			KSE: base,
			PhaseGetter: &decorateKSEPhaseGetterImpl{
				phaseGetter: phaseGetter,
			},
			PhaseSwitcher: &decorateKSEPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case identifier != nil && phaseGetter == nil && phaseSwitcher == nil:
		return &struct {
			*KSE
			api.Identifier
		}{
			KSE: base,
			Identifier: &decorateKSEIdentifierImpl{
				identifier: identifier,
			},
		}

	case identifier != nil && phaseGetter == nil && phaseSwitcher != nil:
		return &struct {
			*KSE
			api.Identifier
			api.PhaseSwitcher
		}{
			KSE: base,
			Identifier: &decorateKSEIdentifierImpl{
				identifier: identifier,
			},
			PhaseSwitcher: &decorateKSEPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case identifier != nil && phaseGetter != nil && phaseSwitcher != nil:
		return &struct {
			*KSE
			api.Identifier
			api.PhaseGetter
			api.PhaseSwitcher
		}{
			KSE: base,
			Identifier: &decorateKSEIdentifierImpl{
				identifier: identifier,
			},
			PhaseGetter: &decorateKSEPhaseGetterImpl{
				phaseGetter: phaseGetter,
			},
			PhaseSwitcher: &decorateKSEPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}
	}

	return nil
}

type decorateKSEIdentifierImpl struct {
	identifier func() (string, error)
}

func (impl *decorateKSEIdentifierImpl) Identify() (string, error) {
	return impl.identifier()
}

type decorateKSEPhaseGetterImpl struct {
	phaseGetter func() (int, error)
}

func (impl *decorateKSEPhaseGetterImpl) GetPhases() (int, error) {
	return impl.phaseGetter()
}

type decorateKSEPhaseSwitcherImpl struct {
	phaseSwitcher func(int) error
}

func (impl *decorateKSEPhaseSwitcherImpl) Phases1p3p(p0 int) error {
	return impl.phaseSwitcher(p0)
}
