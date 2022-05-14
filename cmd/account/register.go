package account

import (
	"KDRC-Client/global"
	"context"
	"github.com/spf13/cobra"
	"log"
	"net/mail"
)
import openapiclient "github.com/KangDroid-Cloud/CoreNetworkCommunication"

var (
	registerEmail    string
	registerPassword string
	registerNicKName string
)

var registrationCommand = &cobra.Command{
	Use:     "register {args}",
	Short:   "Register to KangDroid-Cloud Service!",
	Example: "./kdrc account register --email {email address} --password {password} --nickName {nickname}",
	Run: func(cmd *cobra.Command, args []string) {
		// Create Request
		registrationRequest := *openapiclient.NewAccountRegisterRequest()
		registrationRequest.Email = &registerEmail
		registrationRequest.Password = &registerPassword
		registrationRequest.NickName = &registerNicKName

		// Setup API Location
		apiClient := *openapiclient.NewAPIClient(&global.DefaultServerConfiguration)

		response, err := apiClient.AccountApi.ApiAccountRegisterPost(context.Background()).AccountRegisterRequest(registrationRequest).Execute()

		if err != nil {
			if response == nil {
				log.Fatalf(err.Error())
			}
			switch response.StatusCode {
			case 400:
				log.Fatalf("Request Prevalidation failed!\n")
			case 409:
				log.Fatalf("Account with email %s already exists!\n", registerEmail)
			default:
				log.Fatalf("Unknown error occurred. Status Code: %d, Error Message: %s\n", response.StatusCode, err.Error())
			}
		} else {
			log.Println("Successfully registered to service! Please login.")
		}
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		// Validate Arguments
		_, emailValidationErr := mail.ParseAddress(registerEmail)

		if emailValidationErr != nil {
			log.Fatalf("Cannot validate email address. Are you using correct email address?")
		}

		if len(registerPassword) < 8 {
			log.Fatalf("Password length is at least more than 8 letters or more!")
		}
	},
}

func init() {
	registrationCommand.Flags().StringVarP(&registerEmail, "email", "E", "", "Email Address")
	registrationCommand.Flags().StringVarP(&registerPassword, "password", "P", "", "Password")
	registrationCommand.Flags().StringVarP(&registerNicKName, "nickName", "N", "", "NickName")

	registrationCommand.MarkFlagRequired("email")
	registrationCommand.MarkFlagRequired("password")
	registrationCommand.MarkFlagRequired("nickName")
}
