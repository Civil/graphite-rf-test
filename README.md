graphite-rf-test: Simulation that helps to decide if it's better to use replication factor = 2 or to have 2 separate copies.
------------------------------------------

This tool runs simulation that helps to decide if you should use Replication Factor 2 or Replication Factor 1 and 2 separate groups of servers (exact copies). It have some assumptions behind it - that you use jump hashing algorithm. As a hash it uses siphash, but result is the same for fnv1a or any other good hashing algorithm.


Usage
-----
~~~~
Usage of ./graphite-rf-test:
  -e int
               experiments per dead server count (default 1000)
  -f string
               path to file with metric names (default "metrics")
  -g           Gnuplot friendly output
  -r           Print raw stats in the end
  -s int
               servers per cluster (default 4)
  -type string
               "rf2", "rf1-2c", "rf1-2c-rnd" (default "rf2")
  -w int
               workers (default 4)
~~~~

Available tests:

1. rf2 --- replication factor 2.
2. rf1-2c --- replication factor 1, same total amount of servers, but 2 separate clusters.
3. rf1-2c-rnd --- replication factor 1, same total amount of servers, 2 separate clusters, 2nd cluster uses different init values for the hash.


Acknowledgement
---------------
This program was originally developed for Booking.com.  With approval
from Booking.com, the code was generalised and published as Open Source
on github, for which the author would like to express his gratitude.

License
-------

This code is licensed under the Apache2 license.
