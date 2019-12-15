KNATIVE_PATH='knative'
KNATIVE_VER='v0.11.0'
REGISTRY_URL='yhyddr'

# get images list

echo "collect image to tmp file ..."

cd $KNATIVE_PATH
rm -rf image.tmp

for line in `grep -RI " image: " *.yaml | grep gcr.io`
do 
    echo ${line}
	if [[ ${line} =~ 'gcr.io' ]]
	then
		if [[ ${line} =~ 'gcr.io/knative-releases/knative.dev' ]]
		then 
			sub_line1=${line##gcr.io/knative-releases/github.com/knative/}
			sub_line2=${sub_line1%%@sha*}
			container_name=knative_${sub_line2//\//_}
            echo 1:${container_name}
			echo ${line#image:} ${REGISTRY_URL}/${container_name}:${KNATIVE_VER} >> image.tmp
		else 
			sub_line1=${line#image:}
			sub_line2=${sub_line1#*/}
			sub_line3=${sub_line2%%:*}
			container_name=knative_${sub_line3//\//_}
            echo 2:${container_name}
            
			echo ${line#image:} ${REGISTRY_URL}/${container_name}:${KNATIVE_VER} >> image.tmp;
		fi 
	fi
done

for line in `grep -RI " value: " *.yaml | grep gcr.io`
do 
    echo ${line}
	if [[ ${line} =~ 'gcr.io' ]]
	then
		if [[ ${line} =~ 'gcr.io/knative-releases/knative.dev' ]]
		then 
			sub_line1=${line##gcr.io/knative-releases/github.com/knative/}
			sub_line2=${sub_line1%%@sha*}
			container_name=knative_${sub_line2//\//_}
            echo 3:${container_name}
			echo ${line#value:} ${REGISTRY_URL}/${container_name}:${KNATIVE_VER} >> image.tmp
		else 
			sub_line1=${line#value:}
			sub_line2=${sub_line1#*/}
			sub_line3=${sub_line2%%:*}
			container_name=knative_${sub_line3//\//_}

            echo 4:$container_name
			echo ${line#value:} ${REGISTRY_URL}/${container_name}:${KNATIVE_VER} >> image.tmp;
		fi 
	fi
done