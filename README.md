# k8secretdir - easier secrets encoding

`go get github.com/Yolean/k8secretdir`

Kubernetes [secret definitions](https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/secrets.md) require all data elements to be base64-encoded.

This little tool is meant to simplify writing those definitions. You give it a path, it recursively walks it and outputs all files base64-encoded in YAML syntax.

If you have following directory tree:

    foo/
	    bar
	    baz

And you run `k8secretdir foo`, it will output

      bar: >
        contents of bar in base64
	  baz: >
	    contents of baz in base64

So you only have to write the beginning of service definition (api version, metadata, resource kind) and pipe output of this command at the end.

Note that it uses only base names of files, if you have two files with the same name in different subdirectories it will output both of them with the same name.
