// 抄自 https://github.com/marmotedu
// 用于将flag分组输出
package flag

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/pflag"
)

type NamedFlagSets struct {
	// flag名称的有序列表
	Order []string
	// flag名称到flagset的映射
	FlagSets map[string]*pflag.FlagSet
}

// FlagSet 返回与给定名称关联的 FlagSet。
// 如果 FlagSet 尚不存在，它会创建一个新的 FlagSet，将其添加到 NamedFlagSets，并记录其创建顺序。
// FlagSet 是使用 pflag 创建的。
// ExitOnError 错误处理
// policy.
// Parameters:
//   - name: The name of the FlagSet to retrieve or create.
//
// Returns:
//   - *pflag.FlagSet: The FlagSet associated with the given name.
func (nfs *NamedFlagSets) FlagSet(name string) *pflag.FlagSet {
	if nfs.FlagSets == nil {
		nfs.FlagSets = map[string]*pflag.FlagSet{}
	}
	if _, ok := nfs.FlagSets[name]; !ok {
		nfs.FlagSets[name] = pflag.NewFlagSet(name, pflag.ExitOnError)
		nfs.Order = append(nfs.Order, name)
	}
	return nfs.FlagSets[name]
}

func PrintSections(w io.Writer, fss NamedFlagSets, cols int) {
	for _, name := range fss.Order {
		fs := fss.FlagSets[name]
		if !fs.HasFlags() {
			continue
		}

		wideFS := pflag.NewFlagSet("", pflag.ExitOnError)
		wideFS.AddFlagSet(fs)

		var zzz string
		if cols > 24 {
			zzz = strings.Repeat("z", cols-24)
			wideFS.Int(zzz, 0, strings.Repeat("z", cols-24))
		}

		var buf bytes.Buffer
		fmt.Fprintf(&buf, "\n%s flags:\n\n%s", strings.ToUpper(name[:1])+name[1:], wideFS.FlagUsagesWrapped(cols))

		if cols > 24 {
			i := strings.Index(buf.String(), zzz)
			lines := strings.Split(buf.String()[:i], "\n")
			fmt.Fprint(w, strings.Join(lines[:len(lines)-1], "\n"))
			fmt.Fprintln(w)
		} else {
			fmt.Fprint(w, buf.String())
		}
	}
}
