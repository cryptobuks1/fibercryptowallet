package skycoin

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	util "github.com/fibercrypto/FiberCryptoWallet/src/util"
)

type SkyFiberPlugin struct {
	Params SkyFiberParams
}

func (p *SkyFiberPlugin) ListSupportedAltcoins() []core.AltcoinMetadata {
	return []core.AltcoinMetadata{
		core.AltcoinMetadata{
			Name:     SkycoinName,
			Ticker:   SkycoinTicker,
			Family:   SkycoinFamily,
			HasBip44: false,
		},
		core.AltcoinMetadata{
			Name:     CoinHoursName,
			Ticker:   CoinHoursTicker,
			Family:   CoinHoursFamily,
			HasBip44: false,
		},
	}
}

func (p *SkyFiberPlugin) ListSupportedFamilies() []string {
	return []string{SkycoinFamily}
}

func (p *SkyFiberPlugin) RegisterTo(manager core.AltcoinManager) {
	for _, info := range p.ListSupportedAltcoins() {
		manager.RegisterAltcoin(info, p)
	}
}

func (p *SkyFiberPlugin) GetName() string {
	return "SkyFiber"
}

func (p *SkyFiberPlugin) GetDescription() string {
	return "FiberCrypto wallet connector for Skycoin and SkyFiber altcoins"
}

func (p *SkyFiberPlugin) LoadWalletEnvs() []core.WalletEnv {
	// TODO: Load wallets at $HOME/.skycoin/wallets
	return nil
}

func NewSkyFiberPlugin(params SkyFiberParams) core.AltcoinPlugin {
	return &SkyFiberPlugin{
		Params: params,
	}
}

func init() {
	util.RegisterAltcoin(NewSkyFiberPlugin(SkycoinMainNetParams))
}