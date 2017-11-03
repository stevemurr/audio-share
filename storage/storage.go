package storage

import "murrman/audio-share/model"

// Service is the storage interface
type Service interface {
	PutAudio(key string, val *model.Audio)
	PutAudioRegion(key string, regions model.Regions)
	DeleteAudioRegion(key string, region model.Region)
	GetAll() map[string]*model.Audio
	GetAudio(key string) *model.Audio
	GetAudioData(key string) []byte
}
