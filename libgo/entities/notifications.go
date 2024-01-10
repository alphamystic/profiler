package entities

import (
  "github.com/alphamystic/profiler/libgo/utils"
)

type Notifications struct {
  Read bool
  NotificationID string
  Title string
  FromID string
  ToID string
  Description string
  utils.TimeStamps
}
