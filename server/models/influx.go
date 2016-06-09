package models

import (
    "github.com/influxdb/influxdb/client/v2"
    "github.com/astaxie/beego"
    "time"
    )
var (
    MyInflux *Influx
)
type (
    Influx struct {
    clnt client.Client
    MyDB string
    }
)

func NewInflux() *Influx {
    cln,_ := client.NewHTTPClient(client.HTTPConfig{
        Addr: "http://localhost:8086",
        Username: beego.AppConfig.String("username"),
        Password: beego.AppConfig.String("password"),
    })
    return &Influx{cln,beego.AppConfig.String("DB")}
}
func (m *Influx) QueryDB(cmd string) (res []client.Result, err error) {
    q := client.Query{
        Command:  cmd,
        Database: m.MyDB,
    }
    if response, err := m.clnt.Query(q); err == nil {
        if response.Error() != nil {
            return res, response.Error()
        }
        res = response.Results
    }
    return res, nil
}
func (m *Influx) WriteData(name string,tags map[string]string,fields map[string]interface{},t time.Time){
                bp,_ := client.NewBatchPoints(client.BatchPointsConfig{
                    Database:  m.MyDB,
                    Precision: "us",
                })
                pt,_ := client.NewPoint(name, tags,fields,t)
                bp.AddPoint(pt)
            // Write the batch
                m.clnt.Write(bp)
}
func init() {
	MyInflux = NewInflux()
}