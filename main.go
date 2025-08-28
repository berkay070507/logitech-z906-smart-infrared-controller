package main

import (
	_ "github.com/joho/godotenv/autoload" // Important to keep this as the first import!

	"github.com/nuriofernandez/logitech-z906-smart-infrared-controller/tuya"

	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/power", power)
	http.HandleFunc("/up", up)
	http.HandleFunc("/down", down)
	http.ListenAndServe(":5566", nil)
}

func power(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Power!")
	tuya.SendRemoteControl("1")
}
func up(w http.ResponseWriter, req *http.Request) {
	fmt.Println("UP!")
	tuya.SendRemoteControl("50")
}
func down(w http.ResponseWriter, req *http.Request) {
	fmt.Println("DOWN !")
	tuya.SendRemoteControl("51")
}
