
# SentinelIP

**SentinelIP** is a real-time IP threat detection service that combines proxy & VPN detection with IP threat intelligence to protect against malicious traffic. It helps identify suspicious IPs, mitigate fraud, and block bot attacks for web applications and APIs.

## Features

- **Real-time Threat Detection**: Detect and flag malicious IP addresses in real-time.
- **Proxy & VPN Detection**: Identify users accessing through proxies or VPNs.
- **IP Threat Intelligence**: Use known threat intelligence to block malicious traffic.
- **API Integration**: Easily integrate with your web applications or APIs for automatic threat detection.

## Installation

### Prerequisites
- Go 1.18+ installed
- Access to the internet for real-time threat data and IP checks.

### Clone the repository
```bash ```
git clone https://github.com/ScribeAegis/SentinelIP.git
cd SentinelIP


### Running Locally

1.  Install dependencies:  go mod tidy
2. Start the application: go run cmd/main.go
3. The application will run locally on `http://localhost:8080/`. You can test the threat detection functionality at: http://localhost:8080/threat-detection

## Usage

### Threat Detection Endpoint

Send a request to the endpoint with an IP address for detection.

-   **URL**: `/threat-detection`
-   **Method**: `GET`
-   **Query Parameter**:
    -   `ip` (string): The IP address to be checked.

Example Request: http://localhost:8080/threat-detection?ip=8.8.8.8

Example Response: 
{
   "ip": "8.8.8.8",
   "status": "safe"
}

## Contributing

We welcome contributions to **SentinelIP**! To contribute:

1.  Fork the repository.
2.  Create a new branch for your changes.
3.  Submit a pull request with a description of your changes.

## License

Distributed under the MIT License. See LICENSE for more information.

## Acknowledgements

-   Go programming language
-   Open-source threat intelligence APIs
-   Community contributions
