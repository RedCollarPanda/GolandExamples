package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

//easyjson:json
type BaseMessage struct {
	Message           string `json:"message"`
	Recipient         string `json:"recipient"`
	Sender            string `json:"sender"`
	MessageID         string `json:"messageId"`
	SenderDeviceID    string `json:"senderDeviceId"`
	RecipientDeviceID string `json:"recipientDeviceId"`
	GmtTimestamp      int    `json:"gmtTimestamp"`
	Type              int    `json:"type"`
}

//easyjson:json
type TextMessage struct {
	Message           string `json:"message"`
	Recipient         string `json:"recipient"`
	Sender            string `json:"sender"`
	MessageID         string `json:"messageId"`
	SenderDeviceID    string `json:"senderDeviceId"`
	RecipientDeviceID string `json:"recipientDeviceId"`
	GmtTimestamp      int    `json:"gmtTimestamp"`
	Type              int    `json:"type"`
}

//easyjson:json
type RecipientReadMessage struct {
	Recipient         string `json:"recipient"`
	Sender            string `json:"sender"`
	MessageID         string `json:"messageId"`
	SenderDeviceID    string `json:"senderDeviceId"`
	RecipientDeviceID string `json:"recipientDeviceId"`
	GmtTimestamp      int    `json:"gmtTimestamp"`
	Type              int    `json:"type"`
}

//easyjson:json
type GroupChatInvite struct {
	InvitedUser       string `json:"invitedUser"`
	Group             string `json:"group"`
	Recipient         string `json:"recipient"`
	Sender            string `json:"sender"`
	MessageID         string `json:"messageId"`
	SenderDeviceID    string `json:"senderDeviceId"`
	RecipientDeviceID string `json:"recipientDeviceId"`
	GmtTimestamp      int    `json:"gmtTimestamp"`
	Type              int    `json:"type"`
}

//easyjson:json
type SystemMessage struct {
	SystemInfo        string `json:"systemInfo"`
	Recipient         string `json:"recipient"`
	Sender            string `json:"sender"`
	MessageID         string `json:"messageId"`
	SenderDeviceID    string `json:"senderDeviceId"`
	RecipientDeviceID string `json:"recipientDeviceId"`
	GmtTimestamp      int    `json:"gmtTimestamp"`
	Type              int    `json:"type"`
}

//easyjson:json
type ServerRecipientAck struct {
	MessageID         string `json:"messageId"`
	SenderDeviceID    string `json:"senderDeviceId"`
	RecipientDeviceID string `json:"recipientDeviceId"`
	GmtTimestamp      int    `json:"gmtTimestamp"`
	Type              int    `json:"type"`
}

//easyjson:json
type SenderServerAck struct {
	MessageID         string `json:"messageId"`
	SenderDeviceID    string `json:"senderDeviceId"`
	RecipientDeviceID string `json:"recipientDeviceId"`
	GmtTimestamp      int    `json:"gmtTimestamp"`
	Type              int    `json:"type"`
}

//easyjson:json
type SenderRecipientAck struct {
	MessageID         string `json:"messageId"`
	SenderDeviceID    string `json:"senderDeviceId"`
	RecipientDeviceID string `json:"recipientDeviceId"`
	GmtTimestamp      int    `json:"gmtTimestamp"`
	Type              int    `json:"type"`
}

var (
	testJson = `[{"message":"testNessage","recipient":"pavel","sender":"kit","messageId":"14c229af-63c2-4647-81a7-9b2201b16748","senderDeviceId":"sdid","recipientDeviceId":"rdid","gmtTimestamp":0,"type":0}, {"recipient":"pavel","sender":"kit","messageId":"09404db7-c63f-4aac-9435-4714e624e000","senderDeviceId":"sdid","recipientDeviceId":"rdid","gmtTimestamp":0,"type":1}, {"invitedUser":"invUser","group":"TestGroup","recipient":"pavel","sender":"kit","messageId":"77bba15c-cdfa-4b86-b267-bb63821c504b","senderDeviceId":"sdid","recipientDeviceId":"rdid","gmtTimestamp":0,"type":2}, {"systemInfo":"systemInfo","recipient":"pavel","sender":"kit","messageId":"c4022a61-1c99-4dd6-8bee-124eb9ad2cb7","senderDeviceId":"sdid","recipientDeviceId":"rdid","gmtTimestamp":0,"type":3}, {"messageId":"14c229af-63c2-4647-81a7-9b2201b16748","senderDeviceId":"sdid","recipientDeviceId":"rdid","gmtTimestamp":0,"type":4}, {"messageId":"14c229af-63c2-4647-81a7-9b2201b16748","senderDeviceId":"sdis","recipientDeviceId":"rdid","gmtTimestamp":0,"type":5}, {"messageId":"14c229af-63c2-4647-81a7-9b2201b16748","senderDeviceId":"sdid","recipientDeviceId":"rdid","gmtTimestamp":0,"type":6}]`
)

func createKeyValuePairs(m map[string]interface{}) string {
	b := new(bytes.Buffer)
	size := len(m)
	b.WriteByte('{')
	for key, value := range m {

		if str, ok := value.(string); ok {
			fmt.Fprintf(b, "\"%s\":\"%s\"", key, str)
		} else if _, ok := value.(float64); ok {
			fmt.Fprintf(b, "\"%s\":%d", key, int(value.(float64)))
		}

		size = size - 1
		if size != 0 {
			b.WriteByte(',')
		}
	}
	b.WriteByte('}')

	return b.String()
}

func prepareMessageString(messagesArray string) string {

	str := strings.Replace(messagesArray, "}, {", "} !!! {", -1)
	str = strings.Replace(str, "[", "", -1)
	str = strings.Replace(str, "]", "", -1)
	return str
}

func easyJsonHandler() {

	finalStr := prepareMessageString(string(testJson))
	array := strings.Split(finalStr, "!!!")

	messages := make([]BaseMessage, 0, 10)
	servRecAcks := make([]ServerRecipientAck, 0, 10)
	senderServerAck := make([]SenderServerAck, 0, 10)
	senderRecipientAck := make([]SenderRecipientAck, 0, 10)

	for _, message := range array {

		if strings.Contains(message, `"type":0`) ||
			strings.Contains(message, `"type":1`) ||
			strings.Contains(message, `"type":2`) ||
			strings.Contains(message, `"type":3`) {

			t := BaseMessage{}
			result := t.UnmarshalJSON([]byte(message))
			if result != nil {

				fmt.Print("Error message parcing: " + message)
				continue
			}

			t.Message = message
			messages = append(messages, t)

		} else if strings.Contains(message, `"type":4`) {

			t := ServerRecipientAck{}
			result := t.UnmarshalJSON([]byte(message))
			if result != nil {

				fmt.Print("Error message parcing: " + message)
				continue
			}

			servRecAcks = append(servRecAcks, t)

		} else if strings.Contains(message, `"type":5`) {
			t := SenderServerAck{}
			result := t.UnmarshalJSON([]byte(message))
			if result != nil {

				fmt.Print("Error message parcing: " + message)
				continue
			}

			senderServerAck = append(senderServerAck, t)

		} else if strings.Contains(message, `"type":6`) {
			t := SenderRecipientAck{}
			result := t.UnmarshalJSON([]byte(message))
			if result != nil {

				fmt.Print("Error message parcing: " + message)
				continue
			}

			senderRecipientAck = append(senderRecipientAck, t)

		} else if strings.Contains(message, `"type":7`) {

		} else if strings.Contains(message, `"type":8`) {

		} else {

		}
	}

}

func jsonNativeHandler() {

	var raw []map[string]interface{}
	var s []string
	json.Unmarshal([]byte(testJson), &raw)
	for _, element := range raw {

		typeId := int(element["type"].(float64))
		switch typeId {
		case 0:
			s = append(s, createKeyValuePairs(element))
			break
		case 1:
			s = append(s, createKeyValuePairs(element))
			break
		case 2:
			s = append(s, createKeyValuePairs(element))
			break

		case 3:
			s = append(s, createKeyValuePairs(element))
			break

		case 4:
			s = append(s, createKeyValuePairs(element))
			break

		case 5:
			s = append(s, createKeyValuePairs(element))
			break

		case 6:
			s = append(s, createKeyValuePairs(element))
			break

		default:

			break

		}
	}

}

func main() {

	/*
		We need to parce json array into separated messages or create structures of json element  by its [type]

		We use :
			reflection 						- jsonNativeHandler
			easyJson with code generation   - easyJsonHandler

		go test -bench . -benchmem
		goos: linux
		goarch: amd64
		pkg: GolandExamples/easyJson
		BenchmarkMyJson-8       	  200000	     10246 ns/op	    9448 B/op	      44 allocs/op
		BenchmarkNativeJson-8   	   30000	     41697 ns/op	   15405 B/op	     414 allocs/op
		PASS
		ok  	GolandExamples/easyJson	3.965s


		aasyJson is way faster and uses less memory.
	*/
}
