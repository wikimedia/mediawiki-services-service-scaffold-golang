module github.com/nikkhn/service-scaffold-golang

go 1.15

// dont need the gopath (like we have in kask) or any of the places its in use
// b/c kask is overriding using dep found in gopath

require (
	gopkg.in/yaml.v2 v2.4.0
	schneider.vip/problem v1.6.0
)
