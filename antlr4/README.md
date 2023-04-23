Based on:
https://github.com/antlr/antlr4/tree/master/docker

Fix errors:
git config --global --unset http.proxy
git config --global --unset https.proxy


Build image:
chmod +x buildDockerImage.sh

Use image:
docker run --rm -u $(id -u ${USER}):$(id -g ${USER}) -v `pwd`:/work antlr/antlr4 -Dlanguage=Go JSON.g4