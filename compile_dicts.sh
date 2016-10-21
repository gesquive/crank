#!/bin/bash

DEST="passgen/data.go"

echo "package passgen" > ${DEST}
echo "" >> ${DEST}

for filename in data/*; do
    echo "var $(basename ${filename}) = []string{" >> ${DEST}
    while read w; do
        echo "    \"$w\"," >> ${DEST}
    done < "${filename}"
    echo "}" >> ${DEST}
    echo "" >> ${DEST}
done
