package saml

import (
	"time"

	"encoding/xml"

	"github.com/beevik/etree"
	"github.com/kr/pretty"
	. "gopkg.in/check.v1"
)

type MetadataTest struct{}

var _ = Suite(&MetadataTest{})

func (s *MetadataTest) TestCanParseMetadata(c *C) {
	buf := []byte(`<?xml version='1.0' encoding='UTF-8'?><md:EntityDescriptor ID='_af805d1c-c2e3-444e-9cf5-efc664eeace6' entityID='https://dev.aa.kndr.org/users/auth/saml/metadata' xmlns:md='urn:oasis:names:tc:SAML:2.0:metadata' xmlns:saml='urn:oasis:names:tc:SAML:2.0:assertion'><md:SPSSODescriptor AuthnRequestsSigned='false' WantAssertionsSigned='false' protocolSupportEnumeration='urn:oasis:names:tc:SAML:2.0:protocol'><md:AssertionConsumerService Binding='urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST' Location='https://dev.aa.kndr.org/users/auth/saml/callback' index='0' isDefault='true'/><md:AttributeConsumingService index='1' isDefault='true'><md:ServiceName xml:lang='en'>Required attributes</md:ServiceName><md:RequestedAttribute FriendlyName='Email address' Name='email' NameFormat='urn:oasis:names:tc:SAML:2.0:attrname-format:basic'/><md:RequestedAttribute FriendlyName='Full name' Name='name' NameFormat='urn:oasis:names:tc:SAML:2.0:attrname-format:basic'/><md:RequestedAttribute FriendlyName='Given name' Name='first_name' NameFormat='urn:oasis:names:tc:SAML:2.0:attrname-format:basic'/><md:RequestedAttribute FriendlyName='Family name' Name='last_name' NameFormat='urn:oasis:names:tc:SAML:2.0:attrname-format:basic'/></md:AttributeConsumingService></md:SPSSODescriptor></md:EntityDescriptor>`)

	metadata := EntityDescriptor{}
	err := xml.Unmarshal(buf, &metadata)
	c.Assert(err, IsNil)
	pretty.Print(metadata)
	var False = false
	var True = true
	c.Assert(metadata, DeepEquals, EntityDescriptor{
		EntityID: "https://dev.aa.kndr.org/users/auth/saml/metadata",
		ID:       "_af805d1c-c2e3-444e-9cf5-efc664eeace6",
		SPSSODescriptors: []SPSSODescriptor{
			SPSSODescriptor{
				XMLName: xml.Name{Space: "urn:oasis:names:tc:SAML:2.0:metadata", Local: "SPSSODescriptor"},
				SSODescriptor: SSODescriptor{
					RoleDescriptor: RoleDescriptor{
						ProtocolSupportEnumeration: "urn:oasis:names:tc:SAML:2.0:protocol",
					},
				},
				AuthnRequestsSigned:  &False,
				WantAssertionsSigned: &False,
				AssertionConsumerServices: []IndexedEndpoint{
					IndexedEndpoint{
						Binding:   "urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST",
						Location:  "https://dev.aa.kndr.org/users/auth/saml/callback",
						Index:     0,
						IsDefault: &True,
					},
				},
				AttributeConsumingServices: []AttributeConsumingService{
					AttributeConsumingService{
						Index:     1,
						IsDefault: &True,
						ServiceNames: []ServiceName{
							ServiceName{Value: "Required attributes"},
						},
						RequestedAttributes: []RequestedAttribute{
							{
								Attribute: Attribute{
									FriendlyName: "Email address",
									Name:         "email",
									NameFormat:   "urn:oasis:names:tc:SAML:2.0:attrname-format:basic",
								},
							},
							{
								Attribute: Attribute{
									FriendlyName: "Full name",
									Name:         "name",
									NameFormat:   "urn:oasis:names:tc:SAML:2.0:attrname-format:basic",
								},
							},
							{
								Attribute: Attribute{
									FriendlyName: "Given name",
									Name:         "first_name",
									NameFormat:   "urn:oasis:names:tc:SAML:2.0:attrname-format:basic",
								},
							},
							{
								Attribute: Attribute{
									FriendlyName: "Family name",
									Name:         "last_name",
									NameFormat:   "urn:oasis:names:tc:SAML:2.0:attrname-format:basic",
								},
							},
						},
					},
				},
			},
		},
	})
}

func (s *MetadataTest) TestCanProduceSPMetadata(c *C) {
	validUntil, _ := time.Parse("2006-02-01T15:04:05.000000", "2013-10-03T00:32:19.104000")
	AuthnRequestsSigned := true
	WantAssertionsSigned := true
	metadata := EntityDescriptor{
		EntityID:   "http://localhost:5000/e087a985171710fb9fb30f30f41384f9/saml2/metadata/",
		ValidUntil: validUntil,
		SPSSODescriptors: []SPSSODescriptor{
			SPSSODescriptor{
				AuthnRequestsSigned:  &AuthnRequestsSigned,
				WantAssertionsSigned: &WantAssertionsSigned,
				SSODescriptor: SSODescriptor{
					RoleDescriptor: RoleDescriptor{
						ProtocolSupportEnumeration: "urn:oasis:names:tc:SAML:2.0:protocol",
						KeyDescriptors: []KeyDescriptor{
							{
								Use: "encryption",
								KeyInfo: KeyInfo{
									Certificate: `MIIB7zCCAVgCCQDFzbKIp7b3MTANBgkqhkiG9w0BAQUFADA8MQswCQYDVQQGEwJVUzELMAkGA1UE
CAwCR0ExDDAKBgNVBAoMA2ZvbzESMBAGA1UEAwwJbG9jYWxob3N0MB4XDTEzMTAwMjAwMDg1MVoX
DTE0MTAwMjAwMDg1MVowPDELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkdBMQwwCgYDVQQKDANmb28x
EjAQBgNVBAMMCWxvY2FsaG9zdDCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA1PMHYmhZj308
kWLhZVT4vOulqx/9ibm5B86fPWwUKKQ2i12MYtz07tzukPymisTDhQaqyJ8Kqb/6JjhmeMnEOdTv
SPmHO8m1ZVveJU6NoKRn/mP/BD7FW52WhbrUXLSeHVSKfWkNk6S4hk9MV9TswTvyRIKvRsw0X/gf
nqkroJcCAwEAATANBgkqhkiG9w0BAQUFAAOBgQCMMlIO+GNcGekevKgkakpMdAqJfs24maGb90Dv
TLbRZRD7Xvn1MnVBBS9hzlXiFLYOInXACMW5gcoRFfeTQLSouMM8o57h0uKjfTmuoWHLQLi6hnF+
cvCsEFiJZ4AbF+DgmO6TarJ8O05t8zvnOwJlNCASPZRH/JmF8tX0hoHuAQ==`,
								},
							},
							{
								Use: "signing",
								KeyInfo: KeyInfo{
									Certificate: `MIIB7zCCAVgCCQDFzbKIp7b3MTANBgkqhkiG9w0BAQUFADA8MQswCQYDVQQGEwJVUzELMAkGA1UE
CAwCR0ExDDAKBgNVBAoMA2ZvbzESMBAGA1UEAwwJbG9jYWxob3N0MB4XDTEzMTAwMjAwMDg1MVoX
DTE0MTAwMjAwMDg1MVowPDELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkdBMQwwCgYDVQQKDANmb28x
EjAQBgNVBAMMCWxvY2FsaG9zdDCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA1PMHYmhZj308
kWLhZVT4vOulqx/9ibm5B86fPWwUKKQ2i12MYtz07tzukPymisTDhQaqyJ8Kqb/6JjhmeMnEOdTv
SPmHO8m1ZVveJU6NoKRn/mP/BD7FW52WhbrUXLSeHVSKfWkNk6S4hk9MV9TswTvyRIKvRsw0X/gf
nqkroJcCAwEAATANBgkqhkiG9w0BAQUFAAOBgQCMMlIO+GNcGekevKgkakpMdAqJfs24maGb90Dv
TLbRZRD7Xvn1MnVBBS9hzlXiFLYOInXACMW5gcoRFfeTQLSouMM8o57h0uKjfTmuoWHLQLi6hnF+
cvCsEFiJZ4AbF+DgmO6TarJ8O05t8zvnOwJlNCASPZRH/JmF8tX0hoHuAQ==`,
								},
							},
						},
					},

					SingleLogoutServices: []Endpoint{{
						Binding:  "urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect",
						Location: "http://localhost:5000/e087a985171710fb9fb30f30f41384f9/saml2/ls/",
					}},
				},

				AssertionConsumerServices: []IndexedEndpoint{{
					Binding:  "urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST",
					Location: "http://localhost:5000/e087a985171710fb9fb30f30f41384f9/saml2/ls/",
					Index:    1,
				}},
			},
		},
	}

	doc := etree.NewDocument()
	doc.SetRoot(metadata.Element())
	doc.Indent(2)
	buf, err := doc.WriteToBytes()
	c.Assert(err, IsNil)
	c.Assert(string(buf), Equals, ""+
		"<md:EntityDescriptor xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\" entityID=\"http://localhost:5000/e087a985171710fb9fb30f30f41384f9/saml2/metadata/\" validUntil=\"2013-03-10T00:32:19.104Z\">\n"+
		"  <md:md:SPSSODescriptor protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\" AuthnRequestsSigned=\"true\" WantAssertionsSigned=\"true\">\n"+
		"    <md:KeyDescriptor use=\"encryption\">\n"+
		"      <ds:KeyInfo xmlns:ds=\"http://www.w3.org/2000/09/xmldsig#\">\n"+
		"        <ds:X509Data>\n"+
		"          <ds:X509Certificate>MIIB7zCCAVgCCQDFzbKIp7b3MTANBgkqhkiG9w0BAQUFADA8MQswCQYDVQQGEwJVUzELMAkGA1UE\n"+
		"CAwCR0ExDDAKBgNVBAoMA2ZvbzESMBAGA1UEAwwJbG9jYWxob3N0MB4XDTEzMTAwMjAwMDg1MVoX\n"+
		"DTE0MTAwMjAwMDg1MVowPDELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkdBMQwwCgYDVQQKDANmb28x\n"+
		"EjAQBgNVBAMMCWxvY2FsaG9zdDCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA1PMHYmhZj308\n"+
		"kWLhZVT4vOulqx/9ibm5B86fPWwUKKQ2i12MYtz07tzukPymisTDhQaqyJ8Kqb/6JjhmeMnEOdTv\n"+
		"SPmHO8m1ZVveJU6NoKRn/mP/BD7FW52WhbrUXLSeHVSKfWkNk6S4hk9MV9TswTvyRIKvRsw0X/gf\n"+
		"nqkroJcCAwEAATANBgkqhkiG9w0BAQUFAAOBgQCMMlIO+GNcGekevKgkakpMdAqJfs24maGb90Dv\n"+
		"TLbRZRD7Xvn1MnVBBS9hzlXiFLYOInXACMW5gcoRFfeTQLSouMM8o57h0uKjfTmuoWHLQLi6hnF+\n"+
		"cvCsEFiJZ4AbF+DgmO6TarJ8O05t8zvnOwJlNCASPZRH/JmF8tX0hoHuAQ==</ds:X509Certificate>\n"+
		"        </ds:X509Data>\n"+
		"      </ds:KeyInfo>\n"+
		"    </md:KeyDescriptor>\n"+
		"    <md:KeyDescriptor use=\"signing\">\n"+
		"      <ds:KeyInfo xmlns:ds=\"http://www.w3.org/2000/09/xmldsig#\">\n"+
		"        <ds:X509Data>\n"+
		"          <ds:X509Certificate>MIIB7zCCAVgCCQDFzbKIp7b3MTANBgkqhkiG9w0BAQUFADA8MQswCQYDVQQGEwJVUzELMAkGA1UE\n"+
		"CAwCR0ExDDAKBgNVBAoMA2ZvbzESMBAGA1UEAwwJbG9jYWxob3N0MB4XDTEzMTAwMjAwMDg1MVoX\n"+
		"DTE0MTAwMjAwMDg1MVowPDELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkdBMQwwCgYDVQQKDANmb28x\n"+
		"EjAQBgNVBAMMCWxvY2FsaG9zdDCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA1PMHYmhZj308\n"+
		"kWLhZVT4vOulqx/9ibm5B86fPWwUKKQ2i12MYtz07tzukPymisTDhQaqyJ8Kqb/6JjhmeMnEOdTv\n"+
		"SPmHO8m1ZVveJU6NoKRn/mP/BD7FW52WhbrUXLSeHVSKfWkNk6S4hk9MV9TswTvyRIKvRsw0X/gf\n"+
		"nqkroJcCAwEAATANBgkqhkiG9w0BAQUFAAOBgQCMMlIO+GNcGekevKgkakpMdAqJfs24maGb90Dv\n"+
		"TLbRZRD7Xvn1MnVBBS9hzlXiFLYOInXACMW5gcoRFfeTQLSouMM8o57h0uKjfTmuoWHLQLi6hnF+\n"+
		"cvCsEFiJZ4AbF+DgmO6TarJ8O05t8zvnOwJlNCASPZRH/JmF8tX0hoHuAQ==</ds:X509Certificate>\n"+
		"        </ds:X509Data>\n"+
		"      </ds:KeyInfo>\n"+
		"    </md:KeyDescriptor>\n"+
		"    <md:EndpointType Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"http://localhost:5000/e087a985171710fb9fb30f30f41384f9/saml2/ls/\"/>\n"+
		"    <md:EndpointType Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"http://localhost:5000/e087a985171710fb9fb30f30f41384f9/saml2/ls/\" index=\"1\"/>\n"+
		"  </md:md:SPSSODescriptor>\n"+
		"</md:EntityDescriptor>\n")
}
