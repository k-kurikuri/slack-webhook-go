package slack

const (
	typeMarkDown = "mrkdwn"
	typeDivider  = "divider"
	typeSection  = "section"
	typeImage    = "image"
)

// WebHookBlockKit is slack's incoming web hook request body format
// FYI: https://api.slack.com/tools/block-kit-builder
type WebHookBlockKit struct {
	Blocks `json:"blocks"`
}

// Blocks is block list
type Blocks []Block

// Block is web hook style
type Block struct {
	Type      string     `json:"type"`
	Text      *Text      `json:"text,omitempty"`
	Accessory *Accessory `json:"accessory,omitempty"`
}

// Text is web hook style
type Text struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
}

// Accessory is web hook style
type Accessory struct {
	Type     string `json:"type"`
	ImageURL string `json:"image_url"`
	AltText  string `json:"alt_text"`
}
