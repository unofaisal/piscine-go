#! /bin/bash
find . -type f -name "*.sh" | sort -r | rev | cut -d '/' -f1 | rev | cut -d '.' -f1