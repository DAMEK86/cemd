package ucmgcp

import "github.com/enbility/cemd/api"

const (
	// Grid maximum allowed feed-in power as percentage value of the cumulated
	// nominal peak power of all electricity producting PV systems was updated
	//
	// Use `PowerLimitationFactor` to get the current data
	//
	// Use Case MGCP, Scenario 2
	DataUpdatePowerLimitationFactor api.EventType = "DataUpdatePowerLimitationFactor"

	// Grid momentary power consumption/production data updated
	//
	// Use `Power` to get the current data
	//
	// Use Case MGCP, Scenario 2
	DataUpdatePower api.EventType = "DataUpdatePower"

	// Total grid feed in energy data updated
	//
	// Use `EnergyFeedIn` to get the current data
	//
	// Use Case MGCP, Scenario 3
	DataUpdateEnergyFeedIn api.EventType = "DataUpdateEnergyFeedIn"

	// Total grid consumed energy data updated
	//
	// Use `EnergyConsumed` to get the current data
	//
	// Use Case MGCP, Scenario 4
	DataUpdateEnergyConsumed api.EventType = "DataUpdateEnergyConsumed"

	// Phase specific momentary current consumption/production phase detail data updated
	//
	// Use `CurrentPerPhase` to get the current data
	//
	// Use Case MGCP, Scenario 5
	DataUpdateCurrentPerPhase api.EventType = "DataUpdateCurrentPerPhase"

	// Phase specific voltage at the grid connection point
	//
	// Use `VoltagePerPhase` to get the current data
	//
	// Use Case MGCP, Scenario 6
	DataUpdateVoltagePerPhase api.EventType = "DataUpdateVoltagePerPhase"

	// Grid frequency data updated
	//
	// Use `Frequency` to get the current data
	//
	// Use Case MGCP, Scenario 7
	DataUpdateFrequency api.EventType = "DataUpdateFrequency"
)
