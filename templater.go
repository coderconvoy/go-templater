package templater

import (
	"errors"
	"fmt"
	"net/http"
	"text/template"
)

func tDict(items ...interface{}) (map[string]interface{}, error) {
	if len(items)%2 != 0 {
		return nil, errors.New("tDict requires even number of arguments")
	}
	res := make(map[string]interface{}, len(items)/2)
	for i := 0; i < len(items)-1; i += 2 {
		k, ok := items[i].(string)
		if !ok {
			return nil, errors.New("tDict keys must be strings")
		}
		res[k] = items[i+1]
	}
	return res, nil
}

func PowerTemplates(glob string) *template.Template {
	t := template.New("")
	fMap := template.FuncMap{
		"tDict": tDict,
	}
	t = t.Funcs(fMap)
	t, err := t.ParseGlob(glob)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return t
}

func Exec(t *template.Template, w http.ResponseWriter, tName string, data interface{}) {
	err := t.ExecuteTemplate(w, tName, data)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, err)
	}
}
