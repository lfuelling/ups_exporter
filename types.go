package main

type Status struct {
	OperatingMode  string
	UPSTemperature float64
	InputVoltage float64
	InputFrequency float64
	InputPhase int
	OutputPhase int
	OutputVoltage float64
	OutputFrequency float64
	OutputCurrent float64
	OutputLoadPercent int
	BatteryVoltage float64
	BatteryCapacityPercent int
	BatteryRemainingBackupTime int
	BatteryGroupNumber int
	RatedVA float64
	RatedOutputFrequency float64
	RatedOutputVoltage float64
	RatedOutputCurrent float64
	RatedBatteryVoltage float64
	AutoReboot bool
	EcoMode bool
	BypassDisabled bool
	ConverterMode bool
	BypassWhenOff bool
	UPSType string
	UPSSerial string
	UPSFirmware string
	SNMPFirmware string
	EquipmentAttached string
}
