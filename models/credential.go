package models

import "time"

type Credential struct {
    ServiceName string    `json:"service-name"`
    User        string    `json:"user"`
    Password    string    `json:"password"`
    CreatedAt   time.Time `json:"created_at"`
}
