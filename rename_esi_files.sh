#!/bin/bash

# This script is meant to shorten the names of the auto-generated files
# for goesi. Having a large quantity of files with long names means
# Windows users are likely to run into this issue:
# https://github.com/golang/go/issues/18468

# The goal of this script is to shorten the length of the files to help
# avoid running into the 32k character limit of the go compiler on Windows.

for file in ../goesi/esi/*.go; do
    newfilename=$file
    newfilename=${newfilename//model_/m_}
    newfilename=${newfilename//delete_/d_}
    newfilename=${newfilename//get_/g_}
    newfilename=${newfilename//post_/p_}
    newfilename=${newfilename//_character/_char}
    newfilename=${newfilename//_corporation/_corp}
    newfilename=${newfilename//_alliance/_alli}
    newfilename=${newfilename//_universe/_uni}

    mv "$file" "$newfilename"
done