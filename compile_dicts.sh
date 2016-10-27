#!/bin/bash

DEST="passgen/data.go"

DICTS="wordlists/common/american-full
wordlists/common/american-common
wordlists/common/british-full"

echo "package passgen" > ${DEST}
echo "" >> ${DEST}

for filename in $DICTS; do
    base_name=$(basename ${filename})
    var_name=$(echo ${base_name} | sed 's/-//g')
    echo "var ${var_name} = []string{" >> ${DEST}
    while read w; do
        echo "    \"$w\"," >> ${DEST}
    done < "${filename}"
    echo "}" >> ${DEST}
    echo "" >> ${DEST}
done
