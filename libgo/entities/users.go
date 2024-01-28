package entities

import (
  "github.com/alphamystic/profiler/libgo/utils"
)

type UserData struct {
  UserID string
  Role string
  Hash string
  Name string
  Email string
  Password string
  Verified bool
  Userhash string
  utils.TimeStamps
}


func CreateUser() error {
  return nil
}

func ListUsers(verified bool)([]ent.UserData,error) {
  return nil,nil
}

func GetUser(uid string) (*ent.UserData,error) {
  return nil,nil
}

func VerifyUser(uid,admId string) error {
  return nil
}

func Authenticate(email,password string)(*ent.UserData,error) {
  return nil,nil
}
