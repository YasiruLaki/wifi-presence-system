#!/bin/bash

sudo wdutil info | grep "RSSI" | awk '{print $2}'