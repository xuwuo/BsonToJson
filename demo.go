// // Package util contains some common functions of GO_SPIDER project.
// package main

// import (
//     "encoding/json"
//     "time"
//     "regexp"
//     "github.com/xuwuo/BsonToJson"
//     "encoding/json"
// )


// func timeCost(start time.Time){
//     terminal:=time.Since(start)
//     println(terminal)
// }

// func demo1() {

//     destbody := `{tatalProperty:17,root:[{street:'张公桥街道',community:'社区居委会',aab010:'8adc81e65ecb03b2015effeb0e7925a0',aab019:'1',totalRelief:'',aad011:'',aad011ForMember:'}]}`
//     reg := regexp.MustCompile(`(\w+)\s*:([^A-Za-z0-9_])`)
//     destbody = reg.ReplaceAllString(destbody, `"$1":$2`)
//     reg = regexp.MustCompile(`([:\[, \{])'`)
//     destbody = reg.ReplaceAllString(destbody, `$1"`)
//     reg = regexp.MustCompile(`'([:\], \}])`)
//     destbody = reg.ReplaceAllString(destbody, `"$1`)
//     reg = regexp.MustCompile(`([:\[,\{])(\w+)\s*:`)
//     destbody = reg.ReplaceAllString(destbody, `$1"$2":`)
//     reg = regexp.MustCompile(`(\d{2}):"(\d{2})":(\d{2})`)
//     destbody = reg.ReplaceAllString(destbody, `$1:$2:$3`)

//     b, _ := json.Marshal(destbody)
//     println(string(b))
    
// }


// func demo2() {
//     defer timeCost(time.Now())
//     destbody := `{tatalProperty:17,root:[{street:'张公桥街道',community:'王浩儿社区居委会',aab010:'8adc81e65ecb03b2015effeb0e7925a0',aab019:'1',totalRelief:'469.0'}]}`
//     // destbody := `{"root":[{"aab010":"8adc81e65ecb03b2015effeb0e7925a0","aab019":"1","aac000":"8adc81e65ecb03b2015effeb0eaf25a1"}],"tatalProperty":17.01}`
//     raw := []byte(destbody)
//     var parsed map[string]interface{}
//     if err := bson.Unmarshal(raw, &parsed); err != nil {
//         panic(err)
//     }
    
//     b, _ := json.Marshal(parsed)
//     println("log:"+string(b))

// }

// func main() {
//     // demo1()
//     demo2()
// }

