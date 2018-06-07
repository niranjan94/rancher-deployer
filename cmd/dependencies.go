package cmd

import (
	"fmt"
	"runtime"
	"github.com/niranjan94/rancher-deployer/cmd/utils"
	"github.com/fatih/color"
	"os"
	"io/ioutil"
)

const KubectlVersion = "v1.10.3"
const RancherCliVersion = "v2.0.0"

//
// Download kubectl & rancher-cli
//
func DownloadDependencies() string {
	dir, _ := ioutil.TempDir("", "rancher-deployer");
	kubectlUrl := fmt.Sprintf(
		"https://storage.googleapis.com/kubernetes-release/release/%s/bin/%s/%s/kubectl",
		KubectlVersion,
		runtime.GOOS,
		runtime.GOARCH,
	)
	kubectlDir := dir + "/kubectl"
	utils.DownloadFile(kubectlDir, kubectlUrl)
	utils.RunCommand("chmod", "+x", kubectlDir)
	color.Green("✔Downloaded kubectl " + KubectlVersion)

	rancherUrl := fmt.Sprintf(
		"https://github.com/rancher/cli/releases/download/%s/rancher-%s-%s-%s.tar.gz",
		RancherCliVersion,
		runtime.GOOS,
		runtime.GOARCH,
		RancherCliVersion,
	)
	rancherDir := dir + "/rancher"
	rancherCliFiles := utils.DownloadExtract(rancherUrl)
	os.Rename(fmt.Sprintf("%s/rancher-%s/rancher", rancherCliFiles, RancherCliVersion), rancherDir)
	utils.RunCommand("chmod", "+x", rancherDir)
	os.Environ()
	color.Green("✔downloaded rancher-cli " + RancherCliVersion)
	defer os.RemoveAll(rancherCliFiles)
	os.Setenv("PATH", dir + ":" + os.Getenv("PATH"))
	return dir
}
