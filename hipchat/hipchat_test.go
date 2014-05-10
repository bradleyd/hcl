package hipchat

import (
  "testing" 
)

func TestRooms(t *testing.T) {
  tken := "12345"
  hc := New(tken)
  rooms := hc.getRooms()
  if len(rooms) == 0 {
    t.Error("Did not fetch rooms")
  }
}
