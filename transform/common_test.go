package transform

import (
	"strings"
	"testing"
	"github.com/beevik/etree"
	"github.com/andreluzz/cas-xog/common"
)

var packageMockFolder string

func init() {
	packageMockFolder = "../" + common.FOLDER_MOCK + "transform/"
}

func TestExecuteToReturnErrorNoHeaderElement(t *testing.T) {
	xog := etree.NewDocument()
	err := Execute(xog, nil, common.DriverFile{})
	if err == nil {
		t.Fatalf("Error executing transformation. Not testing if xog has element head.")
	}
}

func TestExecuteToReturnPage(t *testing.T) {
	file := common.DriverFile{
		Type: common.PAGE,
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "page_full_xog.xml")
	err := Execute(xog, nil, file)

	if err != nil {
		t.Fatalf("Error transforming page XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "page_result.xml") == false {
		t.Errorf("Error transforming page XOG file. Invalid result XML.")
	}
}

func TestExecuteToReturnPageWithoutElementOBSandSecurity(t *testing.T) {
	file := common.DriverFile{
		Type: common.PAGE,
		Elements: []common.Element{
			{
				Action: "remove",
				XPath: "//OBSAssocs",
			},
			{
				Action: "remove",
				XPath: "//Security",
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "page_full_xog.xml")
	err := Execute(xog, nil, file)

	if err != nil {
		t.Fatalf("Error transforming page XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "page_no_element_result.xml") == false {
		t.Errorf("Error transforming page XOG file. Invalid result XML.")
	}
}

func TestExecuteToReturnGroup(t *testing.T) {
	file := common.DriverFile{
		Code: "ObjectAdmin",
		Type: common.GROUP,
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "group_full_xog.xml")
	err := Execute(xog, nil, file)

	if err != nil {
		t.Fatalf("Error transforming group XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "group_result.xml") == false {
		t.Errorf("Error transforming group XOG file. Invalid result XML.")
	}
}

func TestExecuteToReturnGroupWithoutMembers(t *testing.T) {
	file := common.DriverFile{
		Code: "ObjectAdmin",
		Type: common.GROUP,
		Elements: []common.Element{
			{
				Action: "remove",
				XPath: "//members",
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "group_full_xog.xml")
	err := Execute(xog, nil, file)

	if err != nil {
		t.Fatalf("Error transforming group XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "group_no_members_result.xml") == false {
		t.Errorf("Error transforming group XOG file. Invalid result XML.")
	}
}

func TestExecuteToReturnPortletFromQuery(t *testing.T) {
	file := common.DriverFile{
		Code: "apm.appByQuadrant",
		Type: common.PORTLET,
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "portlet_query_full_xog.xml")
	err := Execute(xog, nil, file)

	if err != nil {
		t.Fatalf("Error transforming query portlet XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "portlet_query_result.xml") == false {
		t.Errorf("Error transforming query portlet XOG file. Invalid result XML.")
	}
}

func TestExecuteToReturnPortletFromObject(t *testing.T) {
	file := common.DriverFile{
		Code: "test_cas_xog",
		Type: common.PORTLET,
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "portlet_object_full_xog.xml")
	err := Execute(xog, nil, file)

	if err != nil {
		t.Fatalf("Error transforming object portlet XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "portlet_object_result.xml") == false {
		t.Errorf("Error transforming object portlet XOG file. Invalid result XML.")
	}
}

func TestExecuteToReturnQuery(t *testing.T) {
	file := common.DriverFile{
		Code: "cop.processBottlenecks",
		Type: common.QUERY,
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "query_full_xog.xml")
	err := Execute(xog, nil, file)

	if err != nil {
		t.Fatalf("Error transforming query XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "query_result.xml") == false {
		t.Errorf("Error transforming query XOG file. Invalid result XML.")
	}
}

func readMockResultAndCompare(xog *etree.Document, compareXml string) bool {
	xog.Indent(2)
	xogString, _ := xog.WriteToString()
	xogString = strings.Replace(xogString, " ", "", -1)

	xogProcessedToCompare := etree.NewDocument()
	xogProcessedToCompare.ReadFromFile(packageMockFolder + compareXml)
	xogProcessedToCompare.Indent(2)

	xogProcessedToCompareString, _ := xogProcessedToCompare.WriteToString()
	xogProcessedToCompareString = strings.Replace(xogProcessedToCompareString, " ", "", -1)
	if xogString != xogProcessedToCompareString {
		return false
	}
	return true
}