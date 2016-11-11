#!/bin/bash

/usr/bin/docker run -it \
  -e DISTIL_MONGO_ADDR="cm1.smartgrid.store:27017" \
  -e DISTIL_BTRDB_ADDR="cm1.smartgrid.store:4410" \
  -e SOURCECODE="github.com/immesys/examples/adder" \
  -e REF_PMU_PATH1="/REFSET/LBNL/a6_bus1/L1MAG" \
  -e REF_PMU_PATH2="/REFSET/LBNL/a6_bus1/L2MAG" \
  btrdb/distiller
