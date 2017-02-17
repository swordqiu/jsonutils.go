package jsonutils

import (
    "testing"
)

func TestHexchar2num(t *testing.T) {
    cases := []struct {
        in, want byte
    } {
            {'F', 15},
            {'A', 10},
            {'0', 0},
            {'1', 1},
    }
    for _, c := range cases {
        got, _ := hexchar2num(c.in)
        if got != c.want {
            t.Errorf("Hexchar2num(%c) == %c, want %c", c.in, got, c.want)
        }
    }
    _, e := hexchar2num('G')
    if e == nil {
        t.Errorf("Hexchar2num(G) should raise error")
    }
}

func TestHexstr2byte(t *testing.T) {
    cases := []struct {
        in []byte
        want byte
    } {
            {[]byte{'F', 'F'}, 255},
            {[]byte{'0', '0'}, 0},
            {[]byte{'1', '0'}, 16},
    }
    for _, c := range cases {
        got, _ := hexstr2byte(c.in)
        if got != c.want {
            t.Errorf("hexstr2byte(%s) == %d, want %d", c.in, got, c.want)
        }
    }
}

func TestHexstr2rune(t *testing.T) {
    cases := []struct {
        in []byte 
        want rune
    } {
            {[]byte("00FF"), 255},
            {[]byte("0000"), 0},
            {[]byte("0010"), 16},
    }
    for _, c := range cases {
        got, _ := hexstr2rune(c.in)
        if got != c.want {
            t.Errorf("hexstr2rune(%s) == %d, want %d", c.in, got, c.want)
        }
    }
}

func TestReadString(t *testing.T) {
    cases := []struct {
        in []byte
        want string 
        want_quote bool
    } {
            {[]byte("\"00FF\""), "00FF", true},
            {[]byte("0"), "0", false},
            {[]byte("\"a\\nb\\n\""), "a\nb\n", true},
            {[]byte("123\n22"), "123", false},
            {[]byte("abc:"), "abc", false},
    }
    for _, c := range cases {
        got, quote, _, _ := parseString(c.in, 0)
        if got != c.want || quote != c.want_quote {
            t.Errorf("readString(%s) == %s %v, want %s %v", c.in, got, quote, c.want, c.want_quote)
        }
    }
}

func TestJSONParse(t *testing.T) {
    cases := []struct {
        in, out string
    } {
        {"{'name': '大家好'}", "{\"name\": \"\\xe5\\xa4\\xa7\\xe5\\xae\\xb6\\xe5\\xa5\\xbd\"}"},
        {"\"\\xe5\\xa5\\xbd\"", "\"好\""},
    }
    for _, c := range cases {
        got, _ := ParseString(c.in)
        got2, _ := ParseString(c.out)
        if got.String() != got2.String() {
            t.Errorf("JSONParse: %s(%s) != %s(%s)", c.in, got, c.out, got2)
        }
    }
}
