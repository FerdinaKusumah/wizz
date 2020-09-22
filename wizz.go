package wizz

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/FerdinaKusumah/wizz/connection"
	"github.com/FerdinaKusumah/wizz/models"
	"github.com/FerdinaKusumah/wizz/utils"
	"log"
	"os"
	"time"
)

var (
	bulbIp, method string
	value, r, g, b float64
	sceneId        int64
)

func flags() {
	flag.StringVar(&bulbIp, "bulbIp", "192.168.1.9", "Port buob in local network")
	flag.Float64Var(&value, "value", 20.0, "Value to adjust in bulb")
	flag.Float64Var(&r, "r", 20.0, "Value to adjust in RGB Mode")
	flag.Float64Var(&g, "g", 20.0, "Value to adjust in RGB Mode")
	flag.Float64Var(&b, "b", 1.0, "Value to adjust in RGB Mode")
	flag.Int64Var(&sceneId, "sceneId", 20.0, "Set bulb in specific scene mode")
	flag.StringVar(&method, "method", "GetState", "Available method is GetState, GetConfig, TurnOnLight, TurnOffLight, SetColorTemp, SetBrightness, SetColorRGB, SetColorScene, SetColorWarmWhite, SetColorColdWhite")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func init() {
	flags()
}

func main() {
	flag.Parse()

	var (
		response = new(models.ResponsePayload)
		result   []byte
		err      error
	)
	switch method {
	case "GetState":
		if response, err = GetState(bulbIp); err != nil {
			log.Fatalf(`Unable to read response: %s`, err)
		}
	case "GetConfig":
		if response, err = GetConfig(bulbIp); err != nil {
			log.Fatalf(`Unable to read response: %s`, err)
		}
	case "TurnOnLight":
		if response, err = TurnOnLight(bulbIp); err != nil {
			log.Fatalf(`Unable to read response: %s`, err)
		}
	case "TurnOffLight":
		if response, err = TurnOffLight(bulbIp); err != nil {
			log.Fatalf(`Unable to read response: %s`, err)
		}
	case "SetColorTemp":
		if response, err = SetColorTemp(bulbIp, value); err != nil {
			log.Fatalf(`Unable to read response: %s`, err)
		}
	case "SetBrightness":
		if response, err = SetBrightness(bulbIp, value); err != nil {
			log.Fatalf(`Unable to read response: %s`, err)
		}
	case "SetColorRGB":
		if response, err = SetColorRGB(bulbIp, r, g, b); err != nil {
			log.Fatalf(`Unable to read response: %s`, err)
		}
	case "SetColorScene":
		if response, err = SetColorScene(bulbIp, sceneId); err != nil {
			log.Fatalf(`Unable to read response: %s`, err)
		}
	case "SetColorWarmWhite":
		if response, err = SetColorWarmWhite(bulbIp, value); err != nil {
			log.Fatalf(`Unable to read response: %s`, err)
		}
	case "SetColorColdWhite":
		if response, err = SetColorColdWhite(bulbIp, value); err != nil {
			log.Fatalf(`Unable to read response: %s`, err)
		}
	default:
		if response, err = GetState(bulbIp); err != nil {
			log.Fatalf(`Unable to read response: %s`, err)
		}
	}
	if result, err = json.Marshal(response); err != nil {
		log.Fatalf(`Unable to convert to json string: %s`, err)
	}
	fmt.Println(string(result))
	os.Exit(0)
}

func GetState(bulbIp string) (*models.ResponsePayload, error) {
	var (
		response = new(models.ResponsePayload)
		err      error
		payload  = &models.RequestPayload{
			Method: "getPilot",
		}
	)
	if response, err = connection.SendUdpMessage(bulbIp, payload); err != nil {
		log.Fatalf(`Unable to send message to udp: %s`, err)
		return nil, err
	}
	time.Sleep(time.Second * 1)
	return response, nil
}

func GetConfig(bulbIp string) (*models.ResponsePayload, error) {
	var (
		response = new(models.ResponsePayload)
		err      error
		payload  = &models.RequestPayload{
			Method: "getSystemConfig",
		}
	)
	if response, err = connection.SendUdpMessage(bulbIp, payload); err != nil {
		log.Fatalf(`Unable to send message to udp: %s`, err)
		return nil, err
	}
	time.Sleep(time.Second * 1)
	return response, nil
}

func TurnOnLight(bulbIp string) (*models.ResponsePayload, error) {
	var (
		response = new(models.ResponsePayload)
		err      error
		payload  = &models.RequestPayload{
			Method: "setPilot",
			Params: models.ParamPayload{
				State: true,
				Speed: 50, // must between 0 - 100
			},
		}
	)
	if response, err = connection.SendUdpMessage(bulbIp, payload); err != nil {
		log.Fatalf(`Unable to send message to udp: %s`, err)
		return nil, err
	}
	time.Sleep(time.Second * 1)
	return response, nil
}

func TurnOffLight(bulbIp string) (*models.ResponsePayload, error) {
	var (
		response = new(models.ResponsePayload)
		err      error
		payload  = &models.RequestPayload{
			Method: "setPilot",
			Params: models.ParamPayload{
				State: false,
				Speed: 50, // must between 0 - 100
			},
		}
	)
	if response, err = connection.SendUdpMessage(bulbIp, payload); err != nil {
		log.Fatalf(`Unable to send message to udp: %s`, err)
		return nil, err
	}
	time.Sleep(time.Second * 1)
	return response, nil
}

func SetColorTemp(bulbIp string, value float64) (*models.ResponsePayload, error) {
	var (
		response = new(models.ResponsePayload)
		err      error
		payload  = &models.RequestPayload{
			Method: "setPilot",
		}
	)

	// normalize the kelvin values - should be removed
	if value < 2500 {
		value = 2500
	}
	if value > 6500 {
		value = 6500
	}

	payload.Params = models.ParamPayload{
		ColorTemp: value,
		State:     true,
		Speed:     50, // must between 0 - 100
	}
	if response, err = connection.SendUdpMessage(bulbIp, payload); err != nil {
		log.Fatalf(`Unable to send message to udp: %s`, err)
		return nil, err
	}
	time.Sleep(time.Second * 1)
	return response, nil
}

func SetBrightness(bulbIp string, value float64) (*models.ResponsePayload, error) {
	var (
		response = new(models.ResponsePayload)
		err      error
		payload  = &models.RequestPayload{
			Method: "setPilot",
		}
	)
	brightnessPercent := utils.HexToPercent(value)
	if brightnessPercent < 10 {
		brightnessPercent = 10
	}

	payload.Params = models.ParamPayload{
		Dimming: int64(brightnessPercent),
		State:   true,
		Speed:   50, // must between 0 - 100
	}
	if response, err = connection.SendUdpMessage(bulbIp, payload); err != nil {
		log.Fatalf(`Unable to send message to udp: %s`, err)
		return nil, err
	}
	time.Sleep(time.Second * 1)
	return response, nil
}

func SetColorRGB(bulbIp string, r, g, b float64) (*models.ResponsePayload, error) {
	var (
		response = new(models.ResponsePayload)
		err      error
		payload  = &models.RequestPayload{
			Method: "setPilot",
		}
	)

	if r > 255 {
		r = 255
	}
	if g > 255 {
		g = 255
	}
	if b > 255 {
		b = 255
	}

	payload.Params = models.ParamPayload{
		R:     r,
		G:     g,
		B:     b,
		State: true,
		Speed: 50, // must between 0 - 100
	}
	if response, err = connection.SendUdpMessage(bulbIp, payload); err != nil {
		log.Fatalf(`Unable to send message to udp: %s`, err)
		return nil, err
	}
	time.Sleep(time.Second * 1)
	return response, nil
}

func SetColorScene(bulbIp string, sceneId int64) (*models.ResponsePayload, error) {
	var (
		response = new(models.ResponsePayload)
		err      error
		exists   bool
		payload  = &models.RequestPayload{
			Method: "setPilot",
			Params: models.ParamPayload{
				SceneId: 1,
				State:   true,
				Speed:   50, // must between 0 - 100
			},
		}
	)
	if _, exists = models.SceneModel[sceneId]; exists == true {
		payload.Params.SceneId = sceneId
	}
	if response, err = connection.SendUdpMessage(bulbIp, payload); err != nil {
		log.Fatalf(`Unable to send message to udp: %s`, err)
		return nil, err
	}
	time.Sleep(time.Second * 1)
	return response, nil
}

func SetColorWarmWhite(bulbIp string, value float64) (*models.ResponsePayload, error) {
	var (
		response = new(models.ResponsePayload)
		err      error
		payload  = &models.RequestPayload{
			Method: "setPilot",
			Params: models.ParamPayload{
				WarmWhite: 0,
				State:     true,
				Speed:     50, // must between 0 - 100
			},
		}
	)
	if value < 0 {
		value = 0
	}

	if value > 256 {
		value = 256
	}
	payload.Params.WarmWhite = value

	if response, err = connection.SendUdpMessage(bulbIp, payload); err != nil {
		log.Fatalf(`Unable to send message to udp: %s`, err)
		return nil, err
	}
	time.Sleep(time.Second * 1)
	return response, nil
}

func SetColorColdWhite(bulbIp string, value float64) (*models.ResponsePayload, error) {
	var (
		response = new(models.ResponsePayload)
		err      error
		payload  = &models.RequestPayload{
			Method: "setPilot",
			Params: models.ParamPayload{
				ColdWhite: 0,
				State:     true,
				Speed:     50, // must between 0 - 100
			},
		}
	)
	if value < 0 {
		value = 0
	}

	if value > 256 {
		value = 256
	}
	payload.Params.ColdWhite = value

	if response, err = connection.SendUdpMessage(bulbIp, payload); err != nil {
		log.Fatalf(`Unable to send message to udp: %s`, err)
		return nil, err
	}
	time.Sleep(time.Second * 1)
	return response, nil
}
