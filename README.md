# Host analyzer CLI
Simple CLI tool to retrieve information about a host.

## Commands

./host-analyzer
NAME:
   hostname analyzer cli - Let's you query IPs, CNAMEs, MX records and Name servers

USAGE:
   hostname analyzer cli [global options] command [command options] [arguments...]

COMMANDS:
   ns       Looks up the Name Servers for a particular host
   ip       Looks up the IP addresses for a particular host
   cname    Looks up the CNAME for a particular Host
   mx       Looks up the MX records for a particular Host
   txt      Looks up the TXT records for a particular Host
   icmp     Performs an ICMP (ping) on a particual Host
   http     Performs an HTTP call and finds HTTP protocol information
   tls      Performs a TLS handshake to find information about the TLS configuration
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help