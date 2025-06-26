package main

import (
	"net/http"
	"notification-service/controller"
	service "notification-service/emailservice"
	"notification-service/router"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	orderNotificationSerice := service.NewOrderConfrimationEmailSender()
	orderNotificationController := controller.NewOrderNotificationController(orderNotificationSerice)
	paymentNotificationSerice := service.NewPaymentConfrimationEmailSender()
	paymentNotificationController := controller.NewPaymentNotificationController(paymentNotificationSerice)
	notificationRouters := router.NotificatioRouter(orderNotificationController, paymentNotificationController)
	server := http.Server{
		Addr:    ":9095",
		Handler: notificationRouters,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
