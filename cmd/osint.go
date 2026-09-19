package cmd

import (
	"ghostline/internal/osint"
	"ghostline/internal/ui"
)

func init() {
	ui.RegisterCategory(ui.Category{
		Name: "OSINT",
		Items: []ui.Item{
			{"Port Scanner",       "scan target hosts for open ports", osint.PortScanner},
			{"Whois Lookup",       "domain registration details",       osint.WhoisLookup},
			{"DNS Lookup",         "A, MX, TXT, NS, CNAME records",     osint.DNSLookup},
			{"Dox Tracker",        "lookup dox information database",   osint.DoxTracker},
			{"Dox Creator",        "create custom doxing profiles",     osint.DoxCreator},
			{"Phone Lookup",       "carrier and location of phone",     osint.PhoneLookup},
			{"Email Lookup",       "OSINT data associated with email",  osint.EmailLookup},
			{"IP Geolocation",     "geolocate any IP + port scan",      osint.IPGeo},
			{"Image Geolocation",  "geospy-style AI location from photo", osint.ImageGeo},
			{"Username Search",    "search username across 40+ sites",  osint.UsernameSearch},
			{"Breach Check",       "check email against breach databases", osint.BreachCheck},
			{"Social Scraper",     "pull public social media data",     osint.SocialScraper},
			{"Domain Reputation",  "threat intel on domains",           osint.DomainRep},
			{"SSL Cert Info",      "inspect SSL certificates (crt.sh)", osint.SSLCert},
			{"Reverse Image",      "reverse image search engine",       osint.ReverseImage},
			{"Metadata Extract",   "pull EXIF from images",             osint.MetadataExtract},
			{"Subdomain Enum",     "enumerate subdomains (crt.sh)",     osint.SubdomainEnum},
			{"Shodan Lookup",      "search shodan + censys + zoomeye",  osint.ShodanLookup},
			{"Wayback Machine",    "historic snapshots of URLs",        osint.Wayback},
			{"Dark Web Search",    "search .onion indexes (tor)",       osint.DarkWeb},
		},
	})
}