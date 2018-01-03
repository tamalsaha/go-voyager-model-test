package main

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	tapi "github.com/appscode/voyager/apis/voyager/v1beta1"
	"fmt"
)

func main()  {
	ing := &tapi.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-voyager",
			Namespace: "default",
		},
		Spec: tapi.IngressSpec{
			Rules:         make([]tapi.IngressRule, 0),
			FrontendRules: make([]tapi.FrontendRule, 0),
		},
	}
	fmt.Println(ing)
}
