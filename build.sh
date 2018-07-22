#!/bin/bash

#Initiate variables
echo "Building cloud application portal"
bin_root=$GOPATH/bin
bin_cloud_appplication_portal=$bin_root/cloud_application_portal
bin_cloud_application=$bin_cloud_appplication_portal/cloud_applications

src_cloud_applications=cloud_applications
src_go_home=github.com/TheTerribleChild/cloud_application_portal

mkdir -p $bin_cloud_appplication_portal
mkdir -p $bin_cloud_application


. ./$src_cloud_applications/build.sh