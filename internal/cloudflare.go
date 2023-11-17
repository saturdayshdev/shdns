package lib

import (
	"context"
	"errors"
	"strings"

	"github.com/cloudflare/cloudflare-go"
)

type CloudflareClient struct {
	API *cloudflare.API
}

func split(name string) (subdomain string, domain string, err error) {
	values := strings.Split(name, ".")
	if len(values) < 2 {
		return "", "", errors.New("invalid domain name")
	}

	return values[0], strings.Join(values[1:], "."), nil
}

func (c *CloudflareClient) GetZoneID(zoneName string) (string, error) {
	zoneId, err := c.API.ZoneIDByName(zoneName)
	if err != nil {
		return "", err
	}

	return zoneId, nil
}

func (c* CloudflareClient) CreateDNSRecord(recType string, recName string, recValue string, recProxied bool) error {
	subdomain, domain, err := split(recName)
	if err != nil {
		return err
	}

	zoneId, err := c.GetZoneID(domain)
	if err != nil {
		return err
	}

	params := cloudflare.CreateDNSRecordParams{Type: recType, Name: subdomain, Content: recValue, Proxied: &recProxied}
	_, err = c.API.CreateDNSRecord(context.Background(), cloudflare.ZoneIdentifier(zoneId), params)
	if err != nil {
		return err
	}

	return nil
}

func InitCloudflareClient(apiKey string, apiEmail string) (*CloudflareClient, error) {
	client, err := cloudflare.New(apiKey, apiEmail)
	if err != nil {
		return nil, err
	}

	return &CloudflareClient{API: client}, nil
}