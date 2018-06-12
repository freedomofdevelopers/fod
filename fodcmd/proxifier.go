package main

import (
	"io"
	"strings"
	"text/template"
)

var xmlTmpl = template.Must(template.New("xml").Parse(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<ProxifierProfile version="101" platform="Windows" product_id="0" product_minver="310">
  <Options>
    <Resolve>
      <AutoModeDetection enabled="false" />
      <ViaProxy enabled="true">
        <TryLocalDnsFirst enabled="false" />
      </ViaProxy>
      <ExclusionList>%ComputerName%; localhost; *.local</ExclusionList>
    </Resolve>
    <Encryption mode="basic" />
    <HttpProxiesSupport enabled="true" />
    <HandleDirectConnections enabled="false" />
    <ConnectionLoopDetection enabled="true" />
    <ProcessServices enabled="false" />
    <ProcessOtherUsers enabled="false" />
  </Options>
  <ProxyList>
    <Proxy id="100" type="HTTPS">
      <Address>fod.backtory.com</Address>
      <Port>8118</Port>
      <Options>48</Options>
    </Proxy>
  </ProxyList>
  <ChainList />
  <RuleList>
    <Rule enabled="true">
      <Name>FOD</Name>
      <Targets>{{.}}</Targets>
      <Action type="Proxy">100</Action>
    </Rule>
    <Rule enabled="true">
      <Name>Localhost</Name>
      <Targets>localhost; 127.0.0.1; %ComputerName%</Targets>
      <Action type="Direct" />
    </Rule>
    <Rule enabled="true">
      <Name>Default</Name>
      <Action type="Direct" />
    </Rule>
  </RuleList>
</ProxifierProfile>`))

func proxifier(target io.Writer, domains ...string) error {
	d := make([]string, len(domains))
	for i := range domains {
		d[i] = "*" + domains[i]
	}

	return xmlTmpl.Execute(target, strings.Join(d, ";"))
}
