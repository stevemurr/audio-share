package ondisk

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"murrman/audio-share/model"
	"os"
	"path/filepath"
)

// OnDisk defines an on disk database
type OnDisk struct {
	Root string
}

// PutAudio writes an audio file to disk and creates a meta data file
func (on *OnDisk) PutAudio(key string, val *model.Audio) {
	// create on disk file name
	wav := filepath.Join(on.Root, key+".wav")
	meta := filepath.Join(on.Root, key+".json")
	// create the file
	f, err := os.Create(wav)
	if err != nil {
		panic(err)
	}
	// write to the file
	_, err = f.Write(val.Data)
	if err != nil {
		panic(err)
	}
	// close this file to reuse 'f'
	f.Close()
	// create the meta data file
	f, err = os.Create(meta)
	if err != nil {
		panic(err)
	}

	// write the meta data file
	if err := json.NewEncoder(f).Encode(val); err != nil {
		panic(err)
	}
	f.Close()
}

// PutAudioRegion --
func (on *OnDisk) PutAudioRegion(key string, regions model.Regions) {
	var a *model.Audio
	meta := filepath.Join(on.Root, key+".json")
	// open meta data file
	f, err := os.Open(meta)
	if err != nil {
		panic(err)
	}
	// decode json
	if err := json.NewDecoder(f).Decode(&a); err != nil {
		log.Println(err)
		panic(err)
	}
	f.Close()
	// update the regions
	a.Regions = regions
	os.Remove(meta)
	f, err = os.Create(meta)
	if err != nil {
		panic(err)
	}
	// write it to disk
	if err := json.NewEncoder(f).Encode(&a); err != nil {
		log.Println(err)
		panic(err)
	}
	f.Close()
}

// DeleteAudioRegion --
func (on *OnDisk) DeleteAudioRegion(key string, region model.Region) {
	var a *model.Audio
	meta := filepath.Join(on.Root, key+".json")
	// open meta data file
	f, err := os.Open(meta)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// decode json
	if err := json.NewDecoder(f).Decode(&a); err != nil {
		panic(err)
	}
	// update the regions
	newRegions := model.Regions{}
	for _, region := range a.Regions {
		if region.TimeStamp != region.TimeStamp {
			newRegions = append(newRegions, region)
		}
	}
	a.Regions = newRegions
	// write it to disk
	if err := json.NewEncoder(f).Encode(&a); err != nil {
		panic(err)
	}
}

// GetAll returns a map of hashes to audio objects
func (on *OnDisk) GetAll() map[string]*model.Audio {
	files, err := filepath.Glob(filepath.Join(on.Root, "*.json"))
	if err != nil {
		panic(err)
	}
	res := map[string]*model.Audio{}
	for _, file := range files {
		var a *model.Audio
		// open meta data file
		f, err := os.Open(file)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		// decode json
		if err := json.NewDecoder(f).Decode(&a); err != nil {
			panic(err)
		}
		res[a.Hash] = a
	}
	return res
}

// GetAudio --
func (on *OnDisk) GetAudio(key string) *model.Audio {
	file := filepath.Join(on.Root, key+".json")
	var a *model.Audio
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	// decode the meta data file
	if err := json.NewDecoder(f).Decode(&a); err != nil {
		panic(err)
	}
	return a
}

// GetAudioData --
func (on *OnDisk) GetAudioData(key string) []byte {
	file := filepath.Join(on.Root, key+".wav")
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return b
}

// New returns a new instance of the ondisk database
func New(root string) *OnDisk {
	// if root folder does not exist
	if _, err := os.Stat(root); err != nil {
		log.Println(err)
		// cant create the root dir
		if err := os.MkdirAll(root, 0755); err != nil {
			panic(err)
		}
	}
	return &OnDisk{
		Root: root,
	}
}
