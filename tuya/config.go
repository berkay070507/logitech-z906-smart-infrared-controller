package tuya

import "os"

var (
	ClientID      = os.Getenv("CLIENT_ID")
	Secret        = os.Getenv("SECRET")
	Host          = os.Getenv("HOST")
	IRDeviceID    = os.Getenv("IR_DEVICE_ID")
	AudioDeviceID = os.Getenv("AUDIO_DEVICE_ID")
)
