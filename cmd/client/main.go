package main

import (
	"fmt"

	"github.com/gongt/wireguard-config-distribute/internal/client"
	"github.com/gongt/wireguard-config-distribute/internal/tools"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	/* wg interface */
	ListenPort uint16 `short:"p" long:"port" description:"wireguard listening port" default:"51820" env:"WIREGUARD_PORT"`
	MTU        uint16 `long:"mtu" description:"wireguard interface MTU" env:"WIREGUARD_MTU"`

	/* config server and self config */
	Server      string `short:"s" long:"server" description:"config server ip:port" required:"true" env:"WIREGUARD_SERVER"`
	NetworkName string `long:"netgroup" description:"a name of local network, all machines in one local network should same" required:"true" env:"WIREGUARD_NETWORK"`
	JoinGroup   string `short:"g" long:"group" description:"join which VPN network" default:"default" env:"WIREGUARD_GROUP"`
	PerferIp    string `long:"perfer-ip" description:"request to use this VPN ip, only last two digist" default-mask:"allocate by server" env:"WIREGUARD_REQUEST_IP"`

	/* Self application config */
	Title    string `long:"title" description:"human readable name of this machine" default-mask:"same with hostname" env:"WIREGUARD_TITLE"`
	Hostname string `long:"hostname" description:"custom hostname (insteadof environment)" default-mask:"use HOSTNAME"`
	HostFile string `long:"hosts-file" description:"watch and read hosts file" default:"/etc/hosts" env:"WIREGUARD_HOSTS_FILE"`

	/* Public IPv4 */
	// Ipv4Only bool `short:"4" long:"ipv4only" description:"disable outgoing IPv6 connection" env:"WIREGUARD_IPV4"`
	Ipv6Only bool `short:"6" long:"ipv6only" description:"disable outgoing IPv4 connection, disable ipv4 auto detect" env:"WIREGUARD_IPV6"`

	PublicIp        string `long:"external-ip" description:"manually set public ipv4 address of this device, disable auto detect" env:"WIREGUARD_PUBLIC_IP"`
	IpServerDsiable bool   `long:"external-ip-noserver" description:"disable detect public ipv4 by talk to wireguard config server" env:"WIREGUARD_SERVER"`
	IpUpnpDsiable   bool   `long:"external-ip-noupnp" description:"disable detect public ipv4 by UPnP/NAT-PMP" env:"WIREGUARD_UPNP"`
	IpHttpDsiable   bool   `long:"external-ip-nohttp" description:"disable detect public ipv4 by request a http api" env:"WIREGUARD_HTTP"`

	/* Local IPv4 */
	InternalIp string `long:"internal-ip" description:"manually set internal ipv4 address of this device" default-mask:"detect from default route" env:"WIREGUARD_PRIVATE_IP"`
}

func main() {
	_, err := flags.NewParser(&opts, flags.HelpFlag|flags.PassDoubleDash).Parse()

	if err != nil {
		serr, ok := err.(*flags.Error)
		if ok {
			switch serr.Type {
			case flags.ErrHelp:
				return
			}
		}
		tools.Die("Failed parse arguments.\n\t%s", err.Error())
	}

	fmt.Println(opts)

	client := client.NewClient()

	client.StartNetwork()

	client.StartCommunication()

	tools.WaitForCtrlC()

	client.Quit()

	fmt.Println("Bye, bye!")
}
