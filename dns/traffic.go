package dns

import (
	"context"
	"github.com/jeessy2/ddns-go/v6/config"
	"github.com/jeessy2/ddns-go/v6/util"
	"github.com/volcengine/volc-sdk-golang/service/dns"
	"os"
)

const (
	trafficEndpoint string = "https://open.volcengineapi.com"
)

type Traffic struct {
	DNS     config.DNS
	Domains config.Domains
	TTL     string
}

func (tra *Traffic) Init(dnsConf *config.DnsConfig, ipv4cache *util.IpCache, ipv6cache *util.IpCache) {
	os.Setenv("VOLC_ACCESSKEY", dnsConf.DNS.ID)
	os.Setenv("VOLC_SECRETKEY", dnsConf.DNS.Secret)
	tra.Domains.Ipv4Cache = ipv4cache
	tra.Domains.Ipv6Cache = ipv6cache
	tra.DNS = dnsConf.DNS
	tra.Domains.GetNewIp(dnsConf)
	if dnsConf.TTL == "" {
		// 默认600s
		tra.TTL = "300"
	} else {
		tra.TTL = dnsConf.TTL
	}

}

func (tra *Traffic) AddUpdateDomainRecords() (domains config.Domains) {
	// 初始化 SDK
	ctx := context.Background()
	client := dns.NewClient(dns.NewVolcCaller())

	request := dns.UpdateRecordRequest{}
	request.RecordID = "31293584"
	request.Host = tra.Domains.Ipv6Domains[0].SubDomain
	a := tra.Domains.Ipv6Addr
	request.Value = &a
	request.Line = "default"
	b := "AAAA"
	request.Type = &b
	var c int64 = 1
	request.Weight = &c
	var d int64 = 300
	request.TTL = &d

	_, err := client.UpdateRecord(ctx, &request)
	if err != nil {
		util.Log("更新失败! %s", err)
	}
	return tra.Domains
}
