#!/bin/bash

set -eu

make build
readonly OUTDIR=testdata/out/tester
mkdir -p $OUTDIR
cd $_

readonly CMD=../../../bin/imgctl
readonly INDIR=../../in

joinpath() {
  echo $@ | tr " " :
}

# generate

$CMD generate $(joinpath $INDIR/{body,eye,eyebrows,mouse}.png)
$CMD generate -d gen $(joinpath $INDIR/{body,eye,eyebrows,mouse}.png)

# scale

$CMD scale -s 50 $INDIR/actor*.png
$CMD scale -s 200 -d sca $INDIR/actor*.png

# crop

$CMD crop -x 20 -y 20 -W 144 -H 144 $INDIR/actor*.png
$CMD crop -d cr -x 20 -y 20 -W 144 -H 144 $INDIR/actor*.png

# flip

$CMD flip $INDIR/actor*.png
$CMD flip -d fl $INDIR/actor*.png

# paste

$CMD paste -r 2 -c 4 -W 144 -H 144 $INDIR/face*.png
$CMD paste -r 2 -c 4 -W 144 -H 144 -d pas $INDIR/face*.png
