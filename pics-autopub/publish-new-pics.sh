#! /bin/bash

count_file=~/Documenti/aqquadro.it/pics-autopub/count
count=`cat $count_file`

gsutil rsync ~/Immagini gs://www.aqquadro.it/pics
notify-send "PICS" "rsync per $count"

## tirare giu il template compilato
gsutil cp gs://www.aqquadro.it/index.html /home/elgoog/Documenti/aqquadro.it/pics-autopub/index.html
notify-send "PICS" "download ultima versione index.html"

## sed dentro template compilato
sed -ri "s/(TOT\s=\s[0-9]*);/TOT = $count;/g" /home/elgoog/Documenti/aqquadro.it/pics-autopub/index.html

## upload sovrascrittura template compilato
gsutil cp /home/elgoog/Documenti/aqquadro.it/pics-autopub/index.html gs://www.aqquadro.it/index.html
notify-send "PICS" "upload versione aggiornata index.html"
