package model

import "time"

// Audio --
type Audio struct {
	Name       string    `json:"name"`
	Data       []byte    `json:"-"`
	UploadedAt time.Time `json:"uploadedAt"`
	Hash       string    `json:"hash"`
	Regions    Regions   `json:"regions"`
}

// Region is a slice of an audio for use with wavesurfer
type Region struct {
	Start     float64   `json:"start"`
	End       float64   `json:"end"`
	Notes     string    `json:"notes"`
	TimeStamp time.Time `json:"time"`
}

// Regions --
type Regions []*Region
