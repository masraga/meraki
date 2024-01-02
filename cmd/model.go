/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// modelCmd represents the model command
var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "creating database model",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var modelScript, repoScript []string
		modelName, _ := cmd.Flags().GetString("name")
		modelFileName := fmt.Sprintf("./models/%s.go", modelName)

		//START: STRUCT CODE
		modelScript = append(modelScript, "package models \n\n")

		//IMPORT LIBRARY HERE
		modelScript = append(modelScript, "import \"go.mongodb.org/mongo-driver/bson/primitive\" \n\n")

		modelScript = append(modelScript, fmt.Sprintf("type %s struct {\n", modelName))

		modelScript = append(modelScript, "\t/*\n")
		modelScript = append(modelScript, "\t\tdefault collection field\n")
		modelScript = append(modelScript, "\t*/\n")
		modelScript = append(modelScript, "\tID primitive.ObjectID `bson:\"_id,omitempty\"`\n")
		modelScript = append(modelScript, "\tIsDeleted bool `bson:\"isDeleted\"`\n")
		modelScript = append(modelScript, "\tDeletedId primitive.ObjectID `bson:\"deletedId,omitempty\"`\n")
		modelScript = append(modelScript, "\tDeletedAt primitive.DateTime `bson:\"deletedAt,omitempty\"`\n")
		modelScript = append(modelScript, "\tCreatedAt primitive.DateTime `bson:\"createdAt\"`\n")
		modelScript = append(modelScript, "\tCreatedId primitive.ObjectID `bson:\"createdId\"`\n")
		modelScript = append(modelScript, "\tUpdatedAt primitive.DateTime `bson:\"updatedAt,omitempty\"`\n")
		modelScript = append(modelScript, "\tUpdatedId primitive.ObjectID `bson:\"updatedId,omitempty\"`\n\n")

		modelScript = append(modelScript, "\t/*\n")
		modelScript = append(modelScript, "\t\tfill your collection field below\n")
		modelScript = append(modelScript, "\t*/")

		modelScript = append(modelScript, "\n}")
		//END: STRUCT CODE

		os.WriteFile(modelFileName, []byte(strings.Join(modelScript, "")), 0664)
		fmt.Println(`model created in:`, modelFileName)

		// ================================================================================================
		repoFileName := fmt.Sprintf("./repositories/%s.go", modelName)

		// START: CREATE REPO SCRIPT
		repoScript = append(repoScript, "package repositories\n\n")

		repoScript = append(repoScript, "import(\n")
		repoScript = append(repoScript, "\t\"time\"\n\n")
		repoScript = append(repoScript, "\t\"github.com/masraga/meraki/models\"\n")
		repoScript = append(repoScript, "\t\"github.com/masraga/meraki/pkg\"\n")
		repoScript = append(repoScript, "\tdriver \"github.com/masraga/meraki/pkg/driver\"\n")
		repoScript = append(repoScript, "\trepositories \"github.com/masraga/meraki/pkg/repositories\"\n")
		repoScript = append(repoScript, "\t\"go.mongodb.org/mongo-driver/bson\"\n")
		repoScript = append(repoScript, "\t\"go.mongodb.org/mongo-driver/bson/primitive\"\n")
		repoScript = append(repoScript, "\t\"go.mongodb.org/mongo-driver/mongo\"\n")
		repoScript = append(repoScript, ")\n\n")

		repoScript = append(repoScript, fmt.Sprintf("type %s struct {\n", modelName))
		repoScript = append(repoScript, "\tDb *driver.MongodbDriver\n")
		repoScript = append(repoScript, "\tRepo *repositories.MongoRepository\n")
		repoScript = append(repoScript, fmt.Sprintf("\tModel *models.%s\n", modelName))
		repoScript = append(repoScript, "}\n\n")

		// func Create
		repoScript = append(repoScript, fmt.Sprintf("func (r *%s) Create(request models.%s) (*mongo.InsertOneResult, error) {\n", modelName, modelName))
		repoScript = append(repoScript, "\trequest.CreatedAt = primitive.NewDateTimeFromTime(time.Now())\n")
		repoScript = append(repoScript, "\trequest.IsDeleted = false \n\n")

		repoScript = append(repoScript, "\tquery, err := r.Repo.InsertOne(request)\n")
		repoScript = append(repoScript, "\tif err != nil {\n")
		repoScript = append(repoScript, "\t\treturn nil, err\n")
		repoScript = append(repoScript, "\t}\n\n")

		repoScript = append(repoScript, "\treturn query, err\n")

		repoScript = append(repoScript, "}\n\n")

		//FUNC FindById
		repoScript = append(repoScript, fmt.Sprintf("func (r *%s) FindById(id string) (*models.%s, error) {\n", modelName, modelName))
		repoScript = append(repoScript, "\terr := r.Repo.FindById(id).Decode(&r.Model)\n")
		repoScript = append(repoScript, "\tif err != nil {\n")
		repoScript = append(repoScript, "\t\treturn nil, err\n")
		repoScript = append(repoScript, "\t}\n\n")
		repoScript = append(repoScript, "\treturn r.Model, nil\n")

		repoScript = append(repoScript, "}\n\n")

		//FUNC UpdateByID
		repoScript = append(repoScript, fmt.Sprintf("func (r *%s) UpdateByID(id string, fieldSet *models.%s) (*mongo.UpdateResult, error) {\n", modelName, modelName))
		repoScript = append(repoScript, "\tfieldSet.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())\n")
		repoScript = append(repoScript, "\tupdateResult, err := r.Repo.UpdateByID(id, fieldSet)\n")
		repoScript = append(repoScript, "\treturn updateResult, err\n")
		repoScript = append(repoScript, "}\n\n")

		//FUNC DeleteByID
		repoScript = append(repoScript, fmt.Sprintf("func (r *%s) DeleteByID(id string) (*mongo.DeleteResult, error) {\n", modelName))
		repoScript = append(repoScript, "\tresult, err := r.Repo.DeleteByID(id)\n")
		repoScript = append(repoScript, "\treturn result, err\n")
		repoScript = append(repoScript, "}\n\n")

		//FUNC SoftDeleteByID
		repoScript = append(repoScript, fmt.Sprintf("func (r *%s) SoftDeleteByID(id string) (*mongo.UpdateResult, error) {\n", modelName))
		repoScript = append(repoScript, "\tr.Model.IsDeleted = true\n")
		repoScript = append(repoScript, "\tr.Model.DeletedAt = primitive.NewDateTimeFromTime(time.Now())\n")
		repoScript = append(repoScript, "\tdeleteResult, err := r.Repo.UpdateByID(id, r.Model)\n")
		repoScript = append(repoScript, "\treturn deleteResult, err\n")
		repoScript = append(repoScript, "}\n\n")

		//FUNC FindOne
		repoScript = append(repoScript, fmt.Sprintf("func (r *%s) FindOne(filter bson.D) (*models.%s, error) {\n", modelName, modelName))
		repoScript = append(repoScript, "\tresult, err := r.Repo.FindOne(filter)\n")
		repoScript = append(repoScript, "\tif err != nil {\n")
		repoScript = append(repoScript, "\t\treturn nil, err\n")
		repoScript = append(repoScript, "\t}\n\n")
		repoScript = append(repoScript, "\tresult.Decode(&r.Model)\n")
		repoScript = append(repoScript, "\treturn r.Model, err\n")
		repoScript = append(repoScript, "}\n\n")

		//FUNC Aggregate
		repoScript = append(repoScript, fmt.Sprintf("func (r *%s) Aggregate(filter bson.D) (*mongo.Cursor, error) {\n", modelName))
		repoScript = append(repoScript, "\tcursor, err := r.Repo.Aggregate(filter)\n")
		repoScript = append(repoScript, "\tif err != nil {\n")
		repoScript = append(repoScript, "\t\treturn nil, err\n")
		repoScript = append(repoScript, "\t}\n\n")
		repoScript = append(repoScript, "\treturn cursor, nil\n")
		repoScript = append(repoScript, "}\n\n")

		//FUNC NewUser
		repoScript = append(repoScript, fmt.Sprintf("func New%s() *%s {\n", modelName, modelName))
		repoScript = append(repoScript, fmt.Sprintf("\tvar model *models.%s\n", modelName))
		repoScript = append(repoScript, fmt.Sprintf("\trepo := repositories.NewMongoRepository(\"%s\")\n", modelName))
		repoScript = append(repoScript, "\tdb := pkg.NewAutoload().Database()\n")
		repoScript = append(repoScript, fmt.Sprintf("\treturn &%s{\n", modelName))
		repoScript = append(repoScript, "\t\tDb: db,\n")
		repoScript = append(repoScript, "\t\tRepo: repo,\n")
		repoScript = append(repoScript, "\t\tModel: model,\n")
		repoScript = append(repoScript, "\t}\n")
		repoScript = append(repoScript, "}\n\n")
		// END: CREATE REPO SCRIPT

		os.WriteFile(repoFileName, []byte(strings.Join(repoScript, "")), 0664)
		fmt.Println(`Repository created in:`, repoFileName)
	},
}

func init() {
	rootCmd.AddCommand(modelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	modelCmd.PersistentFlags().String("name", "", "add model name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
