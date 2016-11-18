
{{$name:=.Name}}
{"Name":"{{$name}}",
"Emails":[
{{range $k,$v:=.Emails}}	
	{{$email:=$v|email_expand}}
	{{if $k}}
		,"{{$email}}"
	{{else}}
		"{{$email}}"
	{{end}}
{{end}}
]
}
