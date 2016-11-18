The name is {{.Name}}.
The age is {{.Age}}.

{{range.Emails}}
	An email is {{.|email_expand}}
{{end}}

{{with.Jobs}}
	{{range .}}
		An employer is {{.Employer}}
		and the role is {{.Role}}
	{{end}}
{{end}}
