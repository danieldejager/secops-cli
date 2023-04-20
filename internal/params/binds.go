package params

var EnvVarsBinds = []struct {
	Key     string
	Env     string
	Default string
}{
	{VTBasePathKey, VTBasePathEnv, "https://www.virustotal.com/api/v3"},
	{VTFilePathKey, VTFilePathEnv, "/files/"},
}
