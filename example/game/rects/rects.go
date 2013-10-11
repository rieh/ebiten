// Copyright 2013 Hajime Hoshi
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package rects

import (
	"github.com/hajimehoshi/go.ebiten"
	"github.com/hajimehoshi/go.ebiten/graphics"
	"github.com/hajimehoshi/go.ebiten/graphics/matrix"
	"image/color"
	"math"
	"math/rand"
	"time"
)

type Rects struct {
	rectTexture       graphics.RenderTarget
	rectTextureInited bool
	offscreen         graphics.RenderTarget
	offscreenInited   bool
	rectBounds        *graphics.Rect
	rectColor         *color.RGBA
}

func New() *Rects {
	return &Rects{
		rectTextureInited: false,
		offscreenInited:   false,
		rectBounds:        &graphics.Rect{},
		rectColor:         &color.RGBA{},
	}
}

func (game *Rects) Init(tf graphics.TextureFactory) {
	game.rectTexture = tf.NewRenderTarget(16, 16)
	game.offscreen = tf.NewRenderTarget(256, 240)
}

func (game *Rects) Update(context ebiten.GameContext) {
	game.rectBounds.X = rand.Intn(context.ScreenWidth())
	game.rectBounds.Y = rand.Intn(context.ScreenHeight())
	game.rectBounds.Width =
		rand.Intn(context.ScreenWidth() - game.rectBounds.X)
	game.rectBounds.Height =
		rand.Intn(context.ScreenHeight() - game.rectBounds.Y)

	game.rectColor.R = uint8(rand.Intn(256))
	game.rectColor.G = uint8(rand.Intn(256))
	game.rectColor.B = uint8(rand.Intn(256))
	game.rectColor.A = uint8(rand.Intn(256))
}

func (game *Rects) rectGeometryMatrix() matrix.Geometry {
	geometryMatrix := matrix.IdentityGeometry()
	scaleX := float64(game.rectBounds.Width) /
		float64(game.rectTexture.Width())
	scaleY := float64(game.rectBounds.Height) /
		float64(game.rectTexture.Height())
	geometryMatrix.Scale(scaleX, scaleY)
	geometryMatrix.Translate(
		float64(game.rectBounds.X), float64(game.rectBounds.Y))
	return geometryMatrix
}

func (game *Rects) rectColorMatrix() matrix.Color {
	colorMatrix := matrix.IdentityColor()
	colorMatrix.Elements[0][0] =
		float64(game.rectColor.R) / float64(math.MaxUint8)
	colorMatrix.Elements[1][1] =
		float64(game.rectColor.G) / float64(math.MaxUint8)
	colorMatrix.Elements[2][2] =
		float64(game.rectColor.B) / float64(math.MaxUint8)
	colorMatrix.Elements[3][3] =
		float64(game.rectColor.A) / float64(math.MaxUint8)
	return colorMatrix
}

func (game *Rects) Draw(g graphics.Context) {
	if !game.rectTextureInited {
		g.SetOffscreen(game.rectTexture.ID())
		g.Fill(255, 255, 255)
		game.rectTextureInited = true
	}

	g.SetOffscreen(game.offscreen.ID())
	if !game.offscreenInited {
		g.Fill(0, 0, 0)
		game.offscreenInited = true
	}
	g.DrawTexture(game.rectTexture.Texture().ID(),
		game.rectGeometryMatrix(),
		game.rectColorMatrix())

	g.SetOffscreen(g.Screen().ID())
	g.DrawTexture(game.offscreen.Texture().ID(),
		matrix.IdentityGeometry(),
		matrix.IdentityColor())
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
