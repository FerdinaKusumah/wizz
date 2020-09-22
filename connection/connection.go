package connection

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/FerdinaKusumah/wizz/models"
	"log"
	"net"
	"strings"
	"time"
)

const Port = "38899"

func SendUdpMessage(host string, message *models.RequestPayload) (*models.ResponsePayload, error) {
	var (
		err             error
		response        = make([]byte, 4096)
		responsePayload = new(models.ResponsePayload)
		remoteAddr      = new(net.UDPAddr)
		conn            = new(net.UDPConn)
		payload         []byte
	)
	time.Sleep(time.Second * 1)
	// doing connection to UDP
	if remoteAddr, err = net.ResolveUDPAddr("udp", fmt.Sprintf(`%s:%s`, host, Port)); err != nil {
		log.Fatalf(`Unable to resolve to udp: %s`, err)
		return nil, err
	}
	if conn, err = net.DialUDP("udp", nil, remoteAddr); err != nil {
		log.Fatalf(`Unable to dial up to udp: %s`, err)
		return nil, err
	}
	defer conn.Close()
	// marshall payload to json string
	if payload, err = json.Marshal(message); err != nil {
		log.Fatalf(`Unable to marshal payload: %s`, err)
	}
	payloadString := string(payload)
	fmt.Println(fmt.Sprintf(`Payload string: %s`, payloadString))
	// send payload to bulb
	if _, err = conn.Write(payload); err != nil {
		log.Fatalf(`Unable to send message to UDP: %s`, err)
		return nil, err
	}
	// read response from bulb
	if _, err = bufio.NewReader(conn).Read(response); err != nil {
		log.Fatalf(`Unable to read message from UDP: %s`, err)
	}
	result := []byte(strings.Trim(string(response), "\x00'"))
	// convert string result to struct again
	if err = json.Unmarshal(result, responsePayload); err != nil {
		log.Fatalf(`'Unable to unmarshall response: %s'`, err)
	}
	return responsePayload, nil
}
