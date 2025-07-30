#! /bin/bash
find . -name "*.sh" | rev | cut -c 4- | rev | cut -d '/' -f 2