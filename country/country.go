package country

type CountryFlag struct {
	PNG string `json:"png"`
	SVG string `json:"svg"`
	Alt string `json:"alt"`
}
type NativeName struct {
	Official string `json:"official"`
	Common   string `json:"common"`
}
type CountryName struct {
	Common     string                `json:"common"`
	Official   string                `json:"official"`
	NativeName map[string]NativeName `json:"nativeName"`
}
type CurrencyDetail struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
type CountryMap struct {
	GoogleMap      string `json:"googleMaps"`
	OpenStreetMaps string `json:"openStreetMaps"`
}
type Country struct {
	Name       CountryName               `json:"name"`
	Flag       CountryFlag               `json:"flags"`
	Currencies map[string]CurrencyDetail `json:"currencies"`
	Capital    []string                  `json:"capital"`
	Region     string                    `json:"region"`
	Languages  map[string]string         `json:"languages"`
	Maps       CountryMap                `json:"maps"`
}
