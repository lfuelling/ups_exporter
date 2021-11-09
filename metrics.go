package main

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func getInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		log.Println(err)
		return -1
	}

	return i
}

func getBool(s string) bool {
	b, err := strconv.ParseBool(s)

	if err != nil {
		log.Println(err)
		return false
	}

	return b
}

func getBaseInfo() ([]string, error) {
	baseOut, err := exec.Command("/var/www/html/web_pages/cgi-bin/baseInfo.cgi").Output()
	if err != nil {
		return nil, err
	}

	splitBaseInfo := strings.Split(string(baseOut), "\n")
	return splitBaseInfo, nil
}

func getRealInfo() ([]string, error) {
	realOut, err := exec.Command("/var/www/html/web_pages/cgi-bin/realInfo.cgi").Output()
	if err != nil {
		return nil, err
	}

	splitStatus := strings.Split(string(realOut), "\n")
	return splitStatus, nil
}

func getMetrics() (Status, error) {
	splitBaseInfo, err := getBaseInfo()
	if err != nil {
		return Status{}, err
	}

	splitRealInfo, err1 := getRealInfo()
	if err1 != nil {
		return Status{}, err1
	}

	status := Status{

		// stuff from realInfo.cgi
		OperatingMode:              splitRealInfo[2],
		UPSTemperature:             getInt(splitRealInfo[3]) / 10,
		AutoReboot:                 getBool(splitRealInfo[4]),
		ConverterMode:              getBool(splitRealInfo[5]),
		EcoMode:                    getBool(splitRealInfo[6]),
		BypassWhenOff:              getBool(splitRealInfo[7]),
		BypassDisabled:             getBool(splitRealInfo[8]),
		BatteryVoltage:             getInt(splitRealInfo[11]) / 10,
		BatteryCapacityPercent:     getInt(splitRealInfo[12]),
		BatteryRemainingBackupTime: getInt(splitRealInfo[13]),
		InputFrequency:             getInt(splitRealInfo[14]),
		InputVoltage:               getInt(splitRealInfo[15]) / 10,
		OutputFrequency:            getInt(splitRealInfo[17]) / 10,
		OutputVoltage:              getInt(splitRealInfo[18]) / 10,
		OutputLoadPercent:          getInt(splitRealInfo[20]),
		OutputCurrent:              getInt(splitRealInfo[38]) / 10,

		// stuff from baseInfo.cgi
		UPSType:              splitBaseInfo[3],
		InputPhase:           getInt(strings.Split(splitBaseInfo[4], "/")[0]),
		OutputPhase:          getInt(strings.Split(splitBaseInfo[4], "/")[1]),
		UPSSerial:            splitBaseInfo[6],
		UPSFirmware:          splitBaseInfo[7],
		BatteryGroupNumber:   getInt(splitBaseInfo[8]),
		RatedVA:              getInt(splitBaseInfo[9]) / 10,
		RatedOutputVoltage:   getInt(splitBaseInfo[10]) / 10,
		RatedOutputFrequency: getInt(splitBaseInfo[11]) / 10,
		RatedOutputCurrent:   getInt(splitBaseInfo[12]) / 10,
		RatedBatteryVoltage:  getInt(splitBaseInfo[13]) / 10,
		SNMPFirmware:         splitBaseInfo[14],
		EquipmentAttached:    splitBaseInfo[15],
	}

	return status, nil
}
