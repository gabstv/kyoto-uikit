package twui

import (
	"html/template"

	"github.com/yuriizinets/kyoto"
)

type MarketingSectionPricing struct {
	Title       string
	Description string
	CardBlock   []MarketingSectionPricingCardBlock

	InternalActiveTab string
}

func (c *MarketingSectionPricing) Actions(p kyoto.Page) kyoto.ActionMap {
	return kyoto.ActionMap{
		"ToggleTabMenu": func(args ...interface{}) {
			c.InternalActiveTab = args[0].(string)
		},
	}
}

func (c *MarketingSectionPricing) Init() {
	if len(c.CardBlock) > 0 {
		c.InternalActiveTab = c.CardBlock[0].Title
	}
}

type MarketingSectionPricingCardBlock struct {
	Title         string
	Inline        bool
	FeaturesTitle string
	Cards         []MarketingSectionPricingCard
}

type MarketingSectionPricingCard struct {
	Title                 string
	Description           string
	PriceTitle            string
	Price                 MarketingSectionPricingCardPrice
	AdditionalInformation MarketingSectionPricingLink
	FreeSample            MarketingSectionPricingCardSample
	Features              []string
	CTALink               MarketingSectionPricingLink
}

type MarketingSectionPricingCardPrice struct {
	Value    int
	Suffix   string // month, year, half-year
	Currency string
}

type MarketingSectionPricingLink struct {
	Href string
	Text string
}

type MarketingSectionPricingCardSample struct {
	Enabled bool

	Sample template.HTML
}
