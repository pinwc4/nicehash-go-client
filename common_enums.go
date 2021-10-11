package nhclient

import "encoding/json"

type algorithm uint

const (
	AlgorithmScrypt algorithm = iota
	AlgorithmSHA256
	AlgorithmX11
	AlgorithmX13
	AlgorithmKeccak
	AlgorithmNeoscrypt
	AlgorithmQubit
	AlgorithmQuark
	AlgorithmLyra2Rev2
	AlgorithmDaggerHashimoto
	AlgorithmDecred
	AlgorithmLBRY
	AlgorithmEquihash
	AlgorithmBlake2s
	AlgorithmLyra2Z
	AlgorithmX16R
	AlgorithmSHA256AsicBoost
	AlgorithmZHash
	AlgorithmGrincuckatoo31
	AlgorithmLyra2Rev3
	AlgorithmCryptonightr
	AlgorithmCuckoocycle
	AlgorithmX16RV2
	AlgorithmRandomXMonero
	AlgorithmEagleSong
	AlgorithmGrincuckatoo32
	AlgorithmHandshake
	AlgorithmKawpow
	AlgorithmBeamV3
	AlgorithmOctopus
	AlgorithmAutolykos
)

func (e algorithm) String() string {
	return map[algorithm]string{
		AlgorithmScrypt:          "SCRYPT",
		AlgorithmSHA256:          "SHA256",
		AlgorithmX11:             "X11",
		AlgorithmX13:             "X13",
		AlgorithmKeccak:          "KECCAK",
		AlgorithmNeoscrypt:       "NEOSCRYPT",
		AlgorithmQubit:           "QUBIT",
		AlgorithmQuark:           "QUARK",
		AlgorithmLyra2Rev2:       "LYRA2REV2",
		AlgorithmDaggerHashimoto: "DAGGERHASHIMOTO",
		AlgorithmDecred:          "DECRED",
		AlgorithmLBRY:            "LBRY",
		AlgorithmEquihash:        "EQUIHASH",
		AlgorithmBlake2s:         "BLAKE2S",
		AlgorithmLyra2Z:          "LYRA2Z",
		AlgorithmX16R:            "X16R",
		AlgorithmSHA256AsicBoost: "SHA256ASICBOOST",
		AlgorithmZHash:           "ZHASH",
		AlgorithmGrincuckatoo31:  "GRINCUCKATOO31",
		AlgorithmLyra2Rev3:       "LYRA2REV3",
		AlgorithmCryptonightr:    "CRYPTONIGHTR",
		AlgorithmCuckoocycle:     "CUCKOOCYCLE",
		AlgorithmX16RV2:          "X16RV2",
		AlgorithmRandomXMonero:   "RANDOMXMONERO",
		AlgorithmEagleSong:       "EAGLESONG",
		AlgorithmGrincuckatoo32:  "GRINCUCKATOO32",
		AlgorithmHandshake:       "HANDSHAKE",
		AlgorithmKawpow:          "KAWPOW",
		AlgorithmBeamV3:          "BEAMV3",
		AlgorithmOctopus:         "OCTOPUS",
		AlgorithmAutolykos:       "AUTOLYKOS",
	}[e]
}

func (e *algorithm) FromString(es string) algorithm {
	return map[string]algorithm{
		"SCRYPT":          AlgorithmScrypt,
		"SHA256":          AlgorithmSHA256,
		"X11":             AlgorithmX11,
		"X13":             AlgorithmX13,
		"KECCAK":          AlgorithmKeccak,
		"NEOSCRYPT":       AlgorithmNeoscrypt,
		"QUBIT":           AlgorithmQubit,
		"QUARK":           AlgorithmQuark,
		"LYRA2REV2":       AlgorithmLyra2Rev2,
		"DAGGERHASHIMOTO": AlgorithmDaggerHashimoto,
		"DECRED":          AlgorithmDecred,
		"LBRY":            AlgorithmLBRY,
		"EQUIHASH":        AlgorithmEquihash,
		"BLAKE2S":         AlgorithmBlake2s,
		"LYRA2Z":          AlgorithmLyra2Z,
		"X16R":            AlgorithmX16R,
		"SHA256ASICBOOST": AlgorithmSHA256AsicBoost,
		"ZHASH":           AlgorithmZHash,
		"GRINCUCKATOO31":  AlgorithmGrincuckatoo31,
		"LYRA2REV3":       AlgorithmLyra2Rev3,
		"CRYPTONIGHTR":    AlgorithmCryptonightr,
		"CUCKOOCYCLE":     AlgorithmCuckoocycle,
		"X16RV2":          AlgorithmX16RV2,
		"RANDOMXMONERO":   AlgorithmRandomXMonero,
		"EAGLESONG":       AlgorithmEagleSong,
		"GRINCUCKATOO32":  AlgorithmGrincuckatoo32,
		"HANDSHAKE":       AlgorithmHandshake,
		"KAWPOW":          AlgorithmKawpow,
		"BEAMV3":          AlgorithmBeamV3,
		"OCTOPUS":         AlgorithmOctopus,
		"AUTOLYKOS":       AlgorithmAutolykos,
	}[es]
}

func (e algorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

func (e *algorithm) UnmarshalJSON(bytes []byte) error {
	var s string
	err := json.Unmarshal(bytes, &s)
	if err != nil {
		return err
	}
	*e = e.FromString(s)
	return nil
}
