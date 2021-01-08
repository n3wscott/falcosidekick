package outputs

import (
	"fmt"
	"log"
	"strconv"

	"github.com/falcosecurity/falcosidekick/types"
	"github.com/kubeless/kubeless/pkg/utils"
	"github.com/kubernetes/client-go/rest"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KubelessCall .
func (c *Client) KubelessCall(falcopayload types.FalcoPayload) {
	str, _ := json.Marshall(falcopayload)
	c.Stats.Kubeless.Add(Total, 1)

	clientset := utils.GetClientOutOfCluster()
	svc, err := clientset.CoreV1().Services(ns).Get(funcName, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	port := strconv.Itoa(int(svc.Spec.Ports[0].Port))

	req := &rest.Request{}
	req = clientset.CoreV1().RESTClient().Post().AbsPath("/api/v1/namespaces/" + c.Config.Kubeless.Namespace + "/services/" + c.Config.Kubeless.function + ":" + port + "/proxy/").Body(str)
	req.SetHeader("Content-Type", "application/json")
	req.SetHeader("event-type", "application/json")
	req.SetHeader("User-Agent", "falcosidekick")
	eventID, err := utils.GetRandString(11)
	if err != nil {
		panic(err)
	}
	req.SetHeader("event-id", eventID)
	req.SetHeader("event-type", "falco")
	req.SetHeader("event-namespace", ns)

	res, err := req.Do().Raw()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	// Setting the success status
	go c.CountMetric(Outputs, 1, []string{"output:kubeless", "status:ok"})
	c.Stats.Kubeless.Add(OK, 1)
	c.PromStats.Outputs.With(map[string]string{"destination": "kubeless", "status": OK}).Inc()
	log.Printf("[INFO] : Kubeless - Publish OK\n")
}
