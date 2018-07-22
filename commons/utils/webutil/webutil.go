package webutil

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/htmlindex"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type CharSet string

const (
	GB2312 CharSet = "GB2312"
	UTF8   CharSet = "UTF-8"
)

func GetURLContentAsString(url string, charset string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		reader, _ := decode(resp.Body, charset, int(resp.ContentLength))
		contents, _ := ioutil.ReadAll(reader)
		return string(contents), nil
	}
	return "", nil
}

func GetURLContentAsNode(url string, charset string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		reader, _ := decode(resp.Body, charset, int(resp.ContentLength))
		node, err := html.Parse(reader)
		return node, err
	}
	return nil, nil
}

type NodeSignature struct {
	Signature string
	Node      *html.Node
}

func PrintNodeSignature(url string, charset string) {
	node, _ := GetURLContentAsNode(url, charset)
	signatures := GenerateNodeSignature(node)
	for i, sig := range signatures {
		count := strings.Count(sig.Signature, "/")
		head := strings.Repeat("|      ", count)
		fmt.Printf("%s====Node====\n", head)
		fmt.Printf("%sData      : %s\n", head, strings.Replace(sig.Node.Data, "\n", "\n"+head, -1))
		fmt.Printf("%sSignature : %s\n", head, sig.Signature)
		fmt.Printf("%sIndex     : %d\n", head, i)
		for _, attr := range sig.Node.Attr {
			fmt.Printf("%s{\n", head)
			fmt.Printf("%s  Key: %s\n%s  Val: %s\n", head, attr.Key, head, attr.Val)
			fmt.Printf("%s}\n", head)
		}
		fmt.Printf("%s============\n", head)
	}
}

func GenerateNodeSignature(node *html.Node) []NodeSignature {

	nodeSignatures := []NodeSignature{}
	var visit func(node *html.Node, signature string)
	visit = func(node *html.Node, signature string) {
		nodeSignatures = append(nodeSignatures, NodeSignature{Signature: signature, Node: node})
		index := 0
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			visit(c, fmt.Sprintf("%s/%d", signature, index))
			index++
		}
	}
	visit(node, "root")
	return nodeSignatures
}

func decode(body io.Reader, charset string, length int) (io.Reader, error) {
	if charset == "" {
		charset = detectContentCharset(body, length)
		fmt.Println(charset)
	}
	e, err := htmlindex.Get(charset)
	if err != nil {
		return nil, err
	}

	body = e.NewDecoder().Reader(body)
	return body, nil

}

func detectContentCharset(body io.Reader, length int) string {
	r := bufio.NewReader(body)
	data, _ := r.Peek(length)
	_, name, _ := charset.DetermineEncoding(data, "")
	return name
}
