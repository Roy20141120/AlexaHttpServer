package myserver

import (
	"github.com/b00giZm/golexa"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"math/rand"
	"fmt"
	"common"
	"log"
)

const (
	INTENT_SWITCH_CHANNEL = "MytvSwitchChannel"
	INTENT_SET_VOLUME = "MytvSetVolume"
	INTENT_WATCH_MOVIE = "MytvWatchMovie"
)

const (
	CHANNEL_NAME = "ChannelName"
	CHANNEL_NUMBER = "ChannelNumber"
	VOLUME_NUMBER = "VolumeNumber"
	MOVIE_NAME = "MovieName"
)

var GREETING = [...]string {
	"welcome to my tv alexa skill, what can i do for you",
	"welcome to my tv, it's really my honor to serve you",
	"nice to meet you, i'm ready to serve you",
	"nice to meet you again, i'm very pleased to serve you"}


type Slot struct {
	Name string
	Value string
}

type Intent struct {
	Name string
	Slots []Slot
}


func handler(writer http.ResponseWriter, request *http.Request) {

	//Chan := make(chan string)
	app := golexa.Default()

	app.OnLaunch(func(alexa *golexa.Alexa, req *golexa.Request, session *golexa.Session) *golexa.Response {

		greetIndex := rand.Intn(100) % len(GREETING)
		randomGreet := GREETING[greetIndex]

		return alexa.Response().AddPlainTextSpeech(randomGreet)
	})

	app.OnIntent(func(alexa *golexa.Alexa, intent *golexa.Intent, req *golexa.Request, session *golexa.Session) *golexa.Response {

		var temp string
		intentName := req.Intent.Name;

		var msg string

		switch {
		case intentName == INTENT_SWITCH_CHANNEL:
			fmt.Println("MytvSwitchChannel")
			temp += "ok, my tv will switch to channel "
		case intentName == INTENT_SET_VOLUME:
			fmt.Println("MytvSetVolume ")
			temp += "ok, my tv will set volume to "
		case intentName == INTENT_WATCH_MOVIE:
			fmt.Println("MytvWatchMovie ")
			temp += "ok, my tv will search "
		default:
			fmt.Println("unknown")
			temp += "Sorry, I am confused"
		}

		for slot := range req.Intent.Slots {
			if len(req.Intent.Slots[slot].Value) > 0 {
				slotName := req.Intent.Slots[slot].Name

				switch {
				case slotName == CHANNEL_NAME:
					fmt.Println("CHANNEL_NAME:", req.Intent.Slots[slot].Value)
					temp += req.Intent.Slots[slot].Value
					//msg = common.Msg{IntentType:common.SWITCH_CHANNEL_BY_NAME, Value:req.Intent.Slots[slot].Value}
					msg += common.SWITCH_CHANNEL_BY_NAME
					msg += ":"
					msg += req.Intent.Slots[slot].Value
				case slotName == CHANNEL_NUMBER:
					fmt.Println("CHANNEL_NUMBER:", req.Intent.Slots[slot].Value)
					temp += req.Intent.Slots[slot].Value
					//msg = common.Msg{IntentType:common.SWITCH_CHANNEL_BY_NUMBER, Value:req.Intent.Slots[slot].Value}
					msg += common.SWITCH_CHANNEL_BY_NUMBER
					msg += ":"
					msg += req.Intent.Slots[slot].Value
				case slotName == VOLUME_NUMBER:
					fmt.Println("VOLUME_NUMBER:", req.Intent.Slots[slot].Value)
					temp += req.Intent.Slots[slot].Value
					//msg = common.Msg{IntentType:common.SET_VOLUME, Value:req.Intent.Slots[slot].Value}
					msg += common.SET_VOLUME
					msg += ":"
					msg += req.Intent.Slots[slot].Value
				case slotName == MOVIE_NAME:
					fmt.Println("MOVIE_NAME:", req.Intent.Slots[slot].Value)
					temp += req.Intent.Slots[slot].Value
					//msg = common.Msg{IntentType:common.WATCH_MOVIE, Value:req.Intent.Slots[slot].Value}
					msg += common.WATCH_MOVIE
					msg += ":"
					msg += req.Intent.Slots[slot].Value
				default:
					fmt.Println("unknown slot name")
				}
			}
		}

		go func(){
			//jsonMsg,_ := json.Marshal(msg)
			//log.Println("event:", msg)
			//log.Println("http send msg to socket")
			common.HttpToSocket <- msg
			log.Println("http send msg to socket success")
		}()

		return alexa.Response().AddPlainTextSpeech(temp)
	})

	result, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Println("handler post error: ", err.Error())
	}

	fmt.Println("handler post result: ", string(result))
	request.Body.Close()

	msg := json.RawMessage(result)
	response,_ := app.Process(msg)
	re,_ := json.Marshal(response)
	writer.Write(re)
}


//func handler(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("hello"))
//	app := golexa.Default()
//	app.OnLaunch(func(a *golexa.Alexa, req *golexa.Request, session *golexa.Session) *golexa.Response {
//		log.Println("my OnLaunch func")
//
//		return a.Response().AddPlainTextSpeech("Welcome to my awesome app")
//	})
//	app.OnIntent(func(a *golexa.Alexa, intent *golexa.Intent, req *golexa.Request, session *golexa.Session) *golexa.Response{
//		log.Println("my OnIntent func")
//		return a.Response().AddPlainTextSpeech("Welcome to my own app")
//	})
//	result , err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		log.Println("handler post error: ", err.Error())
//	}
//
//	log.Println("handler post result : ", string(result))
//
//	r.Body.Close()
//	msg := json.RawMessage(result)
//	response,_ := app.Process(msg)
//	re,_ := json.Marshal(response)
//	w.Write(re)
//
//	//common.HttpToSocket <- msg
//	log.Println("pretend to get msg from alexa")
//	common.HttpToSocket <- common.Msg{Debug:"get debug msg ============"}
//
//	//for test
//	//m := common.Msg{
//	//	Meta:map[string]interface{}{
//	//		"meta":"test",
//	//		"ID":1234,
//	//	},
//	//	Content: "for test",
//	//}
//	//common.HttpToSocket <- m
//	//
//	//m = <- common.SocketToHttp
//	//re,_ := json.Marshal(m)
//	//w.Write(re)
//}

func StartHttpServer(port string)  {
	http.HandleFunc("/", handler)
	http.ListenAndServe(port, nil)
}


func StartHttpsServer(port string)  {
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(port, "server.crt", "server.key", nil)
}
