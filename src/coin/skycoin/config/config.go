package config

import (
	"encoding/json"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"

	local "github.com/fibercrypto/fibercryptowallet/src/main"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
)

const (
	LocalWallet               = "local"
	RemoteWallet              = "remote"
	SectionName               = "skycoin"
	SettingPathToLog          = "log"
	SettingPathToNode         = "node"
	SettingNodeAddress        = "address"
	SettingPathToWalletSource = "walletSource"
)

var (
	sectionManager *local.SectionManager
	log            = logging.MustGetLogger("Skycoin Config")
)

func getMultiPlatformUserDirectory() string {
	usr, err := user.Current()
	if err != nil {
		log.WithError(err).Error()
		return ""
	}
	return filepath.Join(usr.HomeDir, string(os.PathSeparator), ".skycoin", string(os.PathSeparator), "wallets")
}

func RegisterConfig() error {
	cm := local.GetConfigManager()
	node := map[string]string{"address": "https://node.skycoin.com"}
	nodeBytes, err := json.Marshal(node)
	if err != nil {
		return err
	}
	nodeOpt := local.NewOption(SettingPathToNode, []string{}, false, string(nodeBytes))
	walletsDefaultDirectory := getMultiPlatformUserDirectory()
	wltSrc := &walletSource{
		id:     "1",
		Tp:     string(LocalWallet),
		Source: walletsDefaultDirectory,
	}
	wltSrcBytes, err := json.Marshal(wltSrc)
	if err != nil {
		return err
	}

	wltOpt := local.NewOption(string(wltSrc.id), []string{SettingPathToWalletSource}, false, string(wltSrcBytes))

	level := map[string]string{"level": strconv.Itoa(logging.Warning)}
	levelBytes, err := json.Marshal(level)
	if err != nil {
		return err
	}
	logLevelOpt := local.NewOption(SettingPathToLog, []string{}, false, string(levelBytes))

	output := map[string]string{"output": "none"}
	outputBytes, err := json.Marshal(output)
	if err != nil {
		return err
	}

	logOutputOpt := local.NewOption(SettingPathToLog, []string{}, false, string(outputBytes))

	outputFile := map[string]string{"outputFile": ""}
	outputFileBytes, err := json.Marshal(outputFile)
	if err != nil {
		return err
	}

	logOutputFileOpt := local.NewOption(SettingPathToLog, []string{}, false, string(outputFileBytes))

	sectionManager = cm.RegisterSection(SectionName, []*local.Option{nodeOpt, wltOpt, logLevelOpt, logOutputOpt, logOutputFileOpt})
	return nil
}

func GetOption(path string) (string, error) {
	stringList := strings.Split(path, "/")
	return sectionManager.GetValue(stringList[len(stringList)-1], stringList[:len(stringList)-1])
}

func getValues(prefix string) ([]string, error) {
	return sectionManager.GetValues(strings.Split(prefix, "/"))
}

func GetDataRefreshTimeout() uint64 {
	return getFromCache(local.DataRefreshTimeoutKey)
}

func GetDataUpdateTime() uint64 {
	return getFromCache(local.DataUpdateTimeKey)
}

func getFromCache(cacheKeyValue string) uint64 {
	cm := local.GetConfigManager()
	sm := cm.GetSectionManager("global")
	value, err := sm.GetValue("cache", nil)
	if err != nil {
		log.WithError(err).Warn("Couldn't get cache value option for saved settings")
		return 0
	}

	keyValue := make(map[string]string)
	err = json.Unmarshal([]byte(value), &keyValue)
	if err != nil {
		log.WithError(err).Warn("Couldn't unmarshal from options")
		return 0
	}
	strVal, ok := keyValue[cacheKeyValue]
	if !ok {
		return 0
	}
	val, err := strconv.ParseUint(strVal, 10, 64)
	if err != nil {
		log.WithError(err).Warn("Couldn't parse %s to int", strVal)
		return 0
	}
	return val
}

func GetWalletSources() ([]*walletSource, error) {
	wltsString, err := getValues(SettingPathToWalletSource)
	if err != nil {
		return nil, err
	}
	wltSrcs := make([]*walletSource, len(wltsString))
	for i, wlt := range wltsString {
		wltSrcs[i] = new(walletSource)
		err = json.Unmarshal([]byte(wlt), wltSrcs[i])
		if err != nil {
			return nil, err
		}
		if wltSrcs[i].Tp == RemoteWallet {
			node, err := GetNodeSource()
			if err != nil {
				return nil, err
			}
			wltSrcs[i].Source = node[SettingNodeAddress]
		}
	}
	return wltSrcs, nil
}

func GetNodeSource() (map[string]string, error) {

	nodeSettingStr, err := GetOption(SettingPathToNode)
	if err != nil {
		return nil, err
	}
	node := make(map[string]string)
	err = json.Unmarshal([]byte(nodeSettingStr), &node)
	if err != nil {
		return nil, err
	}
	return node, nil
}

type walletSource struct {
	id     string
	Tp     string `json:"SourceType"`
	Source string `json:"Source"`
}
