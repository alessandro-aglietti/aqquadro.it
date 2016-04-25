#! /bin/bash

folder=~/Documenti/aqquadro.it/pics-autopub/pics
count_file=~/Documenti/aqquadro.it/pics-autopub/count
count=`cat $count_file`

inotifywait -m -q -e create -r --format '%w%f' $folder | while read file
  do
    count=$((count+1))
    echo $count > $count_file
    cp "$file" ~/Immagini/$count
    rm "$file"
  done
