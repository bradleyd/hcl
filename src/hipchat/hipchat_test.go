package hipchat

import (
  "testing" 
  "os"
)

func TestRooms(t *testing.T) {
  tken := os.Getenv("TOKEN")
  hc := New(tken)
  err := hc.ListRooms()
  if err != nil {
    t.Error(err)
  }
}

func TestRoomMessage(t *testing.T) {
  tken := os.Getenv("TOKEN")
  hc := New(tken)
  msg := "hello world!"
  room := "test"
  from := os.Getenv("USER")
  res, err := hc.MessageRoom(room, msg, from)
  if err !=nil {
    t.Error(err)
  }

  if res == "" {
    t.Errorf("Expected sent, but got %v", res)
  }
}
