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

// ImageExtList is a function that returns a list of image extensions
func ImageExtList() []string {
	return []string{
		"ase",
		"art",
		"bmp",
		"blp",
		"cd5",
		"cit",
		"cpt",
		"cr2",
		"cut",
		"dds",
		"dib",
		"djvu",
		"exif",
		"gif",
		"gpl",
		"grf",
		"icns",
		"ico",
		"jng",
		"jpeg",
		"jpg",
		"jfif",
		"jp2",
		"jps",
		"max",
		"miff",
		"mng",
		"msp",
		"nitf",
		"ota",
		"pbm",
		"pc1",
		"pc2",
		"pc3",
		"pcf",
		"pdn",
		"pgm",
		"PI1",
		"PI2",
		"PI3",
		"pict",
		"pct",
		"pnm",
		"pns",
		"ppm",
		"psb",
		"psd",
		"pdd",
		"psp",
		"px",
		"pxm",
		"pxr",
		"qfx",
		"raw",
		"rle",
		"sct",
		"rgb",
		"tiff",
		"tif",
		"vtf",
		"xbm",
		"xcf",
		"xpm",
		"3dv",
		"amf",
		"ai",
		"awg",
		"cgm",
		"cdr",
		"cmx",
		"dxf",
		"e2d",
		"egt",
		"eps",
		"fs",
		"gbr",
		"odg",
		"svg",
		"stl",
		"vrml",
		"x3d",
		"sxd",
		"v2d",
		"vnd",
		"wmf",
		"emf",
		"xar",
		"png",
		"webp",
		"jxr",
		"hdp",
		"wdp",
		"cur",
		"ecw",
		"iff",
		"lbm",
		"liff",
		"nrrd",
		"pam",
		"pcx",
		"pgf",
		"sgi",
		"rgba",
		"bw",
		"int",
		"inta",
		"sid",
		"ras",
		"sun",
		"tga",
	}
}

// ImageExtMap is a function that returns a map of image extensions
func ImageExtMap() map[string]bool {
	return map[string]bool{
		"ase":  true,
		"art":  true,
		"bmp":  true,
		"blp":  true,
		"cd5":  true,
		"cit":  true,
		"cpt":  true,
		"cr2":  true,
		"cut":  true,
		"dds":  true,
		"dib":  true,
		"djvu": true,
		"egt":  true,
		"exif": true,
		"gif":  true,
		"gpl":  true,
		"grf":  true,
		"icns": true,
		"ico":  true,
		"iff":  true,
		"jng":  true,
		"jpeg": true,
		"jpg":  true,
		"jfif": true,
		"jp2":  true,
		"jps":  true,
		"lbm":  true,
		"max":  true,
		"miff": true,
		"mng":  true,
		"msp":  true,
		"nitf": true,
		"ota":  true,
		"pbm":  true,
		"pc1":  true,
		"pc2":  true,
		"pc3":  true,
		"pcf":  true,
		"pcx":  true,
		"pdn":  true,
		"pgm":  true,
		"PI1":  true,
		"PI2":  true,
		"PI3":  true,
		"pict": true,
		"pct":  true,
		"pnm":  true,
		"pns":  true,
		"ppm":  true,
		"psb":  true,
		"psd":  true,
		"pdd":  true,
		"psp":  true,
		"px":   true,
		"pxm":  true,
		"pxr":  true,
		"qfx":  true,
		"raw":  true,
		"rle":  true,
		"sct":  true,
		"sgi":  true,
		"rgb":  true,
		"int":  true,
		"bw":   true,
		"tga":  true,
		"tiff": true,
		"tif":  true,
		"vtf":  true,
		"xbm":  true,
		"xcf":  true,
		"xpm":  true,
		"3dv":  true,
		"amf":  true,
		"ai":   true,
		"awg":  true,
		"cgm":  true,
		"cdr":  true,
		"cmx":  true,
		"dxf":  true,
		"e2d":  true,
		"eps":  true,
		"fs":   true,
		"gbr":  true,
		"odg":  true,
		"svg":  true,
		"stl":  true,
		"vrml": true,
		"x3d":  true,
		"sxd":  true,
		"v2d":  true,
		"vnd":  true,
		"wmf":  true,
		"emf":  true,
		"xar":  true,
		"png":  true,
		"webp": true,
		"jxr":  true,
		"hdp":  true,
		"wdp":  true,
		"cur":  true,
		"ecw":  true,
		"liff": true,
		"nrrd": true,
		"pam":  true,
		"pgf":  true,
		"rgba": true,
		"inta": true,
		"sid":  true,
		"ras":  true,
		"sun":  true,
	}
}
