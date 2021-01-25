#!/bin/bash
pid=`cat master-test.pid`
echo $pid
kill -INT $pid