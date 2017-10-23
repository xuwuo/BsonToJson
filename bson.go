// Copyright 2017-10-23 The Go BsonToJson Project
// Use of this source code is governed by the license for
// The Go BsonToJson Project, found in the LICENSE file.
//
// Authors:
//   2017-10-23 xuwuo <xuwuo@126.com>

package main

import (
    "time"
)

func timeCost(start time.Time){
    terminal:=time.Since(start)
    println(terminal)
}

func main() {
    defer timeCost(time.Now())
    destbody := `{tatalProperty:17,root:[{street:'张公桥街道',community:'王浩儿社区居委会',aab010:'8adc81e65ecb03b2015effeb0e7925a0',aab019:'1',totalRelief:'469.0'}]}`
    // destbody := `{"root":[{"aab010":"8adc81e65ecb03b2015effeb0e7925a0","aab019":"1","aac000":"8adc81e65ecb03b2015effeb0eaf25a1"}],"tatalProperty":17.01}`
    raw := []byte(destbody)
    var parsed map[string]interface{}
    if err := bson.Unmarshal(raw, &parsed); err != nil {
        panic(err)
    }
    

}