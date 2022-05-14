package cmd

import (
    "KDRC-Client/global"
    "context"
    "encoding/json"
    openapi "github.com/KangDroid-Cloud/CoreNetworkCommunication"
    "github.com/spf13/cobra"
    "log"
)

var (
    loginEmail    string
    loginPassword string
)

var loginCmd = &cobra.Command{
    Use:     "login",
    Short:   "Login to Cloud Service",
    Example: "kdrc login --email {email} --password {password}",
    Run: func(cmd *cobra.Command, args []string) {
        // Create Request
        loginRequest := openapi.LoginRequest{
            Email:    &loginEmail,
            Password: &loginPassword,
        }

        // API
        apiClient := *openapi.NewAPIClient(&global.DefaultServerConfiguration)

        // Request it
        loginResponse, rawResponse, err := apiClient.AuthApi.ApiAuthLoginPost(context.Background()).LoginRequest(loginRequest).Execute()

        if err != nil {
            if rawResponse == nil {
                log.Fatalf(err.Error())
            }
            switch rawResponse.StatusCode {
            case 401:
                log.Fatalf("Please check your email add password again.")
            default:
                log.Fatalf("Unknown error occurred. Status Code: %d, Error Message: %s\n", rawResponse.StatusCode, err.Error())
            }
        }

        // Save Login Response.
        handleLoginResponse(loginResponse)
    },
}

func handleLoginResponse(loginResponse *openapi.LoginResponse) {
    marshalResult, err := json.Marshal(loginResponse)

    if err != nil {
        log.Fatalf("Failed to save login response. Error: %s\n", err.Error())
    }

    global.SaveTokenToDisk(marshalResult)
}

func init() {
    loginCmd.Flags().StringVarP(&loginEmail, "email", "E", "", "Email Address")
    loginCmd.Flags().StringVarP(&loginPassword, "password", "P", "", "Password")

    loginCmd.MarkFlagRequired("email")
    loginCmd.MarkFlagRequired("password")
}
