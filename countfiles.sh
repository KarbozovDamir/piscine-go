
find . -name '*.sh' -type f -print | wc -l | rev | cut -d '.' -f2 | cut -d '/' -f1 | rev