
{{$name:=.Name}}
The name is {{$name}}.
The age is {{.Age}}.

{{range.Emails}}
	Name is {{$name}}, email is {{.|email_expand}}
{{end}}

{{with.Jobs}}
	{{range .}}
		An employer is {{.Employer}}
		and the role is {{.Role}}
	{{end}}
{{end}}
