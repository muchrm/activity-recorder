package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"influxproject/models"
	"time"
)

type (
	XYZType struct {
		X    string
		Y    string
		Z    string
	}
	ArrayData struct{
			Name string
			Time string
			Device string
			Type string
			Field string
			DataName string
	}
	DataControls struct{
		beego.Controller
	}
	Response struct {
		Status string `json:"status"`
	}
	DataChanel struct {
		Name   string
		Tags   map[string]string
		Fields map[string]interface{}
		Time   time.Time
	}
)

var (
	C chan DataChanel
)

func post_accel(value ArrayData){
	var field XYZType
	if err := json.Unmarshal([]byte(value.Field), &field); err != nil {
		return
	} else{
		var data DataChanel
		data.Name = "Accelerometer"
		data.Tags = map[string]string{
				"username": value.Name,
				"type": value.Type,
				"device": value.Device,
			}
		data.Fields = map[string]interface{}{
				"x": field.X,
				"y": field.Y,
				"z": field.Z,
			}
		timeStampString := value.Time
		layout := "2006-01-02T15:04:05.000Z"
		timeStamp, _ := time.Parse(layout, timeStampString)
		data.Time = timeStamp
		//log.Print(data)
		C <- data
	}
}

func (this *DataControls) Post() {
	var obj []ArrayData
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &obj); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("{\"status\":\"fail\"}"))
		return
	} else {
		for _,value := range obj {
				post_accel(value)
		}
		this.Data["json"] = Response{"success"}
		this.ServeJSON()
	}
}

func WaitInsert() {
	for {
		r := <-C
		go models.MyInflux.WriteData(r.Name, r.Tags, r.Fields, r.Time)
	}
}
func init() {
	C = make(chan DataChanel, 10)
	go WaitInsert()
}
