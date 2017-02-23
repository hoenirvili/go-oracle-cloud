package response

// IpAddressAssociation IP address association
// to associate an IP address reservation
type IpAddressAssociation struct {
	Name                 string   `json:"name"`
	Uri                  string   `json:"uri"`
	Description          string   `json:"description,omitempty"`
	Tags                 []string `json:"tags,omitempty"`
	Vnic                 string   `json:"vnic"`
	IpAddressReservation string   `json:"ipAddressReservation"`
}

type AllIpAddressAssociation struct {
	Result []IpAddressAssociation `json:"result,omitempty"`
}

// An IP network allows you to define an IP subnet in your account.
// The size of the IP subnet and the set IP addresses in the subnet
// are determined by the IP address prefix that you specify while creating
// the IP network. These IP addresses aren't part of the common pool
// of Oracle-provided IP addresses used by the shared network.
// When you add an instance to an IP network, the instance is assigned an
// IP address in that subnet. You can assign IP addresses to instances
// either statically or dynamically, depending on your business needs.
// So you have complete control over the IP addresses assigned to your instances
type Ip struct {
	// Description of the object.
	Description string `json:"description,omitempty"`

	// Specify the size of the IP subnet.
	// It is a range of IPv4 addresses assigned in the
	// virtual network, in CIDR address prefix format.
	// While specifying the IP address prefix take care
	// of the following points:
	//
	// * These IP addresses aren't part of the common
	// pool of Oracle-provided IP addresses used by the shared network.
	//
	// * There's no conflict with the range of IP addresses used in another
	// IP network, the IP addresses used your on-premises network,
	// or with the range of private IP addresses used in the shared network.
	// If IP networks with overlapping IP subnets are linked to an IP exchange,
	// packets going to and from those IP networks are dropped.
	//
	// * The upper limit of the CIDR block size for an IP network is /16.
	//
	// Note: The first IP address of any IP network is reserved for
	// the default gateway, the DHCP server, and the
	// DNS server of that IP network.
	IpAddressPrefix string `json:"ipAddressPrefix"`

	// Specify the IP network exchange to which the IP network belongs.
	// You can add an IP network to only one IP network exchange,
	// but an IP network exchange can include multiple IP networks.
	// An IP network exchange enables access between IP networks that
	// have non-overlapping addresses, so that instances on these
	// networks can exchange packets with each other without NAT.
	IpNetworkExchange string `json:"ipNetworkExchange,omitempty"`

	// Object names can contain only alphanumeric, underscore (_),
	// dash (-), and period (.) characters. Object names are case-sensitive.
	Name string `json:"name"`

	// If true, enable public internet access using NAPT for VNICs without any public IP reservation.
	PublicNaptEnabledFlag bool `json:"publicNaptEnabledFlag"`

	// Tags associated with the object.
	Tags []string `json:"tags,omitempty"`

	// Uniform Resource Identifier
	Uri string `json:"uri"`
}

type AllIp struct {
	Result []Ip `json:"result,omitempty"`
}

type IpReservation struct {
	Account    string   `json:"account"`
	Ip         string   `json:"ip"`
	Name       string   `json:"name"`
	Parentpool string   `json:"parentpool"`
	Quota      string   `json:"quota"`
	Tags       []string `json:"tags,omitempty"`
	Uri        string   `json:"uri"`
	Used       bool     `json:"used"`
}

type AllIpReservation struct {
	Result []IpReservation `json:"result,omitmepty"`
}

// IpAssociation is a link between an IP reservation
// and the vcable of an instance. A vcable is an attachment
// point to a specific network interface of an instance.
// A vcable is created automatically when an instance is created
// and is deleted when the instance is deleted.
type IpAssociation struct {
	Account     string `json:"account"`
	Ip          string `json:"ip"`
	Name        string `json:"name"`
	Parentpool  string `json:"parentpool"`
	Reservation string `json:"reservation"`
	Uri         string `json:"uri"`
	Vcable      string `json:"vcable"`
}

type AllIpAssociation struct {
	Result []IpAssociation `json:"result,omitempty"`
}
