#!/bin/bash
for srv in 4 6 8 10 16; do
	for type in rf1-2c rf1-2c-rnd rf2 cmp_al cmp_cl; do 
		gnuplot -c ./experiment_${type}.gnuplot ${srv}
	done
done
