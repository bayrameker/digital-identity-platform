package models

type Permission struct {
    UserID     string `json:"user_id"`
    DataID     string `json:"data_id"`
    Permission string `json:"permission"` // "read", "write", "delete", "owner"
}
