// Scry Info.  All rights reserved.
// license that can be found in the license file.

package settings

type ScryInfo struct {
	Chain    Chain    `yaml:"chain",json:"chain"`
	Config   Config   `yaml:"config",json:"config"`
}

type Chain struct {
	Contracts Contracts `yaml:"contracts",json:"contracts"`
	Ethereum  Ethereum  `yaml:"ethereum",json:"ethereum"`
}
type Contracts struct {
	TokenAddr        string `yaml:"tokenAddr",json:"tokenAddr"`
	ProtocolAddr     string `yaml:"protocolAddr",json:"protocolAddr"`
	DeployerKeyJson  string `yaml:"deployerKeyJson",json:"deployerKeyJson"`
	DeployerPassword string `yaml:"deployerPassword",json:"deployerPassword"`
}
type Ethereum struct {
	EthNode string `yaml:"ethNode",json:"ethNode"`
}

type Config struct {
	WSPort         string `yaml:"wsPort",json:"wsPort"`
	UIResourcesDir string `yaml:"uiResourcesDir",json:"uiResourcesDir"`
	AppId          string `yaml:"appId",json:"appId"`
}
