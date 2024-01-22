package entities

import (
  "github.com/alphamystic/profiler/libgo/utils"
)
type PhishingLinks struct {
  Confirmed bool
  PLID string
  Link string
  Domain string
  RedirectDomain string // refers to the end domain
  PLIPAddress string
  AssociateUrls string // store this as a token converted from/to csv
  Target string
  utils.TimeStamps
}
