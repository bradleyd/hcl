package hipchat

import (
  "testing" 
  "os"
)

func TestRooms(t *testing.T) {
  tken := os.Getenv("TOKEN")
  hc := New(tken)
  err := hc.listRooms()
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
  err := hc.messageRoom(room, msg, from)
  if err !=nil {
    t.Error(err)
  }
}
