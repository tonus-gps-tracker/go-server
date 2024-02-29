package service

import (
	"errors"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/tonus-gps-tracker/server/internal/common"
	"github.com/tonus-gps-tracker/server/internal/dto"
	"github.com/tonus-gps-tracker/server/internal/influxdb/repository"
)

const delimiter = ","

type GpsTrackerService struct {
	locationRepository *repository.LocationRepository
}

func (service *GpsTrackerService) Save(body string) error {
	log.Println(body)

	lines := strings.Split(string(body), "\n")

	for _, line := range lines {
		log.Println(line)
		lineData := strings.Split(line, delimiter)
		log.Println(lineData)

		if (len(lineData) - 1) != reflect.TypeOf(dto.GPSData{}).NumField() {
			return errors.New("incompatible number of fields")
		}

		timestamp, err := strconv.ParseInt(lineData[0], 10, 64)
		if err != nil {
			log.Printf("[ERROR] GpsTrackerService_Save, strconv.ParseInt: %s\n", err)
			continue
		}

		gpsData := dto.GPSData{
			Latitude:    common.StringToFloat32(lineData[1]),
			Longitude:   common.StringToFloat32(lineData[2]),
			Altitude:    common.StringToFloat32(lineData[3]),
			Speed:       common.StringToFloat32(lineData[4]),
			NSatellites: common.StringToInt(lineData[5]),
		}

		service.locationRepository.Save(time.Unix(timestamp, 0), gpsData)
	}

	return nil
}

func NewGpsTrackerService(locationRepository *repository.LocationRepository) *GpsTrackerService {
	return &GpsTrackerService{
		locationRepository: locationRepository,
	}
}
