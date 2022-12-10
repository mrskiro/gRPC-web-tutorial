import {
  createConnectTransport,
  createPromiseClient,
} from "@bufbuild/connect-web"

import { NotificationService } from "./gen/proto/v1/notification_connectweb"

const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
})

export const notificationClient = createPromiseClient(
  NotificationService,
  transport
)
