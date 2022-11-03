package models

const K8sUMLVirtualBase = `
@startuml {{ .Name | default "UML Output" }}

{{ if .Header}}
  header {{ .Header }}
{{ end }}

{{ if .Title }}
  title {{ .Title }}
{{ end }}

{{ partial "K8sVirtual" .K8s }}

@enduml
`

const K8sUMLInfraBase = `
@startuml {{ .Name | default "UML Output" }}

{{ if .Header}}
  header {{ .Header }}
{{ end }}

{{ if .Title }}
  title {{ .Title }}
{{ end }}

{{ partial "K8sInfra" .K8s }}

@enduml`

const K8sInfra = `
Package {{ .Name | replace "-" "_" }} <<Kubernetes Cluster>>{
  {{ if .Nodes }}
    {{ range $i, $n := .Nodes }}
       {{ partial "K8sNode" $n }}
    {{ end }}
  {{ end }}
 }
 `

const K8sVirtual = `
node {{ .Name | replace "-" "_"  }} <<Kubernetes Cluster>>{
	{{ if .Namespaces }}
	  {{ range $i, $n := .Namespaces }}
     {{ partial "NamespaceModel" $n }}
	  {{ end }}
	{{ end }}
   }
`

const NamespaceModel = `
package {{ .Namespace.Name| replace "-" "_"  }} <<Namespace>>{
	{{ if .Deployments }}
	  {{ range $i, $d := .Deployments }}
    {{ partial "DeploymentModel" $d }}
	  {{ end }}
	{{ end }}
   }
`

const DeploymentModel = `
frame {{ .ObjectMeta.Name | replace "-" "_"  }} as "{{ .ObjectMeta.Name }} rep {{ .Spec.Replicas }}" <<Deployment>> {
	{{ partial "PodModel" .Spec.Template }}
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

const K8sNode = `
node {{ .Node.ObjectMeta.Name | replace "-" "_" }} <<Node>> [
    {{ .Node.ObjectMeta.Name }}
    {{- if .Node.Spec.ProviderID }}{{ .Node.Spec.ProviderID }}{{- end}}
    Unschedulable: {{ .Node.Spec.Unschedulable }}
    {{ if .Node.ObjectMeta.Labels }}
      labels
      ---
      {{ range $key, $value := .Node.ObjectMeta.Labels }}
        {{ $key }}: {{ $value }}
        ...
      {{ end }}
    {{ end }}
    ===
    {{ if .Node.Spec.Taints }}
      {{ partial "K8sTaints" .Node.Spec.Taints }}
    {{ end }}
]
`

const K8sTaints = `
{{ range $i, $t := . }}
  - {{ $t.key }}: {{ $t.value }}
    effect: {{ $t.effect }}
  ...
{{ end }}
`