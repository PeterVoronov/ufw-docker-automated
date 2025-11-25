package ufwhandler

type TrackedContainer struct {
	Name                 string
	IPAddressMap         map[string]string
	Labels               map[string]string
	UfwInboundRules      []UfwRule
	UfwOutboundRules     []UfwRule
	UfwHostOutboundRules []UfwRule
}

type UfwRule struct {
	CIDR    string
	Port    string
	Proto   string
	Comment string
}
