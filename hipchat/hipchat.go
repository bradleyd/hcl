package hipchat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
        "bytes"
        "net/url"
)

type Room struct {
	Room_id int    `json:"room_id"`
	Name    string `json:"name"`
        Topic   string `json:"topic"`
}

type Hipchat struct {
	token string
	url   string
}

type HipchatFeed struct {
	Rooms []Room `json:"rooms"`
}

type MessageResponse struct {
  Status string `json:"status"`
}
func New(token string) *Hipchat {
	return &Hipchat{token, "https://api.hipchat.com/v1"}
}
func postMessage(path, room, message string) ([]byte, error) {
        data := url.Values{}
        data.Set("room_id", room)
        data.Add("message", message)
        data.Add("from", "hcl")
        data.Add("notify", "1")
        data.Add("color", "green")

	req, err := http.NewRequest("POST", path, bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println("postMessage ERROR: ", err)
	}

        // for v2 only
        //authHeader := fmt.Sprintf("Bearer %s", token)
        //req.Header.Add("Authentication", authHeader)
        req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client do ERROR: ", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read body ERROR: ", err)
	}
	return body, nil
}


func getContent(path string) ([]byte, error) {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		fmt.Println("getContent ERROR: ", err)
	}

        // for v2 only
        //authHeader := fmt.Sprintf("Bearer %s", token)
        //req.Header.Add("Authentication", authHeader)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client do ERROR: ", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read body ERROR: ", err)
	}
	return body, nil
}

func (h Hipchat) listRooms() (error) {
	var hf HipchatFeed
	queryParams := fmt.Sprintf("auth_token=%s", h.token)
	url := h.url + "/rooms/list?" + queryParams
	fmt.Println(url)
	content, err := getContent(url)
	err = json.Unmarshal(content, &hf)
	for _, item := range hf.Rooms {
          fmt.Printf("Name: %v Topic: %v\n", item.Name, item.Topic)
	}
	if err != nil {
		return err
	}
	return nil
}

func (h Hipchat) messageRoom(room,message,from string) error {
        var mr MessageResponse
	queryParams := fmt.Sprintf("auth_token=%s", h.token)
	url := h.url + "/rooms/message?" + queryParams
	fmt.Println(url)
	content, err := postMessage(url, room, message)
	err = json.Unmarshal(content, &mr)

        if err !=nil {
          return err
        }

        fmt.Printf("%s\n", mr)
        return nil
}
