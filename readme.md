![Go logo](https://go.dev/images/go-logo-blue.svg)

# ImageEdit package
![dino] ![dinoPIX]

![Build Report Card](https://img.shields.io/badge/Build-issues-red)
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
      ImageEdit [args] infile=[path/filename.png] outfile=[path/filename.png] function=[FX | FY | ...] pixels=[int]

Arguments:
      infile      : path to photo to edit
      outfile     : path to save new edited photo
      function   : name of edit function
                    [FX]   [FY]   [RRC]
                    [FXY]  [RRY]  [PIX]
                    [RRX]  [RRR]
      pixels      : number of pixels to edit
      help        : print usage instructions

Example:
      C:/user> ImageEdit infile=./filetoedit.png outfile=./newfilename.png function=RRR pixels=50
~~~
---
### Arguments
| Key | Example Value |
|-|-|
| infile | `./forward/path/to/file` |
| outfile | `../../backward/path/to/file` |
| function| `FX` `FY` `FXY` `RRX` `RRY` `RRR` `RRC` `PIX`|
| *pixels | `33`|
---
**requires integer*
### Usage examples
~~~
C:\path\to\ImageEditfolder> go run main.go infile=./path/picture.png outfile=./path/newpicture.png function=FX

C:\path\to\ImageEditfolder> go run main.go infile=./path/picture.png outfile=./path/newpicture.png function=RRR pixels=33
~~~
---
# Function Examples

  ![dino]

- Flip right

  ![dinoFY]

- Flip down

  ![dinoFX]

- rotate

  ![dinoFXY]

- Round robin y-axis 33px

  ![dinoRRY]

- Round robin x-axis 33px

  ![dinoRRX]

- Round robin columns 3px

  ![dinoRRC]

- Round robin rows 3px

  ![dinoRRR]

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
