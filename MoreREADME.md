## To-Do List
* Colored printing implementation: :heavy_check_mark: 
* Output implementation: :heavy_check_mark:
* Skipping implementation (Optional)
* Shortening and code refactoration: :heavy_check_mark:
* List of payloads in separate files: :heavy_check_mark:
* http+https with `www` prefix for more edge casese

## FAQ
1. Does CRLF injection only affect HTTP/1? 
 
* No, I found CRLF injection on two different http/2 enabled website with this previous tool CRLFi, however it was slow. This is fast. Also checkout this: [CRLF injection in HTTP2](https://security.stackexchange.com/questions/235046/does-http-2-prevent-security-vulnerabilites-like-crlf-injection)
