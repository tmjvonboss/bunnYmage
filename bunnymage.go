package bunnYmage

import (
	"bytes"
	"errors"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"golang.org/x/image/webp"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"math"
	"os"
	"path/filepath"
	"strings"
)

func ReadImageFromDisk(location string) (*BunnyImage, error) {
	exists, extension, name := ImageExists(location)
	if !exists {
		return nil, errors.New("file doesn't exist")
	}
	fImage, err := os.Open(location)
	if err != nil {
		return nil, err
	}
	bunImage := new(BunnyImage)
	var decodedImage image.Image
	var decodeErr error
	switch extension {
	case ".png":
		bunImage.ImageType = PNG
		decodedImage, decodeErr = png.Decode(fImage)
	case ".gif":
		bunImage.ImageType = GIF
		decodedImage, decodeErr = gif.Decode(fImage)
	case ".bmp":
		bunImage.ImageType = BMP
		decodedImage, decodeErr = bmp.Decode(fImage)
	case ".tiff":
		bunImage.ImageType = TIFF
		decodedImage, decodeErr = tiff.Decode(fImage)
	case ".wepb":
		bunImage.ImageType = WEBP
		decodedImage, decodeErr = webp.Decode(fImage)
	case ".jpeg":
		fallthrough
	case ".jpg":
		fallthrough
	default:
		bunImage.ImageType = JPG
		decodedImage, decodeErr = jpeg.Decode(fImage)
	}
	_ = fImage.Close()
	if decodeErr != nil {
		return nil, decodeErr
	}
	bunImage.Image = decodedImage
	bunImage.FilePath = filepath.Dir(location)
	bunImage.FileName = name
	return bunImage, nil
}

func (bunImage BunnyImage) WriteImageToDisk(location string) error {
	var f *os.File
	var fileErr, writeErr error
	if location == "" {
		location = bunImage.FilePath + bunImage.FileName + bunImage.ImageType.Extension()
	}
	if _, err := os.Stat(location); os.IsNotExist(err) {
		f, fileErr = os.Create(location)
	} else {
		f, fileErr = os.Open(location)
	}
	if fileErr != nil {
		return fileErr
	}
	switch bunImage.ImageType {
	case PNG:
		writeErr = png.Encode(f, bunImage.Image)
	case GIF:
		writeErr = gif.Encode(f, bunImage.Image, nil)
	case BMP:
		writeErr = bmp.Encode(f, bunImage.Image)
	case TIFF:
		writeErr = tiff.Encode(f, bunImage.Image, nil)
	case WEBP:
		writeErr = errors.New(".webp not supported")
	case JPG:
		fallthrough
	default:
		writeErr = jpeg.Encode(f, bunImage.Image, nil)
	}
	_ = f.Close()
	return writeErr
}

func (bunImage BunnyImage) ConvertAndWriteToDisk(location string, toType FileType) error {
	var f *os.File
	var fileErr, writeErr error
	if location == "" {
		location = bunImage.FilePath + bunImage.FileName + toType.Extension()
	}
	if _, err := os.Stat(location); os.IsNotExist(err) {
		f, fileErr = os.Create(location)
	} else {
		f, fileErr = os.Open(location)
	}
	if fileErr != nil {
		return fileErr
	}
	switch toType {
	case PNG:
		writeErr = png.Encode(f, bunImage.Image)
	case GIF:
		writeErr = gif.Encode(f, bunImage.Image, nil)
	case BMP:
		writeErr = bmp.Encode(f, bunImage.Image)
	case TIFF:
		writeErr = tiff.Encode(f, bunImage.Image, nil)
	case WEBP:
		writeErr = errors.New(".webp not supported")
	case JPG:
		fallthrough
	default:
		writeErr = jpeg.Encode(f, bunImage.Image, nil)
	}
	_ = f.Close()
	return writeErr
}

func ReadImageFromByteArray(data []byte, name, extension string) (*BunnyImage, error) {
	extension = strings.ToLower(extension)
	fImage := bytes.NewReader(data)
	bunImage := new(BunnyImage)
	var decodedImage image.Image
	var decodeErr error
	switch extension {
	case ".png":
		bunImage.ImageType = PNG
		decodedImage, decodeErr = png.Decode(fImage)
	case ".gif":
		bunImage.ImageType = GIF
		decodedImage, decodeErr = gif.Decode(fImage)
	case ".bmp":
		bunImage.ImageType = BMP
		decodedImage, decodeErr = bmp.Decode(fImage)
	case ".tiff":
		bunImage.ImageType = TIFF
		decodedImage, decodeErr = tiff.Decode(fImage)
	case ".wepb":
		bunImage.ImageType = WEBP
		decodedImage, decodeErr = webp.Decode(fImage)
	case ".jpeg":
		fallthrough
	case ".jpg":
		fallthrough
	default:
		bunImage.ImageType = JPG
		decodedImage, decodeErr = jpeg.Decode(fImage)
	}
	if decodeErr != nil {
		return nil, decodeErr
	}
	bunImage.Image = decodedImage
	bunImage.FilePath = name
	bunImage.FileName = name
	return bunImage, nil
}

func (bunImage BunnyImage) WriteImageToByteArray() ([]byte, error) {
	var b bytes.Buffer
	var writeErr error
	f := io.Writer(&b)
	switch bunImage.ImageType {
	case PNG:
		writeErr = png.Encode(f, bunImage.Image)
	case GIF:
		writeErr = gif.Encode(f, bunImage.Image, nil)
	case BMP:
		writeErr = bmp.Encode(f, bunImage.Image)
	case TIFF:
		writeErr = tiff.Encode(f, bunImage.Image, nil)
	case WEBP:
		writeErr = errors.New(".webp not supported")
	case JPG:
		fallthrough
	default:
		writeErr = jpeg.Encode(f, bunImage.Image, nil)
	}
	if writeErr == nil {
		return b.Bytes(), writeErr
	}
	return nil, writeErr
}

func (bunImage BunnyImage) ConvertAndWriteToByteArray(toType FileType) ([]byte, error) {
	var b bytes.Buffer
	var writeErr error
	f := io.Writer(&b)
	switch toType {
	case PNG:
		writeErr = png.Encode(f, bunImage.Image)
	case GIF:
		writeErr = gif.Encode(f, bunImage.Image, nil)
	case BMP:
		writeErr = bmp.Encode(f, bunImage.Image)
	case TIFF:
		writeErr = tiff.Encode(f, bunImage.Image, nil)
	case WEBP:
		writeErr = errors.New(".webp not supported")
	case JPG:
		fallthrough
	default:
		writeErr = jpeg.Encode(f, bunImage.Image, nil)
	}
	if writeErr == nil {
		return b.Bytes(), writeErr
	}
	return nil, writeErr
}

func ReadImageFromReader(fImage *io.Reader, name, extension string) (*BunnyImage, error) {
	extension = strings.ToLower(extension)
	bunImage := new(BunnyImage)
	var decodedImage image.Image
	var decodeErr error
	switch extension {
	case ".png":
		bunImage.ImageType = PNG
		decodedImage, decodeErr = png.Decode(*fImage)
	case ".gif":
		bunImage.ImageType = GIF
		decodedImage, decodeErr = gif.Decode(*fImage)
	case ".bmp":
		bunImage.ImageType = BMP
		decodedImage, decodeErr = bmp.Decode(*fImage)
	case ".tiff":
		bunImage.ImageType = TIFF
		decodedImage, decodeErr = tiff.Decode(*fImage)
	case ".wepb":
		bunImage.ImageType = WEBP
		decodedImage, decodeErr = webp.Decode(*fImage)
	case ".jpeg":
		fallthrough
	case ".jpg":
		fallthrough
	default:
		bunImage.ImageType = JPG
		decodedImage, decodeErr = jpeg.Decode(*fImage)
	}
	if decodeErr != nil {
		return nil, decodeErr
	}
	bunImage.Image = decodedImage
	bunImage.FilePath = name
	bunImage.FileName = name
	return bunImage, nil
}

func (bunImage BunnyImage) WriteImageToWriter(f *io.Writer) (*io.Writer, error) {
	var writeErr error
	switch bunImage.ImageType {
	case PNG:
		writeErr = png.Encode(*f, bunImage.Image)
	case GIF:
		writeErr = gif.Encode(*f, bunImage.Image, nil)
	case BMP:
		writeErr = bmp.Encode(*f, bunImage.Image)
	case TIFF:
		writeErr = tiff.Encode(*f, bunImage.Image, nil)
	case WEBP:
		writeErr = errors.New(".webp not supported")
	case JPG:
		fallthrough
	default:
		writeErr = jpeg.Encode(*f, bunImage.Image, nil)
	}
	if writeErr == nil {
		return f, writeErr
	}
	return nil, writeErr
}

func (bunImage BunnyImage) ConvertAndWriteToWriter(f *io.Writer, toType FileType) (*io.Writer, error) {
	var writeErr error
	switch toType {
	case PNG:
		writeErr = png.Encode(*f, bunImage.Image)
	case GIF:
		writeErr = gif.Encode(*f, bunImage.Image, nil)
	case BMP:
		writeErr = bmp.Encode(*f, bunImage.Image)
	case TIFF:
		writeErr = tiff.Encode(*f, bunImage.Image, nil)
	case WEBP:
		writeErr = errors.New(".webp not supported")
	case JPG:
		fallthrough
	default:
		writeErr = jpeg.Encode(*f, bunImage.Image, nil)
	}
	if writeErr == nil {
		return f, writeErr
	}
	return nil, writeErr
}

func GenerateGIFAndWriteToDisk(interval float64, location string, images ...*BunnyImage) error {
	actualInterval := int(math.RoundToEven(100.0 * interval))
	var gifs []BunnyImage
	for _, i := range images {
		if i.ImageType == GIF {
			gifs = append(gifs, *i)
		} else {
			conBytes, err := i.ConvertAndWriteToByteArray(GIF)
			if err == nil {
				conImage, err := ReadImageFromByteArray(conBytes, i.FileName, ".gif")
				if err == nil {
					gifs = append(gifs, *conImage)
				}
			}
		}
	}
	if len(gifs) > 1 {
		outGif := &gif.GIF{}
		for _, inGif := range gifs {
			outGif.Image = append(outGif.Image, inGif.Image.(*image.Paletted))
			outGif.Delay = append(outGif.Delay, actualInterval)
		}
		var f *os.File
		var fileErr error
		if _, err := os.Stat(location); os.IsNotExist(err) {
			f, fileErr = os.Create(location)
		} else {
			f, fileErr = os.Open(location)
		}
		if fileErr != nil {
			return fileErr
		}
		err := gif.EncodeAll(f, outGif)
		_ = f.Close()
		return err
	} else {
		return errors.New("no images to merge into GIF")
	}
}

func GenerateGIFAndWriteToByteArray(interval float64, images ...*BunnyImage) ([]byte, error) {
	actualInterval := int(math.RoundToEven(100.0 * interval))
	var gifs []BunnyImage
	for _, i := range images {
		if i.ImageType == GIF {
			gifs = append(gifs, *i)
		} else {
			conBytes, err := i.ConvertAndWriteToByteArray(GIF)
			if err == nil {
				conImage, err := ReadImageFromByteArray(conBytes, i.FileName, ".gif")
				if err == nil {
					gifs = append(gifs, *conImage)
				}
			}
		}
	}
	if len(gifs) > 1 {
		outGif := &gif.GIF{}
		for _, inGif := range gifs {
			outGif.Image = append(outGif.Image, inGif.Image.(*image.Paletted))
			outGif.Delay = append(outGif.Delay, actualInterval)
		}
		var b bytes.Buffer
		f := io.Writer(&b)
		err := gif.EncodeAll(f, outGif)
		return b.Bytes(), err
	} else {
		return nil, errors.New("no images to merge into GIF")
	}
}

func GenerateGIFAndWriteToWriter(interval float64, f *io.Writer, images ...*BunnyImage) (*io.Writer, error) {
	actualInterval := int(math.RoundToEven(100.0 * interval))
	var gifs []BunnyImage
	for _, i := range images {
		if i.ImageType == GIF {
			gifs = append(gifs, *i)
		} else {
			conBytes, err := i.ConvertAndWriteToByteArray(GIF)
			if err == nil {
				conImage, err := ReadImageFromByteArray(conBytes, i.FileName, ".gif")
				if err == nil {
					gifs = append(gifs, *conImage)
				}
			}
		}
	}
	if len(gifs) > 1 {
		outGif := &gif.GIF{}
		for _, inGif := range gifs {
			outGif.Image = append(outGif.Image, inGif.Image.(*image.Paletted))
			outGif.Delay = append(outGif.Delay, actualInterval)
		}
		err := gif.EncodeAll(*f, outGif)
		return f, err
	} else {
		return nil, errors.New("no images to merge into GIF")
	}
}