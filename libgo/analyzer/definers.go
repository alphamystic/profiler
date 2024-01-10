package analyzer


// use this pcap files for testing and trying out different functionality
// all analysis are written into an anlysis
type Analyzer interface {
	Analyze() (*Analysis,error) // live analyssis can be self or network spooofed
	PCAPAnalyze() (*Analysis,error)
}

// we are filtering everything and everyone then generate an analysis
// from this we should decode/generate what we find malicious
// we let filter decide the protocol then let it decode an anlysis of it
// at call/INitiation, the protocol is specified that way they can be called at async each on it,s own routine
type Filter struct {
	Protocol string //for layer types we will know surpported protocols
	// we can have an IOC or yara rule here
	YRS []dfn.YaraRule
}

// @TODO Set count for each Interface
type Analysis struct {
	InterfaceName string
	SourceIP net.IP // should be equal to "MY" IP but different for .pcap or dns level packet analysis
	DestIP net.IP
	SourceMac net.HardwareAddr
	DestinationMac net.HardwareAddr
	MIMEType string // should tell us what kind of file rather data is flowing
	// what if it's a domain name for source
	Protocol string
	DataSize int
	URLS []string // remeber to map the URL's to the IP
	IPs []map[string]string // domain:Ip
	Malicious bool
	TimeStamp time.Time
  Hash string
}

func PrintAnlysis(anls *Analysis) {
  utils.PrintTextInASpecificColorInBold("magenta","---------------------------------------------------------------------------------------------------")
  utils.PrintTextInASpecificColor("blue",fmt.Sprintf("Interface Name   %s",anls.InterfaceName))
  utils.PrintTextInASpecificColor("blue",fmt.Sprintf("Source IP   %s",anls.SourceIP.ToString()))
  utils.PrintTextInASpecificColor("blue",fmt.Sprintf("Source IP   %s",anls.DestIP.ToString()))
	utils.PrintTextInASpecificColor("blue",fmt.Sprintf("Source Mac   %s",anls.SourceMac.ToString()))
	utils.PrintTextInASpecificColor("blue",fmt.Sprintf("DestinationMac   %s",anls.DestinationMac))
	utils.PrintTextInASpecificColor("blue",fmt.Sprintf("MIMEType   %s",anls.MIMEType))
	utils.PrintTextInASpecificColor("blue",fmt.Sprintf("Protocol   %s",anls.Protocol))
	utils.PrintTextInASpecificColor("blue",fmt.Sprintf("DataSize   %d",anls.DataSize))
	utils.PrintTextInASpecificColor("blue","IP's Found:")
	// find a way to print up the IP's
	utils.PrintTextInASpecificColor("blue",fmt.Sprintf("Malicious   %s",anls.Malicious))
	utils.PrintTextInASpecificColor("blue",fmt.Sprintf("TimeStamp   %s",anls.TimeStamp.ToString()))
	utils.PrintTextInASpecificColor("blue",fmt.Sprintf("Hash (Sha256 Hash of the packet in question)   %s",anls.Hash))
  utils.PrintTextInASpecificColorInBold("magenta","---------------------------------------------------------------------------------------------------")
}

func (pf *Filter) Analyze() (*Analysis,error) {}

func (pf *Filter) PCAPAnalyze() (*Analysis,error) {}

// any port number less than 1000 is non standard
func (pf *Filter) IsNonStandardPort(packet gopacket.Packet) bool {
	// assume it's a tcp layer type (turn this to a switch pf.Protocol)
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		tcp, _ := tcpLayer.(*layers.TCP)
		//if tcp.SrcPort != 80 && tcp.DstPort != 80 && tcp.SrcPort != 443 && tcp.DstPort != 443 && tcp.SrcPort != 3000 && tcp.DstPort != 3000 && tcp.SrcPort != 44566 && tcp.DstPort != 44566 {
    // this assumes the packe is inbound
    if tcp.SrcPort >= 1000 {
			return true
		}
	}
	return false
}
