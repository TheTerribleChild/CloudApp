package sourceconnector

import (
	"errors"
	"fmt"
	"github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/novel_application/connectors/service"
	"github.com/TheTerribleChild/cloud_appplication_portal/commons/utils/fileutil"
	"plugin"
)

//Connector Used to connect to various sources
type Connector interface {
	GetConnectorID() string
	GetConnectorName() string
	GetChapterList(string) (connectorservice.NovelSourceData, error)
	GetChapterContent(connectorservice.ChapterSourceMetadata) (connectorservice.ChapterSourceData, error)
}

var sourceConnectorLocation = "./sourceconnectors/"
var sourceConnectorMap map[string]Connector
var sourceConnectorIDList []string

//GetSourceConnector Gets a source connector by name
func GetSourceConnector(connectorName string) (Connector, error) {
	if len(connectorName) == 0 {
		return nil, errors.New("connectorName parameter cannot be null")
	}

	if connector, ok := sourceConnectorMap[connectorName]; ok {
		return connector, nil
	}
	return nil, fmt.Errorf("source connector '%s' not found", connectorName)
}

//GetSourceConnectorIDs Get list of source connector loaded
func GetSourceConnectorIDs() []string {
	return sourceConnectorIDList
}

func init() {
	sourceConnectorMap = make(map[string]Connector)
	loadSourceConnectors()
}

func loadSourceConnectors() error {
	files, err := fileutil.GetAllFileInDirectory(sourceConnectorLocation, ".so")

	if err != nil {
		return err
	}
	for _, file := range files {
		connector, err := loadPlugin(sourceConnectorLocation + file)
		if err == nil {
			connectorID := connector.GetConnectorID()
			sourceConnectorIDList = append(sourceConnectorIDList, connectorID)
			sourceConnectorMap[connectorID] = connector
			fmt.Println(connector.GetConnectorName() + " successfully loaded.")
		} else {
			fmt.Println(err)
		}
	}
	return nil
}

func loadPlugin(connectorFileName string) (Connector, error) {
	if len(connectorFileName) == 0 {
		return nil, errors.New("connectorFileName parameter cannot be null")
	}

	if connector, ok := sourceConnectorMap[connectorFileName]; ok {
		return connector, nil
	}
	fmt.Println("Loading plugin " + connectorFileName)
	plug, err := plugin.Open(connectorFileName)
	if err != nil {
		return nil, fmt.Errorf("unable to load plugin '%s'. %s", connectorFileName, err)
	}
	symConnector, err := plug.Lookup("Connector")
	if err != nil {
		return nil, fmt.Errorf("unable to load plugin '%s'. %s", connectorFileName, err)
	}
	connector, ok := symConnector.(Connector)

	if !ok {
		return nil, fmt.Errorf("unable to load plugin '%s'. Unexpected type from module symbol", connectorFileName)
	}
	fmt.Println("here")
	sourceConnectorMap[connectorFileName] = connector

	return connector, nil
}
