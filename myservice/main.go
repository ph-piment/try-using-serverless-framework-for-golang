package main

import (
        "encoding/json"
        "fmt"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/lambda"
)

// for response
type Response struct {
        Message string `json:"message"`
        Ok      bool   `json:"ok"`
}

// data to input to Lambda
type Data struct {
        Key1 string `json:"key1"`
        Key2 string `json:"key2"`
        Key3 string `json:"key3"`
}

// function to call Lambda
func Handler() (Response, error) {
        // Data that is to send to the Lambda
        //payload := Data{}
        reqPayload := Data{
                Key1: "1",
                Key2: "2",
                Key3: "3",
        }
        // and transform it to json bytes.
        jsonBytes, err := json.Marshal(reqPayload)
        if err != nil {
                return Response{
                        Message: err.Error(),
                        Ok:      false,
                }, err
        }

        // To make session
        profile_name := "hogehoge"
        mySession := session.Must(session.NewSessionWithOptions(session.Options{Profile:profile_name}))

        // To set region you use.
        region := "fugafuga"
        svc := lambda.New(mySession, aws.NewConfig().WithRegion(region))

        // set the Lambda's arn.
        arn := "hogefuga"
        input := &lambda.InvokeInput{
                // To set your Labmda's arn
                FunctionName: aws.String(arn),
                Payload:      jsonBytes,
                // InvocationType Customization is available.
                // RequestResoponse is synchronous mode.(default)
                // Event is asynchronous mode.
                //InvocationType: aws.String("Event"),
                //InvocationType: aws.String("RequestResponse"),
        }

        // To call Lambda
        res, err := svc.Invoke(input)
        if err != nil {
                return Response{
                        Message: err.Error(),
                        Ok:      false,
                }, err
        }
        resPayload := string(res.Payload)

        return Response{
                Message: resPayload,
                Ok:      true,
        }, nil
}

func main() {
        ret, _ := Handler()
        fmt.Println(ret)
}
