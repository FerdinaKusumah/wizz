# wizz
Phillips Wizz Local Connection, see main.go for detail example.

Install package ``go get github.com/FerdinaKusumah/wizz``

## Available methods
1. `getState` get current state bulb
2. `getConfig` get current config
3. `turnOnLight` turn on light bulb
4. `turnOffLight` turn off light bulb
5. `setColorTemp` set bulb color temperature
6. `setBrightness` set brightness light bulb
7. `setColorRGB` set RGB color light bulb
8. `setColorScene` set color scheme based on available scheme
9. `setColorWarmWhite` set color warm light bulb
10. `setColorColdWhite` set color cold light bulb

Reference project from [https://github.com/sbidy/wiz_light](https://github.com/sbidy/wiz_light)

## Example
```go
package main

import (
	"encoding/json"
	"fmt"
	phillipWhizz "github.com/FerdinaKusumah/wizz"
	responseModel "github.com/FerdinaKusumah/wizz/models"
	"log"
)

func main() {
	var (
		bulbIp   = "192.168.1.10"
		response = new(responseModel.ResponsePayload)
		result   []byte
		err      error
	)
	//--------- Example Get Bulb Config -----------
	if response, err = phillipWhizz.GetConfig(bulbIp); err != nil {
		log.Fatalf(`Unable to read response: %s`, err)
	}
	if result, err = json.Marshal(response); err != nil {
		log.Fatalf(`Unable to convert to json string: %s`, err)
	}
	fmt.Println(string(result))

	// --------- Example Turn Off light -----------
	if response, err = phillipWhizz.TurnOffLight(bulbIp); err != nil {
		log.Fatalf(`Unable to read response: %s`, err)
	}
	if result, err = json.Marshal(response); err != nil {
		log.Fatalf(`Unable to convert to json string: %s`, err)
	}
	fmt.Println(string(result))
	...
}
```
