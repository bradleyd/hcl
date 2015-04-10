package hipchat

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "bytes"
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

type MessagePayload struct {
  Color string `json:"color"`
  Message string `json:"message"`
}

func New(token string) *Hipchat {
  return &Hipchat{token, "https://api.hipchat.com/v2"}
}

//func request(method string, url string, data []byte) ([]byte, error) {
  //var response []byte
  //h := &http.Client{}

  //body := bytes.NewBuffer(data)
  //req, err := http.NewRequest(method, url, body)
  //if err != nil {
    //return response, err
  //}

   ////for v2 only
  //authHeader := fmt.Sprintf("Bearer %s", token)
  //req.Header.Add("Authorization", authHeader)
  //req.Header.Add("Accept", "application/json")
  //req.Header.Add("Content-Type", "application/json")
  
  //res := &http.Response{}


  //defer res.Body.Close()

  //return response, err
//}
// post a message to a room
// todo create messagePayload from config
// takes POST url, room name, message, and auth token
func postMessage(path, room, message, token string) (string, error) {
  // create payload 
  json_data := &MessagePayload {
    Color: "green",
    Message: message,
  }
  json_bytes, err := json.Marshal(json_data)
  if nil != err {
    return "JSON marshall error", err
  }

  req, err := http.NewRequest("POST", path, bytes.NewReader(json_bytes))
  if err != nil {
    fmt.Println("postMessage ERROR: ", err)
  }

  // for v2 only
  authHeader := fmt.Sprintf("Bearer %s", token)
  req.Header.Add("Authorization", authHeader)
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Content-Type", "application/json")

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
  fmt.Println(body)
  return resp.Status, nil
}

// post a message to a user
// todo create messagePayload from config
// takes POST url, room name, message, and auth token
func postUserMessage(path, user, message, token string) (string, error) {
  // create payload 
  json_data := &MessagePayload {
    Color: "green",
    Message: message,
  }
  json_bytes, err := json.Marshal(json_data)
  if nil != err {
    return "JSON marshall error", err
  }

  req, err := http.NewRequest("POST", path, bytes.NewReader(json_bytes))
  if err != nil {
    fmt.Println("postUserMessage ERROR: ", err)
  }

  // for v2 only
  authHeader := fmt.Sprintf("Bearer %s", token)
  req.Header.Add("Authorization", authHeader)
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Content-Type", "application/json")

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
  fmt.Println(body)
  return resp.Status, nil
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

func (h Hipchat) ListRooms() (error) {
  var hf HipchatFeed
  queryParams := fmt.Sprintf("auth_token=%s", h.token)
  url := h.url + "/rooms/list?" + queryParams
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

func (h Hipchat) MessageRoom(room,message,from string) (string, error) {
  url := fmt.Sprintf("%s/room/%s/notification", h.url, room)
  content, err := postMessage(url, room, message, h.token)
  return content, err
}

func (h Hipchat) MessageUser(user, message, from string) (string, error) {
  fmt.Println("i am in func")
  url := fmt.Sprintf("%s/user/%s/message", h.url, user)
  content, err := postUserMessage(url, user, message, h.token)
  return content, err
}
