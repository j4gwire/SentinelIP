package proxy

// Detector struct to manage proxy detection
type Detector struct{}

// NewDetector creates and returns a new instance of Detector
func NewDetector() *Detector {
	return &Detector{}
}

// IsProxy checks if an IP is a proxy (mock)
func (d *Detector) IsProxy(ip string) (bool, error) {
	// For now, let's just return true if the IP ends in ".2" (as an example)
	if ip[len(ip)-2:] == ".2" {
		return true, nil
	}
	return false, nil
}
