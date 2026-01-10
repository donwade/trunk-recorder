#!/bin/bash

if [ ! -d build ]; then
    ./install.sh | tee build.log
else
    cd build
    make -j4
    if [ $? == 0 ]; then
        set -x
        yes "$PISSWORD" | sudo -S  make install
        if [ ${PIPESTATUS[1]} == 0 ]; then
            trunk-recorder --config /home/dwade/my-trunkRecorder/trunk-recorder/config.json
        fi
    fi
fi

