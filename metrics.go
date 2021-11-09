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
		UPSTemperature:             float64(getInt(splitRealInfo[3])) / 10,
		AutoReboot:                 getBool(splitRealInfo[4]),
		ConverterMode:              getBool(splitRealInfo[5]),
		EcoMode:                    getBool(splitRealInfo[6]),
		BypassWhenOff:              getBool(splitRealInfo[7]),
		BypassDisabled:             getBool(splitRealInfo[8]),
		BatteryVoltage:             float64(getInt(splitRealInfo[11])) / 10,
		BatteryCapacityPercent:     float64(getInt(splitRealInfo[12])),
		BatteryRemainingBackupTime: float64(getInt(splitRealInfo[13])),
		InputFrequency:             float64(getInt(splitRealInfo[14])) / 10,
		InputVoltage:               float64(getInt(splitRealInfo[15])) / 10,
		OutputFrequency:            float64(getInt(splitRealInfo[17])) / 10,
		OutputVoltage:              float64(getInt(splitRealInfo[18])) / 10,
		OutputLoadPercent:          float64(getInt(splitRealInfo[20])),
		BypassFrequency:            float64(getInt(splitRealInfo[21])) / 10,
		LoadLevelS:                 float64(getInt(splitRealInfo[28])),
		BypassVoltageS:             float64(getInt(splitRealInfo[29])) / 10,
		BypassV23:                  float64(getInt(splitRealInfo[30])) / 10,
		LoadLevelT:                 float64(getInt(splitRealInfo[35])),
		BypassVoltageT:             float64(getInt(splitRealInfo[36])) / 10,
		BypassV31:                  float64(getInt(splitRealInfo[37])) / 10,
		OutputCurrent:              float64(getInt(splitRealInfo[38])) / 10,
		OutputCurrentS:             float64(getInt(splitRealInfo[39])) / 10,
		OutputCurrentT:             float64(getInt(splitRealInfo[40])) / 10,
		OutputApparentPower:        float64(getInt(splitRealInfo[48])),
		OutputActivePower:          float64(getInt(splitRealInfo[49])),
		BatteryChargingCurrent:     float64(getInt(splitRealInfo[50])),
		BatteryDischargingCurrent:  float64(getInt(splitRealInfo[51])),

		// stuff from baseInfo.cgi
		UPSType:              splitBaseInfo[3],
		InputPhase:           float64(getInt(strings.Split(splitBaseInfo[4], "/")[0])),
		OutputPhase:          float64(getInt(strings.Split(splitBaseInfo[4], "/")[1])),
		UPSSerial:            splitBaseInfo[6],
		UPSFirmware:          splitBaseInfo[7],
		BatteryGroupNumber:   float64(getInt(splitBaseInfo[8])),
		RatedVA:              float64(getInt(splitBaseInfo[9])) / 10,
		RatedOutputFrequency: float64(getInt(splitBaseInfo[11])) / 10,
		RatedOutputVoltage:   float64(getInt(splitBaseInfo[10])) / 10,
		RatedOutputCurrent:   float64(getInt(splitBaseInfo[12])) / 10,
		RatedBatteryVoltage:  float64(getInt(splitBaseInfo[13])) / 10,
		SNMPFirmware:         splitBaseInfo[14],
		EquipmentAttached:    splitBaseInfo[15],
	}

	return status, nil
}
