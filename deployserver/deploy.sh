kill -9 $(pgrep webserver)
cd /www/gowebFirstPage/
git pull https://github.com/Esdeath/gowebFirstPage.git
cd webserver/
./webserver &