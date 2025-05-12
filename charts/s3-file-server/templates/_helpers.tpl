{{- define "s3-file-server.name" -}}
s3-file-server
{{- end }}

{{- define "s3-file-server.fullname" -}}
{{ include "s3-file-server.name" . }}
{{- end }}