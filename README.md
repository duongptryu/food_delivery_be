# food_delivery_be
#This project is backend for app food delivery
#Apply architecture distributed system
#Easy to convert to microservice

# Câu thần chú
 - **Nginx proxy**: docker run -d -p 80:80 -p 443:443 --net my-net --name nginx-proxy -e ENABLE_IPV6=true -v ~/nginx-certs:/etc/nginx/certs:ro \
   -v ~/nginx/vhost.d:/etc/nginx/vhost.d \
   -v ~/nginx-logs:/var/log/nginx \
   -v /usr/share/nginx/html \
   -v /var/run/docker.sock:/tmp/docker.sock:ro \
   --privileged=true \
   jwilder/nginx-proxy

https://github.com/nginx-proxy/nginx-proxy

 - **Let encrypt:** docker run -d \
   --net my-net \
   -v ~/nginx-certs:/etc/nginx/certs:rw \
   --volumes-from nginx-proxy \
   -v /var/run/docker.sock:/var/run/docker.sock:ro \
   --privileged=true \
   jrcs/letsencrypt-nginx-proxy-companion
 
https://github.com/jwilder/docker-letsencrypt-nginx-proxy-companion

 - **mysql:** docker run -d -p 3306:3306 --name mysql -network my-net -e MYSQL_ROOT_PASSWORD=root mysql

 - **Database string:** root:root@tcp(localhost:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local
# Database
 - https://gist.github.com/viettranx/660e983c727c5be00aed3b67ccec8714