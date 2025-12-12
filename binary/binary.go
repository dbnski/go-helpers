package binary

const hexChars = "0123456789ABCDEF"

func HexifyBytesToString(b []byte) string {
    if len(b) == 0 {
        return ""
    }

    var o []byte
    return string(HexifyBytes(o, b))
}

func HexifyBytes(o []byte, b []byte) []byte {
    if len(b) == 0 {
        o = o[:0]
        return o
    }

    size := len(b) * 5 - 1
    if cap(o) < size {
        o = make([]byte, size)
    }
    o = o[:size]

    j := 0
    for i, v := range b {
        o[j+0] = '0'
        o[j+1] = 'x'
        o[j+2] = hexChars[v>>4]
        o[j+3] = hexChars[v&0x0F]

        if i < size - 1 {
            o[j+4] = ' '
        }
        j += 5
    }

    return o
}

func HexifyToString(b []byte) string {
    if len(b) == 0 {
        return ""
    }

    var o []byte
    return string(Hexify(o, b))
}

func Hexify(o []byte, b []byte) []byte {
    if len(b) == 0 {
        o = o[:0]
        return o
    }

    size := len(b) * 2 + 2
    if cap(o) < size {
        o = make([]byte, size)
    }
    o = o[:size]

    o[0] = '0'
    o[1] = 'x'

    j := 2
    for _, v := range b {
        o[j+0] = hexChars[v>>4]
        o[j+1] = hexChars[v&0x0F]
        j += 2
    }

    return o
}