// THIS FILE IS GENERATED. SEE ./scripts/gen_stdlib_map.sh
// Generated by: go version go1.21.1 darwin/arm64
package gomod

// isStandardlibPackge determines whether a package is in the standard library
// or not. At this point, it checks whether the package name is one of those
// that is found from running "go list std" in the latest released go version.
func isStandardlibPackge(pkg string) bool {
	_, ok := standardLibraryMap[pkg]
	return ok
}

var contained = struct{}{}

// This list is calculated from "go list std".
var standardLibraryMap = map[string]interface{}{
	"archive/tar":                         contained,
	"archive/zip":                         contained,
	"bufio":                               contained,
	"bytes":                               contained,
	"cmp":                                 contained,
	"compress/bzip2":                      contained,
	"compress/flate":                      contained,
	"compress/gzip":                       contained,
	"compress/lzw":                        contained,
	"compress/zlib":                       contained,
	"container/heap":                      contained,
	"container/list":                      contained,
	"container/ring":                      contained,
	"context":                             contained,
	"crypto":                              contained,
	"crypto/aes":                          contained,
	"crypto/cipher":                       contained,
	"crypto/des":                          contained,
	"crypto/dsa":                          contained,
	"crypto/ecdh":                         contained,
	"crypto/ecdsa":                        contained,
	"crypto/ed25519":                      contained,
	"crypto/elliptic":                     contained,
	"crypto/hmac":                         contained,
	"crypto/internal/alias":               contained,
	"crypto/internal/bigmod":              contained,
	"crypto/internal/boring":              contained,
	"crypto/internal/boring/bbig":         contained,
	"crypto/internal/boring/bcache":       contained,
	"crypto/internal/boring/sig":          contained,
	"crypto/internal/edwards25519":        contained,
	"crypto/internal/edwards25519/field":  contained,
	"crypto/internal/nistec":              contained,
	"crypto/internal/nistec/fiat":         contained,
	"crypto/internal/randutil":            contained,
	"crypto/md5":                          contained,
	"crypto/rand":                         contained,
	"crypto/rc4":                          contained,
	"crypto/rsa":                          contained,
	"crypto/sha1":                         contained,
	"crypto/sha256":                       contained,
	"crypto/sha512":                       contained,
	"crypto/subtle":                       contained,
	"crypto/tls":                          contained,
	"crypto/x509":                         contained,
	"crypto/x509/internal/macos":          contained,
	"crypto/x509/pkix":                    contained,
	"database/sql":                        contained,
	"database/sql/driver":                 contained,
	"debug/buildinfo":                     contained,
	"debug/dwarf":                         contained,
	"debug/elf":                           contained,
	"debug/gosym":                         contained,
	"debug/macho":                         contained,
	"debug/pe":                            contained,
	"debug/plan9obj":                      contained,
	"embed":                               contained,
	"embed/internal/embedtest":            contained,
	"encoding":                            contained,
	"encoding/ascii85":                    contained,
	"encoding/asn1":                       contained,
	"encoding/base32":                     contained,
	"encoding/base64":                     contained,
	"encoding/binary":                     contained,
	"encoding/csv":                        contained,
	"encoding/gob":                        contained,
	"encoding/hex":                        contained,
	"encoding/json":                       contained,
	"encoding/pem":                        contained,
	"encoding/xml":                        contained,
	"errors":                              contained,
	"expvar":                              contained,
	"flag":                                contained,
	"fmt":                                 contained,
	"go/ast":                              contained,
	"go/build":                            contained,
	"go/build/constraint":                 contained,
	"go/constant":                         contained,
	"go/doc":                              contained,
	"go/doc/comment":                      contained,
	"go/format":                           contained,
	"go/importer":                         contained,
	"go/internal/gccgoimporter":           contained,
	"go/internal/gcimporter":              contained,
	"go/internal/srcimporter":             contained,
	"go/internal/typeparams":              contained,
	"go/parser":                           contained,
	"go/printer":                          contained,
	"go/scanner":                          contained,
	"go/token":                            contained,
	"go/types":                            contained,
	"hash":                                contained,
	"hash/adler32":                        contained,
	"hash/crc32":                          contained,
	"hash/crc64":                          contained,
	"hash/fnv":                            contained,
	"hash/maphash":                        contained,
	"html":                                contained,
	"html/template":                       contained,
	"image":                               contained,
	"image/color":                         contained,
	"image/color/palette":                 contained,
	"image/draw":                          contained,
	"image/gif":                           contained,
	"image/internal/imageutil":            contained,
	"image/jpeg":                          contained,
	"image/png":                           contained,
	"index/suffixarray":                   contained,
	"internal/abi":                        contained,
	"internal/bisect":                     contained,
	"internal/buildcfg":                   contained,
	"internal/bytealg":                    contained,
	"internal/cfg":                        contained,
	"internal/coverage":                   contained,
	"internal/coverage/calloc":            contained,
	"internal/coverage/cformat":           contained,
	"internal/coverage/cmerge":            contained,
	"internal/coverage/decodecounter":     contained,
	"internal/coverage/decodemeta":        contained,
	"internal/coverage/encodecounter":     contained,
	"internal/coverage/encodemeta":        contained,
	"internal/coverage/pods":              contained,
	"internal/coverage/rtcov":             contained,
	"internal/coverage/slicereader":       contained,
	"internal/coverage/slicewriter":       contained,
	"internal/coverage/stringtab":         contained,
	"internal/coverage/test":              contained,
	"internal/coverage/uleb128":           contained,
	"internal/cpu":                        contained,
	"internal/dag":                        contained,
	"internal/diff":                       contained,
	"internal/fmtsort":                    contained,
	"internal/fuzz":                       contained,
	"internal/goarch":                     contained,
	"internal/godebug":                    contained,
	"internal/godebugs":                   contained,
	"internal/goexperiment":               contained,
	"internal/goos":                       contained,
	"internal/goroot":                     contained,
	"internal/goversion":                  contained,
	"internal/intern":                     contained,
	"internal/itoa":                       contained,
	"internal/lazyregexp":                 contained,
	"internal/lazytemplate":               contained,
	"internal/nettrace":                   contained,
	"internal/obscuretestdata":            contained,
	"internal/oserror":                    contained,
	"internal/pkgbits":                    contained,
	"internal/platform":                   contained,
	"internal/poll":                       contained,
	"internal/profile":                    contained,
	"internal/race":                       contained,
	"internal/reflectlite":                contained,
	"internal/safefilepath":               contained,
	"internal/saferio":                    contained,
	"internal/singleflight":               contained,
	"internal/syscall/execenv":            contained,
	"internal/syscall/unix":               contained,
	"internal/sysinfo":                    contained,
	"internal/testenv":                    contained,
	"internal/testlog":                    contained,
	"internal/testpty":                    contained,
	"internal/trace":                      contained,
	"internal/txtar":                      contained,
	"internal/types/errors":               contained,
	"internal/unsafeheader":               contained,
	"internal/xcoff":                      contained,
	"internal/zstd":                       contained,
	"io":                                  contained,
	"io/fs":                               contained,
	"io/ioutil":                           contained,
	"log":                                 contained,
	"log/internal":                        contained,
	"log/slog":                            contained,
	"log/slog/internal":                   contained,
	"log/slog/internal/benchmarks":        contained,
	"log/slog/internal/buffer":            contained,
	"log/slog/internal/slogtest":          contained,
	"log/syslog":                          contained,
	"maps":                                contained,
	"math":                                contained,
	"math/big":                            contained,
	"math/bits":                           contained,
	"math/cmplx":                          contained,
	"math/rand":                           contained,
	"mime":                                contained,
	"mime/multipart":                      contained,
	"mime/quotedprintable":                contained,
	"net":                                 contained,
	"net/http":                            contained,
	"net/http/cgi":                        contained,
	"net/http/cookiejar":                  contained,
	"net/http/fcgi":                       contained,
	"net/http/httptest":                   contained,
	"net/http/httptrace":                  contained,
	"net/http/httputil":                   contained,
	"net/http/internal":                   contained,
	"net/http/internal/ascii":             contained,
	"net/http/internal/testcert":          contained,
	"net/http/pprof":                      contained,
	"net/internal/socktest":               contained,
	"net/mail":                            contained,
	"net/netip":                           contained,
	"net/rpc":                             contained,
	"net/rpc/jsonrpc":                     contained,
	"net/smtp":                            contained,
	"net/textproto":                       contained,
	"net/url":                             contained,
	"os":                                  contained,
	"os/exec":                             contained,
	"os/exec/internal/fdtest":             contained,
	"os/signal":                           contained,
	"os/user":                             contained,
	"path":                                contained,
	"path/filepath":                       contained,
	"plugin":                              contained,
	"reflect":                             contained,
	"reflect/internal/example1":           contained,
	"reflect/internal/example2":           contained,
	"regexp":                              contained,
	"regexp/syntax":                       contained,
	"runtime":                             contained,
	"runtime/cgo":                         contained,
	"runtime/coverage":                    contained,
	"runtime/debug":                       contained,
	"runtime/internal/atomic":             contained,
	"runtime/internal/math":               contained,
	"runtime/internal/sys":                contained,
	"runtime/internal/wasitest":           contained,
	"runtime/metrics":                     contained,
	"runtime/pprof":                       contained,
	"runtime/race":                        contained,
	"runtime/trace":                       contained,
	"slices":                              contained,
	"sort":                                contained,
	"strconv":                             contained,
	"strings":                             contained,
	"sync":                                contained,
	"sync/atomic":                         contained,
	"syscall":                             contained,
	"testing":                             contained,
	"testing/fstest":                      contained,
	"testing/internal/testdeps":           contained,
	"testing/iotest":                      contained,
	"testing/quick":                       contained,
	"testing/slogtest":                    contained,
	"text/scanner":                        contained,
	"text/tabwriter":                      contained,
	"text/template":                       contained,
	"text/template/parse":                 contained,
	"time":                                contained,
	"time/tzdata":                         contained,
	"unicode":                             contained,
	"unicode/utf16":                       contained,
	"unicode/utf8":                        contained,
	"unsafe":                              contained,
	"vendor/golang.org/x/crypto/chacha20": contained,
	"vendor/golang.org/x/crypto/chacha20poly1305":  contained,
	"vendor/golang.org/x/crypto/cryptobyte":        contained,
	"vendor/golang.org/x/crypto/cryptobyte/asn1":   contained,
	"vendor/golang.org/x/crypto/hkdf":              contained,
	"vendor/golang.org/x/crypto/internal/alias":    contained,
	"vendor/golang.org/x/crypto/internal/poly1305": contained,
	"vendor/golang.org/x/net/dns/dnsmessage":       contained,
	"vendor/golang.org/x/net/http/httpguts":        contained,
	"vendor/golang.org/x/net/http/httpproxy":       contained,
	"vendor/golang.org/x/net/http2/hpack":          contained,
	"vendor/golang.org/x/net/idna":                 contained,
	"vendor/golang.org/x/net/nettest":              contained,
	"vendor/golang.org/x/net/route":                contained,
	"vendor/golang.org/x/sys/cpu":                  contained,
	"vendor/golang.org/x/text/secure/bidirule":     contained,
	"vendor/golang.org/x/text/transform":           contained,
	"vendor/golang.org/x/text/unicode/bidi":        contained,
	"vendor/golang.org/x/text/unicode/norm":        contained,
}
