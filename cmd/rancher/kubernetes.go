package rancher

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"io/ioutil"
	"os"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

type K8s struct {
	client *kubernetes.Clientset
}

func NewK8sClient(kubeConfigPath string, fromFile bool) *K8s {
	k8s := new(K8s)

	if !fromFile {
		tempFile, err := ioutil.TempFile("", "rancher-")
		if err != nil {
			panic(err.Error())
		}
		ioutil.WriteFile(tempFile.Name(), []byte(kubeConfigPath), 0644)
		kubeConfigPath = tempFile.Name()
		defer os.Remove(kubeConfigPath)
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		panic(err.Error())
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	k8s.client = clientSet
	return k8s
}

func (k *K8s) UpdateDeployment(namespace string, deploymentName string, patch string)  {
	_, err := k.client.AppsV1beta2().Deployments(namespace).Patch(deploymentName, types.MergePatchType, []byte(patch))
	if err != nil {
		panic(err.Error())
	}
}

func (k *K8s) UpdateDeploymentImage(namespace string, deploymentName string, image string, tag string)  {
	deployment, err := k.client.AppsV1beta2().Deployments(namespace).Get(deploymentName, v1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	for i :=0; i < len(deployment.Spec.Template.Spec.Containers) ; i++ {
		deployment.Spec.Template.Spec.Containers[i].Image = image + ":" + tag
	}

	deployment, err = k.client.AppsV1beta2().Deployments(namespace).Update(deployment)
	if err != nil {
		panic(err.Error())
	}
}