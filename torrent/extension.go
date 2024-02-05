package torrent

import (
	"bittorrent/peer-wire"
	"math"
	"math/rand"
)

var extensions = peer_wire.Extensions{
	peer_wire.ExtMetadataName: peer_wire.ExtMetadataID,
}

// prepare client's handshake msg for send
func extensionHandshakeMsg(metaSz int64) *peer_wire.Msg {
	return &peer_wire.Msg{
		Kind:       peer_wire.Extended,
		ExtendedID: 0,
		ExtendedMsg: struct {
			ExtMap peer_wire.Extensions `bencode:"m"`
			MetaSz int64                `bencode:"metadata_size" empty:"omit"`
		}{
			ExtMap: extensions,
			MetaSz: metaSz,
		},
	}
}

type freqMap map[int64]int

func newFreqMap() freqMap {
	return make(map[int64]int)
}

func (f freqMap) add(n int64) {
	var count int
	var ok bool
	if count, ok = f[n]; !ok {
		f[n] = 1
	} else {
		f[n] = count + 1
	}
}

func (f freqMap) remove(n int64) {
	if _, ok := f[n]; ok {
		f[n]--
		if f[n] == 0 {
			delete(f, n)
		}
	}
}

func (f freqMap) len() int {
	return len(f)
}

func (f freqMap) initKey(n int64) {
	if _, ok := f[n]; ok {
		return
	}
	f[n] = 0
}

func (f freqMap) max() (max int64) {
	var maxFreq int
	for k, v := range f {
		if v > maxFreq {
			maxFreq = v
			max = k
		}
	}
	return
}

func (f freqMap) min() (min int64) {
	var minFreq = math.MaxInt64
	for k, v := range f {
		if v < minFreq {
			minFreq = v
			min = k
		}
	}
	return
}

// pick a random key among the keys whose value = `val`
func (f freqMap) pickRandom(val int) int64 {
	equals := []int64{}
	for k, v := range f {
		if v == val {
			equals = append(equals, k)
		}
	}
	if len(equals) == 0 {
		panic("fmap: pickRandom")
	}
	return equals[rand.Intn(len(equals))]
}
func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
