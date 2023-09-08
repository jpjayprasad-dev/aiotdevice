package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"github.com/spf13/viper"
)

type Record struct {
	DateTime  string `csv:"datetime"`
	DeviceID  string `csv:"device_id"`
	DataPoint string `csv:"datapoint"`
	Value     string `csv:"value"`
}

type Control struct {
	DateTime     string `json:"datetime"`
	RoomID       int    `json:"room_id"`
	DeviceID     string `json:"device_id"`
	ControlPoint string `json:"controlpoint"`
	Value        string `json:"value"`
}

func main() {

	// Config
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	input_file := viper.GetString("room.infile")
	output_file := viper.GetString("room.outfile")
	port := viper.GetString("room.port")
	room_id := viper.GetString("room.id")

	records, err := ReadCsv(input_file)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/:device_id/data", func(c *gin.Context) {
		device_id := c.Param("device_id")
		var record_to_send Record
		for i, record := range records {
			fmt.Printf("Datapoint: %s, Value: %s\n", record.DataPoint, record.Value)
			if record.DeviceID == device_id {
				record_to_send = record
				records = append(records[:i], records[i+1:]...)
				break
			}
		}

		c.JSON(200, gin.H{
			"datetime":  record_to_send.DateTime,
			"datapoint": record_to_send.DataPoint,
			"value":     record_to_send.Value,
			"device_id": device_id,
			"room_id":   room_id,
		})
	})

	var controls []Control

	r.POST("/control", func(c *gin.Context) {
		var control Control

		if err := c.ShouldBindJSON(&control); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		control.DateTime = time.Now().Format(time.RFC850)
		controls = append(controls, control)
		WriteJson(output_file, controls)
		c.JSON(http.StatusCreated, control)
	})

	r.Run(":" + port)

}

// ReadCsv accepts a file and returns array of Record
func ReadCsv(filename string) ([]Record, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return []Record{}, err
	}
	defer f.Close()

	// Read the CSV file into a slice of Record structs
	var records []Record
	if err := gocsv.UnmarshalFile(f, &records); err != nil {
		panic(err)
	}

	return records, nil
}

// WriteJson accepts array of controls and writes it back to
// a json file
func WriteJson(filename string, controls []Control) error {
	file, err := json.MarshalIndent(controls, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		return err
	}
	return nil
}
