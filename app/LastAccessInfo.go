package main

import (
    "time"
)

type LastAccessInfo struct {
  UserName        string    `firestore:"user_name"`
  LastAccess      time.Time `firestore:"last_access"`
}

func InitLastAccessInfo(userName string, lastAccess time.Time)(info LastAccessInfo){
  info.UserName = userName
  info.LastAccess = lastAccess
  return
}
