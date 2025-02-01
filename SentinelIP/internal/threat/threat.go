package threat

// Detector struct to manage threat detection
type Detector struct{}

// NewDetector creates and returns a new instance of Detector
func NewDetector() *Detector {
	return &Detector{}
}

// IsThreat checks if an IP is associated with a known threat (mock)
func (d *Detector) IsThreat(ip string) (bool, error) {
	// For now, let's just return true if the IP ends in ".1" (as an example)
	if ip[len(ip)-2:] == ".1" {
		return true, nil
	}
	return false, nil
}
