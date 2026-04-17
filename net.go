package golib

import (
	"errors"
	"net"
)

func GetLocalIPs(includeLoopback bool) ([]net.IP, error) {
	ips := []net.IP{}
	ifaces, err := net.Interfaces()
	if err != nil {
		return ips, err
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		if (iface.Flags&net.FlagLoopback != 0) && !includeLoopback {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch addr := addr.(type) {
			case *net.IPNet:
				ip = addr.IP
			case *net.IPAddr:
				ip = addr.IP
			}
			if ip == nil || (ip.IsLoopback() && !includeLoopback) {
				continue
			}

			if ipv4 := ip.To4(); ipv4 == nil && !IsSupportedIPv6(ip) {
				continue
			}
			ips = append(ips, ip)
		}
	}

	return ips, nil
}

func GetInterfaceIPs(includeLoopback bool) ([]net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	ips := []net.IP{}
	for _, addr := range addrs {
		ipnet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if !includeLoopback && ipnet.IP.IsLoopback() {
			continue
		}
		ips = append(ips, ipnet.IP)
	}
	return ips, nil
}

func IsSupportedIPv6(ip net.IP) bool {
	if len(ip) != net.IPv6len ||
		isZeros(ip[:12]) ||
		ip[0] == 0xFE && ip[1]&0xC0 == 0xC0 ||
		ip.IsLinkLocalUnicast() ||
		ip.IsLinkLocalMulticast() {
		return false
	}
	return true
}

func isZeros(ip net.IP) bool {
	for i := 0; i < len(ip); i++ {
		if ip[i] != 0 {
			return false
		}
	}
	return true
}

func FindInterface(dstIP net.IP) (*net.Interface, net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagLoopback == 0 && iface.HardwareAddr != nil {
			if addrs, err := iface.Addrs(); err == nil {
				for _, addr := range addrs {
					if ipnet, ok := addr.(*net.IPNet); ok && ipnet.Contains(dstIP) {
						return &iface, ipnet.IP, nil
					}
				}
			}
		}
	}
	return nil, nil, errors.New("no interface")
}

func GetInterface() (*net.Interface, net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagLoopback == 0 && iface.HardwareAddr != nil {
			if addrs, err := iface.Addrs(); err == nil {
				for _, addr := range addrs {
					switch a := addr.(type) {
					case *net.IPNet:
						if ip := a.IP.To4(); ip != nil {
							return &iface, ip, nil
						}
					case *net.IPAddr:
						if ip := a.IP.To4(); ip != nil {
							return &iface, ip, nil
						}
					}
				}
			}
		}
	}
	return nil, nil, errors.New("no interface")
}

func GetHardwareAddr() (net.HardwareAddr, error) {
	iface, _, err := GetInterface()
	if err != nil {
		return nil, err
	}
	return iface.HardwareAddr, nil
}

func GetIP() (net.IP, error) {
	_, ip, err := GetInterface()
	return ip, err
}

func GetInternetIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:1")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return conn.LocalAddr().(*net.UDPAddr).IP, nil
}

func GetIPv4Interfaces() ([]net.Interface, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	result := make([]net.Interface, 0, len(ifaces))
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagMulticast != 0 && HasIPv4Address(&iface) {
			result = append(result, iface)
		}
	}
	return result, nil
}

func HasIPv4Address(iface *net.Interface) bool {
	addrs, err := iface.Addrs()
	if err != nil {
		return false
	}
	for _, addr := range addrs {
		ip, _, err := net.ParseCIDR(addr.String())
		if err != nil {
			continue
		}
		if len(ip.To4()) == net.IPv4len && !ip.IsUnspecified() {
			return true
		}
	}
	return false
}

func ListenTCP(addr string) (*net.TCPListener, uint16, error) {
	if addr == "" {
		addr = "127.0.0.1:0"
	}
	laddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, 0, err
	}
	listener, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		return nil, 0, err
	}
	return listener, uint16(listener.Addr().(*net.TCPAddr).Port), nil
}
