package main

type Status struct {
	OperatingMode  string
	UPSTemperature int
	InputVoltage int
	InputFrequency int
	InputPhase int
	OutputPhase int
	OutputVoltage int
	OutputFrequency int
	OutputCurrent int
	OutputLoadPercent int
	BatteryVoltage int
	BatteryCapacityPercent int
	BatteryRemainingBackupTime int
	BatteryGroupNumber int
	RatedVA int
	RatedOutputFrequency int
	RatedOutputVoltage int
	RatedOutputCurrent int
	RatedBatteryVoltage int
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
