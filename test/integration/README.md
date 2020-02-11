# Exercise tests

This directory contains additional test suites beyond the unit tests already. Whereas the unit tests run very quickly (since they don't connect to db), the tests in this directory are only run manually. 

Make sure mongo db is running locally.

The test packages are:

## integration
These tests insert and fetch records to mongo db connected locally. 