package main

//https://www.cnblogs.com/junneyang/p/6211190.html

import (
	"encoding/json"
	"fmt"
	"github.com/k-sone/snmpgo"
	"io/ioutil"
	"net/http"
	//"strconv"
	"time"
)

func SnmpWalk(ip string, community string, targetOid string) []byte {
	snmp, err := snmpgo.NewSNMP(snmpgo.SNMPArguments{
		Version:   snmpgo.V2c,
		Address:   ip + ":161",
		Retries:   1,
		Timeout:   time.Second * 1,
		Community: community,
	})
	if err != nil {
		fmt.Println(err)
		return []byte("{\"code\":-1}")
	}

	oids, err := snmpgo.NewOids([]string{targetOid})
	if err != nil {
		fmt.Println(err)
		return []byte("{\"code\":-1}")
	}

	if err = snmp.Open(); err != nil {
		fmt.Println(err)
		return []byte("{\"code\":-1}")
	}
	defer snmp.Close()

	pdu, err := snmp.GetBulkWalk(oids, 0, 10)
	if err != nil {
		fmt.Println(err)
		return []byte("{\"code\":-1}")
	}
	if pdu.ErrorStatus() != snmpgo.NoError {
		fmt.Println(pdu.ErrorStatus(), pdu.ErrorIndex())
	}

	// //datas := make(map[int]string)
	// datas := make(map[string]string)
	// for _, v := range pdu.VarBinds() {
	// 	//fmt.Println(v.Oid, v.Variable)
	// 	// key, err := strconv.Atoi(v.Oid.String()[len(targetOid)+1:]) //key只要最后的数字
	// 	// if err !=nil{
	// 	//     return []byte("{\"code\":-2}")
	// 	// }
	// 	key := v.Oid.String()[len(targetOid)+1:]
	// 	datas[key] = v.Variable.String()
	// }
	// json_datas, _ := json.Marshal(datas)

	datas := make(map[string]interface{})
	datas["code"] = 0
	datas["data"] = make(map[string]string)
	for _, v := range pdu.VarBinds() {
		key := v.Oid.String()[len(targetOid)+1:]
		datas["data"].(map[string]string)[key] = v.Variable.String()
	}
	json_datas, _ := json.Marshal(datas)

	return json_datas
}

func ApiSnmpWalk(w http.ResponseWriter, r *http.Request) {
	//读取内容
	body, _ := ioutil.ReadAll(r.Body)
	//json转map
	var req map[string]string
	err := json.Unmarshal(body, &req)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(req)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(SnmpWalk(req["ip"], req["community"], req["targetOid"]))
}

func main() {
	http.HandleFunc("/api/snmpwalk", ApiSnmpWalk)
	fmt.Println("服务器已启动：0.0.0.0:8085/api/snmpwalk")
	http.ListenAndServe("0.0.0.0:8085", nil)
}
