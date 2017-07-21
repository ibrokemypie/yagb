package internal

import (
	"encoding/json"
	"fmt"
	"time"
)

//Header is the initial slug for the i3bar
type Header struct {
	Version     int  `json:"version"`
	StopSignal  int  `json:"stop_signal,omitempty"`
	ContSignal  int  `json:"cont_signal,omitempty"`
	ClickEvents bool `json:"click_events,omitempty"`
}

//Block contains the fields each module can use
type Block struct {
	FullText            string `json:"full_text"`
	ShortText           string `json:"short_text,omitempty"`
	Color               string `json:"color,omitempty"`
	Background          string `json:"background,omitempty"`
	Border              string `json:"border,omitempty"`
	MinWidth            int    `json:"min_width,omitempty"`
	Align               string `json:"align,omitempty"`
	Name                string `json:"name"`
	Urgent              bool   `json:"urgent,omitempty"`
	Separator           bool   `json:"separator,omitempty"`
	SeparatorBlockWidth int    `json:"separator_block_width,omitempty"`
	Markup              string `json:"markup,omitempty"`
}

//Print outputs data in the i3bar streaming json format
func Print(blocks *[]Block) {
	h, _ := json.Marshal(header)
	fmt.Println(string(h))
	fmt.Println("[")
	for {
		j, _ := json.Marshal(blocks)
		fmt.Printf(string(j))
		fmt.Println(",")
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
}
