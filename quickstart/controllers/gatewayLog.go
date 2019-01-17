package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
	"github.com/gpmgo/gopm/modules/log"
	"strconv"
	"time"
)

type GatewayLogController struct {
	beego.Controller
}

type dataForShow struct {
	Time string
	RxTxInfo string
	Payload string
}

type EventLog struct {
	Type string
	Payload interface{}
}

type data struct {
	Result dataForShow
}

func (c *GatewayLogController) Get() {

	var ch = "lora:as:device:bbbbbbbbbbbbbbbb:pubsub:event"

	pool := newPool()

	co := pool.Get()
	defer co.Close()
	psc := redis.PubSubConn{Conn: co}
	if err := psc.Subscribe(ch); err != nil {
		log.Warn("redis subscribe error: ", err)
	}

	eventsChan := make(chan EventLog, 5)

	go func() {
		send(psc, eventsChan)
		close(eventsChan)
	}()

	for el := range eventsChan {
		res := receive(el)

		fmt.Println("************")
		fmt.Println(res)
		fmt.Println("************")

		fmt.Fprint(c.Ctx.ResponseWriter, res)
	}

}

func receive(e EventLog) []byte {
	show := parsePayload(e.Payload)
	fmt.Println("*************")
	fmt.Println("data for show: ", show)
	fmt.Println("*************")

	var d data
	d.Result = show
	b, err := json.Marshal(d)
	if err != nil {
		fmt.Println(err)
	}

	return b
}

func send(connection redis.PubSubConn, eventsChan chan EventLog) {
	for {
		switch v := connection.Receive().(type) {
		case redis.Message:
			el, err := redisMessageToEventLog(v)
			if err != nil {
				log.Warn("decode message error", err)
			}else {
				eventsChan <- el
			}
		default:
			log.Info("%s", "this is not event message, waiting...")
		}
	}
}

func redisMessageToEventLog(msg redis.Message) (EventLog, error) {
	var el EventLog
	if err := json.Unmarshal(msg.Data, &el); err != nil {
		return el, errors.New("gob decode error")
	}

	return el, nil
}

func newPool() *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 1200,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "192.168.3.157:6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func getData(freq, spred, rssi float64, payload string) dataForShow {
	var res dataForShow

	hour, minute, second := time.Now().Clock()
	res.Time = strconv.Itoa(hour) + ":" + strconv.Itoa(minute) + ":" + strconv.Itoa(second)

	f := strconv.FormatFloat(freq, 'f', 0, 64)
	s := strconv.FormatFloat(spred, 'f', 0, 64)
	r := strconv.FormatFloat(rssi, 'f', 0, 64)
	res.RxTxInfo = f + "," + s + "," + r

	res.Payload = payload

	return res
}

func parsePayload(p interface{}) dataForShow {
	var frequency, spreadFactor, rssi float64
	var payload string

	if v, ok := p.(map[string]interface{}); ok {
		txInfo := v["txInfo"].(map[string]interface{})
		fmt.Printf("frequency: %f\n", txInfo["frequency"])
		frequency = txInfo["frequency"].(float64)

		dataRate := txInfo["dataRate"].(map[string]interface{})
		fmt.Printf("spreadFactor: %f\n", dataRate["spreadFactor"])
		spreadFactor = dataRate["spreadFactor"].(float64)

		rxInfo := v["rxInfo"].([]interface{})
		for _, u := range rxInfo {
			uu := u.(map[string]interface{})
			fmt.Printf("rssi: %f\n", uu["rssi"])
			rssi = uu["rssi"].(float64)
		}

		fmt.Printf("data: %s\n", v["data"])
		payload = v["data"].(string)
	}

	return getData(frequency, spreadFactor, rssi, payload)
}