
func GetAndUpdateCount() (int, error) {

	db := dynamo.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
	table := db.Table("MyVisitorCounter")

	var result MyVisitorCounter
	err := table.Update("path", "getcrazygolang").
		SetExpr("VisitorCount = VisitorCount + ?", 1).
		Value(&result)

	return result.VisitorCount, err

}

func HandleRequest(ctx context.Context) (string, error) {
	count, err := GetAndUpdateCount()
	return fmt.Sprintln(count), err
}

func main() {
	lambda.Start(HandleRequest)
}
