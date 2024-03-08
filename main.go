package main

import (
	"fmt"
	"time"

	rtctokenbuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/RtcTokenBuilder"
)

func main() {
	// Replace "YOUR_AGORA_APP_ID" with your actual Agora app ID
	appId := "87f323efeda84280a6a4f35ad5e93f80"
	// Replace "YOUR_AGORA_APP_CERTIFICATE" with your actual Agora app certificate
	appCertificate := "afb25a096bdf427da031b8154b9981a9"

	channelName := "7d72365eb983485397e3e3f9d460bdda"
	uid := uint32(2882341273)
	uidStr := "2882341273"
	expireTimeInSeconds := uint32(3600)
	currentTimestamp := uint32(time.Now().UTC().Unix())
	expireTimestamp := currentTimestamp + expireTimeInSeconds

	fmt.Println("App Id:", appId)
	fmt.Println("App Certificate:", appCertificate)

	result, err := rtctokenbuilder.BuildTokenWithUID(appId, appCertificate, channelName, uid, rtctokenbuilder.RoleAttendee, expireTimestamp)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Token with uid: %s\n", result)
	}

	result, err = rtctokenbuilder.BuildTokenWithUserAccount(appId, appCertificate, channelName, uidStr, rtctokenbuilder.RoleAttendee, expireTimestamp)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Token with userAccount: %s\n", result)
	}
}
