package logic

import (
	"fmt"
	"github.com/spf13/cobra"
	"zerok8scron/internal/svc"
)

// Hello print "hello SvcName" once per minute
func Hello(_ *cobra.Command, _ []string) error {

	fmt.Printf("srvName : %s , hello \n", svc.GetSvcCtx().Config.Name)

	return nil
}
