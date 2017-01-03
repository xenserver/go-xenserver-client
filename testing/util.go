package testing

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	xen "go-xenserver-client"
	"math/rand"
	"testing"
	"time"
)

const (
	TestVMNamePrefix = "xen-client-test-"
	TestVMCPUs       = 1
	TestVMMemory     = 2048
	TestVMImage      = "ubuntu64.iso"
	TestVMGUESTOS    = "Red Hat Enterprise Linux 7"
)

var (
	testUrl      string
	testUsername string
	testPassword string
)

type TestEnvironment struct {
	client *xen.XenAPIClient
	vm     *xen.VM
}

func newTestEnvironment(t *testing.T) TestEnvironment {
	client := xen.NewXenAPIClient(testUrl, testUsername, testPassword)
	err := client.Login()
	assert.NoError(t, err)
	vms, err := client.GetVMByNameLabel(TestVMGUESTOS)
	assert.NoError(t, err)
	assert.NotEmpty(t, vms)
	cfg := xen.VMConfig{
		Name_label: fmt.Sprintf("%s-%s", TestVMNamePrefix, getRandTestName()),
		VM_ref:     vms[0].Ref,
		Other_config: map[string]string{
			"install-repository": "cdrom",
			"install-methods":    "cdrom,nfs,http,ftp",
			"default_template":   "true",
		},
		CPUMax:    TestVMCPUs,
		MemoryMax: TestVMMemory,
		Image:     TestVMImage,
	}

	vm, err := client.CreateVM(cfg)
	assert.NoError(t, err)

	err = vm.Start(false, false)
	assert.NoError(t, err)

	return TestEnvironment{
		client: client,
		vm:     vm,
	}
}

func (env *TestEnvironment) Close() {
	env.client.Close()
}

func getRandTestName() string {
	rand.Seed(time.Now().UTC().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 5)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
