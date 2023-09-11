// Code generated by serde. DO NOT EDIT.

//go:build durable

package coroutine

import reflect "reflect"
import strings "strings"
import syscall "syscall"
import constant "go/constant"
import semver "golang.org/x/mod/semver"
import packages "golang.org/x/tools/go/packages"
import token "go/token"
import sort "sort"
import exec "os/exec"
import scanner "go/scanner"
import astutil "golang.org/x/tools/go/ast/astutil"
import doc "go/doc"
import runtime "runtime"
import slog "log/slog"
import strconv "strconv"
import os "os"
import unicode "unicode"
import regexp "regexp"
import serde "github.com/stealthrocket/coroutine/serde"
import types "go/types"
import log "log"
import crypto "crypto"
import typeutil "golang.org/x/tools/go/types/typeutil"
import comment "go/doc/comment"
import ast "go/ast"
import sync "sync"
import io "io"
import bytes "bytes"
import rand "math/rand"
import syntax "regexp/syntax"
import atomic "sync/atomic"
import time "time"
import fs "io/fs"
import big "math/big"
import bufio "bufio"
import build "go/build"
import unsafe "unsafe"
import base64 "encoding/base64"
import constraint "go/build/constraint"
import json "encoding/json"
import parser "go/parser"
import objectpath "golang.org/x/tools/go/types/objectpath"

func init() {
	serde.RegisterType[runtime.Frames]()
	serde.RegisterType[strconv.NumError]()
	serde.RegisterType[json.UnmarshalFieldError]()
	serde.RegisterType[types.Func]()
	serde.RegisterType[syscall.Flock_t]()
	serde.RegisterType[os.ProcessState]()
	serde.RegisterType[reflect.MapIter]()
	serde.RegisterType[syscall.SockaddrUnix]()
	serde.RegisterType[syscall.IPMreqn]()
	serde.RegisterType[syscall.Timeval]()
	serde.RegisterType[syscall.TCPInfo]()
	serde.RegisterType[syscall.Sysinfo_t]()
	serde.RegisterType[scanner.Error]()
	serde.RegisterType[ast.Comment]()
	serde.RegisterType[doc.Func]()
	serde.RegisterType[atomic.Uint32]()
	serde.RegisterType[atomic.Int64]()
	serde.RegisterType[syscall.Time_t]()
	serde.RegisterType[syntax.Error]()
	serde.RegisterType[uint8]()
	serde.RegisterType[time.Weekday]()
	serde.RegisterType[Stack]()
	serde.RegisterType[int16]()
	serde.RegisterType[syscall.Signal]()
	serde.RegisterType[syscall.ICMPv6Filter]()
	serde.RegisterType[syscall.WaitStatus]()
	serde.RegisterType[syscall.RawSockaddrNetlink]()
	serde.RegisterType[syscall.Ucred]()
	serde.RegisterType[types.ArgumentError]()
	serde.RegisterType[reflect.StructField]()
	serde.RegisterType[time.Timer]()
	serde.RegisterType[syscall.RawSockaddrInet6]()
	serde.RegisterType[syscall.NetlinkRouteRequest]()
	serde.RegisterType[syscall.SockFprog]()
	serde.RegisterType[ast.File]()
	serde.RegisterType[typeutil.Hasher]()
	serde.RegisterType[sync.Pool]()
	serde.RegisterType[slog.Logger]()
	serde.RegisterType[syscall.Timespec]()
	serde.RegisterType[base64.Encoding]()
	serde.RegisterType[big.RoundingMode]()
	serde.RegisterType[build.Package]()
	serde.RegisterType[syscall.FdSet]()
	serde.RegisterType[syscall.NlMsgerr]()
	serde.RegisterType[unicode.CaseRange]()
	serde.RegisterType[ast.StarExpr]()
	serde.RegisterType[build.Directive]()
	serde.RegisterType[strings.Reader]()
	serde.RegisterType[types.Slice]()
	serde.RegisterType[packages.Package]()
	serde.RegisterType[serde.ID]()
	serde.RegisterType[atomic.Uint64]()
	serde.RegisterType[syscall.Termios]()
	serde.RegisterType[unicode.SpecialCase]()
	serde.RegisterType[json.Number]()
	serde.RegisterType[ast.IncDecStmt]()
	serde.RegisterType[ast.IfStmt]()
	serde.RegisterType[big.Rat]()
	serde.RegisterType[types.Interface]()
	serde.RegisterType[slog.JSONHandler]()
	serde.RegisterType[exec.Cmd]()
	serde.RegisterType[ast.GenDecl]()
	serde.RegisterType[build.ImportMode]()
	serde.RegisterType[syntax.InstOp]()
	serde.RegisterType[int32]()
	serde.RegisterType[fs.PathError]()
	serde.RegisterType[token.File]()
	serde.RegisterType[types.PkgName]()
	serde.RegisterType[types.BasicInfo]()
	serde.RegisterType[ast.Ellipsis]()
	serde.RegisterType[bytes.Reader]()
	serde.RegisterType[json.InvalidUnmarshalError]()
	serde.RegisterType[ast.Ident]()
	serde.RegisterType[uint]()
	serde.RegisterType[io.LimitedReader]()
	serde.RegisterType[slog.LevelVar]()
	serde.RegisterType[syscall.SockaddrInet4]()
	serde.RegisterType[syscall.SockaddrLinklayer]()
	serde.RegisterType[ast.TypeSpec]()
	serde.RegisterType[types.ChanDir]()
	serde.RegisterType[ast.IndexExpr]()
	serde.RegisterType[constraint.TagExpr]()
	serde.RegisterType[build.Context]()
	serde.RegisterType[crypto.Hash]()
	serde.RegisterType[sync.Cond]()
	serde.RegisterType[os.File]()
	serde.RegisterType[ast.Field]()
	serde.RegisterType[ast.BadDecl]()
	serde.RegisterType[runtime.Pinner]()
	serde.RegisterType[runtime.MemProfileRecord]()
	serde.RegisterType[slog.Value]()
	serde.RegisterType[os.Process]()
	serde.RegisterType[types.Struct]()
	serde.RegisterType[doc.Note]()
	serde.RegisterType[syscall.PtraceRegs]()
	serde.RegisterType[packages.ModuleError]()
	serde.RegisterType[ast.UnaryExpr]()
	serde.RegisterType[types.Array]()
	serde.RegisterType[ast.AssignStmt]()
	serde.RegisterType[ast.CommentMap]()
	serde.RegisterType[json.Delim]()
	serde.RegisterType[ast.ValueSpec]()
	serde.RegisterType[strings.Builder]()
	serde.RegisterType[uint32]()
	serde.RegisterType[runtime.StackRecord]()
	serde.RegisterType[slog.Attr]()
	serde.RegisterType[syscall.Inet6Pktinfo]()
	serde.RegisterType[syscall.Inet4Pktinfo]()
	serde.RegisterType[types.MethodSet]()
	serde.RegisterType[comment.DocLink]()
	serde.RegisterType[reflect.StringHeader]()
	serde.RegisterType[types.Info]()
	serde.RegisterType[ast.BadStmt]()
	serde.RegisterType[ast.Object]()
	serde.RegisterType[typeutil.MethodSetCache]()
	serde.RegisterType[ast.BlockStmt]()
	serde.RegisterType[ast.ChanType]()
	serde.RegisterType[serde.Generator]()
	serde.RegisterType[time.Time]()
	serde.RegisterType[syscall.SocketControlMessage]()
	serde.RegisterType[sort.IntSlice]()
	serde.RegisterType[json.UnsupportedTypeError]()
	serde.RegisterType[ast.FuncDecl]()
	serde.RegisterType[serde.Deserializer]()
	serde.RegisterType[uint64]()
	serde.RegisterType[scanner.ErrorList]()
	serde.RegisterType[comment.ListItem]()
	serde.RegisterType[comment.Link]()
	serde.RegisterType[syntax.Regexp]()
	serde.RegisterType[syntax.EmptyOp]()
	serde.RegisterType[slog.Source]()
	serde.RegisterType[reflect.SliceHeader]()
	serde.RegisterType[types.TypeParam]()
	serde.RegisterType[types.Label]()
	serde.RegisterType[types.TypeList]()
	serde.RegisterType[ast.Package]()
	serde.RegisterType[syntax.Op]()
	serde.RegisterType[uintptr]()
	serde.RegisterType[os.LinkError]()
	serde.RegisterType[ast.ParenExpr]()
	serde.RegisterType[ast.ReturnStmt]()
	serde.RegisterType[big.ErrNaN]()
	serde.RegisterType[comment.Doc]()
	serde.RegisterType[types.TypeAndValue]()
	serde.RegisterType[types.BasicKind]()
	serde.RegisterType[serde.Serializer]()
	serde.RegisterType[sync.Mutex]()
	serde.RegisterType[syscall.Ustat_t]()
	serde.RegisterType[syscall.IfAddrmsg]()
	serde.RegisterType[exec.Error]()
	serde.RegisterType[types.Tuple]()
	serde.RegisterType[ast.EmptyStmt]()
	serde.RegisterType[comment.Code]()
	serde.RegisterType[typeutil.Map]()
	serde.RegisterType[ast.CaseClause]()
	serde.RegisterType[syscall.Statfs_t]()
	serde.RegisterType[syscall.ProcAttr]()
	serde.RegisterType[syscall.RtGenmsg]()
	serde.RegisterType[syscall.Utsname]()
	serde.RegisterType[unicode.Range16]()
	serde.RegisterType[types.Initializer]()
	serde.RegisterType[time.Month]()
	serde.RegisterType[parser.Mode]()
	serde.RegisterType[ast.TypeSwitchStmt]()
	serde.RegisterType[constraint.AndExpr]()
	serde.RegisterType[comment.Heading]()
	serde.RegisterType[ast.FuncType]()
	serde.RegisterType[Storage]()
	serde.RegisterType[runtime.Func]()
	serde.RegisterType[syscall.Credential]()
	serde.RegisterType[sort.StringSlice]()
	serde.RegisterType[os.SyscallError]()
	serde.RegisterType[reflect.SelectCase]()
	serde.RegisterType[types.SelectionKind]()
	serde.RegisterType[ast.KeyValueExpr]()
	serde.RegisterType[types.Union]()
	serde.RegisterType[scanner.Scanner]()
	serde.RegisterType[build.NoGoError]()
	serde.RegisterType[packages.OverlayJSON]()
	serde.RegisterType[types.Signature]()
	serde.RegisterType[ast.CallExpr]()
	serde.RegisterType[ast.ArrayType]()
	serde.RegisterType[doc.Example]()
	serde.RegisterType[types.Nil]()
	serde.RegisterType[sync.WaitGroup]()
	serde.RegisterType[slog.Kind]()
	serde.RegisterType[syscall.RawSockaddrAny]()
	serde.RegisterType[syscall.Tms]()
	serde.RegisterType[token.FileSet]()
	serde.RegisterType[ast.LabeledStmt]()
	serde.RegisterType[int]()
	serde.RegisterType[float32]()
	serde.RegisterType[atomic.Uintptr]()
	serde.RegisterType[comment.Italic]()
	serde.RegisterType[bool]()
	serde.RegisterType[sync.RWMutex]()
	serde.RegisterType[packages.Error]()
	serde.RegisterType[big.Int]()
	serde.RegisterType[ast.SwitchStmt]()
	serde.RegisterType[comment.Plain]()
	serde.RegisterType[sync.Once]()
	serde.RegisterType[io.SectionReader]()
	serde.RegisterType[io.OffsetWriter]()
	serde.RegisterType[syscall.Errno]()
	serde.RegisterType[json.RawMessage]()
	serde.RegisterType[token.Token]()
	serde.RegisterType[runtime.MemStats]()
	serde.RegisterType[json.Encoder]()
	serde.RegisterType[syscall.InotifyEvent]()
	serde.RegisterType[types.Builtin]()
	serde.RegisterType[big.Accuracy]()
	serde.RegisterType[runtime.PanicNilError]()
	serde.RegisterType[types.Pointer]()
	serde.RegisterType[ast.SendStmt]()
	serde.RegisterType[constant.Kind]()
	serde.RegisterType[comment.LinkDef]()
	serde.RegisterType[ast.ChanDir]()
	serde.RegisterType[slog.TextHandler]()
	serde.RegisterType[reflect.StructTag]()
	serde.RegisterType[packages.LoadMode]()
	serde.RegisterType[types.Package]()
	serde.RegisterType[token.Pos]()
	serde.RegisterType[types.Selection]()
	serde.RegisterType[rand.Rand]()
	serde.RegisterType[token.Position]()
	serde.RegisterType[atomic.Int32]()
	serde.RegisterType[slog.Level]()
	serde.RegisterType[unicode.RangeTable]()
	serde.RegisterType[packages.ErrorKind]()
	serde.RegisterType[ast.ForStmt]()
	serde.RegisterType[time.Location]()
	serde.RegisterType[syscall.RtMsg]()
	serde.RegisterType[types.TypeName]()
	serde.RegisterType[ast.Scope]()
	serde.RegisterType[bufio.Reader]()
	serde.RegisterType[atomic.Bool]()
	serde.RegisterType[slog.HandlerOptions]()
	serde.RegisterType[time.Duration]()
	serde.RegisterType[syscall.Rlimit]()
	serde.RegisterType[packages.Module]()
	serde.RegisterType[ast.BranchStmt]()
	serde.RegisterType[bufio.Scanner]()
	serde.RegisterType[string]()
	serde.RegisterType[syscall.IPMreq]()
	serde.RegisterType[json.SyntaxError]()
	serde.RegisterType[base64.CorruptInputError]()
	serde.RegisterType[constraint.SyntaxError]()
	serde.RegisterType[comment.Paragraph]()
	serde.RegisterType[io.PipeReader]()
	serde.RegisterType[runtime.BlockProfileRecord]()
	serde.RegisterType[ast.CommentGroup]()
	serde.RegisterType[ast.RangeStmt]()
	serde.RegisterType[syntax.Prog]()
	serde.RegisterType[comment.List]()
	serde.RegisterType[complex128]()
	serde.RegisterType[serde.Typedef]()
	serde.RegisterType[atomic.Value]()
	serde.RegisterType[syscall.SysProcAttr]()
	serde.RegisterType[syscall.NetlinkRouteAttr]()
	serde.RegisterType[syscall.Iovec]()
	serde.RegisterType[ast.CompositeLit]()
	serde.RegisterType[ast.CommClause]()
	serde.RegisterType[complex64]()
	serde.RegisterType[io.PipeWriter]()
	serde.RegisterType[fs.FileMode]()
	serde.RegisterType[json.MarshalerError]()
	serde.RegisterType[regexp.Regexp]()
	serde.RegisterType[types.Var]()
	serde.RegisterType[ast.TypeAssertExpr]()
	serde.RegisterType[ast.DeclStmt]()
	serde.RegisterType[syscall.NlMsghdr]()
	serde.RegisterType[syscall.RawSockaddrInet4]()
	serde.RegisterType[syscall.SockaddrInet6]()
	serde.RegisterType[ast.ObjKind]()
	serde.RegisterType[objectpath.Path]()
	serde.RegisterType[exec.ExitError]()
	serde.RegisterType[types.Config]()
	serde.RegisterType[syscall.IPv6Mreq]()
	serde.RegisterType[syscall.Timex]()
	serde.RegisterType[reflect.ChanDir]()
	serde.RegisterType[reflect.Method]()
	serde.RegisterType[unicode.Range32]()
	serde.RegisterType[json.InvalidUTF8Error]()
	serde.RegisterType[ast.SelectorExpr]()
	serde.RegisterType[objectpath.Encoder]()
	serde.RegisterType[build.MultiplePackageError]()
	serde.RegisterType[comment.Printer]()
	serde.RegisterType[unsafe.Pointer]()
	serde.RegisterType[reflect.Value]()
	serde.RegisterType[syscall.Msghdr]()
	serde.RegisterType[syscall.Dirent]()
	serde.RegisterType[constraint.OrExpr]()
	serde.RegisterType[semver.ByVersion]()
	serde.RegisterType[types.Context]()
	serde.RegisterType[ast.FieldList]()
	serde.RegisterType[types.Chan]()
	serde.RegisterType[runtime.TypeAssertionError]()
	serde.RegisterType[json.UnmarshalTypeError]()
	serde.RegisterType[strings.Replacer]()
	serde.RegisterType[types.ImportMode]()
	serde.RegisterType[int64]()
	serde.RegisterType[bytes.Buffer]()
	serde.RegisterType[syscall.RtAttr]()
	serde.RegisterType[syscall.SysProcIDMap]()
	serde.RegisterType[sort.Float64Slice]()
	serde.RegisterType[types.Error]()
	serde.RegisterType[uint16]()
	serde.RegisterType[reflect.ValueError]()
	serde.RegisterType[ast.SelectStmt]()
	serde.RegisterType[ast.DeferStmt]()
	serde.RegisterType[syscall.Cmsghdr]()
	serde.RegisterType[syscall.NetlinkMessage]()
	serde.RegisterType[syscall.RtNexthop]()
	serde.RegisterType[ast.ImportSpec]()
	serde.RegisterType[ast.IndexListExpr]()
	serde.RegisterType[syscall.SockFilter]()
	serde.RegisterType[ast.MapType]()
	serde.RegisterType[ast.FuncLit]()
	serde.RegisterType[big.Float]()
	serde.RegisterType[syscall.Fsid]()
	serde.RegisterType[types.Scope]()
	serde.RegisterType[syscall.SockaddrNetlink]()
	serde.RegisterType[syscall.RawSockaddr]()
	serde.RegisterType[json.Decoder]()
	serde.RegisterType[types.Basic]()
	serde.RegisterType[syntax.Inst]()
	serde.RegisterType[doc.Type]()
	serde.RegisterType[byte]()
	serde.RegisterType[doc.Package]()
	serde.RegisterType[doc.Value]()
	serde.RegisterType[types.Named]()
	serde.RegisterType[types.StdSizes]()
	serde.RegisterType[types.Instance]()
	serde.RegisterType[runtime.Frame]()
	serde.RegisterType[syscall.IPv6MTUInfo]()
	serde.RegisterType[reflect.Kind]()
	serde.RegisterType[reflect.SelectDir]()
	serde.RegisterType[packages.Config]()
	serde.RegisterType[syscall.NlAttr]()
	serde.RegisterType[types.Const]()
	serde.RegisterType[float64]()
	serde.RegisterType[time.Ticker]()
	serde.RegisterType[ast.BasicLit]()
	serde.RegisterType[ast.GoStmt]()
	serde.RegisterType[ast.ExprStmt]()
	serde.RegisterType[big.Word]()
	serde.RegisterType[int8]()
	serde.RegisterType[slog.Record]()
	serde.RegisterType[syscall.Utimbuf]()
	serde.RegisterType[scanner.Mode]()
	serde.RegisterType[syscall.Stat_t]()
	serde.RegisterType[syscall.RawSockaddrLinklayer]()
	serde.RegisterType[syscall.IfInfomsg]()
	serde.RegisterType[types.TypeParamList]()
	serde.RegisterType[ast.BadExpr]()
	serde.RegisterType[constraint.NotExpr]()
	serde.RegisterType[time.ParseError]()
	serde.RegisterType[syscall.Linger]()
	serde.RegisterType[types.Checker]()
	serde.RegisterType[types.Map]()
	serde.RegisterType[os.ProcAttr]()
	serde.RegisterType[json.UnsupportedValueError]()
	serde.RegisterType[types.Term]()
	serde.RegisterType[syntax.Flags]()
	serde.RegisterType[comment.Parser]()
	serde.RegisterType[ast.InterfaceType]()
	serde.RegisterType[ast.SliceExpr]()
	serde.RegisterType[ast.StructType]()
	serde.RegisterType[rand.Zipf]()
	serde.RegisterType[syntax.ErrorCode]()
	serde.RegisterType[bufio.ReadWriter]()
	serde.RegisterType[doc.Mode]()
	serde.RegisterType[astutil.Cursor]()
	serde.RegisterType[sync.Map]()
	serde.RegisterType[syscall.Rusage]()
	serde.RegisterType[syscall.RawSockaddrUnix]()
	serde.RegisterType[syscall.EpollEvent]()
	serde.RegisterType[ast.BinaryExpr]()
	serde.RegisterType[ast.MergeMode]()
	serde.RegisterType[rune]()
	serde.RegisterType[log.Logger]()
	serde.RegisterType[bufio.Writer]()
	serde.RegisterType[Frame]()
}
