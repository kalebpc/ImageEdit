![Go logo](https://go.dev/images/go-logo-blue.svg)

# ImageEdit package
![dino] ![dinoPIX]

![Build Report Card](https://img.shields.io/badge/Build-compiles-brightgreen)
### Going Forward
- More image functions
---
# Cli User Interface: no args
Build the ImageEdit.exe binary
~~~
C:\path\to\ImageEditfolder> go build
~~~
Run ImageEdit
~~~
C:\path\to\ImageEditfolder> ./ImageEdit
~~~
File picker will open folder to /%USERPROFILE%/pictures

NOTE: if pictures folder is not in this path, change path in `ui.go` file
~~~
Pick a file:

> drwxrwxrwx      foo.png
> drwxrwxrwx      bar.png
~~~
Use arrows to highlight image file

Hit `enter`

Hit `q`

Type filename with extension

~~~
  You selected: C:\Users\%USERPROFILE%\pictures\foo.png

Enter name of Outfile:          fooPIX.png
~~~
Hit `enter`

Type name of function
~~~
  You selected: C:\Users\%USERPROFILE%\pictures\foo.png

Enter name of Outfile:          fooPIX.png
Enter name of function to run:  PIX
~~~
Hit `enter`

Type number of pixels
~~~
  You selected: C:\Users\%USERPROFILE%\pictures\foo.png

Enter name of Outfile:          fooPIX.png
Enter name of function to run:  PIX
Enter number of pixels:         15
~~~
Hit `enter`
~~~
  You selected: C:\Users\%USERPROFILE%\pictures\foo.png

Enter name of Outfile:          fooPIX.png
Enter name of function to run:  PIX
Enter number of pixels:         15
New Image Created!
~~~
Look in %USERPROFILE%\pictures folder to find newly pixelated 'fooPIX.png' image

---
# Cli User Interface: passing args
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
      ImageEdit [args] infile=[path/filename.png] outfile=[path/filename.png] function=[FX | FY | ...] pixels=[int]

Arguments:
      infile      : path to photo to edit
      outfile     : path to save new edited photo
      function    : name of edit function
                    [FX]   [FY]   [RRC]
                    [FXY]  [RRY]
                    [RRX]  [RRR]
      pixels      : number of pixels to edit
      help        : print usage instructions

Example:
      C:\path\to\ImageEditfolder> ImageEdit infile=./filetoedit.png outfile=./newfilename.png function=RRR pixels=50
~~~
---
### Arguments
| Key | Example Value |
|-|-|
| *infile | `./forward/path/to/file` |
| *outfile | `../../backward/path/to/file` |
| function| `FX` `FY` `FXY` `RRX` `RRY` `RRR` `RRC` `PIX`|
| **pixels | `33`|
---
**requires exact file path relative to ImageEdit folder or .exe location*

***requires integer*
### Usage examples
~~~
C:\path\to\ImageEditfolder> go run main.go infile=./path/picture.png outfile=./path/newpicture.png function=FX

C:\path\to\ImageEditfolder> go run main.go infile=./path/picture.png outfile=./path/newpicture.png function=RRR pixels=33
  
  * `pixels` not required for flip function *
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
