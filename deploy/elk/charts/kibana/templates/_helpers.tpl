{{- define "kibana.elasticsearch.hosts" -}}
{{- range  $i, $host := .Values.elasticsearch.hosts -}}
{{- if $i }}
{{- print ", "  -}}
{{- end -}}
{{ . | quote}}
{{- end -}}
{{- end -}}