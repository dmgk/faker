#!/bin/sh

for f in ../yaml/*.yml; do
  out="../locales/$(basename $f | tr A-Z a-z | sed 's/yml$/go/')"
  
  echo "Generating $out"

  go run generate.go $f |
    sed -r -e 's/map\[interface\s*\{}\]/map\[string\]/g' |
    sed -r -e 's/\[\]interface\s*\{}/\[\]string/g' |
    sed -r -e 's/string\{/string\{\n/g' |
    sed -r -e 's/("\w+":)/\n\1/g' |
    sed -r -e 's/ *$//g' |
    sed -r -e 's/([^\{])},$/\1,\n},/g' |
    sed -r -e 's/([^\{])},$/\1,\n},/g' |
    sed -r -e 's/([^\{])},$/\1,\n},/g' |
    sed -r -e 's/([^\{])}}$/\1,\n},\n}/g' |
    sed -r -e 's/([^\{])}}$/\1,\n},\n}/g' |
    sed -r -e 's/([^\{])}}$/\1,\n},\n}/g' |
    sed -r -e 's/"},/",\n},\n/g' |
    sed -r -e 's/"buzzwords":\[]string/"buzzwords":\[]\[]string/' |
    sed -r -e 's/"bs":\[]string/"bs":\[]\[]string/' |
    sed -e '/^\s*$/d' | gofmt > $out
done

