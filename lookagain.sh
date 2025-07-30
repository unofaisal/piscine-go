#! /bin/bash
find . -type f -name "*.sh" | rev | cut -c 4- | rev | cut -d '/' -f 2