package inmem

import "murrman/audio-share/model"

// InMem is in memory storage for audio share
type InMem struct {
	DB map[string]*model.Audio
}

// PutAudio puts audio into database
func (inmem *InMem) PutAudio(key string, val *model.Audio) {
	inmem.DB[key] = val
}

func (inmem *InMem) PutAudioRegion(key string, regions model.Regions) {
	inmem.DB[key].Regions = regions
}

func (inmem *InMem) GetAll() map[string]*model.Audio {
	return inmem.DB
}

func (inmem *InMem) GetAudio(key string) *model.Audio {
	return inmem.DB[key]
}

func (inmem *InMem) GetAudioData(key string) []byte {
	return inmem.DB[key].Data
}

func (inmem *InMem) DeleteAudioRegion(key string, region model.Region) {
	newRegions := model.Regions{}
	for _, region := range inmem.DB[key].Regions {
		if region.TimeStamp != region.TimeStamp {
			newRegions = append(newRegions, region)
		}
	}
}

func New() *InMem {
	return &InMem{
		DB: make(map[string]*model.Audio),
	}
}
