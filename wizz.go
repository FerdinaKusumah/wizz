package wizz

import (
	"github.com/FerdinaKusumah/wizz/connection"
	"github.com/FerdinaKusumah/wizz/models"
	"log"
	"time"
)

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
	time.Sleep(time.Second * 2)
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
	time.Sleep(time.Second * 2)
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
	time.Sleep(time.Second * 2)
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
	time.Sleep(time.Second * 2)
	return response, nil
}

func SetColorTemp(bulbIp string, colorTempKelvin int64) (*models.ResponsePayload, error) {
	var (
		response = new(models.ResponsePayload)
		err      error
		payload  = &models.RequestPayload{
			Method: "setPilot",
		}
	)

	// normalize the kelvin values - should be removed
	if colorTempKelvin < 2500 {
		colorTempKelvin = 2500
	}
	if colorTempKelvin > 6500 {
		colorTempKelvin = 6500
	}

	payload.Params = models.ParamPayload{
		ColorTemp: colorTempKelvin,
		State:     true,
		Speed:     50, // must between 0 - 100
	}
	if response, err = connection.SendUdpMessage(bulbIp, payload); err != nil {
		log.Fatalf(`Unable to send message to udp: %s`, err)
		return nil, err
	}
	time.Sleep(time.Second * 2)
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
	time.Sleep(time.Second * 2)
	return response, nil
}

func SetColorRGB(bulbIp string, r, g, b int64) (*models.ResponsePayload, error) {
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
	time.Sleep(time.Second * 2)
	return response, nil
}

func SetColorScene(bulbIp string, scheneId int64) (*models.ResponsePayload, error) {
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
	if _, exists = models.SceneModel[scheneId]; exists == true {
		payload.Params.SceneId = scheneId
	}
	if response, err = connection.SendUdpMessage(bulbIp, payload); err != nil {
		log.Fatalf(`Unable to send message to udp: %s`, err)
		return nil, err
	}
	time.Sleep(time.Second * 2)
	return response, nil
}

func SetColorWarmWhite(bulbIp string, value int64) (*models.ResponsePayload, error) {
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
	time.Sleep(time.Second * 2)
	return response, nil
}

func SetColorColdWhite(bulbIp string, value int64) (*models.ResponsePayload, error) {
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
	time.Sleep(time.Second * 2)
	return response, nil
}
