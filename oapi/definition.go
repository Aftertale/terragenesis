package oapi

type Definition struct {
	openapi string 'yaml:"openapi"' 'json:"openapi"'
	info Info 'yaml:"info"' 'json:"info"'
	ExternalDocs ExternalDocs 'yaml:"externalDocs"' 'json:"externalDocs"'
	servers []Server 'yaml:"servers"' 'json:"servers"'
	tags []Tag 'yaml:"tags"' 'json:"tags"'
	paths string 'yaml:"paths"' 'json:"paths"'
	components string 'yaml:"components"' 'json:"components"'
}

type Info struct {
	title string 'yaml:"title"' 'json:"title"'
	description string 'yaml:"description"' 'json:"description"'
	termsOfService string 'yaml:"termsOfService"' 'json:"termsOfService"'
	contact Contact 'yaml:"contact"' 'json:"contact"'
	license License 'yaml:"license"' 'json:"license"'
	version string 'yaml:"version"' 'json:"version"'
}

type Contact struct {
  Email string 'yaml:"email"' 'json:"email"'
}

type License struct {
  Name string 'yaml:"name"' 'json:"name"'
  URL string 'yaml:"url"' 'json:"url"'
}

type ExternalDocs struct {
  Description string 'yaml:"description"' 'json:"description"'
  URL string 'yaml:"url"' 'json:"url"'
}

type Server struct {
  URL string 'yaml:"url"' 'json:"url"'
}

type Tag struct {
  Name string
  
}