#! /bin/sh
java -jar ~/swagger2markup-cli-1.3.3.jar convert -i ./swagger.json -f FB2B双光谱筒型摄像机API接口v1.0.13 -c ./config.properties
cat appendix.adoc >> FB2B双光谱筒型摄像机API接口v1.0.13.adoc
asciidoctor -n -b html5 -a icons -a toc2 -a theme=flask -a toc-title=目录 -a toclevels=3  FB2B双光谱筒型摄像机API接口v1.0.13.adoc
