package hipchat

import (
  "net/http" 
  "io/ioutil"

)

type Hipchat struct {
  token string
  url string
}


func New(token string) *Hipchat {
  return &Hipchat{token}
}

func getContent(path string) ([]byte, error) {
  // build url 

  req, err := http.NewRequest("GET", h.url, nil)
  if err != nil {
   //
  }
  
  client := &http.Client{}
  resp, err := client.Do(req)
  
  if err != nil {
    //
  }
  
  defer resp.Body.Close()
  
  body, err := ioutil.ReadAll(resp.Body)
  
  if err != nil {
    //
  }

  return body, nil

}

func (h Hipchat) getRooms() []string {

}
