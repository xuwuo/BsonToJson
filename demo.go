// Package util contains some common functions of GO_SPIDER project.
package main

import (
    "encoding/json"
    "regexp"
)


func timeCost(start time.Time){
    terminal:=time.Since(start)
    println(terminal)
}

func main() {

    destbody := `{tatalProperty:17,root:[{street:'张公桥街道',community:'社区居委会',aab010:'8adc81e65ecb03b2015effeb0e7925a0',aab019:'1',totalRelief:'',aad011:'',aad011ForMember:'}]}`
    reg := regexp.MustCompile(`(\w+)\s*:([^A-Za-z0-9_])`)
    destbody = reg.ReplaceAllString(destbody, `"$1":$2`)
    reg = regexp.MustCompile(`([:\[, \{])'`)
    destbody = reg.ReplaceAllString(destbody, `$1"`)
    reg = regexp.MustCompile(`'([:\], \}])`)
    destbody = reg.ReplaceAllString(destbody, `"$1`)
    reg = regexp.MustCompile(`([:\[,\{])(\w+)\s*:`)
    destbody = reg.ReplaceAllString(destbody, `$1"$2":`)
    reg = regexp.MustCompile(`(\d{2}):"(\d{2})":(\d{2})`)
    destbody = reg.ReplaceAllString(destbody, `$1:$2:$3`)

    b, _ := json.Marshal(destbody)
    println(string(b))
    
}


