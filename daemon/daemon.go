package daemon
import (
	"github.com/globaldce/globaldce-gateway/mainchain"
	"github.com/globaldce/globaldce-gateway/wire"
	"github.com/globaldce/globaldce-gateway/wallet"
	//"path/filepath"
	//"fmt"
	//"os"
	//"github.com/globaldce/globaldce-gateway/applog"
)
	var (

		MainwalletFilePath=""
		MainwalletFileKey []byte
		AppPath=""
		AppLocalPort=""
		Walletloaded=false
		Miningrequested=false
		Miningrunning=false
		HotMining=false
		Managingwalletrequested=false
		Seed=false
		SyncingMinNbPeers=0
		AppIsClosing=false
		Mn *mainchain.Maincore
		Wireswarm *wire.Swarm
		Wlt *wallet.Wallet
	)

