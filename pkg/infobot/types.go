package infobot

import (
	"time"
)

var (
	DefaultGreetingKey     = "greeting"
	DefaultGreetingChannel = ""
)

// TeamSettings holds the settings for a team
type TeamSettings struct {
	NonAdminCreate bool `json:"non_admin_create"`
	NonAdminEdit   bool `json:"non_admin_edit"`
	NonAdminDelete bool `json:"non_admin_delete"`
	//GreetingEnabled bool   `json:"greeting_enabled"`
	//GreetingChannel string `json:"greeting_channel"`
	//GreetingKey     string `json:"greeting_key"`
	revision int
}

// NewTeamSettings returns a new TeamSettings struct
func NewTeamSettings() *TeamSettings {
	return &TeamSettings{
		NonAdminCreate: false,
		NonAdminEdit:   false,
		NonAdminDelete: false,
		//GreetingEnabled: false,
		//GreetingChannel: DefaultGreetingChannel,
		//GreetingKey:     DefaultGreetingKey,
		revision: 0,
	}
}

// Info holds the information related to a particular key/value pair
type Info struct {
	Key         string   `json:"key"`
	Value       string   `json:"value"`
	CreatedBy   string   `json:"created_by"`
	CreatedTime int64    `json:"created_time"`
	Actions     []Action `json:"actions"`
	revision    int
}

// NewInfo returns a new Info struct
func NewInfo(Key, Value, CreatedBy string) *Info {
	var CreatedTime = time.Now().UTC().Unix()

	return &Info{
		Key:         Key,
		Value:       Value,
		CreatedBy:   CreatedBy,
		CreatedTime: CreatedTime,
		Actions: []Action{
			{
				ByUser:     CreatedBy,
				ActionType: ActionCreate.String(),
				Timestamp:  CreatedTime,
				NewValue:   Value,
			},
		},
		revision: 0,
	}
}

// Action holds one action record for an Info
type Action struct {
	ByUser     string `json:"by_user"`
	ActionType string `json:"action_type"`
	Timestamp  int64  `json:"timestamp"`
	NewValue   string `json:"new_value"`
}

// NewAction returns a new Action struct
func NewAction(ByUser string, Type ActionType, NewValue string) *Action {
	var Timestamp = time.Now().UTC().Unix()

	return &Action{
		ByUser:     ByUser,
		ActionType: Type.String(),
		Timestamp:  Timestamp,
		NewValue:   NewValue,
	}
}

type ActionType int

const (
	ActionUnknown ActionType = iota
	ActionCreate
	ActionEdit
)

var ActionTypeStringMap = map[ActionType]string{
	ActionUnknown: "unknown",
	ActionCreate:  "create",
	ActionEdit:    "edit",
}

var ActionTypeStringRevMap = map[string]ActionType{
	"unknown": ActionUnknown,
	"create":  ActionCreate,
	"edit":    ActionEdit,
}

func (a ActionType) String() string {
	return ActionTypeStringMap[a]
}
