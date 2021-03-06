// Copyright (C)2018 by Lei Peng <pyp126@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
package core

import (
	"log"
	"math/rand"
	"time"
)

var r *rand.Rand // Rand for this package.

const DefaultTimeFormat string = "2006-01-02 15:04:05"
const DefaultDateFormat string = "2006-01-02"

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomString(strlen int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, strlen)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return string(result)
}

func UnixTime(str string, format ...string) int {
	realFormat := DefaultTimeFormat
	if len(format) > 0 {
		realFormat = format[0]
	}
	if str == "" {
		return 0
	}
	local, err := time.LoadLocation(Conf.TimeZone)
	if err != nil {
		if Conf.Debug {
			log.Println(err)
		}
		return 0
	}
	tm, err := time.ParseInLocation(realFormat, str, local)
	if err != nil {
		return 0
	}
	return int(tm.In(local).Unix())
}

func FromUnixTime(t int64, format ...string) string {
	realFormat := DefaultTimeFormat
	if len(format) > 0 {
		realFormat = format[0]
	}
	if t == 0 {
		return "1970-01-01 00:00:00"
	}
	local, err := time.LoadLocation(Conf.TimeZone)
	if err != nil {
		if Conf.Debug {
			log.Println(err)
		}
		return time.Unix(int64(t), 0).Format(realFormat)
	}
	return time.Unix(int64(t), 0).In(local).Format(realFormat)
}

func NowTime() string {
	return FromUnixTime(time.Now().Unix())
}
