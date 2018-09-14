select 
  {{.Columns}} 
from 
  {{.Table}}
{{ if eq .Where "" }}
{{ else }}
where {{ .Where }}
{{ end }};