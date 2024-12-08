package internal

var RoleScopeMapping map[string][]string

func InitScope(scope map[string][]string) {
	RoleScopeMapping = scope
}
