//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package cjk

import (
	"reflect"
	"testing"

	"github.com/khlieng/dispatch/Godeps/_workspace/src/github.com/blevesearch/bleve/analysis"
	"github.com/khlieng/dispatch/Godeps/_workspace/src/github.com/blevesearch/bleve/registry"
)

func TestCJKAnalyzer(t *testing.T) {
	tests := []struct {
		input  []byte
		output analysis.TokenStream
	}{
		{
			input: []byte("こんにちは世界"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:     []byte("こん"),
					Type:     analysis.Double,
					Position: 1,
					Start:    0,
					End:      6,
				},
				&analysis.Token{
					Term:     []byte("んに"),
					Type:     analysis.Double,
					Position: 2,
					Start:    3,
					End:      9,
				},
				&analysis.Token{
					Term:     []byte("にち"),
					Type:     analysis.Double,
					Position: 3,
					Start:    6,
					End:      12,
				},
				&analysis.Token{
					Term:     []byte("ちは"),
					Type:     analysis.Double,
					Position: 4,
					Start:    9,
					End:      15,
				},
				&analysis.Token{
					Term:     []byte("は世"),
					Type:     analysis.Double,
					Position: 5,
					Start:    12,
					End:      18,
				},
				&analysis.Token{
					Term:     []byte("世界"),
					Type:     analysis.Double,
					Position: 6,
					Start:    15,
					End:      21,
				},
			},
		},
		{
			input: []byte("一二三四五六七八九十"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:     []byte("一二"),
					Type:     analysis.Double,
					Position: 1,
					Start:    0,
					End:      6,
				},
				&analysis.Token{
					Term:     []byte("二三"),
					Type:     analysis.Double,
					Position: 2,
					Start:    3,
					End:      9,
				},
				&analysis.Token{
					Term:     []byte("三四"),
					Type:     analysis.Double,
					Position: 3,
					Start:    6,
					End:      12,
				},
				&analysis.Token{
					Term:     []byte("四五"),
					Type:     analysis.Double,
					Position: 4,
					Start:    9,
					End:      15,
				},
				&analysis.Token{
					Term:     []byte("五六"),
					Type:     analysis.Double,
					Position: 5,
					Start:    12,
					End:      18,
				},
				&analysis.Token{
					Term:     []byte("六七"),
					Type:     analysis.Double,
					Position: 6,
					Start:    15,
					End:      21,
				},
				&analysis.Token{
					Term:     []byte("七八"),
					Type:     analysis.Double,
					Position: 7,
					Start:    18,
					End:      24,
				},
				&analysis.Token{
					Term:     []byte("八九"),
					Type:     analysis.Double,
					Position: 8,
					Start:    21,
					End:      27,
				},
				&analysis.Token{
					Term:     []byte("九十"),
					Type:     analysis.Double,
					Position: 9,
					Start:    24,
					End:      30,
				},
			},
		},
		{
			input: []byte("一 二三四 五六七八九 十"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:     []byte("一"),
					Type:     analysis.Single,
					Position: 1,
					Start:    0,
					End:      3,
				},
				&analysis.Token{
					Term:     []byte("二三"),
					Type:     analysis.Double,
					Position: 2,
					Start:    4,
					End:      10,
				},
				&analysis.Token{
					Term:     []byte("三四"),
					Type:     analysis.Double,
					Position: 3,
					Start:    7,
					End:      13,
				},
				&analysis.Token{
					Term:     []byte("五六"),
					Type:     analysis.Double,
					Position: 5,
					Start:    14,
					End:      20,
				},
				&analysis.Token{
					Term:     []byte("六七"),
					Type:     analysis.Double,
					Position: 6,
					Start:    17,
					End:      23,
				},
				&analysis.Token{
					Term:     []byte("七八"),
					Type:     analysis.Double,
					Position: 7,
					Start:    20,
					End:      26,
				},
				&analysis.Token{
					Term:     []byte("八九"),
					Type:     analysis.Double,
					Position: 8,
					Start:    23,
					End:      29,
				},
				&analysis.Token{
					Term:     []byte("十"),
					Type:     analysis.Single,
					Position: 10,
					Start:    30,
					End:      33,
				},
			},
		},
		{
			input: []byte("abc defgh ijklmn opqrstu vwxy z"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:     []byte("abc"),
					Type:     analysis.AlphaNumeric,
					Position: 1,
					Start:    0,
					End:      3,
				},
				&analysis.Token{
					Term:     []byte("defgh"),
					Type:     analysis.AlphaNumeric,
					Position: 2,
					Start:    4,
					End:      9,
				},
				&analysis.Token{
					Term:     []byte("ijklmn"),
					Type:     analysis.AlphaNumeric,
					Position: 3,
					Start:    10,
					End:      16,
				},
				&analysis.Token{
					Term:     []byte("opqrstu"),
					Type:     analysis.AlphaNumeric,
					Position: 4,
					Start:    17,
					End:      24,
				},
				&analysis.Token{
					Term:     []byte("vwxy"),
					Type:     analysis.AlphaNumeric,
					Position: 5,
					Start:    25,
					End:      29,
				},
				&analysis.Token{
					Term:     []byte("z"),
					Type:     analysis.AlphaNumeric,
					Position: 6,
					Start:    30,
					End:      31,
				},
			},
		},
		{
			input: []byte("あい"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:     []byte("あい"),
					Type:     analysis.Double,
					Position: 1,
					Start:    0,
					End:      6,
				},
			},
		},
		{
			input: []byte("あい   "),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:     []byte("あい"),
					Type:     analysis.Double,
					Position: 1,
					Start:    0,
					End:      6,
				},
			},
		},
		{
			input: []byte("test"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:     []byte("test"),
					Type:     analysis.AlphaNumeric,
					Position: 1,
					Start:    0,
					End:      4,
				},
			},
		},
		{
			input: []byte("test   "),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:     []byte("test"),
					Type:     analysis.AlphaNumeric,
					Position: 1,
					Start:    0,
					End:      4,
				},
			},
		},
		{
			input: []byte("あいtest"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:     []byte("あい"),
					Type:     analysis.Double,
					Position: 1,
					Start:    0,
					End:      6,
				},
				&analysis.Token{
					Term:     []byte("test"),
					Type:     analysis.AlphaNumeric,
					Position: 3,
					Start:    6,
					End:      10,
				},
			},
		},
		{
			input: []byte("testあい    "),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:     []byte("test"),
					Type:     analysis.AlphaNumeric,
					Position: 1,
					Start:    0,
					End:      4,
				},
				&analysis.Token{
					Term:     []byte("あい"),
					Type:     analysis.Double,
					Position: 2,
					Start:    4,
					End:      10,
				},
			},
		},
		{
			input: []byte("あいうえおabcかきくけこ"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:     []byte("あい"),
					Type:     analysis.Double,
					Position: 1,
					Start:    0,
					End:      6,
				},
				&analysis.Token{
					Term:     []byte("いう"),
					Type:     analysis.Double,
					Position: 2,
					Start:    3,
					End:      9,
				},
				&analysis.Token{
					Term:     []byte("うえ"),
					Type:     analysis.Double,
					Position: 3,
					Start:    6,
					End:      12,
				},
				&analysis.Token{
					Term:     []byte("えお"),
					Type:     analysis.Double,
					Position: 4,
					Start:    9,
					End:      15,
				},
				&analysis.Token{
					Term:     []byte("abc"),
					Type:     analysis.AlphaNumeric,
					Position: 6,
					Start:    15,
					End:      18,
				},
				&analysis.Token{
					Term:     []byte("かき"),
					Type:     analysis.Double,
					Position: 7,
					Start:    18,
					End:      24,
				},
				&analysis.Token{
					Term:     []byte("きく"),
					Type:     analysis.Double,
					Position: 8,
					Start:    21,
					End:      27,
				},
				&analysis.Token{
					Term:     []byte("くけ"),
					Type:     analysis.Double,
					Position: 9,
					Start:    24,
					End:      30,
				},
				&analysis.Token{
					Term:     []byte("けこ"),
					Type:     analysis.Double,
					Position: 10,
					Start:    27,
					End:      33,
				},
			},
		},
		{
			input: []byte("あいうえおabんcかきくけ こ"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:     []byte("あい"),
					Type:     analysis.Double,
					Position: 1,
					Start:    0,
					End:      6,
				},
				&analysis.Token{
					Term:     []byte("いう"),
					Type:     analysis.Double,
					Position: 2,
					Start:    3,
					End:      9,
				},
				&analysis.Token{
					Term:     []byte("うえ"),
					Type:     analysis.Double,
					Position: 3,
					Start:    6,
					End:      12,
				},
				&analysis.Token{
					Term:     []byte("えお"),
					Type:     analysis.Double,
					Position: 4,
					Start:    9,
					End:      15,
				},
				&analysis.Token{
					Term:     []byte("ab"),
					Type:     analysis.AlphaNumeric,
					Position: 6,
					Start:    15,
					End:      17,
				},
				&analysis.Token{
					Term:     []byte("ん"),
					Type:     analysis.Single,
					Position: 7,
					Start:    17,
					End:      20,
				},
				&analysis.Token{
					Term:     []byte("c"),
					Type:     analysis.AlphaNumeric,
					Position: 8,
					Start:    20,
					End:      21,
				},
				&analysis.Token{
					Term:     []byte("かき"),
					Type:     analysis.Double,
					Position: 9,
					Start:    21,
					End:      27,
				},
				&analysis.Token{
					Term:     []byte("きく"),
					Type:     analysis.Double,
					Position: 10,
					Start:    24,
					End:      30,
				},
				&analysis.Token{
					Term:     []byte("くけ"),
					Type:     analysis.Double,
					Position: 11,
					Start:    27,
					End:      33,
				},
				&analysis.Token{
					Term:     []byte("こ"),
					Type:     analysis.Single,
					Position: 13,
					Start:    34,
					End:      37,
				},
			},
		},
		{
			input: []byte("一 روبرت موير"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:     []byte("一"),
					Type:     analysis.Single,
					Position: 1,
					Start:    0,
					End:      3,
				},
				&analysis.Token{
					Term:     []byte("روبرت"),
					Type:     analysis.AlphaNumeric,
					Position: 2,
					Start:    4,
					End:      14,
				},
				&analysis.Token{
					Term:     []byte("موير"),
					Type:     analysis.AlphaNumeric,
					Position: 3,
					Start:    15,
					End:      23,
				},
			},
		},
		{
			input: []byte("一 رُوبرت موير"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:     []byte("一"),
					Type:     analysis.Single,
					Position: 1,
					Start:    0,
					End:      3,
				},
				&analysis.Token{
					Term:     []byte("رُوبرت"),
					Type:     analysis.AlphaNumeric,
					Position: 2,
					Start:    4,
					End:      16,
				},
				&analysis.Token{
					Term:     []byte("موير"),
					Type:     analysis.AlphaNumeric,
					Position: 3,
					Start:    17,
					End:      25,
				},
			},
		},
		{
			input: []byte("𩬅艱鍟䇹愯瀛"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:     []byte("𩬅艱"),
					Type:     analysis.Double,
					Position: 1,
					Start:    0,
					End:      7,
				},
				&analysis.Token{
					Term:     []byte("艱鍟"),
					Type:     analysis.Double,
					Position: 2,
					Start:    4,
					End:      10,
				},
				&analysis.Token{
					Term:     []byte("鍟䇹"),
					Type:     analysis.Double,
					Position: 3,
					Start:    7,
					End:      13,
				},
				&analysis.Token{
					Term:     []byte("䇹愯"),
					Type:     analysis.Double,
					Position: 4,
					Start:    10,
					End:      16,
				},
				&analysis.Token{
					Term:     []byte("愯瀛"),
					Type:     analysis.Double,
					Position: 5,
					Start:    13,
					End:      19,
				},
			},
		},
		{
			input: []byte("一"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:     []byte("一"),
					Type:     analysis.Single,
					Position: 1,
					Start:    0,
					End:      3,
				},
			},
		},
		{
			input: []byte("一丁丂"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:     []byte("一丁"),
					Type:     analysis.Double,
					Position: 1,
					Start:    0,
					End:      6,
				},
				&analysis.Token{
					Term:     []byte("丁丂"),
					Type:     analysis.Double,
					Position: 2,
					Start:    3,
					End:      9,
				},
			},
		},
	}

	cache := registry.NewCache()
	for _, test := range tests {
		analyzer, err := cache.AnalyzerNamed(AnalyzerName)
		if err != nil {
			t.Fatal(err)
		}
		actual := analyzer.Analyze(test.input)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf("expected %v, got %v", test.output, actual)
		}
	}
}
