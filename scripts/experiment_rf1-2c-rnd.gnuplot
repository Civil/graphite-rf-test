#!/usr/bin/env gnuplot
set terminal pngcairo size 800,600 enhanced font "Helvetica, 10"

set title sprintf("Probability of lossing data and percentage data loss in worst case for %s servers. RF=1, 2 clusters, randomized", ARG1)
set output sprintf("experiment_%ssrv_rf1-2c-rnd.png", ARG1)
set multiplot
set xrange[1.1:4.9]
set yrange [0:110]
set xlabel "Servers lost" 
set ylabel "%"
set xtics 1 scale 0.1 # scale 0.1 === no tics
set ytics 0, 10, 100
unset key
plot sprintf("experiment_%ssrv_rf1-2c-rnd", ARG1) u ($2-0.18):6 with impulses lw 60 lc 3 t "Max Loss", "" u ($2+0.18):8 with impulses lw 60 lc 6 t "Chance To Loose Data", "" u ($2-0.18):6:(gprintf("%.2f%%", $6)) with labels center offset 0,1 notitle, "" u ($2+0.18):8:(gprintf("%.2f%%", $8)) with labels center offset 0,1 notitle

unset title
unset ytics
unset xtics
unset xlabel
unset ylabel
unset border
set key right bmargin

plot [][0:1] 2 t "Max Loss" lw 10 lc 3, 2 t "Chance To Loose Data" lw 10 lc 6
