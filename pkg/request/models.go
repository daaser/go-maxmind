// Package request provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package request

// Defines values for RequestEventType.
const (
	RequestEventTypeAccountCreation RequestEventType = "account_creation"

	RequestEventTypeAccountLogin RequestEventType = "account_login"

	RequestEventTypeEmailChange RequestEventType = "email_change"

	RequestEventTypePasswordReset RequestEventType = "password_reset"

	RequestEventTypePurchase RequestEventType = "purchase"

	RequestEventTypeRecurringPurchase RequestEventType = "recurring_purchase"

	RequestEventTypeReferral RequestEventType = "referral"

	RequestEventTypeSurvey RequestEventType = "survey"
)

// Defines values for RequestPaymentProcessor.
const (
	RequestPaymentProcessorAdyen RequestPaymentProcessor = "adyen"

	RequestPaymentProcessorAltapay RequestPaymentProcessor = "altapay"

	RequestPaymentProcessorAmazonPayments RequestPaymentProcessor = "amazon_payments"

	RequestPaymentProcessorAuthorizenet RequestPaymentProcessor = "authorizenet"

	RequestPaymentProcessorBalanced RequestPaymentProcessor = "balanced"

	RequestPaymentProcessorBeanstream RequestPaymentProcessor = "beanstream"

	RequestPaymentProcessorBluepay RequestPaymentProcessor = "bluepay"

	RequestPaymentProcessorBraintree RequestPaymentProcessor = "braintree"

	RequestPaymentProcessorCcnow RequestPaymentProcessor = "ccnow"

	RequestPaymentProcessorChasePaymentech RequestPaymentProcessor = "chase_paymentech"

	RequestPaymentProcessorCielo RequestPaymentProcessor = "cielo"

	RequestPaymentProcessorCollector RequestPaymentProcessor = "collector"

	RequestPaymentProcessorCompropago RequestPaymentProcessor = "compropago"

	RequestPaymentProcessorConceptPayments RequestPaymentProcessor = "concept_payments"

	RequestPaymentProcessorConekta RequestPaymentProcessor = "conekta"

	RequestPaymentProcessorCuentadigital RequestPaymentProcessor = "cuentadigital"

	RequestPaymentProcessorDalpay RequestPaymentProcessor = "dalpay"

	RequestPaymentProcessorDibs RequestPaymentProcessor = "dibs"

	RequestPaymentProcessorDigitalRiver RequestPaymentProcessor = "digital_river"

	RequestPaymentProcessorEbs RequestPaymentProcessor = "ebs"

	RequestPaymentProcessorEcomm365 RequestPaymentProcessor = "ecomm365"

	RequestPaymentProcessorElavon RequestPaymentProcessor = "elavon"

	RequestPaymentProcessorEpay RequestPaymentProcessor = "epay"

	RequestPaymentProcessorEprocessingNetwork RequestPaymentProcessor = "eprocessing_network"

	RequestPaymentProcessorEway RequestPaymentProcessor = "eway"

	RequestPaymentProcessorFirstData RequestPaymentProcessor = "first_data"

	RequestPaymentProcessorGlobalPayments RequestPaymentProcessor = "global_payments"

	RequestPaymentProcessorHipay RequestPaymentProcessor = "hipay"

	RequestPaymentProcessorIngenico RequestPaymentProcessor = "ingenico"

	RequestPaymentProcessorInternetsecure RequestPaymentProcessor = "internetsecure"

	RequestPaymentProcessorIntuitQuickbooksPayments RequestPaymentProcessor = "intuit_quickbooks_payments"

	RequestPaymentProcessorIugu RequestPaymentProcessor = "iugu"

	RequestPaymentProcessorLemonPay RequestPaymentProcessor = "lemon_pay"

	RequestPaymentProcessorMastercardPaymentGateway RequestPaymentProcessor = "mastercard_payment_gateway"

	RequestPaymentProcessorMercadopago RequestPaymentProcessor = "mercadopago"

	RequestPaymentProcessorMerchantEsolutions RequestPaymentProcessor = "merchant_esolutions"

	RequestPaymentProcessorMirjeh RequestPaymentProcessor = "mirjeh"

	RequestPaymentProcessorMollie RequestPaymentProcessor = "mollie"

	RequestPaymentProcessorMonerisSolutions RequestPaymentProcessor = "moneris_solutions"

	RequestPaymentProcessorNmi RequestPaymentProcessor = "nmi"

	RequestPaymentProcessorOpenpaymx RequestPaymentProcessor = "openpaymx"

	RequestPaymentProcessorOptimalPayments RequestPaymentProcessor = "optimal_payments"

	RequestPaymentProcessorOrangepay RequestPaymentProcessor = "orangepay"

	RequestPaymentProcessorOther RequestPaymentProcessor = "other"

	RequestPaymentProcessorPacnetServices RequestPaymentProcessor = "pacnet_services"

	RequestPaymentProcessorPayfast RequestPaymentProcessor = "payfast"

	RequestPaymentProcessorPaygate RequestPaymentProcessor = "paygate"

	RequestPaymentProcessorPayone RequestPaymentProcessor = "payone"

	RequestPaymentProcessorPaypal RequestPaymentProcessor = "paypal"

	RequestPaymentProcessorPayplus RequestPaymentProcessor = "payplus"

	RequestPaymentProcessorPaystation RequestPaymentProcessor = "paystation"

	RequestPaymentProcessorPaytrace RequestPaymentProcessor = "paytrace"

	RequestPaymentProcessorPaytrail RequestPaymentProcessor = "paytrail"

	RequestPaymentProcessorPayture RequestPaymentProcessor = "payture"

	RequestPaymentProcessorPayu RequestPaymentProcessor = "payu"

	RequestPaymentProcessorPayulatam RequestPaymentProcessor = "payulatam"

	RequestPaymentProcessorPinpayments RequestPaymentProcessor = "pinpayments"

	RequestPaymentProcessorPrincetonPaymentSolutions RequestPaymentProcessor = "princeton_payment_solutions"

	RequestPaymentProcessorPsigate RequestPaymentProcessor = "psigate"

	RequestPaymentProcessorQiwi RequestPaymentProcessor = "qiwi"

	RequestPaymentProcessorQuickpay RequestPaymentProcessor = "quickpay"

	RequestPaymentProcessorRaberil RequestPaymentProcessor = "raberil"

	RequestPaymentProcessorRede RequestPaymentProcessor = "rede"

	RequestPaymentProcessorRedpagos RequestPaymentProcessor = "redpagos"

	RequestPaymentProcessorRewardspay RequestPaymentProcessor = "rewardspay"

	RequestPaymentProcessorSagepay RequestPaymentProcessor = "sagepay"

	RequestPaymentProcessorSimplifyCommerce RequestPaymentProcessor = "simplify_commerce"

	RequestPaymentProcessorSkrill RequestPaymentProcessor = "skrill"

	RequestPaymentProcessorSmartcoin RequestPaymentProcessor = "smartcoin"

	RequestPaymentProcessorSpsDecidir RequestPaymentProcessor = "sps_decidir"

	RequestPaymentProcessorStripe RequestPaymentProcessor = "stripe"

	RequestPaymentProcessorTelerecargas RequestPaymentProcessor = "telerecargas"

	RequestPaymentProcessorTowah RequestPaymentProcessor = "towah"

	RequestPaymentProcessorUsaEpay RequestPaymentProcessor = "usa_epay"

	RequestPaymentProcessorVerepay RequestPaymentProcessor = "verepay"

	RequestPaymentProcessorVindicia RequestPaymentProcessor = "vindicia"

	RequestPaymentProcessorVirtualCardServices RequestPaymentProcessor = "virtual_card_services"

	RequestPaymentProcessorVme RequestPaymentProcessor = "vme"

	RequestPaymentProcessorWorldpay RequestPaymentProcessor = "worldpay"
)

// Defines values for ResponseScoreDispositionAction.
const (
	ResponseScoreDispositionActionAccept ResponseScoreDispositionAction = "accept"

	ResponseScoreDispositionActionManualReview ResponseScoreDispositionAction = "manual_review"

	ResponseScoreDispositionActionReject ResponseScoreDispositionAction = "reject"
)

// Defines values for ResponseScoreDispositionReason.
const (
	ResponseScoreDispositionReasonCustomRule ResponseScoreDispositionReason = "custom_rule"

	ResponseScoreDispositionReasonDefault ResponseScoreDispositionReason = "default"
)

// Request defines model for Request.
type Request struct {
	// This object contains account information for the end-user on the site where the event took place.
	Account    *RequestAccount    `json:"account,omitempty"`
	Billing    *RequestAddress    `json:"billing,omitempty"`
	CreditCard *RequestCreditCard `json:"credit_card,omitempty"`

	// Custom Inputs are optional inputs to the minFraud service that must first be defined for your account. Select “Custom Inputs” from the Account Portal in order to do so. See our Custom Inputs documentation for more information.
	CustomInputs *RequestCustomInputs `json:"custom_inputs,omitempty"`

	// This object contains information about the device used in the transaction.
	Device RequestDevice `json:"device"`
	Email  *RequestEmail `json:"email,omitempty"`

	// This object contains general information related to the event being scored.
	Event    *RequestEvent    `json:"event,omitempty"`
	Order    *RequestOrder    `json:"order,omitempty"`
	Payment  *RequestPayment  `json:"payment,omitempty"`
	Shipping *RequestShipping `json:"shipping,omitempty"`

	// This is an array of shopping cart item objects. A shopping cart should consist of an array of one or more item objects.
	ShoppingCart *RequestShoppingCart `json:"shopping_cart,omitempty"`
}

// This object contains account information for the end-user on the site where the event took place.
type RequestAccount struct {
	// A unique user ID associated with the end-user in your system. If your system allows the login name for the account to be changed, this should not be the login name for the account, but rather should be an internal ID that does not change. This is not your MaxMind user ID.
	UserId *string `json:"user_id,omitempty"`

	// An MD5 hash as a hexadecimal string of the username or login name associated with the account.
	UsernameMd5 *string `json:"username_md5,omitempty"`
}

// RequestAddress defines model for Request.Address.
type RequestAddress struct {
	// The first line of the user’s billing address.
	Address *string `json:"address,omitempty"`

	// The second line of the user’s billing address.
	Address2 *string `json:"address_2,omitempty"`

	// The city of the user’s billing address.
	City *string `json:"city,omitempty"`

	// The company of the end user as provided in their billing information.
	Company *string `json:"company,omitempty"`

	// The two character [ISO 3166-1 alpha-2 country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the user’s billing address.
	Country *string `json:"country,omitempty"`

	// The first name of the end user as provided in their billing information.
	FirstName *string `json:"first_name,omitempty"`

	// The last name of the end user as provided in their billing information.
	LastName *string `json:"last_name,omitempty"`

	// The country code for phone number associated with the user’s billing address.
	PhoneCountryCode *string `json:"phone_country_code,omitempty"`

	// The phone number without the country code for the user’s billing address.
	PhoneNumber *string `json:"phone_number,omitempty"`

	// The postal code of the user’s billing address.
	Postal *string `json:"postal,omitempty"`

	// The [ISO 3166-2 subdivision code](http://en.wikipedia.org/wiki/ISO_3166-2) for the user’s billing address.
	Region *string `json:"region,omitempty"`
}

// RequestCreditCard defines model for Request.CreditCard.
type RequestCreditCard struct {
	// The address verification system (AVS) check result, as returned to you by the credit card processor. The minFraud service supports the standard AVS codes.
	AvsResult *string `json:"avs_result,omitempty"`

	// The name of the issuing bank as provided by the end user.
	BankName *string `json:"bank_name,omitempty"`

	// The phone country code for the issuing bank as provided by the end user.
	BankPhoneCountryCode *string `json:"bank_phone_country_code,omitempty"`

	// The phone number, without the country code, for the issuing bank as provided by the end user.
	BankPhoneNumber *string `json:"bank_phone_number,omitempty"`

	// The card verification value (CVV) code as provided by the payment processor.
	CvvResult *string `json:"cvv_result,omitempty"`

	// The issuer ID number for the credit card. This is the first 6 digits of the credit card number. It identifies the issuing bank.
	IssuerIdNumber *string `json:"issuer_id_number,omitempty"`

	// The last four digits of the credit card number.
	Last4Digits *string `json:"last_4_digits,omitempty"`

	// A token uniquely identifying the card. The token should consist of non-space printable ASCII characters. If the token is all digits, it must be more than 19 characters long. The token must not be a primary account number (PAN) or a simple transformation of it. If you have a valid token that looks like a PAN but is not one, you may prefix that token with a fixed string, e.g., `token-`.
	Token *string `json:"token,omitempty"`
}

// Custom Inputs are optional inputs to the minFraud service that must first be defined for your account. Select “Custom Inputs” from the Account Portal in order to do so. See our Custom Inputs documentation for more information.
type RequestCustomInputs interface{}

// This object contains information about the device used in the transaction.
type RequestDevice struct {
	// The HTTP “Accept-Language” header of the device used in the transaction.
	AcceptLangauge *string `json:"accept_langauge,omitempty"`

	// The IP address associated with the device used by the customer in the transaction. The IP address must be in IPv4 or IPv6 presentation format, i.e., dotted-quad notation or the IPv6 hexadecimal-colon notation.
	IpAddress string `json:"ip_address"`

	// The HTTP “User-Agent” header of the browser used in the transaction.
	UserAgent *string `json:"user_agent,omitempty"`
}

// RequestEmail defines model for Request.Email.
type RequestEmail struct {
	// This field must be either be a valid email address or an MD5 of the email used in the transaction.
	Address *string `json:"address,omitempty"`

	// The domain of the email address used in the transaction.
	Domain *string `json:"domain,omitempty"`
}

// This object contains general information related to the event being scored.
type RequestEvent struct {
	// Your internal ID for the shop, affiliate, or merchant this order is coming from. Required for minFraud users who are resellers, payment providers, gateways and affiliate networks.
	ShopId *string `json:"shop_id,omitempty"`

	// The date and time the event occurred. The string must be in the [RFC 3339](https://tools.ietf.org/html/rfc3339) date-time format, e.g., “2012-04-12T23:20:50.52Z”. If this field is not in the request, the current time will be used.
	Time *string `json:"time,omitempty"`

	// Your internal ID for the transaction. We can use this to locate a specific transaction in our logs, and it will also show up in email alerts and notifications from us to you.
	TransactionId *string `json:"transaction_id,omitempty"`

	// The type of event being scored.
	Type *RequestEventType `json:"type,omitempty"`
}

// The type of event being scored.
type RequestEventType string

// RequestOrder defines model for Request.Order.
type RequestOrder struct {
	// The ID of the affiliate where the order is coming from.
	AffiliateId *string `json:"affiliate_id,omitempty"`

	// The total order amount for the transaction before taxes and discounts.
	Amount *float64 `json:"amount,omitempty"`

	// The [ISO 4217 currency code] (http://en.wikipedia.org/wiki/ISO_4217) for the currency used in the transaction.
	Currency *string `json:"currency,omitempty"`

	// The discount code applied to the transaction. If multiple discount codes were used, please separate them with a comma.
	DiscountCode *string `json:"discount_code,omitempty"`

	// Whether the purchaser included a gift message.
	HasGiftMessage *bool `json:"has_gift_message,omitempty"`

	// Whether order was marked as a gift by the purchaser.
	IsGift *bool `json:"is_gift,omitempty"`

	// The URI of the referring site for this order. Needs to be absolute and have a URI scheme such as `https://`.
	ReferrerUri *string `json:"referrer_uri,omitempty"`

	// The ID of the sub-affiliate where the order is coming from.
	SubaffiliateId *string `json:"subaffiliate_id,omitempty"`
}

// RequestPayment defines model for Request.Payment.
type RequestPayment struct {
	// The decline code as provided by your payment processor. If the transaction was not declined, do not include this field.
	DeclineCode *string `json:"decline_code,omitempty"`

	// If your payment processor is missing from this list, please contact [support@maxmind.com}(mailto:support@maxmind.com).
	Processor *RequestPaymentProcessor `json:"processor,omitempty"`

	// The authorization outcome from the payment processor. If the transaction has not yet been approved or denied, do not include this field.
	WasAuthorized *bool `json:"was_authorized,omitempty"`
}

// If your payment processor is missing from this list, please contact [support@maxmind.com}(mailto:support@maxmind.com).
type RequestPaymentProcessor string

// RequestShipping defines model for Request.Shipping.
type RequestShipping struct {
	// The first line of the user’s billing address.
	Address *string `json:"address,omitempty"`

	// The second line of the user’s billing address.
	Address2 *string `json:"address_2,omitempty"`

	// The city of the user’s billing address.
	City *string `json:"city,omitempty"`

	// The company of the end user as provided in their billing information.
	Company *string `json:"company,omitempty"`

	// The two character [ISO 3166-1 alpha-2 country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the user’s billing address.
	Country *string `json:"country,omitempty"`

	// The first name of the end user as provided in their billing information.
	FirstName *string `json:"first_name,omitempty"`

	// The last name of the end user as provided in their billing information.
	LastName *string `json:"last_name,omitempty"`

	// The country code for phone number associated with the user’s billing address.
	PhoneCountryCode *string `json:"phone_country_code,omitempty"`

	// The phone number without the country code for the user’s billing address.
	PhoneNumber *string `json:"phone_number,omitempty"`

	// The postal code of the user’s billing address.
	Postal *string `json:"postal,omitempty"`

	// The [ISO 3166-2 subdivision code](http://en.wikipedia.org/wiki/ISO_3166-2) for the user’s billing address.
	Region *string `json:"region,omitempty"`
}

// This is an array of shopping cart item objects. A shopping cart should consist of an array of one or more item objects.
type RequestShoppingCart []RequestShoppingCartItem

// RequestShoppingCartItem defines model for Request.ShoppingCartItem.
type RequestShoppingCartItem struct {
	// The category of the item.
	Category *string `json:"category,omitempty"`

	// Your internal ID for the item.
	ItemId *string `json:"item_id,omitempty"`

	// The per-unit price of this item in the shopping cart. This should use the same currency as the order currency.
	Price *float64 `json:"price,omitempty"`

	// The quantity of the item in the shopping cart.
	Quantity *int `json:"quantity,omitempty"`
}

// ResponseFactors defines model for Response.Factors.
type ResponseFactors struct {
	// Embedded struct due to allOf(#/components/schemas/Response.Score)
	ResponseScore `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	// This object contains minFraud response data associated with the billing address. If the billing address was not provided in the request or could not be parsed, this object will not be present in the response.
	BillingAddress *ResponseFactorsBillingAddress `json:"billing_address,omitempty"`

	// This object contains information about the device that MaxMind believes is associated with the IP address passed in the request.
	Device *ResponseFactorsDevice `json:"device,omitempty"`

	// This object contains minFraud response data associated with the shipping address. If the shipping address was not provided in the request or could not be parsed, this object will not be present in the response.
	Email     *ResponseFactorsEmail     `json:"email,omitempty"`
	IpAddress *ResponseFactorsIpAddress `json:"ip_address,omitempty"`

	// This object contains minFraud response data associated with the shipping address. If the shipping address was not provided in the request or could not be parsed, this object will not be present in the response.
	ShippingAddress *ResponseFactorsShippingAddress `json:"shipping_address,omitempty"`
}

// This object contains minFraud response data associated with the billing address. If the billing address was not provided in the request or could not be parsed, this object will not be present in the response.
type ResponseFactorsBillingAddress struct {
	// The distance in kilometers from the address to the IP location.
	DistanceToIpLocation *int `json:"distance_to_ip_location,omitempty"`

	// This field is true if the address is in the IP country. The field is false when the address is not in the IP country. If the IP address could not be geolocated, the field will not be included in the response.
	IsInIpCountry *bool `json:"is_in_ip_country,omitempty"`

	// This field is true if the postal code provided with the address is in the city for the address. The field is false when the postal code is not in the city.
	IsPostalInCity *bool `json:"is_postal_in_city,omitempty"`

	// The approximate latitude associated with the address.
	Latitude *float64 `json:"latitude,omitempty"`

	// The approximate longitude associated with the address.
	Longitude *float64 `json:"longitude,omitempty"`
}

// This object contains information about the device that MaxMind believes is associated with the IP address passed in the request.
type ResponseFactorsDevice struct {
	// A number from 0.01 to 99 representing the confidence that the `/device/id` refers to a unique device as opposed to a cluster of similar devices. A confidence of 0.01 indicates very low confidence that the device is unique, whereas 99 indicates very high confidence.
	Confidence *string `json:"confidence,omitempty"`

	// A UUID that MaxMind uses for the device associated with this IP address. Note that many devices cannot be uniquely identified because they are too common (for example, all iPhones of a given model and OS release). In these cases, the minFraud service will simply not return a UUID for that device. This is only available if you are using the [Device Tracking Add-on](https://dev.maxmind.com/minfraud/device/).
	Id *string `json:"id,omitempty"`

	// The date and time of the last sighting of the device. The value is formatted according to [RFC 3339](https://tools.ietf.org/html/rfc3339).
	LastSeen *string `json:"last_seen,omitempty"`
}

// This object contains minFraud response data associated with the shipping address. If the shipping address was not provided in the request or could not be parsed, this object will not be present in the response.
type ResponseFactorsEmail struct {
	// The distance in kilometers from the shipping address to billing address.
	DistanceToBillingAddress *int `json:"distance_to_billing_address,omitempty"`

	// The distance in kilometers from the address to the IP location.
	DistanceToIpLocation *int `json:"distance_to_ip_location,omitempty"`

	// This field is true if the shipping address is an address associated with fraudulent transactions. The field is false when the address is not associated with increased risk.
	IsHighRisk *bool `json:"is_high_risk,omitempty"`

	// This field is true if the shipping address is in the IP country. The field is false when the address is not in the IP country. If the IP address could not be geolocated, then the field will not be included in the response.
	IsInIpCountry *bool `json:"is_in_ip_country,omitempty"`

	// This field is true if the postal code provided with the address is in the city for the address. The field is false when the postal code is not in the city.
	IsPostalInCity *bool `json:"is_postal_in_city,omitempty"`

	// The approximate latitude associated with the address.
	Latitude *float64 `json:"latitude,omitempty"`

	// The approximate longitude associated with the address.
	Longitude *float64 `json:"longitude,omitempty"`
}

// ResponseFactorsIpAddress defines model for Response.Factors.IpAddress.
type ResponseFactorsIpAddress struct {
	// Embedded struct due to allOf(#/components/schemas/Response.Score.IpAddress)
	ResponseScoreIpAddress `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	City *struct {
		// Embedded fields due to inline allOf schema
		Confidence *int `json:"confidence,omitempty"`
		// Embedded struct due to allOf(#/components/schemas/Response.Factors.IpAddress.GeoName)
		ResponseFactorsIpAddressGeoName `yaml:",inline"`
	} `json:"city,omitempty"`
	Continent *struct {
		// Embedded fields due to inline allOf schema
		Code *string `json:"code,omitempty"`
		// Embedded struct due to allOf(#/components/schemas/Response.Factors.IpAddress.GeoName)
		ResponseFactorsIpAddressGeoName `yaml:",inline"`
	} `json:"continent,omitempty"`
	Country *struct {
		// Embedded fields due to inline allOf schema
		Confidence *int `json:"confidence,omitempty"`

		// This value is true if the IP country is high risk.
		IsHighRisk *bool   `json:"is_high_risk,omitempty"`
		IsoCode    *string `json:"iso_code,omitempty"`
		// Embedded struct due to allOf(#/components/schemas/Response.Factors.IpAddress.GeoName)
		ResponseFactorsIpAddressGeoName `yaml:",inline"`
	} `json:"country,omitempty"`
	Location *struct {
		AccuracyRadius *int     `json:"accuracy_radius,omitempty"`
		AverageIncome  *int     `json:"average_income,omitempty"`
		Latitude       *float64 `json:"latitude,omitempty"`

		// The date and time of the transaction in the time zone associated with the IP address. The value is formatted according to [RFC 3339](https://tools.ietf.org/html/rfc3339). For instance, the local time in Boston might be returned as `2015-04-27T19:17:24-04:00`.
		LocalTime         *string  `json:"local_time,omitempty"`
		Longitude         *float64 `json:"longitude,omitempty"`
		MetroCode         *int     `json:"metro_code,omitempty"`
		PopulationDensity *int     `json:"population_density,omitempty"`
		TimeZone          *string  `json:"time_zone,omitempty"`
	} `json:"location,omitempty"`
	Postal *struct {
		Code       *string `json:"code,omitempty"`
		Confidence *int    `json:"confidence,omitempty"`
	} `json:"postal,omitempty"`
	RegisteredCountry *struct {
		// Embedded fields due to inline allOf schema
		IsoCode *string `json:"iso_code,omitempty"`
		// Embedded struct due to allOf(#/components/schemas/Response.Factors.IpAddress.GeoName)
		ResponseFactorsIpAddressGeoName `yaml:",inline"`
	} `json:"registered_country,omitempty"`
	RepresentedCountry *struct {
		// Embedded fields due to inline allOf schema
		IsoCode *string `json:"iso_code,omitempty"`
		Type    *string `json:"type,omitempty"`
		// Embedded struct due to allOf(#/components/schemas/Response.Factors.IpAddress.GeoName)
		ResponseFactorsIpAddressGeoName `yaml:",inline"`
	} `json:"represented_country,omitempty"`
	Subdivisons *struct {
		// Embedded fields due to inline allOf schema
		Confidence *int    `json:"confidence,omitempty"`
		IsoCode    *string `json:"iso_code,omitempty"`
		// Embedded struct due to allOf(#/components/schemas/Response.Factors.IpAddress.GeoName)
		ResponseFactorsIpAddressGeoName `yaml:",inline"`
	} `json:"subdivisons,omitempty"`
	Traits *struct {
		AutonomousSystemNumber       *int    `json:"autonomous_system_number,omitempty"`
		AutonomousSystemOrganization *string `json:"autonomous_system_organization,omitempty"`
		Domain                       *string `json:"domain,omitempty"`
		IpAddress                    *string `json:"ip_address,omitempty"`
		IsAnonymousProxy             *bool   `json:"is_anonymous_proxy,omitempty"`
		IsSatelliteProvider          *bool   `json:"is_satellite_provider,omitempty"`
		Isp                          *string `json:"isp,omitempty"`
		Organization                 *string `json:"organization,omitempty"`
		UserType                     *string `json:"user_type,omitempty"`
	} `json:"traits,omitempty"`
}

// ResponseFactorsIpAddressGeoName defines model for Response.Factors.IpAddress.GeoName.
type ResponseFactorsIpAddressGeoName struct {
	GeonameId *float32 `json:"geoname_id,omitempty"`
	Names     *struct {
		De   *string `json:"de,omitempty"`
		En   *string `json:"en,omitempty"`
		Es   *string `json:"es,omitempty"`
		Fr   *string `json:"fr,omitempty"`
		Ja   *string `json:"ja,omitempty"`
		PtBR *string `json:"pt-BR,omitempty"`
		Ru   *string `json:"ru,omitempty"`
		ZhCN *string `json:"zh-CN,omitempty"`
	} `json:"names,omitempty"`
}

// This object contains minFraud response data associated with the shipping address. If the shipping address was not provided in the request or could not be parsed, this object will not be present in the response.
type ResponseFactorsShippingAddress struct {
	// The distance in kilometers from the shipping address to billing address.
	DistanceToBillingAddress *int `json:"distance_to_billing_address,omitempty"`

	// The distance in kilometers from the address to the IP location.
	DistanceToIpLocation *int `json:"distance_to_ip_location,omitempty"`

	// This field is true if the shipping address is an address associated with fraudulent transactions. The field is false when the address is not associated with increased risk.
	IsHighRisk *bool `json:"is_high_risk,omitempty"`

	// This field is true if the shipping address is in the IP country. The field is false when the address is not in the IP country. If the IP address could not be geolocated, then the field will not be included in the response.
	IsInIpCountry *bool `json:"is_in_ip_country,omitempty"`

	// This field is true if the postal code provided with the address is in the city for the address. The field is false when the postal code is not in the city.
	IsPostalInCity *bool `json:"is_postal_in_city,omitempty"`

	// The approximate latitude associated with the address.
	Latitude *float64 `json:"latitude,omitempty"`

	// The approximate longitude associated with the address.
	Longitude *float64 `json:"longitude,omitempty"`
}

// ResponseInsights defines model for Response.Insights.
type ResponseInsights struct {
	// Embedded struct due to allOf(#/components/schemas/Response.Factors)
	ResponseFactors `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Subscores *ResponseInsightsSubscores `json:"subscores,omitempty"`
}

// ResponseInsightsSubscores defines model for Response.Insights.Subscores.
type ResponseInsightsSubscores struct {
	// The risk associated with the AVS result.
	AvsResult *float64 `json:"avs_result"`

	// The risk associated with the billing address.
	BillingAddress *float64 `json:"billing_address"`

	// The risk associated with the distance between the billing address and the location for the given IP address.
	BillingAddressDistanceToIpLocation *float64 `json:"billing_address_distance_to_ip_location"`

	// The risk associated with the browser attributes such as the User-Agent and Accept-Language.
	Browser *float64 `json:"browser"`

	// Individualized risk of chargeback for the given IP address on your account and shop ID.This is only available to users sending chargeback data to MaxMind.
	Chargeback *float64 `json:"chargeback"`

	// The risk associated with the country the transaction originated from.
	Country *float64 `json:"country"`

	// The risk associated with the combination of IP country, card issuer country, billing country, and shipping country.
	CountryMismatch *float64 `json:"country_mismatch"`

	// The risk associated with the CVV result.
	CvvResult *float64 `json:"cvv_result"`

	// The risk associated with the device.
	Device *float64 `json:"device"`

	// The risk associated with the device.
	EmailAddress *float64 `json:"email_address"`

	// The general risk associated with the email domain.
	EmailDomain *float64 `json:"email_domain"`

	// The risk associated with the email address local part (the part of the email address before the @ symbol).
	EmailLocalPart *float64 `json:"email_local_part"`

	// Please use [`email_address`]() instead.
	EmailTenure *float64 `json:"email_tenure"`

	// Please use [`risk_score`]() instead.
	IpTenure *float64 `json:"ip_tenure"`

	// The risk associated with the particular issuer ID number (IIN) given the billing location and the history of usage of the IIN on your account and shop ID.
	IssuerIdNumber *float64 `json:"issuer_id_number"`

	// The risk associated with the particular order amount for your account and shop ID.
	OrderAmount *float64 `json:"order_amount"`

	// The risk associated with the particular phone number.
	PhoneNumber *float64 `json:"phone_number"`

	// The risk associated with the shipping address.
	ShippingAddress *float64 `json:"shipping_address"`

	// The risk associated with the distance between the shipping address and the location for the given IP address.
	ShippingAddressDistanceToIpLocation *float64 `json:"shipping_address_distance_to_ip_location"`

	// The risk associated with the local time of day of the transaction in the IP address location.
	TimeOfDay *float64 `json:"time_of_day"`
}

// ResponseScore defines model for Response.Score.
type ResponseScore struct {
	// This object contains information about how a request was handled by the [custom rules](https://www.maxmind.com/en/minfraud-custom-rules) you have defined. If your account does not have any custom rules defined, then this object will not be present in the response.
	Disposition *ResponseScoreDisposition `json:"disposition,omitempty"`

	// The approximate US dollar value of the funds remaining on your MaxMind account.
	FundsRemaining *float64 `json:"funds_remaining,omitempty"`

	// This is the minFraud ID, a UUID that identifies the minFraud response. Use this ID to search your minFraud logs or when making support requests to MaxMind.
	Id        *string                 `json:"id,omitempty"`
	IpAddress *ResponseScoreIpAddress `json:"ip_address,omitempty"`

	// The approximate number of queries remaining for the service before your account runs out of funds.
	QueriesRemaining *int32 `json:"queries_remaining,omitempty"`

	// This field contains the risk score, from 0.01 to 99. A higher score indicates a higher risk of fraud. For example, a score of 20 indicates a 20% chance that a transaction is fraudulent. We never return a risk score of 0, since all transactions have the possibility of being fraudulent. Likewise we never return a risk score of 100.
	RiskScore *float64 `json:"risk_score,omitempty"`

	// Array of warnings.
	Warnings *ResponseScoreWarnings `json:"warnings,omitempty"`
}

// This object contains information about how a request was handled by the [custom rules](https://www.maxmind.com/en/minfraud-custom-rules) you have defined. If your account does not have any custom rules defined, then this object will not be present in the response.
type ResponseScoreDisposition struct {
	// This describes how the request was handled.
	Action *ResponseScoreDispositionAction `json:"action,omitempty"`

	// This describes why the `action` was set to a particular value.
	Reason *ResponseScoreDispositionReason `json:"reason,omitempty"`
}

// This describes how the request was handled.
type ResponseScoreDispositionAction string

// This describes why the `action` was set to a particular value.
type ResponseScoreDispositionReason string

// ResponseScoreIpAddress defines model for Response.Score.IpAddress.
type ResponseScoreIpAddress struct {
	// This field contains the risk associated with the IP address. The value ranges from 0.01 to 99. A higher score indicates a higher risk.
	Risk *float64 `json:"risk,omitempty"`
}

// Array of warnings.
type ResponseScoreWarnings []struct {
	// This value is a machine-readable code identifying the warning.
	Code *string `json:"code,omitempty"`

	// A JSON Pointer to the input field that the warning is associated with. For instance, if the warning was about the billing city, this would be `/billing/city`. If it was for the price in the second shopping cart item, it would be `/shopping_cart/1/price`.
	InputPointer *string `json:"input_pointer,omitempty"`

	// This field provides a human-readable explanation of the warning. The description may change at any time and should not be matched against.
	Warning *string `json:"warning,omitempty"`
}

// PostFactorsJSONBody defines parameters for PostFactors.
type PostFactorsJSONBody Request

// PostInsightsJSONBody defines parameters for PostInsights.
type PostInsightsJSONBody Request

// PostScoreJSONBody defines parameters for PostScore.
type PostScoreJSONBody Request

// PostFactorsJSONRequestBody defines body for PostFactors for application/json ContentType.
type PostFactorsJSONRequestBody PostFactorsJSONBody

// PostInsightsJSONRequestBody defines body for PostInsights for application/json ContentType.
type PostInsightsJSONRequestBody PostInsightsJSONBody

// PostScoreJSONRequestBody defines body for PostScore for application/json ContentType.
type PostScoreJSONRequestBody PostScoreJSONBody
