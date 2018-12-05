FROM busybox

ADD playground /
ADD /html/*gohtml /html/
ADD /data/*data /data/
CMD ["/playground"]

