
{{$comma:=sequence "" ","}}
{{range $}}{{$comma.Next}}{{.}}{{end}}

{{$comma:=sequence "" ","}}
{{$colour:=cycle "black" "white" "red"}}
{{range $}}{{$comma.Next}}{{.}} in {{$colour.Next}}{{end}}
