package gcm

const (
	MessagePriorityNormal = "normal"
	MessagePriorityHigh   = "high"
)

// Message is used by the application server to send a message to
// the GCM server. See the documentation for GCM Architectural
// Overview for more information:

//https://developers.google.com/cloud-messaging/http-server-ref
type Message struct {
	//Targets
	To              string   `json:"to,omitempty"`
	RegistrationIDs []string `json:"registration_ids"`

	//Options
	Priority              string `json:"priority,omitempty"`
	CollapseKey           string `json:"collapse_key,omitempty"`
	DelayWhileIdle        bool   `json:"delay_while_idle,omitempty"`
	TimeToLive            int    `json:"time_to_live,omitempty"`
	RestrictedPackageName string `json:"restricted_package_name,omitempty"`
	DryRun                bool   `json:"dry_run,omitempty"`

	//Payload
	//TODO add notification before opensourcing
	Data         map[string]interface{} `json:"data,omitempty"`
	Notification Notification           `json:"notification,omitempty"`

	//Private field to be able to add extra key values
	//i.e To add notification or thread info for logging
	extra map[string]interface{}
}

func (msg *Message) SetExtra(extra map[string]interface{}) {
	msg.extra = extra
}

func (msg *Message) Extra() map[string]interface{} {
	return msg.extra
}

//https://developers.google.com/cloud-messaging/http-server-ref#table2
type Notification struct {
	Title                 string `json:"title,omitempty"`          //Android (Required), IOs (Optional)
	Body                  string `json:"body,omitempty"`           //Android (Optional), IOs (Optional)
	Icon                  string `json:"icon,omitempty"`           //Android (Required)
	Sound                 string `json:"sound,omitempty"`          //Android (Optional), IOs (Optional)
	Badge                 string `json:"badge,omitempty"`          //IOs (Optional)
	Tag                   string `json:"tag,omitempty"`            //Android (Optional)
	Color                 string `json:"color,omitempty"`          //Android (Optional)
	ClickAction           string `json:"click_action,omitempty"`   //Android (Optional), IOs(Optional)
	BodyLocalizationKey   string `json:"body_loc_key,omitempty"`   //Android (Optional), IOs(Optional)
	BodyLocalizationArgs  string `json:"body_loc_args,omitempty"`  //Android (Optional), IOs(Optional)
	TitleLocalizationKey  string `json:"title_loc_key,omitempty"`  //Android (Optional), IOs(Optional)
	TitleLocalizationArgs string `json:"title_loc_args,omitempty"` //Android (Optional), IOs(Optional)
}

// NewMessage returns a new Message with the specified payload, expiration, priority
// and registration IDs.
func NewMessage(registrationIDs []string, data map[string]interface{}, priority string, timeToLive int) *Message {
	return &Message{
		RegistrationIDs: registrationIDs,

		Priority:   priority,
		TimeToLive: timeToLive,

		Data: data}
}
