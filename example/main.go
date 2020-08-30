package example

import (
	"encoding/json"
	"fmt"
	"github.com/FerdinaKusumah/wizz/handlers"
	"github.com/FerdinaKusumah/wizz/models"
	"log"
)

func example() {
	var (
		bulbIp   = "192.168.1.10"
		response = new(models.ResponsePayload)
		result   []byte
		err      error
	)
	//--------- Example Get Bulb Config -----------
	if response, err = handlers.GetConfig(bulbIp); err != nil {
		log.Fatalf(`Unable to read response: %s`, err)
	}
	if result, err = json.Marshal(response); err != nil {
		log.Fatalf(`Unable to convert to json string: %s`, err)
	}
	fmt.Println(string(result))

	// --------- Example Turn Off light -----------
	//if response, err = handlers.TurnOffLight(bulbIp); err != nil {
	//	log.Fatalf(`Unable to read response: %s`, err)
	//}
	//if result, err = json.Marshal(response); err != nil {
	//	log.Fatalf(`Unable to convert to json string: %s`, err)
	//}
	//fmt.Println(string(result))

	// --------- Example Turn On light -----------
	//if response, err = handlers.TurnOnLight(bulbIp); err != nil {
	//	log.Fatalf(`Unable to read response: %s`, err)
	//}
	//if result, err = json.Marshal(response); err != nil {
	//	log.Fatalf(`Unable to convert to json string: %s`, err)
	//}
	//fmt.Println(string(result))

	// --------- Example Set Brightness -----------
	//if response, err = handlers.SetBrightness(bulbIp, 255); err != nil {
	//	log.Fatalf(`Unable to read response: %s`, err)
	//}
	//if result, err = json.Marshal(response); err != nil {
	//	log.Fatalf(`Unable to convert to json string: %s`, err)
	//}
	//fmt.Println(string(result))

	// --------- Example Set Color Temperature -----------
	//if response, err = handlers.SetColorTemp(bulbIp, 3000); err != nil {
	//	log.Fatalf(`Unable to read response: %s`, err)
	//}
	//if result, err = json.Marshal(response); err != nil {
	//	log.Fatalf(`Unable to convert to json string: %s`, err)
	//}
	//fmt.Println(string(result))

	// --------- Example Set RGB Color -----------
	//if response, err = handlers.SetColorRGB(bulbIp, 153, 153, 255); err != nil {
	//	log.Fatalf(`Unable to read response: %s`, err)
	//}
	//if result, err = json.Marshal(response); err != nil {
	//	log.Fatalf(`Unable to convert to json string: %s`, err)
	//}
	//fmt.Println(string(result))

	// --------- Example Set Scene Color -----------
	//if response, err = handlers.SetColorScene(bulbIp, 12); err != nil {
	//	log.Fatalf(`Unable to read response: %s`, err)
	//}
	//if result, err = json.Marshal(response); err != nil {
	//	log.Fatalf(`Unable to convert to json string: %s`, err)
	//}
	//fmt.Println(string(result))

	// --------- Example Set Warm White -----------
	//if response, err = handlers.SetColorWarmWhite(bulbIp, 12); err != nil {
	//	log.Fatalf(`Unable to read response: %s`, err)
	//}
	//if result, err = json.Marshal(response); err != nil {
	//	log.Fatalf(`Unable to convert to json string: %s`, err)
	//}
	//fmt.Println(string(result))

	// --------- Example Set Cold White -----------
	//if response, err = setColorColdWhite(bulbIp, 12); err != nil {
	//	log.Fatalf(`Unable to read response: %s`, err)
	//}
	//if result, err = json.Marshal(response); err != nil {
	//	log.Fatalf(`Unable to convert to json string: %s`, err)
	//}
	//fmt.Println(string(result))
}
