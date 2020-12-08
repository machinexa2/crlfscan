# CRLFscan
<!-- ![CRLFI](lib/CRLFI.png) -->
<!-- [![Code Grade](https://www.code-inspector.com/project/15089/status/svg)](https://frontend.code-inspector.com/public/project/15089/CRLFi/dashboard)
[![Code Quality](https://www.code-inspector.com/project/15089/score/svg)](https://frontend.code-inspector.com/public/project/15089/CRLFi/dashboard)
-->
## Description
Fast and Automated CRLF injection scanner :zap:. Scans for CRLF injection in parameters, paths and domains. This tool is reimplementation of my tool previously known as CRLFi and was developed in python. Due to concurrency issues, I reimplemented it. The source code of this tool is partially similar to that of CRLFi, and its available at archive.  

This tool is quite different than other crlf injection tools because it finds crlf issues in paths using special path generator, and one by one parameters replacer, and has support for weird cases of CRLF injection. Not only that, it replaces parameters one by one rather than all at a time. This cost more number of requests and I prefer accuracy over speed. Thus, a skipper that checks for same paths may increase speed, but it will take time to write in Go.

## Features
1. Automatic and concurrent scanning of CRLF injections :skull:.  
(Note all features are not implemented, Features list at MoreREADME.md)

## Installation
* ```go get -u -v github.com/machinexa2/crlfscan```

## Usage
```
Usage of crlfscan:
  -o string
        File to write output to (optional)
  -threads int
        No of threads (default 20)
  -url string
        Single url to scan
  -v    Show Verbose output (Default false)
  -version
        Show version of nuclei
  -wordlist string
        List of urls to scan
```

## Example
* ```crlfscan -u google.com```
* ```crlfscan -wordlist testurls```

## Bugs/Issues
1. Most features are implemented, but not all

