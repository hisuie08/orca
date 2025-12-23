package fake

import (
	"orca/infra/applier"
	"orca/infra/inspector"
	"orca/testdata"
)

var (
	ComposeInspector = inspector.FakeComposeInspector{
		Mock: testdata.TestComposeDir,
	}
	ConfigReader = inspector.FakeConfigReader{
		Mock: testdata.TestConfigYaml,
	}
	DockerInspector = inspector.FakeDockerInspector{}
)

var (
	ComposeWriter = applier.FakeDotOrcaDumper{ FakeDir: map[string][]byte{}}
)
