package main

import (
	"io"
	"strings"
	"text/template"
)

var pacTmpl = template.Must(template.New("pac").Funcs(template.FuncMap{
	"escape": jsEscape,
}).Parse(`var FindProxyForURL = function(init, profiles) {
    return function(url, host) {
        "use strict";
        var result = init, scheme = url.substr(0, url.indexOf(":"));
        do {
            result = profiles[result];
            if (typeof result === "function") result = result(url, host, scheme);
        } while (typeof result !== "string" || result.charCodeAt(0) === 43);
        return result;
    };
}("+auto switch", {
    "+auto switch": function(url, host, scheme) {
        "use strict";{{ range $d := . }}
        if (/(?:^|\.){{$d|escape}}$/.test(host)) return "+fod";{{end}}
        return "DIRECT";
    },
    "+fod": function(url, host, scheme) {
        "use strict";
        if (/^127\.0\.0\.1$/.test(host) || /^::1$/.test(host) || /^localhost$/.test(host)) return "DIRECT";
        return "PROXY fodev.org:8118";
    }
});
`))

func jsEscape(in string) string {
	return strings.Replace(in[1:], ".", "\\.", -1)
}

func pacFile(target io.Writer, domains ...string) error {
	return pacTmpl.Execute(target, domains)
}
