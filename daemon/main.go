/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package main

import (
	"aks-imex-operator/client/imds_generated"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

// type VMSSVMsAPI interface {
// 	Get(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, options *VirtualMachineScaleSetVMsClientGetOptions) (armcomputevmssservice.VirtualMachineScaleSetVMsClientGetResponse, error)
// }

// type AzVMsClient struct {
// 	vmsAPI VMSSVMsAPI
// }

// func newAzVMsClient() (*AzVMsClient, error) {
// 	authorizer, err := auth.NewAuthorizer(cfg, env)
// 	if err != nil {
// 		return nil, err
// 	}

// 	azClientConfig := cfg.GetAzureClientConfig(authorizer, env)
// 	azClientConfig.UserAgent = auth.GetUserAgentExtension()
// 	cred, err := auth.NewCredential(cfg, azClientConfig.Authorizer)
// 	if err != nil {
// 		return nil, err
// 	}

// 	agentPoolClient, err := armcontainerservice.NewAgentPoolsClient(cfg.SubscriptionID, cred, opts)
// 	if err != nil {
// 		return nil, err
// 	}
// 	klog.V(5).Infof("Created agent pool client %v using token credential", agentPoolClient)

// 	return &AzVMsClient{
// 		vmsAPI: agentPoolClient,
// 	}, nil
// }

func main() {
	my_pod_name := os.Getenv("MY_POD_NAME")
	my_pod_namespace := os.Getenv("MY_POD_NAMESPACE")

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Create the Azure VM client.
	// azVMClient, err := newAzVMsClient()
	// if err != nil {
	// 	panic(err.Error())
	// }

	for {
		// get pods in all the namespaces by omitting namespace
		// Or specify namespace to get pods in particular namespace
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// Examples for error handling:
		// - Use helper functions e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		_, err = clientset.CoreV1().Pods("default").Get(context.TODO(), "example-xxxxx", metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Pod example-xxxxx not found in default namespace\n")
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found example-xxxxx pod in default namespace\n")
		}

		_, err = clientset.CoreV1().Pods(my_pod_namespace).Get(context.TODO(), my_pod_name, metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Self pod not found in namespace %s\n", my_pod_namespace)
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found pod %s in namespace %s\n", my_pod_name, my_pod_namespace)
		}

		// Here, we get the VMSS that this pod's node is a member of.
		// First, query IMDS to discover information about the node's VM.
		var PTransport = &http.Transport{Proxy: nil}

		client := http.Client{Transport: PTransport}

		req, _ := http.NewRequest("GET", "http://169.254.169.254/metadata/instance", nil)
		req.Header.Add("Metadata", "True")

		q := req.URL.Query()
		q.Add("format", "json")
		q.Add("api-version", "2023-07-01")
		req.URL.RawQuery = q.Encode()

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Errored when sending request to the server")
			return
		}

		defer resp.Body.Close()
		resp_body, _ := io.ReadAll(resp.Body)
		var imds_object imds_generated.Instance
		err = json.Unmarshal(resp_body, &imds_object)
		if err != nil {
			fmt.Println("Errored when parsing the IMDS response as an Instance")
			return
		}

		fmt.Printf("IMDS response: %s\n", resp_body)
		fmt.Printf("VMSS name from IMDS: %s\n", *imds_object.Compute.VMScaleSetName)
		fmt.Printf("VM name from IMDS: %s\n", *imds_object.Compute.Name)

		// Get the VM from the VMSS that this pod is running on and log its network interfaces.

		time.Sleep(10 * time.Second)
	}
}
