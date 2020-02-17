package bridgesig

import (
	"github.com/pkg/errors"
)

const CBridgeSigSz = 65

func DecodeECDSASig(sig []byte) (
	v byte,
	r []byte,
	s []byte,
	err error,
) {
	if len(sig) != CBridgeSigSz {
		err = errors.New("wrong input")
		return
	}
	v = byte(sig[64] + 27)
	r = sig[:32]
	s = sig[32:64]
	return
}
