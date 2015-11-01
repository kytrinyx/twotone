# TwoTone

`twotone` is a command-line tool that will take a black/white PNG and generate a new one
with a new background/foreground color.

## Installation

```
go get -u github.com/kytrinyx/twotone
```

## Usage

```
export TAN=EAE0CC
export BROWN=2D232A
twotone -bg=$TAN -fg=$BROWN -in=fixtures/test.png -out=fixtures/out.png
```

## Known Issues

**only handles RGBA**

I only needed to handle simple PNGs. There's a good chance that you'll stumble on PNGs that aren't handled (NRGBA).
That should be easy to fix, but I haven't bothered yet.

**no smoothing of edges**

That probably has a name, but I'm not a graphics person. If I need it I'll probably spend a few minutes figuring out the fix for it.
