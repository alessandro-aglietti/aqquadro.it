#! /bin/bash

folder=~/Documenti/aqquadro.it/pics-autopub/pics
count_file=~/Documenti/aqquadro.it/pics-autopub/count
count=`cat $count_file`

inotifywait -m -q -e create -r --format '%w%f' $folder | while read file
  do
    notify-send "PICS" "rilevata $file"
    count=$((count+1))
    echo $count > $count_file
    cp "$file" ~/Immagini/$count
    rm "$file"
    ~/Documenti/aqquadro.it/pics-autopub/publish-new-pics.sh
  done
