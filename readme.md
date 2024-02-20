![Go logo](https://go.dev/images/go-logo-blue.svg)

# ImageEdit package
![dino] ![dinoPIX]

### Going Forward
- More image functions
---
# Cli User Interface
- Run in cmd line with required arguments
~~~
C:\path\to\ImageEditfolder> go run main.go infile=./file/path outfile=./newfile/path ...
~~~
- Build the ImageEdit.exe, show help instructions
~~~
C:\path\to\ImageEditfolder> go build
C:\path\to\ImageEditfolder> ./ImageEdit --help
Usage:
      ImageEdit [args] infile=[path/filename.png] outfile=[path/filename.png] function=[Flipx | Flipy | ...] pixels=[int]

Arguments:
      infile      : path to photo to edit
      outfile     : path to save new edited photo
      function    : name of edit function
                    [Flipx]        [Flipy]           [Roundrobincolumns]
                    [Rotate]       [Roundrobiny]     [Pixelate]
                    [Roundrobinx]  [Roundrobinrows]  [Rgbfilter]
      pixels      : number of pixels to edit
      help        : print usage instructions

Example:
      C:/user> ImageEdit infile=./filetoedit.png outfile=./newfilename.png function=Roundrobinrows pixels=50
~~~
---
### Arguments
| Key | Example Value |
|-|-|
| infile | `./forward/path/to/file` |
| outfile | `../../backward/path/to/file` |
| function| `Flipx` `Flipy` `Rotate` `Roundrobinx` `Roundrobiny` `Roundrobinrows` `Roundrobincolumns` `Pixelate` `Rgbfilter`|
| pixels | `33`|
---
### Usage examples
~~~
C:\path\to\ImageEditfolder> go run main.go infile=./path/picture.png outfile=./path/newpicture.png function=Flipx pixels=

C:\path\to\ImageEditfolder> go run main.go infile=./path/picture.png outfile=./path/newpicture.png function=Roundrobinrows pixels=33
~~~
---
# Function Examples

  ![dino]

- Flip right

  ![dinoFY]

- Rotate

  ![dinoFXY]

- Round robin y-axis 33px

  ![dinoRRY]

- Round robin columns 3px

  ![dinoRRC]

- Pixelate 3px

  ![dinoPIX]

---

[dino]:./assets/dino.png
[dinoFX]:./assets/flip/dinoFX.png
[dinoFY]:./assets/flip/dinoFY.png
[dinoRRX]:./assets/roundrobin/dinoRRX.png
[dinoRRY]:./assets/roundrobin/dinoRRY.png
[dinoRRR]:./assets/roundrobin/dinoRRR.png
[dinoRRC]:./assets/roundrobin/dinoRRC.png
[dinoFXY]:./assets/flip/dinoFXY.png
[dinoPIX]:./assets/pixelate/dinoPIX.png
