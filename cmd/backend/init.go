/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package backend

import (
	"encoding/json"
	"fmt"
	"os"
	"password_manager/cmd/utils"

	"github.com/spf13/cobra"
)

// intCmd represents the int command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "file for your credentials",
	Long: `Json file for your credentials`,
	Run: func(cmd *cobra.Command, args []string) {
		createExistingFile()	
	},
}

func init() {

	initCmd.Flags().StringVarP(&fileName, "filename", "f", "", "file name for your credentials")
	initCmd.Flags().StringVarP(&author, "author", "a", "", "author of the file")
	initCmd.Flags().BoolVarP(&jsonfile, "jsonfile", "j", false, "insert true to the flag argument to create json file in any way")
	initCmd.Flags().BoolVarP(&dbfile, "dbfile", "d", false, "insert true to the flag argument to create dbfile in any way")


	// Here you will define your flags and configuration settings.
	BackendCmd.AddCommand(initCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// intCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// intCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createJsonFile() {

	userCredentials := utils.Storage{}
	
	userCredentials.Author = author

	finalJson, err := json.MarshalIndent(userCredentials, "", "\t")
	if err != nil {
		panic(err)
	}

	if fileName != "" {
		os.WriteFile(fileName + ".json", finalJson, 0666)
		}else{
		os.WriteFile("credentials.json", finalJson, 0666)
	}

}

func createExistingFile() {
	_ , err := os.ReadFile("credentials.json")
	if err != nil {
		createJsonFile()
	}else if err == nil && jsonfile{
		createJsonFile()
	}else {
		fmt.Println("File is already existing")
		os.Exit(1)
	}
}

  
  func main() {
	
  }