package client

import "log"

type DeployManifestSource int

const (
	DeployManifestSourceUnknown DeployManifestSource = iota
	DeployManifestSourceText
)

func (t DeployManifestSource) String() string {
	return [...]string{"UNKNOWN", "text"}[t]
}

func ParseDeployManifestSource(s string) (DeployManifestSource, error) {
	switch s {
	default:
		log.Println("[WARN] Unknown deploy manifest source:", s)
		return DeployManifestSourceUnknown, nil
	case "text":
		return DeployManifestSourceText, nil
	}
}

func (t DeployManifestSource) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *DeployManifestSource) UnmarshalText(text []byte) error {
	source, err := ParseDeployManifestSource(string(text))
	if err != nil {
		return err
	}
	*t = source
	return nil
}
