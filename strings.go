package main

import (
	"strconv"
)

func getHelpString() string {
	return "" +
		// temperature
		"# HELP ups_temperature UPS Temperature.\n" +
		"# TYPE ups_temperature gauge\n" +

		// input
		"# HELP ups_input_voltage Input voltage.\n" +
		"# TYPE ups_input_voltage gauge\n" +
		"# HELP ups_input_frequency Input voltage.\n" +
		"# TYPE ups_input_frequency gauge\n" +
		"# HELP ups_input_phases Number of input phases.\n" +
		"# TYPE ups_input_phases gauge\n" +

		// output
		"# HELP ups_output_phases Number of output phases.\n" +
		"# TYPE ups_output_phases gauge\n" +
		"# HELP ups_output_voltage Output voltage.\n" +
		"# TYPE ups_output_voltage gauge\n" +
		"# HELP ups_output_frequency Output frequency.\n" +
		"# TYPE ups_output_frequency gauge\n" +
		"# HELP ups_output_current Output current.\n" +
		"# TYPE ups_output_current gauge\n" +
		"# HELP ups_output_load_percent Output load in percent.\n" +
		"# TYPE ups_output_load_percent gauge\n" +

		// battery
		"# HELP ups_battery_voltage Battery voltage.\n" +
		"# TYPE ups_battery_voltage gauge\n" +
		"# HELP ups_battery_capacity_percent Battery capacity in percent.\n" +
		"# TYPE ups_battery_capacity_percent gauge\n" +
		"# HELP ups_battery_remaining_backup_time Battery remaining backup time.\n" +
		"# TYPE ups_battery_remaining_backup_time gauge\n" +
		"# HELP ups_battery_group_number Battery groups.\n" +
		"# TYPE ups_battery_group_number gauge\n" +

		// rated values
		"# HELP ups_rated_va Rated VA of the USV.\n" +
		"# TYPE ups_rated_va gauge\n" +
		"# HELP ups_rated_output_frequency Rated output_frequency of the USV.\n" +
		"# TYPE ups_rated_output_frequency gauge\n" +
		"# HELP ups_rated_output_voltage Rated output_voltage of the USV.\n" +
		"# TYPE ups_rated_output_voltage gauge\n" +
		"# HELP ups_rated_output_current Rated output_current of the USV.\n" +
		"# TYPE ups_rated_output_current gauge\n" +
		"# HELP ups_rated_battery_voltage Rated battery_voltage of the USV.\n" +
		"# TYPE ups_rated_battery_voltage gauge\n" +

		// state
		"# HELP ups_auto_reboot Auto reboot enabled.\n" +
		"# TYPE ups_auto_reboot gauge\n" +
		"# HELP ups_eco_mode ECO Mode enabled.\n" +
		"# TYPE ups_eco_mode gauge\n" +
		"# HELP ups_bypass_disabled Bypass disabled.\n" +
		"# TYPE ups_bypass_disabled gauge\n" +
		"# HELP ups_bypass_when_off Bypass when off.\n" +
		"# TYPE ups_bypass_when_off gauge\n" +
		"# HELP ups_converter_mode Converter mode.\n" +
		"# TYPE ups_converter_mode gauge\n" +

		// exporter
		"# HELP fetch_success metrics collection successful\n" +
		"# TYPE fetch_success gauge\n" +
		"# HELP fetch_duration metrics collection duration\n" +
		"# TYPE fetch_duration gauge\n"
}

func getFetchSuccessString(upsType string, upsSerial string, state int) string {
	return "fetch_success{type=\"" + upsType + "\",serial=\"" + upsSerial + "\"} " + strconv.Itoa(state) + "\n"
}

func getFetchDurationString(upsType string, upsSerial string, duration int64) string {
	return "fetch_duration{type=\"" + upsType + "\",serial=\"" + upsSerial + "\"} " + strconv.FormatInt(duration, 10) + "\n"
}

func getMetricsString(s Status) string {
	intVal := map[bool]int{true: 1}

	return `ups_temperature{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(s.UPSTemperature) + "\n" +
		// input
		`ups_input_voltage{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(s.InputVoltage) + "\n" +
		`ups_input_frequency{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(s.InputFrequency) + "\n" +
		`ups_input_phases{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(s.InputPhase) + "\n" +
		// output
		`ups_output_voltage{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(s.OutputVoltage) + "\n" +
		`ups_output_frequency{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(s.OutputFrequency) + "\n" +
		`ups_output_phases{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(s.OutputPhase) + "\n" +
		`ups_output_current{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(s.OutputCurrent) + "\n" +
		`ups_output_load_percent{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(s.OutputLoadPercent) + "\n" +
		// battery
		`ups_battery_voltage{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(s.BatteryVoltage) + "\n" +
		`ups_battery_capacity_percent{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(s.BatteryCapacityPercent) + "\n" +
		`ups_battery_remaining_backup_time{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(s.BatteryRemainingBackupTime) + "\n" +
		`ups_battery_group_number{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(s.BatteryGroupNumber) + "\n" +
		// rated
		`ups_rated_va{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(s.RatedVA) + "\n" +
		`ups_rated_output_frequency{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(s.RatedOutputFrequency) + "\n" +
		`ups_rated_output_voltage{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(s.RatedOutputVoltage) + "\n" +
		`ups_rated_output_current{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(s.RatedOutputCurrent) + "\n" +
		`ups_rated_battery_voltage{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(s.RatedBatteryVoltage) + "\n" +
		// state
		`ups_auto_reboot{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(intVal[s.AutoReboot]) + "\n" +
		`ups_eco_mode{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(intVal[s.EcoMode]) + "\n" +
		`ups_bypass_disabled{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(intVal[s.BypassDisabled]) + "\n" +
		`ups_bypass_when_off{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(intVal[s.BypassWhenOff]) + "\n" +
		`ups_converter_mode{type="` + s.UPSType + `",serial="` + s.UPSSerial + `"} ` + strconv.Itoa(intVal[s.ConverterMode]) + "\n"

}
