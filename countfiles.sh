
find . -name '*.sh' -print | wc -l | rev | cut -d '.' -f2 | cut -d '/' -f1 | rev