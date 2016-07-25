#!/usr/bin/env gnuplot
set terminal pngcairo size 800,600 enhanced font "Helvetica, 10"

set title sprintf("Comparation of chance to loose data for different schemas for %s servers.", ARG1)
set output sprintf("experiment_%ssrv_cmp_cl.png", ARG1)
set multiplot
set xrange[1.1:4.9]
set yrange [0:110]
set xlabel "Servers lost" 
set ylabel "%"
set xtics 1 scale 0.1 # scale 0.1 === no tics
set ytics 0, 10, 100
unset key
plot sprintf("experiment_%ssrv_rf2", ARG1) u ($2-0.22):8 with impulses lw 25 lc 3 t "RF2", sprintf("experiment_%ssrv_rf1-2c-rnd", ARG1) u ($2):8 with impulses lw 25 lc 6 t "RF1-RND", sprintf("experiment_%ssrv_rf1-2c", ARG1) u ($2+0.22):8 with impulses lw 25 lc rgb "#00191970" t "RF1", sprintf("experiment_%ssrv_rf2", ARG1) u ($2-0.22):8:(gprintf("%.1f%%", $8)) with labels center offset 0,1 notitle, sprintf("experiment_%ssrv_rf1-2c-rnd", ARG1) u ($2):8:(gprintf("%.1f%%", $8)) with labels center offset 0,1 notitle, sprintf("experiment_%ssrv_rf1-2c", ARG1) u ($2+0.22):8:(gprintf("%.1f%%", $8)) with labels center offset 0,1 notitle

unset title
unset ytics
unset xtics
unset xlabel
unset ylabel
unset border
set key right bmargin

plot [][0:1] 2 t "RF2" lw 10 lc 3, 2 t "RF1-RND" lw 10 lc 6, 2 t "RF1" lw 10 lc rgb "#00191970"
