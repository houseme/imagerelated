/*
 *  Copyright image-related Author(https://houseme.github.io/imagerelated/). All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 *  You can obtain one at https://github.com/houseme/imagerelated.
 */

package imagerelated

import (
	"testing"
)

func TestImageType(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "TestImageType",
			args: args{
				p: "testdata/golang.png",
			},
			want:    "image/png",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ImageType(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("ImageType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
			if got != tt.want {
				t.Errorf("ImageType() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// benchmark
// BenchmarkImageType-16    	   79828	     14316 ns/op
func BenchmarkImageType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ImageType("testdata/golang.png")
	}
}

func TestImageExt(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "TestImageExt",
			args: args{
				filePath: "testdata/golang.png",
			},
			want:    "png",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ImageExt(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ImageExt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ImageExt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// benchmark
// BenchmarkImageExt-16    	   79828	     14316 ns/op
func BenchmarkImageExt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ImageExt("testdata/golang.png")
	}
}
