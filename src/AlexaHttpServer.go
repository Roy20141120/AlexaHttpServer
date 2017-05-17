package main
//
//import (
//	"github.com/b00giZm/golexa"
//	"net/http"
//	"io/ioutil"
//	"fmt"
//	"encoding/json"
//)
//
//func handler(w http.ResponseWriter, r *http.Request) {
//	app := golexa.Default()
//	app.OnLaunch(func(a *golexa.Alexa, req *golexa.Request, session *golexa.Session) *golexa.Response {
//		fmt.Println("my OnLaunch func")
//		return a.Response().AddPlainTextSpeech("Welcome to my awesome app")
//	})
//	app.OnIntent(func(a *golexa.Alexa, intent *golexa.Intent, req *golexa.Request, session *golexa.Session) *golexa.Response{
//		fmt.Println("my OnIntent func")
//		return a.Response().AddPlainTextSpeech("Welcome to my own app")
//	})
//	result , err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		fmt.Println("handler post error: ", err.Error())
//	}
//
//	fmt.Println("handler post result : ", string(result))
//
//	r.Body.Close()
//	msg := json.RawMessage(result)
//	response,_ := app.Process(msg)
//	re,_ := json.Marshal(response)
//	w.Write(re)
//}
//func main() {
//	http.HandleFunc("/", handler)
//	http.ListenAndServe(":80", nil)
//}
