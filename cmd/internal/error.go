package internal

import (
	"errors"
	"fmt"
	"orca/errs"
	"os"
)

func HandleError(err error,silent bool)int{
switch {
	case errors.Is(err, errs.ErrAlreadyInitialized):
		if !silent {
			fmt.Fprintln(os.Stderr, "このディレクトリは既に初期化されています")
			fmt.Fprintln(os.Stderr, "--force オプションで再生成してください")
		}
		return 0
	case errors.Is(err, errs.ErrNotInitialized):
		if !silent {
			fmt.Fprintln(os.Stderr, "このディレクトリは初期化されていません")
			fmt.Fprintln(os.Stderr, "orca init を先に実行してください")
		}
		//return exitcode.NotInitialized
		return 1
	case errors.Is(err, errs.ErrInvalidConfig):
		if !silent {
			fmt.Fprintln(os.Stderr, "orca.ymlの読み込みに失敗しました")
		}
		return 1
	case errors.Is(err, errs.ErrPlanDirty):
		if !silent {
			fmt.Fprintln(os.Stderr, "plan が最新ではありません")
			fmt.Fprintln(os.Stderr, "orca plan を再実行してください")
		}
		//return exitcode.PlanDirty
		return 1

	case errors.Is(err, errs.ErrDryRunViolation):
		if !silent {
			fmt.Fprintln(os.Stderr, "dry-run のため操作は実行されませんでした")
		}
		//return exitcode.OK
		return 0

	default:
		if !silent {
			fmt.Fprintln(os.Stderr, err.Error())
		}
		//return exitcode.GeneralError
		return 1
	}
}