// Copyright 2017-10-23 The Go BsonToJson Project
// Use of this source code is governed by the license for
// The Go BsonToJson Project, found in the LICENSE file.
//
// Authors:
//   2017-10-23 xuwuo <xuwuo@126.com>

package BsonToJson

import (
    "bytes"
    "reflect"
    "errors"
    "strings"
)

const (
    MAP = iota
    LIST
    KEYEND
    NEXT
    STRING
    NUMBER
    BOOL
    NULL
)

type State struct {
    data   []byte
    i      int
    end      int
    v      interface{}
}

func Unmarshal(data []byte, v interface{}) error {
    s := &State{data, 0, len(data), v}
    s.v,_ = s.ReadMap()

    rv := reflect.ValueOf(v)
    if rv.Kind() != reflect.Ptr || rv.IsNil() {
        return errors.New("Need a pointer, got " + reflect.TypeOf(v).String())
    }
    for rv.Kind() == reflect.Ptr {
        rv = rv.Elem()
    }
    sv := reflect.ValueOf(s.v)
    for sv.Kind() == reflect.Ptr {
        sv = sv.Elem()
    }
    var (
        rvt = rv.Type()
        svt = sv.Type()
    )
    if !svt.AssignableTo(rvt) {
        if rv.Kind() != reflect.Slice && sv.Kind() != reflect.Slice {
            return errors.New("Cannot assign " + svt.String() + " to " + rvt.String())
        }
        if sv.Len() == 0 {
            return nil
        }
        var (
            mapi  map[string]interface{}
            mapt  = reflect.TypeOf(mapi)
            svte  = svt.Elem()
            rvte  = rvt.Elem()
            ismap bool
        )
        _, ismap = sv.Index(0).Interface().(map[string]interface{})
        if !(ismap && mapt.AssignableTo(rvte)) {
            return errors.New("Cannot assign " + svte.String() + " to " + rvte.String())
        }
        var (
            ssv = reflect.MakeSlice(rvt, sv.Len(), sv.Cap())
        )
        for i := 0; i < sv.Len(); i++ {
            v := sv.Index(i).Interface().(map[string]interface{})
            ssv.Index(i).Set(reflect.ValueOf(v))
        }
        sv = ssv
    }
    rv.Set(sv)
    return nil
}

func (s *State) ReadMap() (m map[string]interface{}, err error) {
    var key string
    m = make(map[string]interface{})
    for {  
        if (s.i+2) > s.end {break}

        if s.data[s.i] == '}' {
            // if s.data[s.i+1] == ',' && s.data[s.i+2] == '{' {
            //     return m, nil
            // }else if s.data[s.i+1] == ']' {
            //     s.i++
            // }else{
            //    s.i++
            // }
            return m, nil
        }
        if s.data[s.i] == '{' || (s.data[s.i] == ',' && s.data[s.i+1] != '{'){
                // println("====="+string(s.data[s.i])+string(s.data[s.i+1])+string(s.data[s.i+2]))
            key,_ = s.ReadKey()
        }

        s.i++
        if s.data[s.i] == ':'{
            if s.data[s.i+1] == '[' {
                s.i++
                m[key],_ = s.ReadList()

                if s.i+2 < s.end{//}],
                    continue
                }else if s.data[s.i] == ']' {// && s.data[s.i+1] == '}'{
                    return m,nil
                }
            // }else if s.data[s.i+1] == ']' {
            }else{
                m[key],_ = s.ReadValue()
            }
        }

    }

    return m, nil
}

func (s *State) ReadList() (list []interface{}, err error) {
    list = []interface{}{}
    for { 
        if (s.i+2) > s.end {break}
        if s.data[s.i] == ']' {
            // println(len(list)
            return list, nil
        }else if s.data[s.i] == '}' && s.data[s.i+1] == ']'{
            s.i++
            return list, nil
        // }else if s.data[s.i] == '[' {
        }else if s.data[s.i] == '}' && s.data[s.i+1] == ',' {
            s.i++
        }
        s.i++
        m,_ := s.ReadMap()
        list = append(list,m)
    }

    // println(len(list))
    return list, nil
}

func (s *State) ReadKey() (string , error) {
    var (
        c       byte
        buf     *bytes.Buffer
    )
    buf = new(bytes.Buffer)
    for {
        s.i++
        c = s.data[s.i]
        switch {
        case c == '\'':
            continue
        case c == '"':
            continue
        case c == ':':
            s.i--
            return buf.String(),nil
        default:
            buf.Write([]byte{c})
        }
    }

    return "",nil
}

func (s *State) nextType() int {
    for {
        if s.i >= len(s.data) {
            return -1
        }
        c := s.data[s.i]
        switch {
        case c == ' ':
            fallthrough
        case c == '\t':
            s.i++
            break
        case c == '"':
            return STRING
        case '0' <= c && c <= '9' || c == '-':
            return NUMBER
        case c == 't' || c == 'T' || c == 'f' || c == 'F':
            return BOOL
        case c == 'n':
            return NULL
        default:
            return -1
        }
    }
    return -1
}

func (s *State) ReadValue() (interface{} , error) {
    var (
        c       byte
        buf     *bytes.Buffer
    )
    buf = new(bytes.Buffer)

    i := false
    for {
        s.i++
        c = s.data[s.i]
        switch {
        case c == '\'':
            i=true
            continue
        case c == '"':
            i=true
            continue
        case !i && ('0' <= c && c <= '9' || c == '-'):
            return s.ReadNumber(),nil
        case !i && (c == 't' || c == 'T' || c == 'f' || c == 'F'):
            if strings.ToLower(string(s.data[s.i:s.i+4])) == "true" {
                s.i += 4
                return true,nil
            } else if strings.ToLower(string(s.data[s.i:s.i+5])) == "false" {
                s.i += 5
                return false,nil
            }
        case !i && (c == 'n' || c == 'N'):
            if strings.ToLower(string(s.data[s.i:s.i+4])) == "null" {
                s.i += 4
                return nil,nil
            }else{
                buf.Write([]byte{c})
            }
        case c == ',':
            fallthrough
        case c == '}':
            i=false
            // s.i--
            return buf.String(),nil
        default:
            buf.Write([]byte{c})
        }
    }

    return "",nil
}

func (s *State) ReadNumber() (interface{}) {
    var c byte
    var val int64 = 0
    var valf float64 = 0
    var mult int64 = 1
    if s.data[s.i] == '-' {
        mult = -1
        s.i++
    }
    var more = true
    var places int = 0
    for more {
        c = s.data[s.i]
        switch {
        case '0' <= c && c <= '9':
            if places != 0 {
                places *= 10
            }
            val = val*10 + int64(c-'0')
        case '}' == c:
            more = false
        case ']' == c:
            more = false
        case ',' == c:
            s.i--
            more = false
        case ' ' == c || '\t' == c:
            more = false
        case '.' == c:
            valf = float64(val)
            val = 0
            places = 1
        default:
            return nil
        }

        if s.i >= s.end-1 {
            more = false
        }
        s.i++
    }
    s.i--
    if places > 0 {
        return (valf + float64(val)/float64(places)) * float64(mult)
    } else {
        return val * mult
    }
    return nil
}
