<pre><code>
$ curl -vs 127.0.0.1:12345/cache/testkey -XPUT -dtestvalue
* timeout on name lookup is not supported
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 12345 (#0)
> PUT /cache/testkey HTTP/1.1
> Host: 127.0.0.1:12345
> User-Agent: curl/7.47.1
> Accept: */*
> Content-Length: 9
> Content-Type: application/x-www-form-urlencoded
>
} [9 bytes data]
* upload completely sent off: 9 out of 9 bytes
< HTTP/1.1 200 OK
< Date: Thu, 08 Aug 2019 01:22:51 GMT
< Content-Length: 0
<
* Connection #0 to host 127.0.0.1 left intact


$ curl -s 127.0.0.1:12345/cache/testkey
testvalue

$ curl -s 127.0.0.1:12345/status
{"Count":1,"KeySize":7,"ValueSize":9}

$ curl -s 127.0.0.1:12345/cache/testkey -XDELETE

$ curl -s 127.0.0.1:12345/status
{"count":0,"key_size":0,"value_size":0}

</code></pre>