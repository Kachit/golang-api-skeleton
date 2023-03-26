package infrastructure

import (
	"github.com/speps/go-hashids/v2"
)

type HashIds struct {
	*hashids.HashID
}

func NewHashIds(cfg *Config) *HashIds {
	return &HashIds{NewHashID(cfg.HashIds.Salt, cfg.HashIds.Length)}
}

func (h HashIds) EncodeUint64(num uint64) (string, error) {
	n := int64(num)
	return h.HashID.EncodeInt64([]int64{n})
}

func (h HashIds) DecodeUint64(hash string) (uint64, error) {
	nums, err := h.HashID.DecodeInt64WithError(hash)
	if err != nil {
		return 0, err
	}
	return uint64(nums[0]), nil
}

func NewHashID(salt string, len int) *hashids.HashID {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = len
	h, _ := hashids.NewWithData(hd)
	return h
}
