package core

import (
	"encoding/hex"
	"fmt"
	"github.com/ledgerwatch/erigon-lib/chain/networkname"
	libcommon "github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon-lib/common/hexutility"
	"github.com/ledgerwatch/erigon/core/systemcontracts"
	"github.com/ledgerwatch/erigon/core/types"
	"github.com/ledgerwatch/erigon/params"
)

func init() {
	// Initialise SystemContractCodeLookup
	for _, chainName := range []string{networkname.BSCChainName, networkname.ChapelChainName, networkname.BorMainnetChainName, networkname.MumbaiChainName, networkname.AmoyChainName, networkname.BorDevnetChainName} {
		byChain := map[libcommon.Address][]libcommon.CodeRecord{}
		systemcontracts.SystemContractCodeLookup[chainName] = byChain
		// Apply genesis with the block number 0
		genesisBlock := GenesisBlockByChainName(chainName)
		allocToCodeRecords(genesisBlock.Alloc, byChain, 0)
		// Process upgrades
		chainConfig := params.ChainConfigByChainName(chainName)
		if chainConfig.RamanujanBlock != nil {
			blockNum := chainConfig.RamanujanBlock.Uint64()
			if blockNum != 0 {
				addCodeRecords(systemcontracts.RamanujanUpgrade[chainName], blockNum, 0, byChain)
			}
		}
		if chainConfig.NielsBlock != nil {
			blockNum := chainConfig.NielsBlock.Uint64()
			if blockNum != 0 {
				addCodeRecords(systemcontracts.NielsUpgrade[chainName], blockNum, 0, byChain)
			}
		}
		if chainConfig.MirrorSyncBlock != nil {
			blockNum := chainConfig.MirrorSyncBlock.Uint64()
			if blockNum != 0 {
				addCodeRecords(systemcontracts.MirrorUpgrade[chainName], blockNum, 0, byChain)
			}
		}
		if chainConfig.BrunoBlock != nil {
			blockNum := chainConfig.BrunoBlock.Uint64()
			if blockNum != 0 {
				addCodeRecords(systemcontracts.BrunoUpgrade[chainName], blockNum, 0, byChain)
			}
		}
		if chainConfig.EulerBlock != nil {
			blockNum := chainConfig.EulerBlock.Uint64()
			if blockNum != 0 {
				addCodeRecords(systemcontracts.EulerUpgrade[chainName], blockNum, 0, byChain)
			}
		}
		if chainConfig.MoranBlock != nil {
			blockNum := chainConfig.MoranBlock.Uint64()
			if blockNum != 0 {
				addCodeRecords(systemcontracts.MoranUpgrade[chainName], blockNum, 0, byChain)
			}
		}
		if chainConfig.GibbsBlock != nil {
			blockNum := chainConfig.GibbsBlock.Uint64()
			if blockNum != 0 {
				addCodeRecords(systemcontracts.GibbsUpgrade[chainName], blockNum, 0, byChain)
			}
		}
		if chainConfig.PlanckBlock != nil {
			blockNum := chainConfig.PlanckBlock.Uint64()
			if blockNum != 0 {
				addCodeRecords(systemcontracts.PlanckUpgrade[chainName], blockNum, 0, byChain)
			}
		}
		if chainConfig.LubanBlock != nil {
			blockNum := chainConfig.LubanBlock.Uint64()
			if blockNum != 0 {
				addCodeRecords(systemcontracts.LubanUpgrade[chainName], blockNum, 0, byChain)
			}
		}
		if chainConfig.PlatoBlock != nil {
			blockNum := chainConfig.PlatoBlock.Uint64()
			if blockNum != 0 {
				addCodeRecords(systemcontracts.PlatoUpgrade[chainName], blockNum, 0, byChain)
			}
		}
		if chainConfig.KeplerTime != nil {
			blockTime := chainConfig.KeplerTime.Uint64()
			if blockTime != 0 {
				addCodeRecords(systemcontracts.KeplerUpgrade[chainName], 0, blockTime, byChain)
			}
		}
	}

	addGnosisSpecialCase()
}

func allocToCodeRecords(alloc types.GenesisAlloc, byChain map[libcommon.Address][]libcommon.CodeRecord, blockNum uint64) {
	for addr, account := range alloc {
		if len(account.Code) > 0 {
			list := byChain[addr]
			codeHash, err := libcommon.HashData(account.Code)
			if err != nil {
				panic(fmt.Errorf("failed to hash system contract code: %s", err.Error()))
			}
			list = append(list, libcommon.CodeRecord{BlockNumber: blockNum, CodeHash: codeHash})
			byChain[addr] = list
		}
	}
}

// some hard coding for gnosis chain here to solve a historical problem with the token contract being re-written
// and losing the history for it in the DB.  Temporary hack until erigon 3 arrives
func addGnosisSpecialCase() {
	byChain := map[libcommon.Address][]libcommon.CodeRecord{}
	systemcontracts.SystemContractCodeLookup[networkname.GnosisChainName] = byChain
	address := libcommon.HexToAddress("0xf8d1677c8a0c961938bf2f9adc3f3cfda759a9d9")
	list := byChain[address]

	oldContractCode := hexutility.FromHex("0x6080604052600436106101b65763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305d2035b81146101bb57806306fdde03146101e4578063095ea7b31461026e5780630b26cf661461029257806318160ddd146102b557806323b872dd146102dc57806330adf81f14610306578063313ce5671461031b5780633644e51514610346578063395093511461035b5780634000aea01461037f57806340c10f19146103b057806342966c68146103d457806354fd4d50146103ec578063661884631461040157806369ffa08a1461042557806370a082311461044c578063715018a61461046d578063726600ce146104825780637d64bcb4146104a35780637ecebe00146104b8578063859ba28c146104d95780638da5cb5b1461051a5780638fcbaf0c1461054b57806395d89b4114610589578063a457c2d71461059e578063a9059cbb146105c2578063b753a98c146105e6578063bb35783b1461060a578063cd59658314610634578063d73dd62314610649578063dd62ed3e1461066d578063f2d5d56b14610694578063f2fde38b146106b8578063ff9e884d146106d9575b600080fd5b3480156101c757600080fd5b506101d0610700565b604080519115158252519081900360200190f35b3480156101f057600080fd5b506101f9610721565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561023357818101518382015260200161021b565b50505050905090810190601f1680156102605780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561027a57600080fd5b506101d0600160a060020a03600435166024356107af565b34801561029e57600080fd5b506102b3600160a060020a0360043516610803565b005b3480156102c157600080fd5b506102ca61085d565b60408051918252519081900360200190f35b3480156102e857600080fd5b506101d0600160a060020a0360043581169060243516604435610863565b34801561031257600080fd5b506102ca610a32565b34801561032757600080fd5b50610330610a56565b6040805160ff9092168252519081900360200190f35b34801561035257600080fd5b506102ca610a5f565b34801561036757600080fd5b506101d0600160a060020a0360043516602435610a65565b34801561038b57600080fd5b506101d060048035600160a060020a0316906024803591604435918201910135610a78565b3480156103bc57600080fd5b506101d0600160a060020a0360043516602435610b89565b3480156103e057600080fd5b506102b3600435610c94565b3480156103f857600080fd5b506101f9610ca1565b34801561040d57600080fd5b506101d0600160a060020a0360043516602435610cd8565b34801561043157600080fd5b506102b3600160a060020a0360043581169060243516610db5565b34801561045857600080fd5b506102ca600160a060020a0360043516610df1565b34801561047957600080fd5b506102b3610e0c565b34801561048e57600080fd5b506101d0600160a060020a0360043516610e23565b3480156104af57600080fd5b506101d0610e37565b3480156104c457600080fd5b506102ca600160a060020a0360043516610e3e565b3480156104e557600080fd5b506104ee610e50565b6040805167ffffffffffffffff9485168152928416602084015292168183015290519081900360600190f35b34801561052657600080fd5b5061052f610e5b565b60408051600160a060020a039092168252519081900360200190f35b34801561055757600080fd5b506102b3600160a060020a0360043581169060243516604435606435608435151560ff60a4351660c43560e435610e6a565b34801561059557600080fd5b506101f9611171565b3480156105aa57600080fd5b506101d0600160a060020a03600435166024356111cb565b3480156105ce57600080fd5b506101d0600160a060020a03600435166024356111d7565b3480156105f257600080fd5b506102b3600160a060020a0360043516602435611202565b34801561061657600080fd5b506102b3600160a060020a036004358116906024351660443561120d565b34801561064057600080fd5b5061052f61121e565b34801561065557600080fd5b506101d0600160a060020a036004351660243561122d565b34801561067957600080fd5b506102ca600160a060020a03600435811690602435166112b4565b3480156106a057600080fd5b506102b3600160a060020a03600435166024356112df565b3480156106c457600080fd5b506102b3600160a060020a03600435166112ea565b3480156106e557600080fd5b506102ca600160a060020a036004358116906024351661130a565b60065474010000000000000000000000000000000000000000900460ff1681565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156107a75780601f1061077c576101008083540402835291602001916107a7565b820191906000526020600020905b81548152906001019060200180831161078a57829003601f168201915b505050505081565b336000818152600560209081526040808320600160a060020a03871680855290835281842086905581518681529151939490939092600080516020611a13833981519152928290030190a350600192915050565b600654600160a060020a0316331461081a57600080fd5b61082381611327565b151561082e57600080fd5b6007805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60045490565b600080600160a060020a038516151561087b57600080fd5b600160a060020a038416151561089057600080fd5b600160a060020a0385166000908152600360205260409020546108b9908463ffffffff61132f16565b600160a060020a0380871660009081526003602052604080822093909355908616815220546108ee908463ffffffff61134116565b600160a060020a0380861660008181526003602090815260409182902094909455805187815290519193928916926000805160206119f383398151915292918290030190a3600160a060020a0385163314610a1c5761094d85336112b4565b905060001981146109b757610968818463ffffffff61132f16565b600160a060020a038616600081815260056020908152604080832033808552908352928190208590558051948552519193600080516020611a13833981519152929081900390910190a3610a1c565b600160a060020a0385166000908152600a602090815260408083203384529091529020541580610a1157506109ea611354565b600160a060020a0386166000908152600a6020908152604080832033845290915290205410155b1515610a1c57600080fd5b610a27858585611358565b506001949350505050565b7fea2aa0a1be11a07ed86d755c93467f4f82362b452371d1ba94d1715123511acb81565b60025460ff1681565b60085481565b6000610a71838361122d565b9392505050565b600084600160a060020a03811615801590610a9c5750600160a060020a0381163014155b1515610aa757600080fd5b610ab186866113ef565b1515610abc57600080fd5b85600160a060020a031633600160a060020a03167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16878787604051808481526020018060200182810382528484828181526020019250808284376040519201829003965090945050505050a3610b3186611327565b15610b7d57610b7233878787878080601f016020809104026020016040519081016040528093929190818152602001838380828437506113fb945050505050565b1515610b7d57600080fd5b50600195945050505050565b600654600090600160a060020a03163314610ba357600080fd5b60065474010000000000000000000000000000000000000000900460ff1615610bcb57600080fd5b600454610bde908363ffffffff61134116565b600455600160a060020a038316600090815260036020526040902054610c0a908363ffffffff61134116565b600160a060020a038416600081815260036020908152604091829020939093558051858152905191927f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d412139688592918290030190a2604080518381529051600160a060020a038516916000916000805160206119f38339815191529181900360200190a350600192915050565b610c9e3382611591565b50565b60408051808201909152600181527f3100000000000000000000000000000000000000000000000000000000000000602082015281565b336000908152600560209081526040808320600160a060020a0386168452909152812054808310610d2c57336000908152600560209081526040808320600160a060020a0388168452909152812055610d61565b610d3c818463ffffffff61132f16565b336000908152600560209081526040808320600160a060020a03891684529091529020555b336000818152600560209081526040808320600160a060020a038916808552908352928190205481519081529051929392600080516020611a13833981519152929181900390910190a35060019392505050565b600654600160a060020a03163314610dcc57600080fd5b80600160a060020a0381161515610de257600080fd5b610dec8383611680565b505050565b600160a060020a031660009081526003602052604090205490565b600654600160a060020a031633146101b657600080fd5b600754600160a060020a0390811691161490565b6000806000fd5b60096020526000908152604090205481565b600260036000909192565b600654600160a060020a031681565b600080600160a060020a038a161515610e8257600080fd5b600160a060020a0389161515610e9757600080fd5b861580610eab575086610ea8611354565b11155b1515610eb657600080fd5b600854604080517fea2aa0a1be11a07ed86d755c93467f4f82362b452371d1ba94d1715123511acb602080830191909152600160a060020a03808f16838501528d166060830152608082018c905260a082018b905289151560c0808401919091528351808403909101815260e090920192839052815191929182918401908083835b60208310610f575780518252601f199092019160209182019101610f38565b51815160209384036101000a6000190180199092169116179052604080519290940182900382207f190100000000000000000000000000000000000000000000000000000000000083830152602283019790975260428083019790975283518083039097018752606290910192839052855192945084935085019190508083835b60208310610ff75780518252601f199092019160209182019101610fd8565b51815160209384036101000a600019018019909216911617905260408051929094018290038220600080845283830180875282905260ff8d1684870152606084018c9052608084018b905294519098506001965060a080840196509194601f19820194509281900390910191865af1158015611077573d6000803e3d6000fd5b50505060206040510351600160a060020a03168a600160a060020a03161415156110a057600080fd5b600160a060020a038a16600090815260096020526040902080546001810190915588146110cc57600080fd5b856110d85760006110dc565b6000195b600160a060020a03808c166000908152600560209081526040808320938e16835292905220819055905085611112576000611114565b865b600160a060020a03808c166000818152600a60209081526040808320948f1680845294825291829020949094558051858152905192939192600080516020611a13833981519152929181900390910190a350505050505050505050565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156107a75780601f1061077c576101008083540402835291602001916107a7565b6000610a718383610cd8565b60006111e383836113ef565b15156111ee57600080fd5b6111f9338484611358565b50600192915050565b610dec338383610863565b611218838383610863565b50505050565b600754600160a060020a031690565b336000908152600560209081526040808320600160a060020a0386168452909152812054611261908363ffffffff61134116565b336000818152600560209081526040808320600160a060020a038916808552908352928190208590558051948552519193600080516020611a13833981519152929081900390910190a350600192915050565b600160a060020a03918216600090815260056020908152604080832093909416825291909152205490565b610dec823383610863565b600654600160a060020a0316331461130157600080fd5b610c9e816116ac565b600a60209081526000928352604080842090915290825290205481565b6000903b1190565b60008282111561133b57fe5b50900390565b8181018281101561134e57fe5b92915050565b4290565b61136182611327565b80156113885750604080516000815260208101909152611386908490849084906113fb565b155b15610dec5761139682610e23565b156113a057600080fd5b60408051600160a060020a0380861682528416602082015280820183905290517f11249f0fc79fc134a15a10d1da8291b79515bf987e036ced05b9ec119614070b9181900360600190a1505050565b6000610a71838361172a565b600083600160a060020a031663a4c0ed367c0100000000000000000000000000000000000000000000000000000000028685856040516024018084600160a060020a0316600160a060020a0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561148c578181015183820152602001611474565b50505050905090810190601f1680156114b95780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909916989098178852518151919790965086955093509150819050838360005b8381101561154757818101518382015260200161152f565b50505050905090810190601f1680156115745780820380516001836020036101000a031916815260200191505b509150506000604051808303816000865af1979650505050505050565b600160a060020a0382166000908152600360205260409020548111156115b657600080fd5b600160a060020a0382166000908152600360205260409020546115df908263ffffffff61132f16565b600160a060020a03831660009081526003602052604090205560045461160b908263ffffffff61132f16565b600455604080518281529051600160a060020a038416917fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5919081900360200190a2604080518281529051600091600160a060020a038516916000805160206119f38339815191529181900360200190a35050565b600160a060020a038216151561169e57611699816117f9565b6116a8565b6116a88282611805565b5050565b600160a060020a03811615156116c157600080fd5b600654604051600160a060020a038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a36006805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b3360009081526003602052604081205482111561174657600080fd5b600160a060020a038316151561175b57600080fd5b3360009081526003602052604090205461177b908363ffffffff61132f16565b3360009081526003602052604080822092909255600160a060020a038516815220546117ad908363ffffffff61134116565b600160a060020a0384166000818152600360209081526040918290209390935580518581529051919233926000805160206119f38339815191529281900390910190a350600192915050565b30316116a882826118a3565b604080517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290518391600091600160a060020a038416916370a0823191602480830192602092919082900301818787803b15801561186a57600080fd5b505af115801561187e573d6000803e3d6000fd5b505050506040513d602081101561189457600080fd5b5051905061121884848361190b565b604051600160a060020a0383169082156108fc029083906000818181858888f1935050505015156116a85780826118d86119c2565b600160a060020a039091168152604051908190036020019082f080158015611904573d6000803e3d6000fd5b5050505050565b60408051600160a060020a03841660248201526044808201849052825180830390910181526064909101909152602081810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb000000000000000000000000000000000000000000000000000000001781528251606093600093909290918491828a5af160005193508392508080156101b65750506000835111156119ba578115156119ba57600080fd5b505050505050565b6040516021806119d2833901905600608060405260405160208060218339810160405251600160a060020a038116ff00ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925a165627a7a72305820da715ff88e0288dbae664bb8af2f148726bdc8c499fecf88153280d022031e780029")
	newContractCode := hexutility.FromHex("0x6080604052600436106101b35763ffffffff60e060020a60003504166305d2035b81146101b857806306fdde03146101e1578063095ea7b31461026b5780630b26cf661461028f57806318160ddd146102b257806323b872dd146102d957806330adf81f14610303578063313ce567146103185780633644e5151461034357806339509351146103585780634000aea01461037c57806340c10f19146103ad57806342966c68146103d157806354fd4d50146103e957806366188463146103fe57806369ffa08a1461042257806370a0823114610449578063715018a61461046a578063726600ce1461047f5780637d64bcb4146104a05780637ecebe00146104b5578063859ba28c146104d65780638da5cb5b146105175780638fcbaf0c1461054857806395d89b4114610586578063a457c2d71461059b578063a9059cbb146105bf578063b753a98c146105e3578063bb35783b14610607578063c6a1dedf14610631578063cd59658314610646578063d505accf1461065b578063d73dd62314610694578063dd62ed3e146106b8578063f2d5d56b146106df578063f2fde38b14610703578063ff9e884d14610724575b600080fd5b3480156101c457600080fd5b506101cd61074b565b604080519115158252519081900360200190f35b3480156101ed57600080fd5b506101f661076c565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610230578181015183820152602001610218565b50505050905090810190601f16801561025d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561027757600080fd5b506101cd600160a060020a03600435166024356107fa565b34801561029b57600080fd5b506102b0600160a060020a0360043516610810565b005b3480156102be57600080fd5b506102c761086a565b60408051918252519081900360200190f35b3480156102e557600080fd5b506101cd600160a060020a0360043581169060243516604435610870565b34801561030f57600080fd5b506102c7610a38565b34801561032457600080fd5b5061032d610a5c565b6040805160ff9092168252519081900360200190f35b34801561034f57600080fd5b506102c7610a65565b34801561036457600080fd5b506101cd600160a060020a0360043516602435610a6b565b34801561038857600080fd5b506101cd60048035600160a060020a0316906024803591604435918201910135610aac565b3480156103b957600080fd5b506101cd600160a060020a0360043516602435610bbd565b3480156103dd57600080fd5b506102b0600435610cc8565b3480156103f557600080fd5b506101f6610cd5565b34801561040a57600080fd5b506101cd600160a060020a0360043516602435610d0c565b34801561042e57600080fd5b506102b0600160a060020a0360043581169060243516610de9565b34801561045557600080fd5b506102c7600160a060020a0360043516610e0e565b34801561047657600080fd5b506102b0610e29565b34801561048b57600080fd5b506101cd600160a060020a0360043516610e40565b3480156104ac57600080fd5b506101cd610e54565b3480156104c157600080fd5b506102c7600160a060020a0360043516610e5b565b3480156104e257600080fd5b506104eb610e6d565b6040805167ffffffffffffffff9485168152928416602084015292168183015290519081900360600190f35b34801561052357600080fd5b5061052c610e78565b60408051600160a060020a039092168252519081900360200190f35b34801561055457600080fd5b506102b0600160a060020a0360043581169060243516604435606435608435151560ff60a4351660c43560e435610e87565b34801561059257600080fd5b506101f6610fc5565b3480156105a757600080fd5b506101cd600160a060020a036004351660243561101f565b3480156105cb57600080fd5b506101cd600160a060020a0360043516602435611032565b3480156105ef57600080fd5b506102b0600160a060020a0360043516602435611054565b34801561061357600080fd5b506102b0600160a060020a0360043581169060243516604435611064565b34801561063d57600080fd5b506102c7611075565b34801561065257600080fd5b5061052c611099565b34801561066757600080fd5b506102b0600160a060020a036004358116906024351660443560643560ff6084351660a43560c4356110a8565b3480156106a057600080fd5b506101cd600160a060020a0360043516602435611184565b3480156106c457600080fd5b506102c7600160a060020a036004358116906024351661120b565b3480156106eb57600080fd5b506102b0600160a060020a0360043516602435611236565b34801561070f57600080fd5b506102b0600160a060020a0360043516611241565b34801561073057600080fd5b506102c7600160a060020a0360043581169060243516611261565b60065474010000000000000000000000000000000000000000900460ff1681565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156107f25780601f106107c7576101008083540402835291602001916107f2565b820191906000526020600020905b8154815290600101906020018083116107d557829003601f168201915b505050505081565b600061080733848461127e565b50600192915050565b600654600160a060020a0316331461082757600080fd5b610830816112c0565b151561083b57600080fd5b6007805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60045490565b600080600160a060020a038516151561088857600080fd5b600160a060020a038416151561089d57600080fd5b600160a060020a0385166000908152600360205260409020546108c6908463ffffffff6112c816565b600160a060020a0380871660009081526003602052604080822093909355908616815220546108fb908463ffffffff6112da16565b600160a060020a038086166000818152600360209081526040918290209490945580518781529051919392891692600080516020611d7283398151915292918290030190a3600160a060020a0385163314610a225761095a853361120b565b905060001981146109c457610975818463ffffffff6112c816565b600160a060020a038616600081815260056020908152604080832033808552908352928190208590558051948552519193600080516020611d92833981519152929081900390910190a3610a22565b600160a060020a0385166000908152600a602090815260408083203384529091529020541580610a175750600160a060020a0385166000908152600a602090815260408083203384529091529020544211155b1515610a2257600080fd5b610a2d8585856112ed565b506001949350505050565b7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c981565b60025460ff1681565b60085481565b336000818152600560209081526040808320600160a060020a03871684529091528120549091610807918590610aa7908663ffffffff6112da16565b61127e565b600084600160a060020a03811615801590610ad05750600160a060020a0381163014155b1515610adb57600080fd5b610ae58686611324565b1515610af057600080fd5b85600160a060020a031633600160a060020a03167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16878787604051808481526020018060200182810382528484828181526020019250808284376040519201829003965090945050505050a3610b65866112c0565b15610bb157610ba633878787878080601f01602080910402602001604051908101604052809392919081815260200183838082843750611330945050505050565b1515610bb157600080fd5b50600195945050505050565b600654600090600160a060020a03163314610bd757600080fd5b60065474010000000000000000000000000000000000000000900460ff1615610bff57600080fd5b600454610c12908363ffffffff6112da16565b600455600160a060020a038316600090815260036020526040902054610c3e908363ffffffff6112da16565b600160a060020a038416600081815260036020908152604091829020939093558051858152905191927f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d412139688592918290030190a2604080518381529051600160a060020a03851691600091600080516020611d728339815191529181900360200190a350600192915050565b610cd233826114ad565b50565b60408051808201909152600181527f3100000000000000000000000000000000000000000000000000000000000000602082015281565b336000908152600560209081526040808320600160a060020a0386168452909152812054808310610d6057336000908152600560209081526040808320600160a060020a0388168452909152812055610d95565b610d70818463ffffffff6112c816565b336000908152600560209081526040808320600160a060020a03891684529091529020555b336000818152600560209081526040808320600160a060020a038916808552908352928190205481519081529051929392600080516020611d92833981519152929181900390910190a35060019392505050565b600654600160a060020a03163314610e0057600080fd5b610e0a828261159c565b5050565b600160a060020a031660009081526003602052604090205490565b600654600160a060020a031633146101b357600080fd5b600754600160a060020a0390811691161490565b6000806000fd5b60096020526000908152604090205481565b600260056000909192565b600654600160a060020a031681565b600080861580610e975750864211155b1515610ea257600080fd5b604080517fea2aa0a1be11a07ed86d755c93467f4f82362b452371d1ba94d1715123511acb6020820152600160a060020a03808d16828401528b166060820152608081018a905260a0810189905287151560c0808301919091528251808303909101815260e0909101909152610f17906115da565b9150610f25828686866116e1565b600160a060020a038b8116911614610f3c57600080fd5b600160a060020a038a1660009081526009602052604090208054600181019091558814610f6857600080fd5b85610f74576000610f78565b6000195b905085610f86576000610f88565b865b600160a060020a03808c166000908152600a60209081526040808320938e1683529290522055610fb98a8a836118e3565b50505050505050505050565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156107f25780601f106107c7576101008083540402835291602001916107f2565b600061102b8383610d0c565b9392505050565b600061103e8383611324565b151561104957600080fd5b6108073384846112ed565b61105f338383610870565b505050565b61106f838383610870565b50505050565b7fea2aa0a1be11a07ed86d755c93467f4f82362b452371d1ba94d1715123511acb81565b600754600160a060020a031690565b600080428610156110b857600080fd5b600160a060020a03808a1660008181526009602090815260409182902080546001810190915582517f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c99281019290925281830193909352928b166060840152608083018a905260a0830182905260c08084018a90528151808503909101815260e090930190529250611149906115da565b9050611157818686866116e1565b600160a060020a038a811691161461116e57600080fd5b61117989898961127e565b505050505050505050565b336000908152600560209081526040808320600160a060020a03861684529091528120546111b8908363ffffffff6112da16565b336000818152600560209081526040808320600160a060020a038916808552908352928190208590558051948552519193600080516020611d92833981519152929081900390910190a350600192915050565b600160a060020a03918216600090815260056020908152604080832093909416825291909152205490565b61105f823383610870565b600654600160a060020a0316331461125857600080fd5b610cd281611a3e565b600a60209081526000928352604080842090915290825290205481565b6112898383836118e3565b60001981141561105f57600160a060020a038084166000908152600a60209081526040808320938616835292905290812055505050565b6000903b1190565b6000828211156112d457fe5b50900390565b818101828110156112e757fe5b92915050565b6112f682610e40565b1561105f5760408051600081526020810190915261131990849084908490611330565b151561105f57600080fd5b600061102b8383611abc565b600083600160a060020a031663a4c0ed3660e060020a028685856040516024018084600160a060020a0316600160a060020a0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156113a8578181015183820152602001611390565b50505050905090810190601f1680156113d55780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909916989098178852518151919790965086955093509150819050838360005b8381101561146357818101518382015260200161144b565b50505050905090810190601f1680156114905780820380516001836020036101000a031916815260200191505b509150506000604051808303816000865af1979650505050505050565b600160a060020a0382166000908152600360205260409020548111156114d257600080fd5b600160a060020a0382166000908152600360205260409020546114fb908263ffffffff6112c816565b600160a060020a038316600090815260036020526040902055600454611527908263ffffffff6112c816565b600455604080518281529051600160a060020a038416917fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5919081900360200190a2604080518281529051600091600160a060020a03851691600080516020611d728339815191529181900360200190a35050565b80600160a060020a03811615156115b257600080fd5b600160a060020a03831615156115d0576115cb82611b8b565b61105f565b61105f8383611b97565b6000600854826040518082805190602001908083835b6020831061160f5780518252601f1990920191602091820191016115f0565b51815160209384036101000a6000190180199092169116179052604080519290940182900382207f190100000000000000000000000000000000000000000000000000000000000083830152602283019790975260428083019790975283518083039097018752606290910192839052855192945084935085019190508083835b602083106116af5780518252601f199092019160209182019101611690565b5181516020939093036101000a6000190180199091169216919091179052604051920182900390912095945050505050565b6000808460ff16601b14806116f957508460ff16601c145b1515611775576040805160e560020a62461bcd02815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f7565000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115611813576040805160e560020a62461bcd02815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f7565000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60408051600080825260208083018085528a905260ff8916838501526060830188905260808301879052925160019360a0808501949193601f19840193928390039091019190865af115801561186d573d6000803e3d6000fd5b5050604051601f190151915050600160a060020a03811615156118da576040805160e560020a62461bcd02815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015290519081900360640190fd5b95945050505050565b600160a060020a0383161515611968576040805160e560020a62461bcd028152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f7265737300000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a03821615156119ee576040805160e560020a62461bcd02815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f7373000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a0380841660008181526005602090815260408083209487168084529482529182902085905581518581529151600080516020611d928339815191529281900390910190a3505050565b600160a060020a0381161515611a5357600080fd5b600654604051600160a060020a038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a36006805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b33600090815260036020526040812054821115611ad857600080fd5b600160a060020a0383161515611aed57600080fd5b33600090815260036020526040902054611b0d908363ffffffff6112c816565b3360009081526003602052604080822092909255600160a060020a03851681522054611b3f908363ffffffff6112da16565b600160a060020a038416600081815260036020908152604091829020939093558051858152905191923392600080516020611d728339815191529281900390910190a350600192915050565b3031610e0a8282611c44565b604080517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290518391600091600160a060020a038416916370a0823191602480830192602092919082900301818787803b158015611bfc57600080fd5b505af1158015611c10573d6000803e3d6000fd5b505050506040513d6020811015611c2657600080fd5b5051905061106f600160a060020a038516848363ffffffff611cac16565b604051600160a060020a0383169082156108fc029083906000818181858888f193505050501515610e0a578082611c79611d41565b600160a060020a039091168152604051908190036020019082f080158015611ca5573d6000803e3d6000fd5b5050505050565b82600160a060020a031663a9059cbb83836040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050600060405180830381600087803b158015611d0f57600080fd5b505af1158015611d23573d6000803e3d6000fd5b505050503d1561105f5760206000803e600051151561105f57600080fd5b604051602180611d51833901905600608060405260405160208060218339810160405251600160a060020a038116ff00ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925a165627a7a72305820b96bb0733a3e45fdddafa592f51114d0cf16cad047ad60b9b91ae91eb772c6940029")

	codeHash, err := libcommon.HashData(oldContractCode)
	if err != nil {
		panic("could not get code hash from old gnosis token contract")
	}
	list = append(list, libcommon.CodeRecord{
		BlockNumber: 0,
		CodeHash:    codeHash,
	})

	codeHash, err = libcommon.HashData(newContractCode)
	if err != nil {
		panic("could not get code hash from new gnosis token contract")
	}
	list = append(list, libcommon.CodeRecord{
		BlockNumber: 21735000,
		CodeHash:    codeHash,
	})

	byChain[address] = list
}

func addCodeRecords(upgrade *systemcontracts.Upgrade, blockNum uint64, blockTime uint64, byChain map[libcommon.Address][]libcommon.CodeRecord) {
	for _, config := range upgrade.Configs {
		list := byChain[config.ContractAddr]
		code, err := hex.DecodeString(config.Code)
		if err != nil {
			panic(fmt.Errorf("failed to decode system contract code: %s", err.Error()))
		}
		codeHash, err := libcommon.HashData(code)
		if err != nil {
			panic(fmt.Errorf("failed to hash system contract code: %s", err.Error()))
		}
		if blockTime == 0 {
			list = append(list, libcommon.CodeRecord{BlockNumber: blockNum, CodeHash: codeHash})
		} else {
			list = append(list, libcommon.CodeRecord{BlockTime: blockTime, CodeHash: codeHash})
		}
		byChain[config.ContractAddr] = list
	}
}
