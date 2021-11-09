package main

import (
	"fmt"
	"strconv"
)

func generateHelpTypeString(name string, helpText string) string {
	return "# HELP " + name + " " + helpText + "\n" +
		"# TYPE " + name + " gauge\n"
}

func getMetricsString(s Status) string {
	intVal := map[bool]int{true: 1}
	return "" +
		// temperature
		generateHelpTypeString("ups_temperature", "UPS Temperature.") +
		`ups_temperature{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + fmt.Sprintf("%.2f", s.UPSTemperature) + "\n" +

		// input
		generateHelpTypeString("ups_input_voltage", "Input voltage.") +
		`ups_input_voltage{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + fmt.Sprintf("%.2f", s.InputVoltage) + "\n" +
		generateHelpTypeString("ups_input_frequency", "Input frequency.") +
		`ups_input_frequency{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + fmt.Sprintf("%.2f", s.InputFrequency) + "\n" +
		generateHelpTypeString("ups_input_phases", "Number of input phases.") +
		`ups_input_phases{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(s.InputPhase) + "\n" +

		// output
		generateHelpTypeString("ups_output_phases", "Number of output phases.") +
		`ups_output_phases{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(s.OutputPhase) + "\n" +
		generateHelpTypeString("ups_output_voltage", "Output voltage.") +
		`ups_output_voltage{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + fmt.Sprintf("%.2f", s.OutputVoltage) + "\n" +
		generateHelpTypeString("ups_output_frequency", "Output frequency.") +
		`ups_output_frequency{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + fmt.Sprintf("%.2f", s.OutputFrequency) + "\n" +
		generateHelpTypeString("ups_output_current", "Output current.") +
		`ups_output_current{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + fmt.Sprintf("%.2f", s.OutputCurrent) + "\n" +
		generateHelpTypeString("ups_output_load_percent", "Output load in percent.") +
		`ups_output_load_percent{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(s.OutputLoadPercent) + "\n" +

		// battery
		generateHelpTypeString("ups_battery_voltage", "Battery voltage.") +
		`ups_battery_voltage{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + fmt.Sprintf("%.2f", s.BatteryVoltage) + "\n" +
		generateHelpTypeString("ups_battery_capacity_percent", "Battery capacity in percent.") +
		`ups_battery_capacity_percent{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(s.BatteryCapacityPercent) + "\n" +
		generateHelpTypeString("ups_battery_remaining_backup_time", "Battery remaining backup time.") +
		`ups_battery_remaining_backup_time{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(s.BatteryRemainingBackupTime) + "\n" +
		generateHelpTypeString("ups_battery_group_number", "Battery groups.") +
		`ups_battery_group_number{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(s.BatteryGroupNumber) + "\n" +

		// rated values
		generateHelpTypeString("ups_rated_va", "Rated VA of the UPS.") +
		`ups_rated_va{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + fmt.Sprintf("%.2f", s.RatedVA) + "\n" +
		generateHelpTypeString("ups_rated_output_frequency", "Rated output frequency of the UPS.") +
		`ups_rated_output_frequency{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + fmt.Sprintf("%.2f", s.RatedOutputFrequency) + "\n" +
		generateHelpTypeString("ups_rated_output_voltage", "Rated output voltage of the UPS.") +
		`ups_rated_output_voltage{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + fmt.Sprintf("%.2f", s.RatedOutputVoltage) + "\n" +
		generateHelpTypeString("ups_rated_output_current", "Rated output current of the UPS.") +
		`ups_rated_output_current{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + fmt.Sprintf("%.2f", s.RatedOutputCurrent) + "\n" +
		generateHelpTypeString("ups_rated_battery_voltage", "Rated battery voltage of the UPS.") +
		`ups_rated_battery_voltage{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + fmt.Sprintf("%.2f", s.RatedBatteryVoltage) + "\n" +

		// state
		generateHelpTypeString("ups_auto_reboot", "Auto reboot enabled.") +
		`ups_auto_reboot{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(intVal[s.AutoReboot]) + "\n" +
		generateHelpTypeString("ups_eco_mode", "ECO mode enabled.") +
		`ups_eco_mode{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(intVal[s.EcoMode]) + "\n" +
		generateHelpTypeString("ups_bypass_disabled", "Bypass disabled.") +
		`ups_bypass_disabled{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(intVal[s.BypassDisabled]) + "\n" +
		generateHelpTypeString("ups_bypass_when_off", "Bypass when off.") +
		`ups_bypass_when_off{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(intVal[s.BypassWhenOff]) + "\n" +
		generateHelpTypeString("ups_converter_mode", "Converter mode active.") +
		`ups_converter_mode{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(intVal[s.ConverterMode]) + "\n"
}

func getFetchSuccessString(upsType string, upsSerial string, state int) string {
	return generateHelpTypeString("fetch_success", "Metrics collection success.") +
		"fetch_success{type=\"" + upsType + "\",serial=\"" + upsSerial + "\"} " + strconv.Itoa(state) + "\n"
}

func getFetchDurationString(upsType string, upsSerial string, duration int64) string {
	return generateHelpTypeString("fetch_duration", "Metrics collection duration.") +
		"fetch_duration{type=\"" + upsType + "\",serial=\"" + upsSerial + "\"} " + strconv.FormatInt(duration, 10) + "\n"
}
