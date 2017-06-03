package mapping

import (
	"testing"

	"github.com/omniscale/imposm3/element"
)

// go test  ./mapping -run TestFilters_test_t0  -v
func TestFilters_test_t0(t *testing.T) {

	/* ./testfilters_test_mapping.yml ..

	   filters:
	     require:
	       boundary: ["administrative","maritime"]
	   mapping:
	     admin_level: ['2','4']
	   type: linestring

	*/
	filterTest(
		// *testing.T
		t,
		// tablename
		"testfilters_test_t0",
		// Accept
		[]element.Tags{
			element.Tags{"admin_level": "2", "boundary": "administrative"},
			element.Tags{"admin_level": "2", "boundary": "maritime"},
			element.Tags{"admin_level": "4", "boundary": "administrative", "name": "N4"},
			element.Tags{"admin_level": "4", "boundary": "maritime", "name": "N4"},
		},
		// Reject
		[]element.Tags{
			element.Tags{"admin_level": "0", "boundary": "administrative"},
			element.Tags{"admin_level": "1", "boundary": "administrative"},
			element.Tags{"admin_level": "2", "boundary": "postal_code"},
			element.Tags{"admin_level": "2", "boundary": ""},
			element.Tags{"admin_level": "2", "boundary": "__nil__"},
			element.Tags{"admin_level": "4", "boundary": "census"},
			element.Tags{"admin_level": "3", "boundary": "administrative", "name": "NX"},
			element.Tags{"admin_level": "2"},
			element.Tags{"admin_level": "4"},
			element.Tags{"admin_level": "❤"},
			element.Tags{"admin_level": "__any__", "boundary": "__any__"},
			element.Tags{"boundary": "administrative"},
			element.Tags{"boundary": "maritime"},
			element.Tags{"name": "maritime"},
		},
	)
}

// go test  ./mapping -run TestFilters_test_t1  -v
func TestFilters_test_t1(t *testing.T) {

	/* ./testfilters_test_mapping.yml ..

	filters:
		require:
			admin_level: ["2","4"]
	mapping:
		boundary:
		- administrative
		- maritime
	type: linestring

	*/

	filterTest(
		// *testing.T
		t,
		// tablename
		"testfilters_test_t1",
		// Accept
		[]element.Tags{
			element.Tags{"admin_level": "2", "boundary": "administrative"},
			element.Tags{"admin_level": "2", "boundary": "maritime"},
			element.Tags{"admin_level": "4", "boundary": "administrative", "name": "N4"},
			element.Tags{"admin_level": "4", "boundary": "maritime", "name": "N4"},
		},
		// Reject
		[]element.Tags{
			element.Tags{"admin_level": "0", "boundary": "administrative"},
			element.Tags{"admin_level": "1", "boundary": "administrative"},
			element.Tags{"admin_level": "2", "boundary": "postal_code"},
			element.Tags{"admin_level": "2", "boundary": ""},
			element.Tags{"admin_level": "2", "boundary": "__nil__"},
			element.Tags{"admin_level": "4", "boundary": "census"},
			element.Tags{"admin_level": "3", "boundary": "administrative", "name": "NX"},
			element.Tags{"admin_level": "2"},
			element.Tags{"admin_level": "4"},
			element.Tags{"admin_level": "❤"},
			element.Tags{"admin_level": "__any__", "boundary": "__any__"},
			element.Tags{"boundary": "administrative"},
			element.Tags{"boundary": "maritime"},
			element.Tags{"name": "maritime"},
		},
	)
}

// go test  ./mapping -run TestFilters_test_t2_building  -v
func TestFilters_test_t2_building(t *testing.T) {

	/* ./testfilters_test_mapping.yml ..
	   filters:
	     reject:
	       building: ["no","none"]
	     require_regexp:
	       'addr:housenumber': '^\d+[a-zA-Z,]*$'
	       building: '^[a-z_]+$'
	   mapping:
	     building:
	     - __any__
	   type: linestring

	*/

	filterTest(
		// *testing.T
		t,
		// tablename
		"testfilters_test_t2_building",
		// Accept
		[]element.Tags{
			element.Tags{"building": "yes", "addr:housenumber": "1a"},
			element.Tags{"building": "house", "addr:housenumber": "131"},
			element.Tags{"building": "residential", "addr:housenumber": "21"},
			element.Tags{"building": "garage", "addr:housenumber": "0"},
			element.Tags{"building": "hut", "addr:housenumber": "99999999"},
			element.Tags{"building": "_", "addr:housenumber": "333"},

			element.Tags{"building": "__any__", "addr:housenumber": "333"},
			element.Tags{"building": "__nil__", "addr:housenumber": "333"},
			element.Tags{"building": "y", "addr:housenumber": "1abcdefg"},
			element.Tags{"building": "tower_block", "addr:housenumber": "1A"},
			element.Tags{"building": "shed", "name": "N4", "addr:housenumber": "1AAA"},
			element.Tags{"building": "office", "name": "N4", "addr:housenumber": "0XYAB,"},
		},
		// Reject
		[]element.Tags{
			element.Tags{"building": "yes", "addr:housenumber": "aaaaa-number"},
			element.Tags{"building": "house", "addr:housenumber": "1-3a"},
			element.Tags{"building": "house", "addr:housenumber": "❤"},
			element.Tags{"building": "house", "addr:housenumber": "two"},
			element.Tags{"building": "residential", "addr:housenumber": "x21"},

			element.Tags{"building": "", "addr:housenumber": "111"},

			element.Tags{"building": "no"},
			element.Tags{"building": "no", "addr:housenumber": "1a"},
			element.Tags{"building": "No", "addr:housenumber": "1a"},
			element.Tags{"building": "NO", "addr:housenumber": "1a"},
			element.Tags{"building": "none"},
			element.Tags{"building": "none", "addr:housenumber": "0"},
			element.Tags{"building": "nONe", "addr:housenumber": "0"},
			element.Tags{"building": "No"},
			element.Tags{"building": "NO"},
			element.Tags{"building": "NONe"},
			element.Tags{"building": "Garage"},
			element.Tags{"building": "Hut"},
			element.Tags{"building": "Farm"},
			element.Tags{"building": "tower-block"},
			element.Tags{"building": "❤"},
			element.Tags{"building": "Ümlåütê"},
			element.Tags{"building": "木"},
			element.Tags{"building": "SheD", "name": "N4"},
			element.Tags{"building": "oFFice", "name": "N4"},
			element.Tags{"admin_level": "2"},
			element.Tags{"admin_level": "4"},
			element.Tags{"boundary": "administrative"},
			element.Tags{"boundary": "maritime"},
			element.Tags{"name": "maritime"},
		},
	)
}

// go test  ./mapping -run TestFilters_test_t3_highway_with_name  -v
func TestFilters_testt3_highway_with_name(t *testing.T) {

	/* ./testfilters_test_mapping.yml ..
		filters:
	      require:
	        name: ["__any__"]
	      reject:
	        highway: ["no","none"]
	    mapping:
	      highway:
	      - __any__
	    type: linestring
	*/

	filterTest(
		// *testing.T
		t,
		// tablename
		"testfilters_test_t3_highway_with_name",
		// Accept
		[]element.Tags{
			element.Tags{"highway": "residential", "name": "N1"},
			element.Tags{"highway": "service", "name": "N2"},
			element.Tags{"highway": "track", "name": "N3"},
			element.Tags{"highway": "unclassified", "name": "N4"},
			element.Tags{"highway": "path", "name": "N5"},
			element.Tags{"highway": "", "name": "🌍🌎🌏"},
			element.Tags{"highway": "_", "name": "N6"},
			element.Tags{"highway": "y", "name": "N7"},
			element.Tags{"highway": "tower_block", "name": "N8"},
			element.Tags{"highway": "shed", "name": "N9"},
			element.Tags{"highway": "office", "name": "N10"},
			element.Tags{"highway": "SheD", "name": "N11"},
			element.Tags{"highway": "oFFice", "name": "N12"},
			element.Tags{"highway": "❤", "name": "❤"},
			element.Tags{"highway": "Ümlåütê", "name": "Ümlåütê"},
			element.Tags{"highway": "木", "name": "木"},
		},
		// Reject
		[]element.Tags{
			element.Tags{"highway": "no", "name": "N1"},
			element.Tags{"highway": "none", "name": "N2"},
			element.Tags{"highway": "yes"},
			element.Tags{"highway": "no"},
			element.Tags{"highway": "none"},
			element.Tags{"highway": "No"},
			element.Tags{"highway": "NO"},
			element.Tags{"highway": "NONe"},
			element.Tags{"highway": "Garage"},
			element.Tags{"highway": "residential"},
			element.Tags{"highway": "path"},
			element.Tags{"highway": "tower-block"},
			element.Tags{"highway": "❤"},
			element.Tags{"highway": "Ümlåütê"},
			element.Tags{"highway": "木"},
			element.Tags{"admin_level": "2"},
			element.Tags{"admin_level": "4"},
			element.Tags{"boundary": "administrative"},
			element.Tags{"boundary": "maritime"},
			element.Tags{"name": "maritime"},
		},
	)
}

// go test  ./mapping -run TestFilters_test_t4_waterway_with_name  -v
func TestFilters_test_t4_waterway_with_name(t *testing.T) {

	/* ./testfilters_test_mapping.yml ..

	   filters:
	     require:
	       name: ["__any__"]
	       waterway:
	       - stream
	       - river
	       - canal
	       - drain
	       - ditch
	     reject:
	       fixme: ['__any__']
	       amenity: ['__any__']
	       shop: ['__any__']
	       building: ['__any__']
	       tunnel: ['yes']
	     reject_regexp:
	       level: '^\D+.*$'
	   mapping:
	     waterway:
	     - __any__
	   type: linestring

	*/

	filterTest(
		// *testing.T
		t,
		// tablename
		"testfilters_test_t4_waterway_with_name",
		// Accept
		[]element.Tags{
			element.Tags{"waterway": "stream", "name": "N1"},
			element.Tags{"waterway": "river", "name": "N2"},
			element.Tags{"waterway": "canal", "name": "N3"},
			element.Tags{"waterway": "drain", "name": "N4"},
			element.Tags{"waterway": "ditch", "name": "N5"},

			element.Tags{"waterway": "stream", "name": "N1", "tunnel": "no"},
			element.Tags{"waterway": "river", "name": "N2", "boat": "no"},
			element.Tags{"waterway": "canal", "name": "N3"},
			element.Tags{"waterway": "ditch", "name": "N4", "level": "3"},

			element.Tags{"waterway": "stream", "name": "__any__"},
			element.Tags{"waterway": "stream", "name": "__nil__"},

			element.Tags{"waterway": "stream", "name": "❤"},
			element.Tags{"waterway": "stream", "name": "木"},
			element.Tags{"waterway": "stream", "name": "Ümlåütê"},
		},
		// Reject
		[]element.Tags{
			element.Tags{"waterway": "ditch", "name": "N1", "fixme": "incomplete"},
			element.Tags{"waterway": "stream", "name": "N1", "amenity": "parking"},
			element.Tags{"waterway": "river", "name": "N2", "shop": "hairdresser"},
			element.Tags{"waterway": "canal", "name": "N3", "building": "house"},
			element.Tags{"waterway": "drain", "name": "N1 tunnel", "tunnel": "yes"},

			element.Tags{"waterway": "river", "name": "N4", "level": "unknown"},
			element.Tags{"waterway": "ditch", "name": "N4", "level": "primary"},

			element.Tags{"waterway": "path", "name": "N5"},
			element.Tags{"waterway": "_", "name": "N6"},
			element.Tags{"waterway": "y", "name": "N7"},
			element.Tags{"waterway": "tower_block", "name": "N8"},
			element.Tags{"waterway": "shed", "name": "N9"},
			element.Tags{"waterway": "office", "name": "N10"},
			element.Tags{"waterway": "SheD", "name": "N11"},
			element.Tags{"waterway": "oFFice", "name": "N12"},
			element.Tags{"waterway": "❤", "name": "❤"},
			element.Tags{"waterway": "Ümlåütê", "name": "Ümlåütê"},
			element.Tags{"waterway": "木", "name": "木"},
			element.Tags{"waterway": "no", "name": "N1"},
			element.Tags{"waterway": "none", "name": "N2"},

			element.Tags{"waterway": "yes"},
			element.Tags{"waterway": "no"},
			element.Tags{"waterway": "none"},
			element.Tags{"waterway": "tower-block"},
			element.Tags{"waterway": "❤"},
			element.Tags{"waterway": "Ümlåütê"},
			element.Tags{"waterway": "木"},

			element.Tags{"waterway": "__nil__", "name": "__nil__"},
			element.Tags{"waterway": "__any__", "name": "__nil__"},

			element.Tags{"waterway": "stream", "name": "__any__", "shop": "__any__"},
			element.Tags{"waterway": "stream", "name": "__nil__", "shop": "__any__"},
			element.Tags{"waterway": "stream", "name": "__any__", "shop": "__nil__"},
			element.Tags{"waterway": "stream", "name": "__nil__", "shop": "__nil__"},
			element.Tags{"waterway": "stream", "name": "__any__", "shop": ""},
			element.Tags{"waterway": "stream", "name": "__nil__", "shop": ""},

			element.Tags{"admin_level": "2"},
			element.Tags{"admin_level": "4"},
			element.Tags{"boundary": "administrative"},
			element.Tags{"boundary": "maritime"},
			element.Tags{"name": "maritime"},
		},
	)
}

// go test  ./mapping -run TestFilters_test_t5_depricated_exclude_tags  -v
func TestFilters_test_t5_depricated_exclude_tags(t *testing.T) {

	/* ./testfilters_test_mapping.yml ..

	   filters:
	     require:
	       waterway:
	       - stream
	       - river
	       - canal
	       - drain
	       - ditch
	     exclude_tags:
	     - ['waterway', 'stream']
	     - ['waterway', 'river']
	     - ['waterway', 'canal']
	     - ['waterway', 'drain']
	     - ['waterway', 'ditch']
	   mapping:
	     waterway:
	     - __any__
	   type: linestring

	comment:
		the `exclude_tags` will be converted to `reject` filter`

	*/

	filterTest(
		// *testing.T
		t,
		// tablename
		"testfilters_test_t5_depricated_exclude_tags",
		// Accept - in this case Must be EMPTY !
		[]element.Tags{},
		// Reject
		[]element.Tags{
			element.Tags{"waterway": "stream", "name": "N1"},
			element.Tags{"waterway": "river", "name": "N2"},
			element.Tags{"waterway": "canal", "name": "N3"},
			element.Tags{"waterway": "drain", "name": "N4"},
			element.Tags{"waterway": "ditch", "name": "N5"},

			element.Tags{"waterway": "stream", "name": "N1", "tunnel": "no"},
			element.Tags{"waterway": "river", "name": "N2", "boat": "no"},
			element.Tags{"waterway": "canal", "name": "N3"},
			element.Tags{"waterway": "ditch", "name": "N4", "level": "3"},

			element.Tags{"waterway": "ditch", "name": "N1", "fixme": "incomplete"},
			element.Tags{"waterway": "stream", "name": "N1", "amenity": "parking"},
			element.Tags{"waterway": "river", "name": "N2", "shop": "hairdresser"},
			element.Tags{"waterway": "canal", "name": "N3", "building": "house"},
			element.Tags{"waterway": "drain", "name": "N1 tunnel", "tunnel": "yes"},

			element.Tags{"waterway": "river", "name": "N4", "level": "unknown"},
			element.Tags{"waterway": "ditch", "name": "N4", "level": "primary"},

			element.Tags{"waterway": "path", "name": "N5"},
			element.Tags{"waterway": "_", "name": "N6"},
			element.Tags{"waterway": "y", "name": "N7"},
			element.Tags{"waterway": "tower_block", "name": "N8"},
			element.Tags{"waterway": "shed", "name": "N9"},
			element.Tags{"waterway": "office", "name": "N10"},
			element.Tags{"waterway": "SheD", "name": "N11"},
			element.Tags{"waterway": "oFFice", "name": "N12"},
			element.Tags{"waterway": "❤", "name": "❤"},
			element.Tags{"waterway": "Ümlåütê", "name": "Ümlåütê"},
			element.Tags{"waterway": "木", "name": "木"},

			element.Tags{"waterway": "no", "name": "N1"},
			element.Tags{"waterway": "none", "name": "N2"},
			element.Tags{"waterway": "yes"},
			element.Tags{"waterway": "no"},
			element.Tags{"waterway": "none"},
			element.Tags{"waterway": "tower-block"},
			element.Tags{"waterway": "❤"},
			element.Tags{"waterway": "Ümlåütê"},
			element.Tags{"waterway": "木"},
			element.Tags{"admin_level": "2"},
			element.Tags{"admin_level": "4"},
			element.Tags{"boundary": "administrative"},
			element.Tags{"boundary": "maritime"},
			element.Tags{"name": "maritime"},
		},
	)
}

func filterTest(t *testing.T, tablename string, accept []element.Tags, reject []element.Tags) {

	var configTestMapping *Mapping
	var err error

	configTestMapping, err = NewMapping("./testfilters_mapping.yml")
	if err != nil {
		panic(err)
	}

	var actualMatch []Match

	elem := element.Way{}
	ls := configTestMapping.LineStringMatcher

	for _, et := range accept {
		elem.Tags = et
		actualMatch = ls.MatchWay(&elem)

		included := false
		for _, mt := range actualMatch {
			if tablename == mt.Table.Name {
				included = true
				break
			}
		}
		if included == false {
			t.Errorf("TestFilter - Not Accepted : (%s) (%+v)  ", tablename, et)
		}
	}

	for _, et := range reject {
		elem.Tags = et
		actualMatch = ls.MatchWay(&elem)

		included := false
		for _, mt := range actualMatch {
			if tablename == mt.Table.Name {
				included = true
				break
			}
		}
		if included == true {
			t.Errorf("TestFilter - Not Rejected : (%s) (%+v)  ", tablename, et)
		}
	}

}
