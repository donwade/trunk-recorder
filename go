#!/bin/bash

ask "clean" 5
if [ $? != 0 ]; then
    rm -rf build
fi

if [ ! -d build ]; then
    ./install.sh | tee build.log
fi

    cd build
    make -j4
    if [ $? == 0 ]; then
        set -x
        yes "$PISSWORD" | sudo -S  make install | tee ../install.log
        if [ ${PIPESTATUS[1]} == 0 ]; then
            trunk-recorder --config /home/dwade/my-trunkRecorder/trunk-recorder/config.json
        fi
    fi

