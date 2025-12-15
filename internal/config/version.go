package config

import mapval "git.sr.ht/~madcapjake/rhi/internal/map"

// RuntimeVersion is the semantic version of this binary.
// Bump this whenever you add new Native functions or change the VM opcodes.
var RuntimeVersion = mapval.NewVersion(0, 4, 0, false)
