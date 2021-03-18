npm run build
docker build -t zcxt_vue:v1.0.0 ./
docker push zcxt_vue:v1.0.0

# docker build -t registry.cn-shenzhen.aliyuncs.com/cetacean/zcxt_vue:v1.0.0 ./
# docker push registry.cn-shenzhen.aliyuncs.com/cetacean/zcxt_vue:v1.0.0