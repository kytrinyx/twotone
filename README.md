# TwoTone

`twotone` is a command-line tool that will take a PNG and generate a flattened one
with the given background and foreground colors.

By default it will leave background transparent.

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

![](https://raw.githubusercontent.com/kytrinyx/twotone/master/twotone.png)

## Known Issues

**no anti-aliasing**

If I need it I'll probably spend a few minutes (OK, hours. Whatevs!) figuring out the fix for it. So far I've been lucky: when I resize it in Keynote.app, it smoothes the edges for me.
