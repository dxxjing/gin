#!/bin/bash
pid=`cat master-test.pid`
echo $pid
kill -USR2 $pid