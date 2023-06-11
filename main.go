package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Credentials struct {
	AccessKeyId     string
	SecretAccessKey string
	SessionToken    string
	Expiration      string
}

type SessionToken struct {
	Credentials Credentials
}

type Profile struct {
	Name   string
	Region string
	Output string
	Iam    string
	OTP    string
}

func Valid(name string, value string) string {
	if len(name) == 0 {
		return value
	}
	return name
}

func getProfile() Profile {

	var iam string = ""
	var code string = ""
	var name string = ""
	var region string = "ap-northeast-2"
	var output string = "json"

	fmt.Print("MFA ARN:")
	fmt.Scanln(&iam)

	fmt.Print("OTP Code:")
	fmt.Scanln(&code)

	fmt.Printf("SET MFA ProfileName(%s):", name)
	fmt.Scanln(&name)

	fmt.Printf("SET MFA Region($s):", region)
	fmt.Scanln(&region)

	return Profile{
		Name:   Valid(name, "mfa"),
		Region: Valid(name, "ap-northeast-2"),
		Output: Valid(output, "json"),
		Iam:    Valid(iam, ""),
		OTP:    Valid(code, ""),
	}
}

func main() {

	var token SessionToken
	var mfa Profile = getProfile()
	var profile = "default"

	cmd := exec.Command("aws", "sts", "get-session-token", "--serial-number", mfa.Iam, "--token-code", mfa.OTP, "--profile", profile)
	tokenStr, err := cmd.Output()

	if err != nil {
		fmt.Printf("please check iam/otp/aws default profile")
		panic(err)
	}

	json.Unmarshal([]byte(tokenStr), &token)

	cmd = exec.Command("aws", "configure", "set", "profile", mfa.Name)
	_, err = cmd.Output()

	if err != nil {
		fmt.Printf("please check %s", mfa.Name)
		panic(err)
	}

	var mfaProfileRegion = fmt.Sprintf("profile.%s.region", mfa.Region)
	cmd = exec.Command("aws", "configure", "set", mfaProfileRegion, mfa.Region)
	_, err = cmd.Output()

	if err != nil {
		fmt.Printf("please check %s", mfa.Region)
		panic(err)
	}

	var mfaProfileAccessKey = fmt.Sprintf("profile.%s.aws_access_key_id", mfa.Name)
	cmd = exec.Command("aws", "configure", "set", mfaProfileAccessKey, token.Credentials.AccessKeyId)
	_, err = cmd.Output()

	if err != nil {
		fmt.Printf("please check %s", token.Credentials.AccessKeyId)
		panic(err)
	}

	var mfaProfileSecretKey = fmt.Sprintf("profile.%s.aws_secret_access_key", mfa.Name)
	cmd = exec.Command("aws", "configure", "set", mfaProfileSecretKey, token.Credentials.SecretAccessKey)
	_, err = cmd.Output()

	if err != nil {
		fmt.Printf("please check %s", token.Credentials.SecretAccessKey)
		panic(err)
	}

	var mfaProfileSession = fmt.Sprintf("profile.%s.aws_session_token", mfa.Name)
	cmd = exec.Command("aws", "configure", "set", mfaProfileSession, token.Credentials.SessionToken)
	_, err = cmd.Output()

	if err != nil {
		fmt.Printf("please check %s", token.Credentials.SessionToken)
		panic(err)
	}

	var mfaProfileOutput = fmt.Sprintf("profile.%s.output", mfa.Name)
	cmd = exec.Command("aws", "configure", "set", mfaProfileOutput, mfa.Output)
	_, err = cmd.Output()

	if err != nil {
		fmt.Printf("please check %s", mfa.Output)
		panic(err)
	}

	fmt.Printf("testing ...")
	fmt.Printf("aws s3 ls --profile %s", mfa.Name)
}
