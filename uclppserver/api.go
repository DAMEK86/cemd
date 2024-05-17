package uclppserver

import (
	"time"

	"github.com/enbility/cemd/api"
)

//go:generate mockery

// interface for the Limitation of Power Production UseCase
type UCLPPServerInterface interface {
	api.UseCaseInterface

	// Scenario 1

	// return the current loadcontrol limit data
	//
	// return values:
	//   - limit: load limit data
	//
	// possible errors:
	//   - ErrDataNotAvailable if no such limit is (yet) available
	//   - and others
	ProductionLimit() (api.LoadLimit, error)

	// set the current loadcontrol limit data
	SetProductionLimit(limit api.LoadLimit) (resultErr error)

	// Scenario 2

	// return Failsafe limit for the produced active (real) power of the
	// Controllable System. This limit becomes activated in "init" state or "failsafe state".
	//
	// return values:
	//   - value: the power limit in W
	//   - changeable: boolean if the client service can change the limit
	FailsafeProductionActivePowerLimit() (value float64, isChangeable bool, resultErr error)

	// set Failsafe limit for the produced active (real) power of the
	// Controllable System. This limit becomes activated in "init" state or "failsafe state".
	//
	// parameters:
	//   - value: the power limit in W
	//   - changeable: boolean if the client service can change the limit
	SetFailsafeProductionActivePowerLimit(value float64, changeable bool) (resultErr error)

	// return minimum time the Controllable System remains in "failsafe state" unless conditions
	// specified in this Use Case permit leaving the "failsafe state"
	//
	// return values:
	//   - value: the power limit in W
	//   - changeable: boolean if the client service can change the limit
	FailsafeDurationMinimum() (duration time.Duration, isChangeable bool, resultErr error)

	// set minimum time the Controllable System remains in "failsafe state" unless conditions
	// specified in this Use Case permit leaving the "failsafe state"
	//
	// parameters:
	//   - duration: has to be >= 2h and <= 24h
	//   - changeable: boolean if the client service can change this value
	SetFailsafeDurationMinimum(duration time.Duration, changeable bool) (resultErr error)

	// Scenario 3

	// this is automatically covered by the SPINE implementation
	//
	// returns true, if the last heartbeat is within 2 minutes, otherwise false
	IsHeartbeatWithinDuration() bool

	// Scenario 4

	// return nominal maximum active (real) power the Controllable System is
	// allowed to produce due to the customer's contract.
	ContractualProductionNominalMax() (float64, error)

	// set nominal maximum active (real) power the Controllable System is
	// allowed to produce due to the customer's contract.
	//
	// parameters:
	//   - value: contractual nominal max power production in W
	SetContractualProductionNominalMax(value float64) (resultErr error)
}
