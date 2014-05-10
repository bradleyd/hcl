package hipchat

import (
  "testing" 
  "os"
)

func TestRooms(t *testing.T) {
  tken := os.Getenv("TOKEN")
  hc := New(tken)
  err := hc.getRooms()
  if err != nil {
    t.Error(err)
  }

  //for _, item := range bar {
    //fmt.Println(item.Name)
  //}
}
