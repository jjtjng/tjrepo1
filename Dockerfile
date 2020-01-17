FROM containers.cisco.com/sto-ccc-cloud9/hardened_alpine

COPY maclookuptest_linux /usr/local/bin/maclookuptest

ENTRYPOINT ["/bin/sh"]
