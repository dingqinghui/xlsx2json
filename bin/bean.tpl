/**
 * @Author: dingQingHui
 * @Description:
 * @File: bean.tpl
 * @Version: 1.0.0
 * @Date: 2024/10/29 16:27
 */

package staticdata
import (
	"os"
	"encoding/json"
	"path/filepath"
)



func jsonUnmarshal(path string, v any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err2 := json.Unmarshal(data, v); err2 != nil {
		return err2
	}
	return nil
}



type StaticBeanData struct {
{{range $i, $v := .List}}    {{$v.Name}}Dict map[{{$v.Type}}]*{{$v.Name}}
    {{$v.Name}}Array []*{{$v.Name}}
{{end}}}


{{range $i, $v := .List}}
func (v *StaticBeanData) Load{{$v.Name}}(path string) {
    fileName:=filepath.Join(path,"{{$v.Name}}.json")
    var {{$v.Name}}Array []*{{$v.Name}}
	if err := jsonUnmarshal(fileName,&{{$v.Name}}Array ); err != nil {
		return
	}
	var {{$v.Name}}Dict = make(map[{{$v.Type}}]*{{$v.Name}})
	for _, bean := range {{$v.Name}}Array{
	    {{$v.Name}}Dict[bean.GetID()] = bean
	}
	v.{{$v.Name}}Dict =  {{$v.Name}}Dict
	v.{{$v.Name}}Array = {{$v.Name}}Array
}
{{end}}


func (v *StaticBeanData) LoadAll(path string){
{{range $i, $v := .List}}   v.Load{{$v.Name}}(path)
{{end}}
}