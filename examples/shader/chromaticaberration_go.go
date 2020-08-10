// Code generated by file2byteslice. DO NOT EDIT.
// (gofmt is fine after generating)

package main

var chromaticaberration_go = []byte("// Copyright 2020 The Ebiten Authors\r\n//\r\n// Licensed under the Apache License, Version 2.0 (the \"License\");\r\n// you may not use this file except in compliance with the License.\r\n// You may obtain a copy of the License at\r\n//\r\n//     http://www.apache.org/licenses/LICENSE-2.0\r\n//\r\n// Unless required by applicable law or agreed to in writing, software\r\n// distributed under the License is distributed on an \"AS IS\" BASIS,\r\n// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\r\n// See the License for the specific language governing permissions and\r\n// limitations under the License.\r\n\r\n// +build ignore\r\n\r\npackage main\r\n\r\nvar Time float\r\nvar Cursor vec2\r\nvar ScreenSize vec2\r\n\r\nfunc Fragment(position vec4, texCoord vec2, color vec4) vec4 {\r\n\tcenter := ScreenSize / 2\r\n\tamount := normalize(center-Cursor).x / 100\r\n\tvar clr vec3\r\n\tclr.r = texture2At(vec2(texCoord.x+amount, texCoord.y)).r\r\n\tclr.g = texture2At(texCoord).g\r\n\tclr.b = texture2At(vec2(texCoord.x-amount, texCoord.y)).b\r\n\treturn vec4(clr, 1.0)\r\n}\r\n")
