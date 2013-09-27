package zmq

import (
	"encoding/json"
	"github.com/pebbe/zmq3"
	"log"
)

func RecvJson(soc *zmq3.Socket, v interface{}) error {
	rawData, err := soc.Recv(0)
	log.Printf("Received data [%s]\n", rawData)

	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(rawData), &v)
}

func SendJson(soc *zmq3.Socket, v interface{}) (int, error) {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return 0, err
	}

	log.Printf("Sending data [%s]\n", jsonData)
	return soc.Send(string(jsonData), 0)
}

func SendJsonNoReply(soc *zmq3.Socket, v interface{}) (int, error) {
	length, err := SendJson(soc, v)
	if err != nil {
		return 0, err
	}

	_, err2 := soc.Recv(0)
	if err2 != nil {
		return length, err2
	}

	return length, err
}
