package hipchat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Room struct {
	Room_id int    `json:"room_id"`
	Name    string `json:"name"`
}

type Hipchat struct {
	token string
	url   string
}

type Foo struct {
	Rooms []Room `json:"rooms"`
}

func New(token string) *Hipchat {
	return &Hipchat{token, "https://api.hipchat.com/v1"}
}

func getContent(path string) ([]byte, error) {
	// build url
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		fmt.Println("getContent ERROR: ", err)
	}

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

func (h Hipchat) getRooms() (error) {
	var f Foo
	queryParams := fmt.Sprintf("auth_token=%s", h.token)
	url := h.url + "/rooms/list?" + queryParams
	fmt.Println(url)
	content, err := getContent(url)
	err = json.Unmarshal(content, &f)
	for _, item := range f.Rooms {
		fmt.Printf("Name: %v\n", item.Name)
	}
	if err != nil {
		return err
	}
	return nil
}
