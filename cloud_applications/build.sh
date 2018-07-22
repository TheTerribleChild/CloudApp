#!/bin/bash

echo "Building cloud applications"

apps=("novel_application")

for app in ${apps[@]}
do
    bin_app=$bin_cloud_application/$app
    src_app=$src_cloud_applications/$app
    mkdir -p $bin_app
    . ./$src_app/build.sh
done

