FROM scratch
ADD snmpwalk_server /
CMD ["/snmpwalk_server -r -s 172.18.0.3:8500"]
