KNATIVE_PATH='knative'
KNATIVE_VER='v0.11.0'
REGISTRY_URL='yhyddr'

rm -rf $KNATIVE_PATH
mkdir $KNATIVE_PATH


# download yaml

echo "download yaml files ..."

cd $KNATIVE_PATH

wget -q https://github.com/knative/serving/releases/download/${KNATIVE_VER}/serving.yaml 
wget -q https://github.com/knative/build/releases/download/${KNATIVE_VER}/build.yaml 
wget -q https://github.com/knative/eventing/releases/download/${KNATIVE_VER}/release.yaml 
wget -q https://github.com/knative/eventing-sources/releases/download/${KNATIVE_VER}/eventing-sources.yaml 
wget -q https://github.com/knative/serving/releases/download/${KNATIVE_VER}/monitoring.yaml 
wget -q https://raw.githubusercontent.com/knative/serving/${KNATIVE_VER}/third_party/config/build/clusterrole.yaml

cd ..


# get images list

echo "collect image to tmp file ..."

cd $KNATIVE_PATH
rm -rf image.tmp

for line in `grep -RI " image: " *.yaml | grep gcr.io`
do 
	if [[ ${line} =~ 'gcr.io' ]]
	then
		if [[ ${line} =~ 'gcr.io/knative-releases/knative.dev' ]]
		then 
			sub_line1=${line##gcr.io/knative-releases/github.com/knative/}
			sub_line2=${sub_line1%%@sha*}
			container_name=knative_${sub_line2//\//_}

			echo ${line#image:} ${REGISTRY_URL}/${container_name}:${KNATIVE_VER} >> image.tmp
		else 
			sub_line1=${line#image:}
			sub_line2=${sub_line1#*/}
			sub_line3=${sub_line2%%:*}
			container_name=knative_${sub_line3//\//_}

			echo ${line#image:} ${REGISTRY_URL}/${container_name}:${KNATIVE_VER} >> image.tmp;
		fi 
	fi
done

for line in `grep -RI " value: " *.yaml | grep gcr.io`
do 
	if [[ ${line} =~ 'gcr.io' ]]
	then
		if [[ ${line} =~ 'gcr.io/knative-releases/knative.dev' ]]
		then 
			sub_line1=${line##gcr.io/knative-releases/github.com/knative/}
			sub_line2=${sub_line1%%@sha*}
			container_name=knative_${sub_line2//\//_}

			echo ${line#value:} ${REGISTRY_URL}/${container_name}:${KNATIVE_VER} >> image.tmp
		else 
			sub_line1=${line#value:}
			sub_line2=${sub_line1#*/}
			sub_line3=${sub_line2%%:*}
			container_name=knative_${sub_line3//\//_}

			echo ${line#value:} ${REGISTRY_URL}/${container_name}:${KNATIVE_VER} >> image.tmp;
		fi 
	fi
done

cd ..


# download image, tag, push

cd $KNATIVE_PATH

while read line 
do 
	origin_image=`echo ${line} | awk '{print $1}'`
	new_image=`echo ${line} | awk '{print $2}'`

    echo "old:" ${origin_image}
    echo "new:" ${new_image}
	docker pull ${origin_image}
	docker tag ${origin_image} ${new_image}
	docker push ${new_image}

done < image.tmp

cd ..

# done

echo "completed..."