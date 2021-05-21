package gotool

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"os"
)
type Graphic struct {
	Context string
	Name string
	Width int
	Height int
	ToImg bool
}

func GraphicNew (p *Graphic) *Graphic {
	return p
}
func (p *Graphic) SetContext(context string) *Graphic {
	p.Context = context
	return p
}
func (p *Graphic) SetName(name string) *Graphic {
	p.Name = name
	return p
}
func (p *Graphic) SetType(toImg bool) *Graphic {
	p.ToImg = toImg
	return p
}
func (p *Graphic) SetWidth(width int) *Graphic {
	p.Width = width
	return p
}
func (p *Graphic) SetHeight(height int) *Graphic {
	p.Height = height
	return p
}

func (p *Graphic) QR() (string, error) {
	if p.Width == 0 {
		p.Width = 200
	}
	if p.Height == 0 {
		p.Height = 200
	}
	qrCode, err := qr.Encode(p.Context, qr.M, qr.Auto)
	if err != nil {
		return "", err
	}
	// Scale the barcode pixels
	qrCode, err = barcode.Scale(qrCode, p.Width, p.Height)
	if err != nil {
		return "", err
	}
	if p.ToImg {
		// create the output file
		file, _ := os.Create(p.Name)
		defer file.Close()
		// encode the barcode as png
		err := png.Encode(file, qrCode)
		if err != nil {
			return "", err
		}
		return p.Name, nil
	}else{
		buf := new(bytes.Buffer)
		base := base64.NewEncoder(base64.StdEncoding, buf)
		err := png.Encode(base, qrCode)
		if err != nil {
			return "", err
		}
		base.Close()
		return fmt.Sprintf("data:image/png;base64,%s",buf.String()), nil
	}
}

func (p *Graphic) BarCode() (string, error) {
	// 创建一个code128编码的 BarcodeIntCS
	cs, err := code128.Encode(p.Context)
	if err != nil {
		return "", err
	}
	if p.Width == 0 {
		p.Width = 350
	}
	if p.Height == 0 {
		p.Height = 70
	}
	qrCode, err := barcode.Scale(cs, p.Width, p.Height)	// 设置图片像素大小
	if err != nil {
		return "", err
	}
	if p.ToImg {
		// 创建一个要输出数据的文件
		file, _ := os.Create(p.Name)
		defer file.Close()
		err := png.Encode(file,qrCode)	// 将code128的条形码编码为png图片
		if err != nil {
			return "", err
		}
		return p.Name, nil
	}else{
		buf := new(bytes.Buffer)
		base := base64.NewEncoder(base64.StdEncoding, buf)
		err := png.Encode(base, qrCode)
		if err != nil {
			return "", err
		}
		base.Close()
		return fmt.Sprintf("data:image/png;base64,%s",buf.String()), nil
	}
}
