![Go logo](https://go.dev/images/go-logo-blue.svg)

# ImageEdit package

![Build Report Card](https://img.shields.io/badge/Build-passing-brightgreen)
### Going Forward
- More image functions
---
### Current User Interface
- Run in cmd line with required arguments
~~~
C:\path\to\ImageEditfolder> go run main.go infile=./file/path outfile=./newfile/path ...
~~~
- Builds to .exe; runs, but have to pass arguments
~~~
C:\path\to\ImageEditfolder> go build
C:\path\to\ImageEditfolder> ./ImageEdit infile=./file/path outfile=./newfile/path ...
~~~
~~~
C:\path\to\ImageEditfolder> go build
C:\path\to\ImageEditfolder> ./ImageEdit --help
Usage:
      ImageEdit [args] infile=[path/filename.png] outfile=[path/filename.png] functions=[FX | FY | ...] pixels=[int]

Arguments:
      infile      : path to photo to edit
      outfile     : path to save new edited photo
      functions    : name of edit functions
                    [FX]   [FY]   [RRC]
                    [FXY]  [RRY]
                    [RRX]  [RRR]
      pixels      : number of pixels to edit
      help        : print usage instructions

Example:
      C:\path\to\ImageEditfolder> ImageEdit infile=./filetoedit.png outfile=./newfilename.png functions=RRR pixels=50
~~~
---
### Image Functions
- Flip left right
  - Before

  ![dino]

  - After

  ![dinoFX]

- Flip up down

  ![dinoFY]

- rotate

  ![dinoFXY]

- Round robin left right 200px

  ![dinoRRX]

- Round robin image top down 100px

  ![dinoRRY]

- Round robin 33px columns

  ![dinoRRYcolumns]

- Round robin 33px rows

  ![dinoRRR]

- Pixelate 13px

  ![dinoPIX]

---
### Usage examples
~~~
C:\path\to\ImageEditfolder> go run main.go infile=./path/picture.png outfile=./path/newpicture.png functions=FX

C:\path\to\ImageEditfolder> go run main.go infile=./path/picture.png outfile=./path/newpicture.png functions=RRR pixels=33
  
  * `pixels` only required when using round robin funcs *
~~~
### Arguments
| Key | Example Value |
|-|-|
| *infile | `./forward/path/to/file` |
| *outfile | `../../backward/path/to/file` |
| functions| `FX` `FY` `FXY` `RRX` `RRY` `RRR` `RRC` `PIX`|
| **pixels | `33`|
---
**requires exact file path relative to ImageEdit folder or .exe location*

***requires integer*

[dino]:./assets/dino.png
[dinoFX]:./assets/flip/dinoFX.png
[dinoFY]:./assets/flip/dinoFY.png
[dinoRRX]:./assets/roundrobin/dinoRRX.png
[dinoRRY]:./assets/roundrobin/dinoRRY.png
[dinoRRR]:./assets/roundrobin/dinoRRR.png
[dinoRRYcolumns]:./assets/roundrobin/dinoRRC.png
[dinoFXY]:./assets/flip/dinoFXY.png
[dinoPIX]:./assets/pixelate/dinoPIX.png
