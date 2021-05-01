package assets

import (
	_ "embed"
	"fmt"
	"strings"
)

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if asset, ok := assets[cannonicalName]; ok {
		return asset, nil
	}
	return nil, fmt.Errorf("asset %s not found", name)
}

//go:embed css/bootstrap-3.3.2.min.css
var cssBootstrap332MinCss []byte

//go:embed css/jquery-ui-1.10.4-smoothness.css
var cssJqueryUi1104SmoothnessCss []byte

//go:embed css/style.css
var cssStyleCss []byte

//go:embed fonts/glyphicons-halflings-regular.eot
var fontsGlyphiconsHalflingsRegularEot []byte

//go:embed fonts/glyphicons-halflings-regular.svg
var fontsGlyphiconsHalflingsRegularSvg []byte

//go:embed fonts/glyphicons-halflings-regular.ttf
var fontsGlyphiconsHalflingsRegularTtf []byte

//go:embed fonts/glyphicons-halflings-regular.woff
var fontsGlyphiconsHalflingsRegularWoff []byte

//go:embed fonts/glyphicons-halflings-regular.woff2
var fontsGlyphiconsHalflingsRegularWoff2 []byte

//go:embed images/github.png
var imagesGithubPng []byte

//go:embed images/hog.png
var imagesHogPng []byte

//go:embed js/angular-1.3.8.js
var jsAngular138Js []byte

//go:embed js/bootstrap-3.3.2.min.js
var jsBootstrap332MinJs []byte

//go:embed js/controllers.js
var jsControllersJs []byte

//go:embed js/filesize-3.1.2.min.js
var jsFilesize312MinJs []byte

//go:embed js/iso88591_map.js
var jsIso88591_mapJs []byte

//go:embed js/jquery-1.11.0.min.js
var jsJquery1110MinJs []byte

//go:embed js/jquery-ui-1.10.4.min.js
var jsJqueryUi1104MinJs []byte

//go:embed js/moment-2.8.4.js
var jsMoment284Js []byte

//go:embed js/punycode.js
var jsPunycodeJs []byte

//go:embed js/sjis_map.js
var jsSjis_mapJs []byte

//go:embed js/strutil.js
var jsStrutilJs []byte

//go:embed templates/index.html
var templatesIndexHtml []byte

//go:embed templates/layout.html
var templatesLayoutHtml []byte

// assets is a table, holding each asset generator, mapped to its name.
var assets = map[string][]byte{
	"assets/css/bootstrap-3.3.2.min.css":              cssBootstrap332MinCss,
	"assets/css/jquery-ui-1.10.4-smoothness.css":      cssJqueryUi1104SmoothnessCss,
	"assets/css/style.css":                            cssStyleCss,
	"assets/fonts/glyphicons-halflings-regular.eot":   fontsGlyphiconsHalflingsRegularEot,
	"assets/fonts/glyphicons-halflings-regular.svg":   fontsGlyphiconsHalflingsRegularSvg,
	"assets/fonts/glyphicons-halflings-regular.ttf":   fontsGlyphiconsHalflingsRegularTtf,
	"assets/fonts/glyphicons-halflings-regular.woff":  fontsGlyphiconsHalflingsRegularWoff,
	"assets/fonts/glyphicons-halflings-regular.woff2": fontsGlyphiconsHalflingsRegularWoff2,
	"assets/images/github.png":                        imagesGithubPng,
	"assets/images/hog.png":                           imagesHogPng,
	"assets/js/angular-1.3.8.js":                      jsAngular138Js,
	"assets/js/bootstrap-3.3.2.min.js":                jsBootstrap332MinJs,
	"assets/js/controllers.js":                        jsControllersJs,
	"assets/js/filesize-3.1.2.min.js":                 jsFilesize312MinJs,
	"assets/js/iso88591_map.js":                       jsIso88591_mapJs,
	"assets/js/jquery-1.11.0.min.js":                  jsJquery1110MinJs,
	"assets/js/jquery-ui-1.10.4.min.js":               jsJqueryUi1104MinJs,
	"assets/js/moment-2.8.4.js":                       jsMoment284Js,
	"assets/js/punycode.js":                           jsPunycodeJs,
	"assets/js/sjis_map.js":                           jsSjis_mapJs,
	"assets/js/strutil.js":                            jsStrutilJs,
	"assets/templates/index.html":                     templatesIndexHtml,
	"assets/templates/layout.html":                    templatesLayoutHtml,
}
