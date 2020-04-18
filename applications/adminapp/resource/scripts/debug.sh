#!/bin/sh
dlv --listen=:10001 --headless=true --api-version=2 --accept-multiclient exec ./AdminService