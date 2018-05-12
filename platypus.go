package platypus

import (
	"flag"
	"fmt"
)

var platypus = []byte{0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x3a, 0x3a, 0x27, 0x27, 0x27, 0x2c, 0x3a, 0x64, 0x4f, 0x78, 0x3a, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x6c, 0x58, 0x78, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x78, 0x57, 0x4e, 0x6c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x63, 0x3a, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x6b, 0x57, 0x4b, 0x3a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x6f, 0x4b, 0x6b, 0x2c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x6c, 0x64, 0x3b, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3b, 0x3a, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x3a, 0x6c, 0x3b, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x27, 0x64, 0x4f, 0x6f, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x3b, 0x6f, 0x6f, 0x6c, 0x3a, 0x3b, 0x2c, 0x2c, 0x2c, 0x2c, 0x2c, 0x3b, 0x63, 0x6f, 0x6f, 0x63, 0x27, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x27, 0x3a, 0x6f, 0x4b, 0x4e, 0x4f, 0x3b, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x27, 0x64, 0x58, 0x58, 0x6b, 0x3b, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2c, 0x4f, 0x57, 0x57, 0x57, 0x4f, 0x27, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x6f, 0x4e, 0x4d, 0x58, 0x6c, 0x2e, 0x27, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x27, 0x27, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x6f, 0x57, 0x4d, 0x4d, 0x57, 0x64, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3b, 0x4b, 0x4d, 0x4d, 0x30, 0x6c, 0x6c, 0x6f, 0x63, 0x2c, 0x27, 0x27, 0x2e, 0x27, 0x6c, 0x6b, 0x3b, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x63, 0x58, 0x4d, 0x4d, 0x57, 0x6b, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x63, 0x4e, 0x4d, 0x57, 0x78, 0x2e, 0x20, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x3b, 0x4f, 0x63, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x63, 0x6b, 0x4b, 0x57, 0x4e, 0x6f, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3b, 0x30, 0x57, 0x58, 0x64, 0x2e, 0x2e, 0x63, 0x6c, 0x63, 0x63, 0x63, 0x63, 0x3a, 0x3b, 0x6c, 0x4f, 0x63, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x2e, 0x20, 0x2e, 0x2e, 0x63, 0x78, 0x4f, 0x78, 0x63, 0x2c, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x27, 0x3a, 0x6f, 0x4f, 0x6b, 0x6f, 0x2c, 0x2e, 0x2e, 0x2c, 0x63, 0x2c, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x3a, 0x6b, 0x63, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3a, 0x4f, 0x64, 0x6f, 0x6f, 0x6f, 0x6f, 0x78, 0x4f, 0x4b, 0x58, 0x78, 0x2e, 0x20, 0x20, 0x20, 0x63, 0x4b, 0x58, 0x30, 0x6b, 0x64, 0x6f, 0x6c, 0x6f, 0x6b, 0x64, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2c, 0x6b, 0x4f, 0x4f, 0x4f, 0x4f, 0x58, 0x57, 0x57, 0x57, 0x4d, 0x4b, 0x6c, 0x27, 0x27, 0x3b, 0x6b, 0x57, 0x57, 0x4e, 0x4b, 0x4f, 0x4f, 0x4f, 0x4f, 0x4f, 0x6c, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x3b, 0x3a, 0x6c, 0x30, 0x57, 0x57, 0x4e, 0x4e, 0x58, 0x6b, 0x63, 0x3a, 0x2c, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3b, 0x58, 0x4d, 0x4d, 0x4d, 0x4f, 0x27, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x27, 0x3b, 0x3a, 0x64, 0x58, 0x4d, 0x4d, 0x4d, 0x4b, 0x6f, 0x3b, 0x2c, 0x2e, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x2c, 0x3a, 0x6c, 0x6b, 0x4f, 0x6b, 0x64, 0x63, 0x6f, 0x58, 0x4d, 0x4d, 0x4d, 0x30, 0x6f, 0x6c, 0x78, 0x4f, 0x4f, 0x78, 0x63, 0x27, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x6c, 0x30, 0x4e, 0x57, 0x4e, 0x4f, 0x3a, 0x2e, 0x20, 0x20, 0x2c, 0x30, 0x4d, 0x4d, 0x57, 0x6b, 0x2e, 0x20, 0x20, 0x2e, 0x6c, 0x4b, 0x57, 0x4e, 0x6b, 0x3a, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x63, 0x30, 0x57, 0x4d, 0x4e, 0x30, 0x6f, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x2c, 0x30, 0x4d, 0x4d, 0x57, 0x6b, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x27, 0x6b, 0x57, 0x4d, 0x57, 0x4f, 0x2c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x6f, 0x4e, 0x4d, 0x4d, 0x4e, 0x64, 0x27, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2c, 0x30, 0x4d, 0x4d, 0x57, 0x6b, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x27, 0x4f, 0x57, 0x4d, 0x4d, 0x4b, 0x3a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x63, 0x58, 0x4d, 0x4d, 0x4d, 0x30, 0x2c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2c, 0x30, 0x4d, 0x4d, 0x57, 0x6b, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x63, 0x4e, 0x4d, 0x4d, 0x4d, 0x30, 0x2c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x78, 0x4d, 0x4d, 0x4d, 0x57, 0x78, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2c, 0x30, 0x4d, 0x4d, 0x57, 0x6b, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2c, 0x30, 0x4d, 0x4d, 0x4d, 0x57, 0x6c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x6b, 0x4d, 0x4d, 0x4d, 0x57, 0x78, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2c, 0x30, 0x4d, 0x4d, 0x57, 0x6b, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x27, 0x30, 0x4d, 0x4d, 0x4d, 0x57, 0x6f, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x64, 0x57, 0x4d, 0x4d, 0x57, 0x6b, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2c, 0x30, 0x4d, 0x4d, 0x57, 0x6b, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3b, 0x4b, 0x4d, 0x4d, 0x4d, 0x4e, 0x63, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2c, 0x30, 0x4d, 0x4d, 0x4d, 0x58, 0x63, 0x2e, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2c, 0x30, 0x4d, 0x4d, 0x57, 0x6b, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x6f, 0x4e, 0x4d, 0x4d, 0x57, 0x78, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x2e, 0x27, 0x3a, 0x6c, 0x64, 0x78, 0x4f, 0x4f, 0x4f, 0x6b, 0x78, 0x64, 0x6f, 0x6f, 0x6c, 0x6c, 0x6c, 0x6f, 0x6f, 0x64, 0x78, 0x6b, 0x4f, 0x4f, 0x6b, 0x64, 0x63, 0x3b, 0x2e, 0x2e, 0x2c, 0x6c, 0x30, 0x57, 0x4d, 0x4d, 0x57, 0x58, 0x4b, 0x6b, 0x6f, 0x6f, 0x6f, 0x63, 0x2e, 0x2c, 0x30, 0x4d, 0x4d, 0x57, 0x6b, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3a, 0x4b, 0x4d, 0x4d, 0x4e, 0x78, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x2e, 0x20, 0xa, 0x2c, 0x64, 0x30, 0x4e, 0x4d, 0x4d, 0x4d, 0x4e, 0x6b, 0x63, 0x2c, 0x2e, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x2e, 0x2c, 0x6c, 0x4f, 0x4e, 0x57, 0x57, 0x58, 0x58, 0x4e, 0x4d, 0x4d, 0x57, 0x58, 0x58, 0x57, 0x4d, 0x4d, 0x4b, 0x6c, 0x2e, 0x2e, 0x20, 0x20, 0x2c, 0x30, 0x4d, 0x4d, 0x57, 0x6b, 0x2e, 0x20, 0x20, 0x20, 0x2e, 0x6c, 0x58, 0x4d, 0x57, 0x4b, 0x78, 0x6f, 0x63, 0x63, 0x63, 0x63, 0x63, 0x63, 0x63, 0x3a, 0x3a, 0x3a, 0x3a, 0x6c, 0x78, 0x78, 0x27, 0xa, 0x58, 0x4d, 0x4d, 0x4d, 0x4d, 0x4d, 0x4e, 0x6f, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2c, 0x6b, 0x57, 0x4d, 0x4d, 0x4d, 0x4d, 0x57, 0x4f, 0x3b, 0x27, 0x6c, 0x4f, 0x4b, 0x58, 0x58, 0x6b, 0x63, 0x2e, 0x2e, 0x3b, 0x4b, 0x4d, 0x4d, 0x4d, 0x6b, 0x2e, 0x2e, 0x27, 0x6c, 0x4f, 0x4e, 0x4b, 0x6b, 0x3a, 0x2e, 0x27, 0x6c, 0x64, 0x64, 0x64, 0x64, 0x64, 0x64, 0x64, 0x64, 0x64, 0x64, 0x78, 0x30, 0x4b, 0x4f, 0x63, 0xa, 0x30, 0x57, 0x4d, 0x4d, 0x4d, 0x4d, 0x57, 0x6b, 0x27, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x3a, 0x30, 0x57, 0x4d, 0x4d, 0x4d, 0x4d, 0x4d, 0x4e, 0x4f, 0x6c, 0x27, 0x2e, 0x2e, 0x27, 0x3a, 0x6f, 0x78, 0x64, 0x6f, 0x6b, 0x4e, 0x4d, 0x4d, 0x4d, 0x4b, 0x64, 0x6f, 0x78, 0x64, 0x6c, 0x3b, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x63, 0x64, 0x3b, 0xa, 0x2e, 0x63, 0x78, 0x30, 0x58, 0x4e, 0x57, 0x57, 0x58, 0x78, 0x6c, 0x3b, 0x27, 0x2e, 0x2e, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x2e, 0x2e, 0x27, 0x3b, 0x6c, 0x6b, 0x4e, 0x4d, 0x57, 0x4e, 0x30, 0x64, 0x6f, 0x4f, 0x4e, 0x4d, 0x57, 0x58, 0x78, 0x6c, 0x3b, 0x2e, 0x20, 0x20, 0x20, 0x2e, 0x63, 0x58, 0x4d, 0x4d, 0x4d, 0x4f, 0x2c, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x63, 0x4b, 0x30, 0xa, 0x20, 0x20, 0x20, 0x2e, 0x2e, 0x2c, 0x3b, 0x63, 0x6c, 0x6f, 0x64, 0x78, 0x64, 0x64, 0x6f, 0x6f, 0x6f, 0x6f, 0x6f, 0x6f, 0x6f, 0x6f, 0x64, 0x78, 0x78, 0x78, 0x64, 0x6f, 0x6c, 0x3a, 0x27, 0x2e, 0x20, 0x20, 0x2e, 0x3b, 0x64, 0x6b, 0x4f, 0x4f, 0x4f, 0x6b, 0x6c, 0x27, 0x2e, 0x20, 0x2e, 0x6c, 0x4e, 0x4d, 0x4d, 0x4d, 0x4b, 0x3a, 0x2e, 0x20, 0x2e, 0x6c, 0x78, 0x6b, 0x6b, 0x6b, 0x6b, 0x6b, 0x6b, 0x6b, 0x78, 0x6b, 0x6b, 0x6b, 0x6b, 0x6b, 0x6b, 0x6b, 0x6b, 0x6b, 0x6b, 0x4f, 0x6b, 0x63, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x3b, 0x6c, 0x6f, 0x6b, 0x58, 0x4d, 0x4d, 0x4d, 0x4d, 0x57, 0x4b, 0x78, 0x6f, 0x6c, 0x63, 0x3b, 0x3b, 0x3b, 0x3a, 0x3a, 0x3a, 0x3a, 0x3a, 0x3a, 0x3a, 0x3a, 0x3a, 0x3a, 0x3a, 0x3a, 0x3a, 0x3a, 0x3a, 0x3b, 0x27, 0x2e, 0x20, 0xa}

func init() {
	var p bool
	flag.BoolVar(&p, "platypus", false, "Why the platypus?")
	flag.Parse()
	if p {
		fmt.Println(string(platypus))
	}
}
