package main
import (
	"strings"
)

func CopyManifests(sourceUrl string, sourceManifests []string, destPath string) error {
	//var mergeManifestAddr map[string]interface{}
	for i := range sourceManifests {
		//var manifestAddr map[string]interface{}
		httpPath := sourceUrl + sourceManifests[i]
		err := CopyHttpToFile(httpPath, destPath + sourceManifests[i])
		if err != nil {
			return err
		}
		//mergeManifestAddr = *Merge(&manifestAddr, &mergeManifestAddr)
	}
	return nil
}
func AppMap() error {
	// You would parse all the applications as they have same structure we just do a couple
	// one is standard NGINX Ingress Controller. The other is Custom Application
	//
	var inputFile= "apps.yaml"
	var inputAddr map[string]interface{}
	err := GetMap(inputFile, &inputAddr)
	if err != nil {
		return err
	}

	root := inputAddr["root"].(string)
	appName := inputAddr["name"].(string)
	destDir := inputAddr["dest-dir"].(string)
	destPath := root + "/" + appName + "/" + destDir + "/";

	source := inputAddr["source"].(map[interface{}]interface{})
	sourceUrl := source["url"].(string)
	sourceManifests := strings.Split(source["manifests"].(string), " ")
	err = CopyManifests(sourceUrl, sourceManifests, destPath)
	if err != nil {
		return err
	}

	fork := inputAddr["fork"].(map[interface{}]interface{})
	forkUrl := fork["url"].(string)
	forkManifests := strings.Split(fork["manifests"].(string), " ")
	err = CopyManifests(forkUrl, forkManifests, destPath)
	if err != nil {
		return err
	}
	////var mergeManifestAddr map[string]interface{}
	//for i := range sourceManifests {
	//	//var manifestAddr map[string]interface{}
	//	httpPath := sourceUrl + sourceManifests[i]
	//	destPath := root + "/" + appName + "/"+ destDir + "/" + sourceManifests[i];
	//	err := CopyHttpToFile(httpPath, destPath)
	//	if err != nil {
	//		return err
	//	}
	//	//mergeManifestAddr = *Merge(&manifestAddr, &mergeManifestAddr)
	//}

	//d, err := yaml.Marshal(mergeManifestAddr)
	//if err != nil {
	//	return err
	//}
	//fmt.Printf("--- t dump:\n%s\n\n", string(d))

	//
	// var sourceAddr map[string]interface{}
	//for k, v := range sourceAddr  {
	//		err := GetMap(f, &secretAddr)
	//	if err != nil {
	//		return err
	//	}
	//}
	//var sourceAddr map[string]interface{}
	////var mapAddr map[string]interface{}
	//gitHubSource := "https://raw.githubusercontent.com/kubernetes/ingress-nginx/nginx-0.10.2/deploy/without-rbac.yaml"
	//
	//err := GetSourceHttp(gitHubSource, &sourceAddr)
	//if err != nil {
	//	return err
	//}
	//
	////mapYamlFile := "apps/ingress-nginx/without-rbac.yaml"
	//
	////err = GetMap(mapYamlFile, &sourceAddr)
	////if err != nil {
	////	return err
	////}
	//
	////result := Merge(&sourceAddr, &mapAddr)
	//result := sourceAddr
	//
	//fmt.Println(yaml.Marshal(result))
	//
	//d, err := yaml.Marshal(result)
	//if err != nil {
	//	return err
	//}
	//fmt.Printf("--- t dump:\n%s\n\n", string(d))

	return nil
}

