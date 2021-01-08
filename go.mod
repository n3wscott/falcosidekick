module github.com/falcosecurity/falcosidekick

go 1.15

replace github.com/kubernetes/client-go => k8s.io/client-go v0.19.0

require (
	cloud.google.com/go/pubsub v1.8.3
	github.com/Azure/azure-event-hubs-go/v3 v3.3.3
	github.com/DataDog/datadog-go v4.2.0+incompatible
	github.com/PagerDuty/go-pagerduty v1.3.0
	github.com/aws/aws-sdk-go v1.35.30
	github.com/emersion/go-sasl v0.0.0-20200509203442-7bfe0ed36a21
	github.com/emersion/go-smtp v0.14.0
	github.com/kubeless/kubeless v1.0.7
	github.com/kubernetes/client-go v0.19.0
	github.com/nats-io/nats.go v1.10.0
	github.com/nats-io/stan.go v0.7.0
	github.com/prometheus/client_golang v1.8.0
	github.com/segmentio/kafka-go v0.4.8
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	golang.org/x/oauth2 v0.0.0-20201109201403-9fd604954f58
	google.golang.org/api v0.35.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/api v0.19.0
	k8s.io/apiextensions-apiserver v0.0.0-20180327033742-750feebe2038
	k8s.io/apimachinery v0.19.0
	k8s.io/klog v1.0.0 // indirect
)
