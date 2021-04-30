package doc

import "fmt"

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Golang Gin API",
	Description: "An example of gin",
}

type s struct{}

func (s *s) ReadDoc() string {
	//	sInfo := SwaggerInfo
	//	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)
	//
	//	t, err := template.New("swagger_info").Funcs(template.FuncMap{
	//		"marshal": func(v interface{}) string {
	//			a, _ := json.Marshal(v)
	//			return string(a)
	//		},
	//	}).Parse(doc)
	//	if err != nil {
	//		return doc
	//	}
	//
	//	var tpl bytes.Buffer
	//	if err := t.Execute(&tpl, sInfo); err != nil {
	//		return doc
	//	}
	//
	//	return tpl.String()
	return ""
}

func init() {
	//swag.Register(swag.Name, &s{})
	fmt.Println(SwaggerInfo)
}
