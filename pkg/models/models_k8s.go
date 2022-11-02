package models

const K8sUMLBase = `
@startuml {{ .Name | default "UML Output" }}

{{ if .Header}}
  header {{ .Header }}
{{ end }}

{{ if .Title }}
  title {{ .Title }}
{{ end }}

{{ k8sVirtual }}

@enduml
`

const K8sVirtual = `
node {{ .Name | quote }} <<Kubernetes Cluster>>{
	{{ if .Namespaces }}
	  {{ range $i, $n := .Namespaces }}
		 {{ NamespaceModel }}
	  {{ end }}
	{{ end }}
   }
`

const NamespaceModel = `
package {{ .Namespace.Name | quote }} <<Namespace>>{
	{{ if .Deployments }}
	  {{ range $i, $d := .Deployments }}
		 {{ DeploymentModel $d }}
	  {{ end }}
	{{ end }}
   }
`

const DeploymentModel = `
frame {{ .ObjectMeta.Name }} as "{{ .ObjectMeta.Name }} rep {{ .Spec.Replicas }}" <<Deployment>> {
	{{ template "pod.tmpl" .Spec.Template }}
  }
`

const PodModel = `
collections {{ if .ObjectMeta.Name }}{{ .ObjectMeta.Name }}{{ else }} Pod {{ end }} <<Pod>> [
  {{- range $i, $c := .Spec.Containers }}
  {{- if $i}}==={{end}}
    {{ if $c.Name }}Name: {{ $c.Name }}{{ end }}
    ....
    Image: {{ $c.Image }}
    ----
    Ports:
    {{- range $idx, $p := $c.Ports }}
    {{ if $idx }}....
    {{end}}
      {{- if $p.Name }}- {{ $p.Name }}{{ end }}
      {{ if $p.ContainerPort }}port: {{ $p.ContainerPort }}{{ end }}
    {{- end }}
  {{- end }}
  {{- if .Spec.Volumes }}
    ====
    Volumes:
    {{- range $i, $v := .Spec.Volumes }}
    - name: {{ $v.Name }}
      {{- if $v.Secret }}
      secret: {{ $v.Secret.Name }}
      {{ end }}
      {{- if $v.ConfigMap }}
      configmap: {{ $v.ConfigMap.Name }}
      {{ end }}
    {{- end }}
  {{- end }}
]
`
